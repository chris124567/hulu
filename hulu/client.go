package hulu

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"lukechampine.com/frand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	contentTypeJSON = "application/json"
	contentTypeForm = "application/x-www-form-urlencoded"
)

type Client struct {
	c           *http.Client
	huluSession string
	huluGUID    string
}

// Returns a Client object that will use the provided Hulu session cookie to
// interact with the Hulu API.
func NewClient(c *http.Client, huluSession string) Client {
	// they look something like 5E95F69687FDD039CD0388A39FC01E5A
	huluGUID := func() (s string) {
		c := []byte("ABCDEF0123456789")
		for i := 0; i < 32; i++ {
			s += string(c[frand.Intn(len(c))])
		}
		return
	}()

	return Client{c, huluSession, huluGUID}
}

// Returns a Client object using a default HTTP client with a timeout of 10s.
func NewDefaultClient(huluSession string) Client {
	return NewClient(&http.Client{
		Timeout: 10 * time.Second,
	}, huluSession)
}

// Makes an HTTP request to a Hulu API endpoint.  The only cookie Hulu validates is
// the session cookie so we just provide it alone.
func (c Client) request(method string, url string, data io.Reader, contentType string) (*http.Response, error) {
	request, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, err
	}
	request.Close = true
	request.Header = StandardHeaders()
	request.Header.Set("Cookie", "_hulu_session="+c.huluSession)
	if method == http.MethodPost && len(contentType) > 0 {
		request.Header.Set("Content-Type", contentType)
	}
	return c.c.Do(request)
}

// Queries the Hulu entity search API endpoint for shows and movies.  This can
// return content that you do not have the right subscription for (like stuff
// requiring an HBO subscription) so be mindful of that.
func (c Client) Search(query string) (s SearchResults, err error) {
	query = url.QueryEscape(query)
	response, err := c.request(http.MethodGet, fmt.Sprintf("https://discover.hulu.com/content/v5/search/entity?language=en&device_context_id=2&search_query=%s&limit=64&include_offsite=true&v=26e1061d-68ec-48bf-be5a-b2f704d37256&schema=1&device_info=web:3.29.0&referralHost=production&keywords=%s&type=entity&limit=64", query, query), nil, "")
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&s)
	return
}

// Returns the season information containing the episode list in a given season
// for a given show.
func (c Client) Season(id string, season int) (s Season, err error) {
	response, err := c.request(http.MethodGet, fmt.Sprintf("https://discover.hulu.com/content/v5/hubs/series/%s/season/%d?limit=999&schema=1&offset=0&device_info=web:3.29.0&referralHost=production", id, season), nil, "")
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&s)
	return
}

// The /config endpoint returns a large hex encoded string.  This string then
// has to be decoded using a hardcoded key from Hulu.  The decoded data is JSON
// containing a bunch of configuration options for the player.  More importantly,
// it contains the KeyID field which is needed to call Playlist.
func (c Client) ServerConfig() (co Config, err error) {
	rv := strconv.Itoa(int(frand.Uint64n(1e6)))
	base := strings.Join([]string{hex.EncodeToString(deejayKey), strconv.Itoa(deejayDeviceID), strconv.Itoa(deejayKeyVersion), rv}, ",")
	nonce := md5.Sum([]byte(base))

	values := url.Values{}
	values.Add("app_version", strconv.Itoa(deejayKeyVersion))
	values.Add("badging", "true")
	values.Add("device", strconv.Itoa(deejayDeviceID))
	values.Add("device_id", c.huluGUID)
	values.Add("encrypted_nonce", hex.EncodeToString(nonce[:]))
	values.Add("language", "en")
	values.Add("region", "US")
	values.Add("rv", rv)
	values.Add("version", strconv.Itoa(deejayKeyVersion))

	response, err := c.request(http.MethodPost, "https://play.hulu.com/config", strings.NewReader(values.Encode()), contentTypeForm)
	if err != nil {
		return
	}
	defer response.Body.Close()

	ciphertext, err := io.ReadAll(hex.NewDecoder(response.Body))
	if err != nil {
		return
	}

	block, err := aes.NewCipher(deejayKey)
	if err != nil {
		return
	}

	dec := cipher.NewCBCDecrypter(block, make([]byte, 16))
	unpad := func(b []byte) []byte {
		if len(b) == 0 {
			return b
		}
		// pks padding is designed so that the value of all the padding bytes is
		// the number of padding bytes repeated so to figure out how many
		// padding bytes there are we can just look at the value of the last
		// byte
		// i.e if there are 6 padding bytes then it will look at like
		// <data> 0x6 0x6 0x6 0x6 0x6 0x6
		count := int(b[len(b)-1])
		return b[0 : len(b)-count]
	}
	plaintext := make([]byte, len(ciphertext))
	dec.CryptBlocks(plaintext, ciphertext)
	err = json.Unmarshal(unpad(plaintext), &co)
	return
}

