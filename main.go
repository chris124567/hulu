package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"

	hulu "github.com/chris124567/hulu/client"
	"github.com/chris124567/hulu/widevine"
	"lukechampine.com/flagg"
)

func main() {
	rootCmd := flagg.Root
	rootCmd.Usage = flagg.SimpleUsage(rootCmd, `Hulu Downloader
It is necessary to specify the HULU_SESSION environment variable because the Hulu API requires this for all requests.

Subcommands:
search [query] - searches Hulu with the provided query and returns titles and their Hulu IDs
season [id] [season number] - lists episode title and IDs of a given show and season
download [id] - prints the MPD url the video is available at and returns the mp4decrypt command necessary to decrypt it
`)

	searchCmd := flagg.New("search", "Search Hulu for a movie or series.")
	searchQuery := searchCmd.String("query", "", "Search query.")

	seasonCmd := flagg.New("season", "Get information about season in a show by its show ID and season number.")
	seasonID := seasonCmd.String("id", "", "ID of series.")
	seasonNumber := seasonCmd.Int("number", 1, "Season number.")

	downloadCmd := flagg.New("download", "Download a show episode or movie by its ID.")
	downloadID := downloadCmd.String("id", "", "ID of movie or episode.")

	tree := flagg.Tree{
		Cmd: rootCmd,
		Sub: []flagg.Tree{
			{Cmd: searchCmd},
			{Cmd: seasonCmd},
			{Cmd: downloadCmd},
		},
	}
	cmd := flagg.Parse(tree)

	huluGUID := os.Getenv("HULU_GUID")
	// if GUID is not provided, use hash of hostname instead
	if huluGUID == "" {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		uuid := uuid.UUID(md5.Sum([]byte(hostname)))
		huluGUID = strings.ReplaceAll(strings.ToUpper(uuid.String()), "-", "")
	}

	huluSession := os.Getenv("HULU_SESSION")
	if huluSession == "" {
		rootCmd.Usage()
		return
	}

	client := hulu.NewDefaultClient(huluSession, huluGUID)
	w := tabwriter.NewWriter(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()

	switch cmd {
	case searchCmd:
		if !flagg.IsDefined(cmd, "query") {
			cmd.Usage()
			return
		}
		results, err := client.Search(*searchQuery)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s\t%s\t\n", "Title", "ID")
		for _, group := range results.Groups {
			for _, result := range group.Results {
				fmt.Fprintf(w, "%s\t%s\n", result.Visuals.Headline.Text, result.MetricsInfo.TargetID)
			}
		}
	case seasonCmd:
		if !flagg.IsDefined(cmd, "id") || !flagg.IsDefined(cmd, "number") {
			cmd.Usage()
			return
		}
		results, err := client.Season(*seasonID, *seasonNumber)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s\t%s\t\n", "Title", "ID")
		for _, item := range results.Items {
			fmt.Fprintf(w, "%s\t%s\t\n", item.Name, item.ID)
		}
	case downloadCmd:
		if !flagg.IsDefined(cmd, "id") {
			cmd.Usage()
			return
		}

		playbackInformation, err := client.PlaybackInformation(*downloadID)
		if err != nil {
			panic(err)
		}

		serverConfig, err := client.ServerConfig()
		if err != nil {
			panic(err)
		}

		playlist, err := client.Playlist(serverConfig.KeyID, playbackInformation.EabID)
		if err != nil {
			panic(err)
		}

		client := &http.Client{
			Timeout: 10 * time.Second,
		}

		// request MPD file
		response, err := client.Get(playlist.StreamURL)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		// parse init data/PSSH from XML
		initData, err := widevine.InitDataFromMPD(response.Body)
		if err != nil {
			panic(err)
		}

		cdm, err := widevine.NewDefaultCDM(initData)
		if err != nil {
			panic(err)
		}

		licenseRequest, err := cdm.GetLicenseRequest()
		if err != nil {
			panic(err)
		}

		request, err := http.NewRequest(http.MethodPost, playlist.WvServer, bytes.NewReader(licenseRequest))
		if err != nil {
			panic(err)
		}
		// hulu actually checks for headers here so this is necessary
		request.Header = hulu.StandardHeaders()
		request.Close = true
		// send license request to license server
		response, err = client.Do(request)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()
		licenseResponse, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		// parse keys from response
		keys, err := cdm.GetLicenseKeys(licenseRequest, licenseResponse)
		if err != nil {
			panic(err)
		}

		command := "mp4decrypt input.mp4 output.mp4"
		for _, key := range keys {
			if key.Type == widevine.License_KeyContainer_CONTENT {
				command += " --key " + hex.EncodeToString(key.ID) + ":" + hex.EncodeToString(key.Value)
			}
		}
		fmt.Println("MPD URL: ", playlist.StreamURL)
		fmt.Println("Decryption command: ", command)
		return
	}
}
