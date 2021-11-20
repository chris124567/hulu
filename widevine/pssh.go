package widevine

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io"
	"net/http"
)

// This function retrieves the PSSH/Init Data from a given MPD file reader.
// Example file: https://bitmovin-a.akamaihd.net/content/art-of-motion_drm/mpds/11331.mpd
func InitDataFromMPD(r io.Reader) ([]byte, error) {
	type mpd struct {
		XMLName                   xml.Name `xml:"MPD"`
		Text                      string   `xml:",chardata"`
		ID                        string   `xml:"id,attr"`
		Profiles                  string   `xml:"profiles,attr"`
		Type                      string   `xml:"type,attr"`
		AvailabilityStartTime     string   `xml:"availabilityStartTime,attr"`
		PublishTime               string   `xml:"publishTime,attr"`
		MediaPresentationDuration string   `xml:"mediaPresentationDuration,attr"`
		MinBufferTime             string   `xml:"minBufferTime,attr"`
		Version                   string   `xml:"version,attr"`
		Ns2                       string   `xml:"ns2,attr"`
		Xmlns                     string   `xml:"xmlns,attr"`
		Bitmovin                  string   `xml:"bitmovin,attr"`
		Period                    struct {
			Text          string `xml:",chardata"`
			AdaptationSet []struct {
				Text            string `xml:",chardata"`
				MimeType        string `xml:"mimeType,attr"`
				Codecs          string `xml:"codecs,attr"`
				Lang            string `xml:"lang,attr"`
				Label           string `xml:"label,attr"`
				SegmentTemplate struct {
					Text           string `xml:",chardata"`
					Media          string `xml:"media,attr"`
					Initialization string `xml:"initialization,attr"`
					Duration       string `xml:"duration,attr"`
					StartNumber    string `xml:"startNumber,attr"`
					Timescale      string `xml:"timescale,attr"`
				} `xml:"SegmentTemplate"`
				ContentProtection []struct {
					Text        string `xml:",chardata"`
					SchemeIdUri string `xml:"schemeIdUri,attr"`
					Value       string `xml:"value,attr"`
					DefaultKID  string `xml:"default_KID,attr"`
					Pssh        string `xml:"pssh"`
				} `xml:"ContentProtection"`
				Representation []struct {
					Text              string `xml:",chardata"`
					ID                string `xml:"id,attr"`
					Bandwidth         string `xml:"bandwidth,attr"`
					Width             string `xml:"width,attr"`
					Height            string `xml:"height,attr"`
					FrameRate         string `xml:"frameRate,attr"`
					AudioSamplingRate string `xml:"audioSamplingRate,attr"`
					ContentProtection []struct {
						Text        string `xml:",chardata"`
						SchemeIdUri string `xml:"schemeIdUri,attr"`
						Value       string `xml:"value,attr"`
						DefaultKID  string `xml:"default_KID,attr"`
						Cenc        string `xml:"cenc,attr"`
						Pssh        struct {
							Text string `xml:",chardata"`
							Cenc string `xml:"cenc,attr"`
						} `xml:"pssh"`
					} `xml:"ContentProtection"`
				} `xml:"Representation"`
				AudioChannelConfiguration struct {
					Text        string `xml:",chardata"`
					SchemeIdUri string `xml:"schemeIdUri,attr"`
					Value       string `xml:"value,attr"`
				} `xml:"AudioChannelConfiguration"`
			} `xml:"AdaptationSet"`
		} `xml:"Period"`
	}

	var mpdPlaylist mpd
	if err := xml.NewDecoder(r).Decode(&mpdPlaylist); err != nil {
		return nil, err
	}

	const widevineSchemeIdURI = "urn:uuid:edef8ba9-79d6-4ace-a3c8-27dcd51d21ed"
	for _, adaptionSet := range mpdPlaylist.Period.AdaptationSet {
		for _, protection := range adaptionSet.ContentProtection {
			if protection.SchemeIdUri == widevineSchemeIdURI && len(protection.Pssh) > 0 {
				return base64.StdEncoding.DecodeString(protection.Pssh)
			}
		}
	}
	for _, adaptionSet := range mpdPlaylist.Period.AdaptationSet {
		for _, representation := range adaptionSet.Representation {
			for _, protection := range representation.ContentProtection {
				if protection.SchemeIdUri == widevineSchemeIdURI && len(protection.Pssh.Text) > 0 {
					return base64.StdEncoding.DecodeString(protection.Pssh.Text)
				}
			}
		}
	}

	return nil, errors.New("no init data found")
}

// This function retrieves certificate data from a given license server.
func GetCertData(client *http.Client, licenseURL string) ([]byte, error) {
	response, err := client.Post(licenseURL, "application/x-www-form-urlencoded", bytes.NewReader([]byte{0x08, 0x04}))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return io.ReadAll(response.Body)
}