// This allows us to get the EAB ID for a given plain ID.  The EAB ID is
// necessary to call Playlist.
func (c Client) PlaybackInformation(id string) (p PlaybackInformation, err error) {
	response, err := c.request(http.MethodGet, fmt.Sprintf("https://discover.hulu.com/content/v5/deeplink/playback?namespace=entity&id=%s&schema=1&device_info=web:3.29.0&referralHost=production", id), nil, "")
	if err != nil {
		return
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&p)
	return
}

// Playlist returns information containing the Widevine license endpoint,
// the MPD file URL, and information relating to subtitles (Hulu calls them
// transcripts).
func (c Client) Playlist(sessionKey int, eabID string) (p Playlist, err error) {
	randUUID := func() (s string) {
		randChars := func(n int) (s string) {
			c := []byte("ABCDEF0123456789")
			for i := 0; i < 4; i++ {
				s += string(c[frand.Intn(len(c))])
			}
			return
		}
		return strings.Join([]string{randChars(8), randChars(4), randChars(4), randChars(4), randChars(12)}, "-")
	}

	playlistRequest := PlaylistRequest{
		DeviceIdentifier:       c.huluGUID + ":d40b",
		DeejayDeviceID:         deejayDeviceID,
		Version:                deejayKeyVersion,
		AllCdn:                 true,
		ContentEabID:           eabID,
		Region:                 "US",
		XlinkSupport:           false,
		DeviceAdID:             randUUID(),
		LimitAdTracking:        false,
		IgnoreKidsBlock:        false,
		Language:               "en",
		GUID:                   c.huluGUID,
		Rv:                     int(frand.Uint64n(1e7)),
		Kv:                     sessionKey,
		Unencrypted:            true,
		IncludeT2RevenueBeacon: "1",
		CpSessionID:            randUUID(),
		NetworkMode:            "wifi",
		PlayIntent:             "resume",
		Playback: PlaylistRequestPlayback{
			Version: 2,
			Video:   PlaylistRequestVideo{Codecs: PlaylistRequestCodecs{Values: []PlaylistRequestValues{{Type: "H264", Profile: "HIGH", Level: "4.1", Framerate: 30}}, SelectionMode: "ONE"}},
			Audio:   PlaylistRequestAudio{Codecs: PlaylistRequestCodecs{Values: []PlaylistRequestValues{{Type: "AAC"}}, SelectionMode: "ONE"}},
			DRM:     PlaylistRequestDRM{Values: []PlaylistRequestValues{{Type: "WIDEVINE", Version: "MODULAR", SecurityLevel: "L3"}}, SelectionMode: "ONE"},
			Manifest: PlaylistRequestManifest{
				Type:              "DASH",
				HTTPS:             true,
				MultipleCdns:      true,
				PatchUpdates:      true,
				HuluTypes:         true,
				LiveDai:           true,
				MultiplePeriods:   true,
				Xlink:             false,
				SecondaryAudio:    true,
				LiveFragmentDelay: 3,
			},
			Segments: PlaylistRequestSegments{Values: []PlaylistRequestValues{{Type: "FMP4", Encryption: &PlaylistRequestEncryption{Mode: "CENC", Type: "CENC"}, HTTPS: true}}, SelectionMode: "ONE"},
		},
	}

	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(playlistRequest); err != nil {
		return
	}

	response, err := c.request(http.MethodPost, "https://play.hulu.com/v6/playlist", &buf, contentTypeJSON)
	if err != nil {
		return
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&p)
	return
}
