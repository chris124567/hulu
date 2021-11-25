package client

import (
	"net/http"
	"time"
)

// These are constants taken from https://player.hulu.com/site/dash/308343-site-curiosity/js/app.js
var (
	deejayKey = []byte{110, 191, 200, 79, 60, 48, 66, 23, 178, 15, 217, 166, 108, 181, 149, 127}
)

const (
	deejayDeviceID   = 190
	deejayKeyVersion = 1
)

func StandardHeaders() http.Header {
	return http.Header{
		http.CanonicalHeaderKey("sec-ch-ua"):          []string{`" Not A;Brand";v="99"}, "Chromium";v="96"}, "Google Chrome";v="96"`},
		http.CanonicalHeaderKey("sec-ch-ua-mobile"):   []string{"?0"},
		http.CanonicalHeaderKey("User-Agent"):         []string{"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36"},
		http.CanonicalHeaderKey("sec-ch-ua-platform"): []string{`"Linux"`},
		http.CanonicalHeaderKey("Accept"):             []string{"*/*"},
		http.CanonicalHeaderKey("Origin"):             []string{"https://www.hulu.com"},
		http.CanonicalHeaderKey("Sec-Fetch-Site"):     []string{"same-site"},
		http.CanonicalHeaderKey("Sec-Fetch-Mode"):     []string{"cors"},
		http.CanonicalHeaderKey("Sec-Fetch-Dest"):     []string{"empty"},
		http.CanonicalHeaderKey("Referer"):            []string{"https://www.hulu.com/"},
		http.CanonicalHeaderKey("Accept-Language"):    []string{"en-US,en;q=0.9"},
	}
}

type SearchResults struct {
	Groups []struct {
		Category string `json:"category"`
		Results  []struct {
			Type        string `json:"_type"`
			MetricsInfo struct {
				TargetID            string `json:"target_id"`
				TargetType          string `json:"target_type"`
				TargetName          string `json:"target_name"`
				SelectionTrackingID string `json:"selection_tracking_id"`
			} `json:"metrics_info"`
			Personalization struct {
				BowieContext string `json:"bowie_context"`
				Eab          string `json:"eab"`
			} `json:"personalization"`
			DeviceContextFailure bool   `json:"device_context_failure"`
			ViewTemplate         string `json:"view_template"`
			Visuals              struct {
				Artwork struct {
					Type       string `json:"_type"`
					Horizontal struct {
						Type        string `json:"_type"`
						ArtworkType string `json:"artwork_type"`
						Image       struct {
							Path   string `json:"path"`
							Accent struct {
								Hue            int    `json:"hue"`
								Classification string `json:"classification"`
							} `json:"accent"`
							ImageID string `json:"image_id"`
						} `json:"image"`
						Text string `json:"text"`
					} `json:"horizontal"`
				} `json:"artwork"`
				Headline struct {
					Text  string  `json:"text"`
					Index [][]int `json:"index"`
				} `json:"headline"`
				Body struct {
					Text  string  `json:"text"`
					Index [][]int `json:"index"`
				} `json:"body"`
				ActionText      string `json:"action_text"`
				PrimaryBranding struct {
					ID      string `json:"id"`
					Name    string `json:"name"`
					Artwork struct {
						BrandWatermarkBottomRight struct {
							Path   string `json:"path"`
							Accent struct {
								Hue            int    `json:"hue"`
								Classification string `json:"classification"`
							} `json:"accent"`
							ImageType string `json:"image_type"`
							ImageID   string `json:"image_id"`
						} `json:"brand.watermark.bottom.right"`
						BrandLogoBottomRight struct {
							Path   string `json:"path"`
							Accent struct {
								Hue            int    `json:"hue"`
								Classification string `json:"classification"`
							} `json:"accent"`
							ImageType string `json:"image_type"`
							ImageID   string `json:"image_id"`
						} `json:"brand.logo.bottom.right"`
					} `json:"artwork"`
				} `json:"primary_branding"`
				ShortSubtitle struct {
					Text  string        `json:"text"`
					Index []interface{} `json:"index"`
				} `json:"short_subtitle"`
			} `json:"visuals"`
			Actions struct {
				Browse struct {
					TargetType  string `json:"target_type"`
					TargetID    string `json:"target_id"`
					TargetName  string `json:"target_name"`
					TargetTheme string `json:"target_theme"`
					Params      struct {
					} `json:"params"`
					Href        string `json:"href"`
					BrowseTheme string `json:"browse_theme"`
					MetricsInfo struct {
						ActionType        string `json:"action_type"`
						TargetID          string `json:"target_id"`
						TargetType        string `json:"target_type"`
						TargetDisplayName string `json:"target_display_name"`
					} `json:"metrics_info"`
					Type string `json:"type"`
				} `json:"browse"`
				ContextMenu struct {
					Actions []struct {
						ActionType  string `json:"action_type"`
						EntityName  string `json:"entity_name"`
						EntityType  string `json:"entity_type"`
						MetricsInfo struct {
							TargetID          string `json:"target_id"`
							TargetType        string `json:"target_type"`
							TargetDisplayName string `json:"target_display_name"`
							Eab               string `json:"eab"`
							Type              string `json:"_type"`
						} `json:"metrics_info"`
						Eab string `json:"eab"`
					} `json:"actions"`
					Header struct {
						Title   string `json:"title"`
						Artwork struct {
							Type       string `json:"_type"`
							Horizontal struct {
								Type        string `json:"_type"`
								ArtworkType string `json:"artwork_type"`
								Image       struct {
									Path   string `json:"path"`
									Accent struct {
										Hue            int    `json:"hue"`
										Classification string `json:"classification"`
									} `json:"accent"`
									ImageID string `json:"image_id"`
								} `json:"image"`
								Text string `json:"text"`
							} `json:"horizontal"`
							Vertical struct {
								Type        string `json:"_type"`
								ArtworkType string `json:"artwork_type"`
								Text        string `json:"text"`
							} `json:"vertical"`
						} `json:"artwork"`
						PrimaryBranding struct {
							ID      string `json:"id"`
							Name    string `json:"name"`
							Artwork struct {
								BrandWatermarkBottomRight struct {
									Path   string `json:"path"`
									Accent struct {
										Hue            int    `json:"hue"`
										Classification string `json:"classification"`
									} `json:"accent"`
									ImageType string `json:"image_type"`
									ImageID   string `json:"image_id"`
								} `json:"brand.watermark.bottom.right"`
								BrandLogoBottomRight struct {
									Path   string `json:"path"`
									Accent struct {
										Hue            int    `json:"hue"`
										Classification string `json:"classification"`
									} `json:"accent"`
									ImageType string `json:"image_type"`
									ImageID   string `json:"image_id"`
								} `json:"brand.logo.bottom.right"`
							} `json:"artwork"`
						} `json:"primary_branding"`
						Action struct {
							ActionType  string `json:"action_type"`
							EntityName  string `json:"entity_name"`
							EntityType  string `json:"entity_type"`
							MetricsInfo struct {
								TargetID          string `json:"target_id"`
								TargetType        string `json:"target_type"`
								TargetDisplayName string `json:"target_display_name"`
								Eab               string `json:"eab"`
								Type              string `json:"_type"`
							} `json:"metrics_info"`
							Browse struct {
								TargetType  string `json:"target_type"`
								TargetID    string `json:"target_id"`
								TargetTheme string `json:"target_theme"`
								Params      struct {
								} `json:"params"`
								Type string `json:"type"`
							} `json:"browse"`
							TargetName string `json:"target_name"`
							Href       string `json:"href"`
						} `json:"action"`
					} `json:"header"`
				} `json:"context_menu"`
			} `json:"actions"`
			EntityMetadata struct {
				GenreNames   []string  `json:"genre_names"`
				PremiereDate time.Time `json:"premiere_date"`
				Rating       struct {
					Code string `json:"code"`
				} `json:"rating"`
				TargetName string `json:"target_name"`
				IsWarm     bool   `json:"is_warm"`
			} `json:"entity_metadata"`
		} `json:"results"`
	} `json:"groups"`
	Metadata struct {
		SearchResultType    string `json:"search_result_type"`
		Explanation         string `json:"explanation"`
		SelectionTrackingID string `json:"selection_tracking_id"`
	} `json:"metadata"`
	DeviceContextFailure bool `json:"device_context_failure"`
}

type Season struct {
	Type     string `json:"_type"`
	ID       string `json:"id"`
	Href     string `json:"href"`
	P13NHref string `json:"p13n_href"`
	Name     string `json:"name"`
	Theme    string `json:"theme"`
	Artwork  struct {
	} `json:"artwork"`
	DeviceContextFailure bool `json:"device_context_failure"`
	Items                []struct {
		Type        string `json:"_type"`
		ID          string `json:"id"`
		Href        string `json:"href"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Artwork     struct {
			VideoHorizontalHero struct {
				Path   string `json:"path"`
				Accent struct {
					Hue            int    `json:"hue"`
					Classification string `json:"classification"`
				} `json:"accent"`
				ImageType string `json:"image_type"`
				ImageID   string `json:"image_id"`
			} `json:"video.horizontal.hero"`
		} `json:"artwork"`
		MetricsInfo struct {
			Type                string `json:"_type"`
			MetricsAssetName    string `json:"metrics_asset_name"`
			AiringType          string `json:"airing_type"`
			ExternalIdentifiers []struct {
				Namespace string `json:"namespace"`
				ID        string `json:"id"`
			} `json:"external_identifiers"`
		} `json:"metrics_info"`
		Personalization struct {
			Eab string `json:"eab"`
		} `json:"personalization"`
		DeviceContextFailure bool `json:"device_context_failure"`
		Browse               struct {
			TargetType  string `json:"target_type"`
			TargetID    string `json:"target_id"`
			TargetTheme string `json:"target_theme"`
			Params      struct {
			} `json:"params"`
			Href        string `json:"href"`
			BrowseTheme string `json:"browse_theme"`
			Type        string `json:"type"`
		} `json:"browse"`
		SeriesID               string `json:"series_id"`
		SeriesName             string `json:"series_name"`
		Season                 string `json:"season"`
		SeasonShortDisplayName string `json:"season_short_display_name"`
		Bundle                 struct {
			Type         string `json:"_type"`
			ID           int    `json:"id"`
			EabID        string `json:"eab_id"`
			NetworkID    string `json:"network_id"`
			NetworkName  string `json:"network_name"`
			Duration     int    `json:"duration"`
			Availability struct {
				Type                string    `json:"_type"`
				StartDate           time.Time `json:"start_date"`
				EndDate             time.Time `json:"end_date"`
				LocationRequirement string    `json:"location_requirement"`
				IsAvailable         bool      `json:"is_available"`
			} `json:"availability"`
			BundleType          string `json:"bundle_type"`
			Rating              string `json:"rating"`
			OpenCreditEndPos    int    `json:"open_credit_end_pos"`
			CloseCreditStartPos int    `json:"close_credit_start_pos"`
			Rights              struct {
				Startover      bool `json:"startover"`
				Recordable     bool `json:"recordable"`
				Offline        bool `json:"offline"`
				ClientOverride bool `json:"client_override"`
			} `json:"rights"`
			CpID        int      `json:"cp_id"`
			AllEtag     string   `json:"all_etag"`
			RightsEtag  string   `json:"rights_etag"`
			AiringsEtag string   `json:"airings_etag"`
			StreamEtag  string   `json:"stream_etag"`
			RightsTTL   int      `json:"rights_ttl"`
			AiringsTTL  int      `json:"airings_ttl"`
			StreamTTL   int      `json:"stream_ttl"`
			PackageID   int      `json:"package_id"`
			AvFeatures  []string `json:"av_features"`
		} `json:"bundle"`
		Number          string `json:"number"`
		PrimaryBranding struct {
			ID      string `json:"id"`
			Name    string `json:"name"`
			Artwork struct {
				BrandWatermark struct {
					Path   string `json:"path"`
					Accent struct {
						Hue            int    `json:"hue"`
						Classification string `json:"classification"`
					} `json:"accent"`
					ImageType string `json:"image_type"`
					ImageID   string `json:"image_id"`
				} `json:"brand.watermark"`
				BrandWatermarkDark struct {
					Path   string `json:"path"`
					Accent struct {
						Hue            int    `json:"hue"`
						Classification string `json:"classification"`
					} `json:"accent"`
					ImageType string `json:"image_type"`
					ImageID   string `json:"image_id"`
				} `json:"brand.watermark.dark"`
				BrandWatermarkTopRight struct {
					Path   string `json:"path"`
					Accent struct {
						Hue            int    `json:"hue"`
						Classification string `json:"classification"`
					} `json:"accent"`
					ImageType string `json:"image_type"`
					ImageID   string `json:"image_id"`
				} `json:"brand.watermark.top.right"`
				BrandLogo struct {
					Path   string `json:"path"`
					Accent struct {
						Hue            int    `json:"hue"`
						Classification string `json:"classification"`
					} `json:"accent"`
					ImageType string `json:"image_type"`
					ImageID   string `json:"image_id"`
				} `json:"brand.logo"`
				NetworkTile struct {
					Path   string `json:"path"`
					Accent struct {
						Hue            int    `json:"hue"`
						Classification string `json:"classification"`
					} `json:"accent"`
					ImageType string `json:"image_type"`
					ImageID   string `json:"image_id"`
				} `json:"network.tile"`
				BrandWatermarkBottomRight struct {
					Path   string `json:"path"`
					Accent struct {
						Hue            int    `json:"hue"`
						Classification string `json:"classification"`
					} `json:"accent"`
					ImageType string `json:"image_type"`
					ImageID   string `json:"image_id"`
				} `json:"brand.watermark.bottom.right"`
				BrandLogoTopRight struct {
					Path   string `json:"path"`
					Accent struct {
						Hue            int    `json:"hue"`
						Classification string `json:"classification"`
					} `json:"accent"`
					ImageType string `json:"image_type"`
					ImageID   string `json:"image_id"`
				} `json:"brand.logo.top.right"`
				BrandLogoBottomRight struct {
					Path   string `json:"path"`
					Accent struct {
						Hue            int    `json:"hue"`
						Classification string `json:"classification"`
					} `json:"accent"`
					ImageType string `json:"image_type"`
					ImageID   string `json:"image_id"`
				} `json:"brand.logo.bottom.right"`
			} `json:"artwork"`
		} `json:"primary_branding"`
		Rating struct {
			Code string `json:"code"`
		} `json:"rating"`
		GenreNames    []string  `json:"genre_names"`
		PremiereDate  time.Time `json:"premiere_date"`
		Duration      int       `json:"duration"`
		IsFirstRun    bool      `json:"is_first_run"`
		SeriesArtwork struct {
			DetailVerticalHero struct {
				Path   string `json:"path"`
				Accent struct {
					Hue            int    `json:"hue"`
					Classification string `json:"classification"`
				} `json:"accent"`
				ImageType string `json:"image_type"`
				ImageID   string `json:"image_id"`
			} `json:"detail.vertical.hero"`
			TitleTreatmentHorizontal struct {
				Path   string `json:"path"`
				Accent struct {
					Hue            int    `json:"hue"`
					Classification string `json:"classification"`
				} `json:"accent"`
				ImageType string `json:"image_type"`
				ImageID   string `json:"image_id"`
			} `json:"title.treatment.horizontal"`
			ProgramTile struct {
				Path   string `json:"path"`
				Accent struct {
					Hue            int    `json:"hue"`
					Classification string `json:"classification"`
				} `json:"accent"`
				ImageType string `json:"image_type"`
				ImageID   string `json:"image_id"`
			} `json:"program.tile"`
			ProgramVerticalTile struct {
				Path   string `json:"path"`
				Accent struct {
					Hue            int    `json:"hue"`
					Classification string `json:"classification"`
				} `json:"accent"`
				ImageType string `json:"image_type"`
				ImageID   string `json:"image_id"`
			} `json:"program.vertical.tile"`
			TitleTreatmentStacked struct {
				Path   string `json:"path"`
				Accent struct {
					Hue            int    `json:"hue"`
					Classification string `json:"classification"`
				} `json:"accent"`
				ImageType string `json:"image_type"`
				ImageID   string `json:"image_id"`
			} `json:"title.treatment.stacked"`
			DetailHorizontalHero struct {
				Path   string `json:"path"`
				Accent struct {
					Hue            int    `json:"hue"`
					Classification string `json:"classification"`
				} `json:"accent"`
				ImageType string `json:"image_type"`
				ImageID   string `json:"image_id"`
			} `json:"detail.horizontal.hero"`
		} `json:"series_artwork"`
		RestrictionLevel string        `json:"restriction_level"`
		Exclusivity      string        `json:"exclusivity"`
		Actions          []interface{} `json:"actions"`
	} `json:"items"`
	Pagination struct {
		CurrentOffset int `json:"current_offset"`
		TotalCount    int `json:"total_count"`
	} `json:"pagination"`
	SeriesGroupingMetadata struct {
		SeriesGroupingType string `json:"series_grouping_type"`
		SeasonNumber       int    `json:"season_number"`
		GroupingName       string `json:"groupingName"`
		Unknown            bool   `json:"unknown"`
	} `json:"series_grouping_metadata"`
}

type PlaybackInformation struct {
	Type   string `json:"_type"`
	Browse struct {
		TargetType  string `json:"target_type"`
		TargetID    string `json:"target_id"`
		TargetTheme string `json:"target_theme"`
		Params      struct {
		} `json:"params"`
		Type string `json:"type"`
	} `json:"browse"`
	EabID            string `json:"eab_id"`
	Href             string `json:"href"`
	ID               string `json:"id"`
	HrefType         string `json:"href_type"`
	RestrictionLevel string `json:"restriction_level"`
}

type PlaylistRequest struct {
	DeviceIdentifier       string                  `json:"device_identifier"`
	DeejayDeviceID         int                     `json:"deejay_device_id"`
	Version                int                     `json:"version"`
	AllCdn                 bool                    `json:"all_cdn"`
	ContentEabID           string                  `json:"content_eab_id"`
	Region                 string                  `json:"region"`
	XlinkSupport           bool                    `json:"xlink_support"`
	DeviceAdID             string                  `json:"device_ad_id"`
	LimitAdTracking        bool                    `json:"limit_ad_tracking"`
	IgnoreKidsBlock        bool                    `json:"ignore_kids_block"`
	Language               string                  `json:"language"`
	GUID                   string                  `json:"guid"`
	Rv                     int                     `json:"rv"`
	Kv                     int                     `json:"kv"`
	Unencrypted            bool                    `json:"unencrypted"`
	IncludeT2RevenueBeacon string                  `json:"include_t2_revenue_beacon"`
	CpSessionID            string                  `json:"cp_session_id"`
	InterfaceVersion       string                  `json:"interface_version"`
	NetworkMode            string                  `json:"network_mode"`
	PlayIntent             string                  `json:"play_intent"`
	Playback               PlaylistRequestPlayback `json:"playback"`
}

type PlaylistRequestValues struct {
	Type          string                     `json:"type,omitempty"`
	Profile       string                     `json:"profile,omitempty"`
	Level         string                     `json:"level,omitempty"`
	Framerate     int                        `json:"framerate,omitempty"`
	Version       string                     `json:"version,omitempty"`
	SecurityLevel string                     `json:"security_level,omitempty"`
	Encryption    *PlaylistRequestEncryption `json:"encryption,omitempty"`
	HTTPS         bool                       `json:"https,omitempty"`
}

type PlaylistRequestCodecs struct {
	Values        []PlaylistRequestValues `json:"values"`
	SelectionMode string                  `json:"selection_mode"`
}

type PlaylistRequestVideo struct {
	Codecs PlaylistRequestCodecs `json:"codecs"`
}

type PlaylistRequestAudio struct {
	Codecs PlaylistRequestCodecs `json:"codecs"`
}

type PlaylistRequestDRM struct {
	Values        []PlaylistRequestValues `json:"values"`
	SelectionMode string                  `json:"selection_mode"`
}

type PlaylistRequestManifest struct {
	Type              string `json:"type"`
	HTTPS             bool   `json:"https"`
	MultipleCdns      bool   `json:"multiple_cdns"`
	PatchUpdates      bool   `json:"patch_updates"`
	HuluTypes         bool   `json:"hulu_types"`
	LiveDai           bool   `json:"live_dai"`
	MultiplePeriods   bool   `json:"multiple_periods"`
	Xlink             bool   `json:"xlink"`
	SecondaryAudio    bool   `json:"secondary_audio"`
	LiveFragmentDelay int    `json:"live_fragment_delay"`
}

type PlaylistRequestEncryption struct {
	Mode string `json:"mode"`
	Type string `json:"type"`
}

type PlaylistRequestSegments struct {
	Values        []PlaylistRequestValues `json:"values"`
	SelectionMode string                  `json:"selection_mode"`
}

type PlaylistRequestPlayback struct {
	Version  int                     `json:"version"`
	Video    PlaylistRequestVideo    `json:"video"`
	Audio    PlaylistRequestAudio    `json:"audio"`
	DRM      PlaylistRequestDRM      `json:"drm"`
	Manifest PlaylistRequestManifest `json:"manifest"`
	Segments PlaylistRequestSegments `json:"segments"`
}
type Playlist struct {
	UseManifestBreaks bool          `json:"use_manifest_breaks"`
	Adstate           string        `json:"adstate"`
	Breaks            []interface{} `json:"breaks"`
	ContentEabID      string        `json:"content_eab_id"`
	TranscriptsUrls   struct {
		Smi struct {
			En string `json:"en"`
		} `json:"smi"`
		Webvtt struct {
			En string `json:"en"`
		} `json:"webvtt"`
		Ttml struct {
			En string `json:"en"`
		} `json:"ttml"`
	} `json:"transcripts_urls"`
	TranscriptsEncryptionKey string `json:"transcripts_encryption_key"`
	VideoMetadata            struct {
		AspectRatio          string      `json:"aspect_ratio"`
		EndCreditsTime       string      `json:"end_credits_time"`
		FrameRate            int         `json:"frame_rate"`
		HasBug               string      `json:"has_bug"`
		HasCaptions          bool        `json:"has_captions"`
		HasNetworkPreRoll    bool        `json:"has_network_pre_roll"`
		Interstitials        string      `json:"interstitials"`
		Language             string      `json:"language"`
		Length               int         `json:"length"`
		Segments             string      `json:"segments"`
		ID                   int         `json:"id"`
		AssetID              int         `json:"asset_id"`
		Markers              interface{} `json:"markers"`
		TranscriptsDefaultOn bool        `json:"transcripts_default_on"`
		RatingBugBig         string      `json:"rating_bug_big"`
		RatingBugSmall       string      `json:"rating_bug_small"`
	} `json:"video_metadata"`
	TranscriptsEncryptionIv string `json:"transcripts_encryption_iv"`
	Breakhash               string `json:"breakhash"`
	AdBreakTimes            []int  `json:"ad_break_times"`
	TranscriptsDefaultOn    bool   `json:"transcripts_default_on"`
	ResumePosition          int    `json:"resume_position"`
	RecordingOffset         int    `json:"recording_offset"`
	InitialPosition         int    `json:"initial_position"`
	DashPrServer            string `json:"dash_pr_server"`
	WvServer                string `json:"wv_server"`
	AudioTracks             []struct {
		Language     string `json:"language"`
		Role         string `json:"role"`
		CodecsString string `json:"codecs_string"`
		Channels     int    `json:"channels"`
	} `json:"audio_tracks"`
	MbrManifest       string `json:"mbr_manifest"`
	StreamURL         string `json:"stream_url"`
	ThumbnailEndpoint string `json:"thumbnail_endpoint"`
	AssetPlaybackType string `json:"asset_playback_type"`
	SauronID          string `json:"sauron_id"`
	ViewTTLMillis     int    `json:"view_ttl_millis"`
	SauronToken       string `json:"sauron_token"`
	SauronTokenTTL    int    `json:"sauron_token_ttl"`
	SauronTokenTTLMs  int    `json:"sauron_token_ttl_ms"`
}

type Config struct {
	PassThroughQos             string   `json:"pass_through_qos"`
	Kinko                      string   `json:"kinko"`
	PackageID                  int      `json:"package_id"`
	API                        string   `json:"api"`
	QosBeacon                  string   `json:"qos_beacon"`
	NielsenAppName             string   `json:"nielsen_app_name"`
	FeedbackCategory           int      `json:"feedbackCategory"`
	PlusPlanID                 int      `json:"plus_plan_id"`
	FirehoseEndpoint           string   `json:"firehose_endpoint"`
	PbAutoresumeTimeout        int      `json:"pb_autoresume_timeout"`
	SashProductDescription     string   `json:"sash_product_description"`
	PlaylistEndpoint           string   `json:"playlist_endpoint"`
	NielsenAppID               string   `json:"nielsen_app_id"`
	PackageGroupID             int      `json:"package_group_id"`
	FlexActionEndpoint         string   `json:"flex_action_endpoint"`
	PlaybackRequestTimeout     int      `json:"playback_request_timeout"`
	Asset                      string   `json:"asset"`
	NoahSignupExceptionMessage []string `json:"noah_signup_exception_message"`
	PbInterval                 int      `json:"pb_interval"`
	PackageGroupIDFrontPorch   int      `json:"package_group_id_front_porch"`
	FlagsContext               struct {
		FlagStateValid bool   `json:"flag_state_valid"`
		UILink         string `json:"ui_link"`
		Key            string `json:"key"`
		SealTokenState string `json:"seal_token_state"`
	} `json:"flags_context"`
	SashProductTitle          string `json:"sash_product_title"`
	TrackTiersDeepPlayerState int    `json:"track_tiers_deep_player_state"`
	UserAccountURL            string `json:"userAccountURL"`
	Pgid                      int    `json:"pgid"`
	GeokResponse              string `json:"geok_response"`
	UserInfoURL               string `json:"user_info_url"`
	ChangePlanURL             string `json:"changePlanURL"`
	Profiles                  struct {
		PromptAfterIdleMs int `json:"prompt_after_idle_ms"`
	} `json:"profiles"`
	AutoplayIdleTimeout          int    `json:"autoplay_idle_timeout"`
	NielsenSfCode                string `json:"nielsen_sf_code"`
	PlusLandingURL               string `json:"plusLandingURL"`
	NielsenAppVersion            string `json:"nielsen_app_version"`
	PlayerProgressReportInterval int    `json:"player_progress_report_interval"`
	ProductInstrumentationV2     struct {
		MetricsAgent struct {
			Endpoint       string `json:"endpoint"`
			MaxHitRetries  int    `json:"max_hit_retries"`
			SamplingRatios struct {
				ServiceCall float64 `json:"service_call"`
			} `json:"sampling_ratios"`
			MsPerEvent        int  `json:"ms_per_event"`
			Enabled           bool `json:"enabled"`
			MaxHitQueueMs     int  `json:"max_hit_queue_ms"`
			EventFilterConfig struct {
			} `json:"event_filter_config"`
			OnlineAssetMaxBeaconQueueMs     int      `json:"online_asset_max_beacon_queue_ms"`
			DownloadedAssetMaxBeaconQueueMs int      `json:"downloaded_asset_max_beacon_queue_ms"`
			EventWhitelist                  []string `json:"event_whitelist"`
			NonInteractiveEvents            []string `json:"non_interactive_events"`
			BucketSize                      int      `json:"bucket_size"`
		} `json:"metrics_agent"`
		ConvivaAgent struct {
			FatalErrors []string `json:"fatal_errors"`
			Staging     bool     `json:"staging"`
			Token       string   `json:"token"`
			Enabled     bool     `json:"enabled"`
			GatewayURL  string   `json:"gateway_url"`
		} `json:"conviva_agent"`
		MetricsTracker struct {
		} `json:"metrics_tracker"`
		RateLimiting struct {
			SegmentDownloadHit int `json:"segment_download_hit"`
		} `json:"rate_limiting"`
		AdobeAgent struct {
			AppMeasurementTrackingServer string `json:"app_measurement_tracking_server"`
			VisitorMcid                  string `json:"visitor_mcid"`
			Enabled                      bool   `json:"enabled"`
			AppMeasurementRsid           string `json:"app_measurement_rsid"`
			VisitorTrackingServer        string `json:"visitor_tracking_server"`
			HeartbeatTrackingServer      string `json:"heartbeat_tracking_server"`
		} `json:"adobe_agent"`
		MoatAgent struct {
			Enabled bool `json:"enabled"`
		} `json:"moat_agent"`
		AdobeAgentV2 struct {
			AppMeasurementTrackingServer string `json:"app_measurement_tracking_server"`
			VisitorMcid                  string `json:"visitor_mcid"`
			Enabled                      bool   `json:"enabled"`
			AppMeasurementRsid           string `json:"app_measurement_rsid"`
			VisitorTrackingServer        string `json:"visitor_tracking_server"`
			HeartbeatTrackingServer      string `json:"heartbeat_tracking_server"`
		} `json:"adobe_agent_v2"`
	} `json:"product_instrumentation_v2"`
	FeedbackURL       string `json:"feedbackURL"`
	IsAnonProxy       bool   `json:"is_anon_proxy"`
	NielsenEnabled    string `json:"nielsen_enabled"`
	EurekaNamespace   string `json:"eureka_namespace"`
	ReportGeocheckURL string `json:"reportGeocheckURL"`
	PbTracker         string `json:"pb_tracker"`
	MetricsAgent      struct {
		MaxBatchesBuffered int  `json:"max_batches_buffered"`
		MaxBatchSize       int  `json:"max_batch_size"`
		Enabled            bool `json:"enabled"`
		EventFilterConfig  struct {
			ServiceCall struct {
				EventRules []struct {
					RuleType   string `json:"rule_type"`
					RuleFilter struct {
						Type      string `json:"type,omitempty"`
						Dimension string `json:"dimension,omitempty"`
						Value     string `json:"value,omitempty"`
						Filters   []struct {
							Type      string `json:"type,omitempty"`
							Dimension string `json:"dimension,omitempty"`
							Value     string `json:"value,omitempty"`
						} `json:"filters,omitempty"`
					} `json:"rule_filter,omitempty"`
				} `json:"event_rules"`
			} `json:"service_call"`
			Log struct {
				EventRules []struct {
					RuleType   string `json:"rule_type"`
					RuleFilter struct {
						Filter struct {
							Type    string `json:"type"`
							Filters []struct {
								Type      string `json:"type"`
								Dimension string `json:"dimension"`
								Value     string `json:"value"`
							} `json:"filters"`
						} `json:"filter"`
						Type string `json:"type"`
					} `json:"rule_filter"`
				} `json:"event_rules"`
			} `json:"log"`
		} `json:"event_filter_config"`
		Endpoint      string `json:"endpoint"`
		FlushInterval int    `json:"flush_interval"`
	} `json:"metrics_agent"`
	Iball                    string `json:"iball"`
	HuluMbr                  string `json:"hulu_mbr"`
	NoahSignupExceptionShows []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"noah_signup_exception_shows"`
	BeaconConfig           string `json:"beacon_config"`
	Sapi                   string `json:"sapi"`
	Csel                   string `json:"csel"`
	ProfileBitrates        []int  `json:"profile_bitrates"`
	PlusLearnMoreURL       string `json:"plusLearnMoreURL"`
	NoahProductDescription string `json:"noah_product_description"`
	SauronAccessToken      string `json:"sauron_access_token"`
	EurekaApplicationID    string `json:"eureka_application_id"`
	GeokLocation           string `json:"geok_location"`
	PlusInviteURL          string `json:"plusInviteURL"`
	HothHost               string `json:"hoth_host"`
	PlaybackRequestRetries int    `json:"playback_request_retries"`
	BadgingConfig          []struct {
		Text  string `json:"text"`
		State string `json:"state"`
		Style string `json:"style"`
	} `json:"badging_config"`
	AdServer     string `json:"ad_server"`
	RtBeacon     string `json:"rt_beacon"`
	EndpointUrls struct {
		PlaylistV4              string `json:"playlist_v4"`
		PlaylistV5              string `json:"playlist_v5"`
		PlaylistV6              string `json:"playlist_v6"`
		UserStateV5             string `json:"user_state_v5"`
		WatchDownloadV1         string `json:"watch_download_v1"`
		DvrRecordingsV1         string `json:"dvr_recordings_v1"`
		UserV1                  string `json:"user_v1"`
		FlexActionV1            string `json:"flex_action_v1"`
		BrowseV5                string `json:"browse_v5"`
		UserBookmarksV1         string `json:"user_bookmarks_v1"`
		GuideV0                 string `json:"guide_v0"`
		UserTastesV5            string `json:"user_tastes_v5"`
		VortexV0                string `json:"vortex_v0"`
		ConvivaV0               string `json:"conviva_v0"`
		DvrRecordingSettingsV1  string `json:"dvr_recording_settings_v1"`
		DvrV1                   string `json:"dvr_v1"`
		ConfigV0                string `json:"config_v0"`
		EmuV0                   string `json:"emu_v0"`
		SauronV1                string `json:"sauron_v1"`
		PlaybackFeaturesV0      string `json:"playback_features_v0"`
		OfflinePlaylistV1       string `json:"offline_playlist_v1"`
		AuthV1                  string `json:"auth_v1"`
		AuthV2                  string `json:"auth_v2"`
		OnboardingV5            string `json:"onboarding_v5"`
		AuthAppleAuthnRequestV0 string `json:"auth_apple_authn_request_v0"`
		GlobalNavV1             string `json:"global_nav_v1"`
	} `json:"endpoint_urls"`
	IapGracefulDegradationEnabled bool   `json:"iap_graceful_degradation_enabled"`
	KeyExpiration                 int    `json:"key_expiration"`
	Beacon                        string `json:"beacon"`
	Key                           string `json:"key"`
	EurekaApplicationName         string `json:"eureka_application_name"`
	DeviceID                      int    `json:"device_id"`
	PackageIDFrontPorch           int    `json:"package_id_front_porch"`
	PlaybackFeaturesEndpoint      string `json:"playback_features_endpoint"`
	NoahSignupExceptionURL        string `json:"noah_signup_exception_url"`
	ExpirationNoticeHours         int    `json:"expiration_notice_hours"`
	ForgotPasswordURL             string `json:"forgotPasswordURL"`
	SauronEndpoint                string `json:"sauron_endpoint"`
	GlobalNavEndpoint             string `json:"global_nav_endpoint"`
	PassThroughMetric             string `json:"pass_through_metric"`
	BanyaSec                      string `json:"banya_sec"`
	Nydus                         string `json:"nydus"`
	Flags                         struct {
		HuluClientStandardPromptTheme  bool `json:"hulu-client-standard-prompt-theme"`
		HuluClientTwoFactorVerify      bool `json:"hulu-client-two-factor-verify"`
		HuluClientGatewayDeviceRanking bool `json:"hulu-client-gateway-device-ranking"`
		HuluWebDemoPlayerVersion       struct {
			HitchMobilePlaybackProdHuluCom string `json:"hitch-mobile-playback.prod.hulu.com"`
			CoviewingProdHuluCom           string `json:"coviewing.prod.hulu.com"`
			EndcardProdHuluCom             string `json:"endcard.prod.hulu.com"`
			LocalhostProdHuluCom           string `json:"localhost.prod.hulu.com"`
			OneplayerProdHuluCom           string `json:"oneplayer.prod.hulu.com"`
			DevelopProdHuluCom             string `json:"develop.prod.hulu.com"`
		} `json:"hulu-web-demo-player-version"`
		HuluWebChromecastSdkPlayerVersion struct {
			Player  string `json:"player"`
			Options struct {
				MultiKey   bool `json:"multi-key"`
				Hdr        bool `json:"hdr"`
				Touchstone bool `json:"touchstone"`
			} `json:"options"`
			Sdk string `json:"sdk"`
		} `json:"hulu-web-chromecast-sdk-player-version"`
		HuluClientRokuInstantSignupEnabled bool `json:"hulu-client-roku-instant-signup-enabled"`
		HuluClientEndCardFg1               bool `json:"hulu-client-end-card-fg1"`
		HuluClientPinProtectionEnabled     bool `json:"hulu-client-pin-protection-enabled"`
		HuluClientPerformanceTracking      bool `json:"hulu-client-performance-tracking"`
		HuluWebSmokeSitePlayerVersion      struct {
			Nonsub string `json:"nonsub"`
			Sub    string `json:"sub"`
		} `json:"hulu-web-smoke-site-player-version"`
		HuluWebSmokeChromecastSdkPlayerVersion struct {
			Player  string `json:"player"`
			Options struct {
				MultipleKey bool `json:"multiple-key"`
				MultiKey    bool `json:"multi-key"`
				Touchstone  bool `json:"touchstone"`
			} `json:"options"`
			Sdk string `json:"sdk"`
		} `json:"hulu-web-smoke-chromecast-sdk-player-version"`
		HuluClientNeverBlockSvodEnabled     bool `json:"hulu-client-never-block-svod-enabled"`
		HuluClientUpdatedLocationPrompt     bool `json:"hulu-client-updated-location-prompt"`
		HuluClientFlexWelcomeEnabled        bool `json:"hulu-client-flex-welcome-enabled"`
		HuluClientEventPurchaseEnabled      bool `json:"hulu-client-event-purchase-enabled"`
		HuluWebDevelopProdSitePlayerOptions struct {
			CreditEndCardDuration string `json:"credit_end_card_duration"`
			SkipButtonDuration    string `json:"skip_button_duration"`
			EnablePinchZoom       bool   `json:"enable_pinch_zoom"`
			EnabledAdobeAgent     bool   `json:"enabled_adobe_agent"`
			EnabledQueuedSeek     bool   `json:"enabled_queued_seek"`
		} `json:"hulu-web-develop-prod-site-player-options"`
		HuluClientExperienceBrandedPageThemeSupport string `json:"hulu-client-experience-branded-page-theme-support"`
		HuluClientSignupOnDeviceEnabled             bool   `json:"hulu-client-signup-on-device-enabled"`
		HuluClientNonNumericSeasons                 bool   `json:"hulu-client-non-numeric-seasons"`
		HuluClientFlexTimeoutsMs                    int    `json:"hulu-client-flex-timeouts-ms"`
		HuluClientDvrRecordingsGroups               bool   `json:"hulu-client-dvr-recordings-groups"`
		HuluClientPlayerBasicsFg1                   bool   `json:"hulu-client-player-basics-fg-1"`
		HuluClientEventPurchasePollingTimeout       int    `json:"hulu-client-event-purchase-polling-timeout"`
		HuluClientNewDvrFeatures                    bool   `json:"hulu-client-new-dvr-features"`
		HuluClientFeatureMultikey                   bool   `json:"hulu-client-feature-multikey"`
		HuluClientPostPurchaseCollectionID          int    `json:"hulu-client-post-purchase-collection-id"`
		HuluClientPlanSelectExtraCopy               struct {
			ShowExtraCopy bool `json:"showExtraCopy"`
		} `json:"hulu-client-plan-select-extra-copy"`
		HuluWebSmokeChromecastPlayerOptions struct {
			Touchstone bool `json:"touchstone"`
		} `json:"hulu-web-smoke-chromecast-player-options"`
		HuluClientCompassViewAllEnabled bool     `json:"hulu-client-compass-view-all-enabled"`
		HuluClientForcedDcsCapabilities []string `json:"hulu-client-forced-dcs-capabilities"`
		HuluWebSitePlayerOptions        struct {
			CreditEndCardDuration string `json:"credit_end_card_duration"`
			EnabledBrightline     bool   `json:"enabled_brightline"`
			EnabledAdobeAgent     bool   `json:"enabled_adobe_agent"`
			EnabledQueuedSeek     bool   `json:"enabled_queued_seek"`
		} `json:"hulu-web-site-player-options"`
		HuluClientInAppAccountManagementAddOnsEnabled bool `json:"hulu-client-in-app-account-management-add-ons-enabled"`
		HuluClientCompassEnabled                      bool `json:"hulu-client-compass-enabled"`
		HuluClientFeaturePxsSurveyEnabled             bool `json:"hulu-client-feature-pxs-survey-enabled"`
		HuluClientTrailheadBannerTheme                bool `json:"hulu-client-trailhead-banner-theme"`
		HuluClientAvMetadataBadgingEnabled            bool `json:"hulu-client-av-metadata-badging-enabled"`
		HuluClientIdleTimeMs                          int  `json:"hulu-client-idle-time-ms"`
		HuluWebDevelopProdSitePlayerVersion           struct {
			Nonsub string `json:"nonsub"`
			Sub    string `json:"sub"`
		} `json:"hulu-web-develop-prod-site-player-version"`
		HuluClientPlayerProgressReportInterval                     int  `json:"hulu-client-player-progress-report-interval"`
		HuluClientDeviceTokenLoggingEnabled                        bool `json:"hulu-client-device-token-logging-enabled"`
		HuluClientInAppAccountManagementEnabled                    bool `json:"hulu-client-in-app-account-management-enabled"`
		HuluClientEventPurchaseIdentityVerificationPollingInterval int  `json:"hulu-client-event-purchase-identity-verification-polling-interval"`
		HuluClientCheckProgramRecordability                        bool `json:"hulu-client-check-program-recordability"`
		HuluClientAutoAccountLinkEnabled                           bool `json:"hulu-client-auto-account-link-enabled"`
		HuluWebSmokeSitePlayerOptions                              struct {
			CreditEndCardDuration string `json:"credit_end_card_duration"`
			EnabledAdobeAgent     bool   `json:"enabled_adobe_agent"`
			EnabledQueuedSeek     bool   `json:"enabled_queued_seek"`
		} `json:"hulu-web-smoke-site-player-options"`
		HuluClientFliptray2              bool `json:"hulu-client-fliptray-2"`
		HuluClientOneplayer              bool `json:"hulu-client-oneplayer"`
		HuluClientFeaturePxsSurveyConfig struct {
			PxsShowPercentage     float64 `json:"pxs_show_percentage"`
			PxsAutoDismissSeconds int     `json:"pxs_auto_dismiss_seconds"`
			PxsShowFrequencyDays  int     `json:"pxs_show_frequency_days"`
		} `json:"hulu-client-feature-pxs-survey-config"`
		HuluWebBrowseFlags struct {
			EditorialActionsEnabled    bool   `json:"editorialActionsEnabled"`
			ContextMenuActionV2Enabled bool   `json:"contextMenuActionV2Enabled"`
			VideoTileEnabled           bool   `json:"videoTileEnabled"`
			EnableWebp                 bool   `json:"enableWebp"`
			VariationName              string `json:"variationName"`
		} `json:"hulu-web-browse-flags"`
		HuluClientDetailsCastAndCrew       bool `json:"hulu-client-details-cast-and-crew"`
		HuluClientPlanSelectChartEnabled   bool `json:"hulu-client-plan-select-chart-enabled"`
		HuluClientEndpointURLConfiguration bool `json:"hulu-client-endpoint-url-configuration"`
		HuluClientDvrMsbd                  bool `json:"hulu-client-dvr-msbd"`
		HuluClientMyStuffDecoupled         bool `json:"hulu-client-my-stuff-decoupled"`
		HuluWebSitePlayerVersion           struct {
			Nonsub string `json:"nonsub"`
			Sub    string `json:"sub"`
		} `json:"hulu-web-site-player-version"`
		HuluClientTealiumEventsEnabled bool `json:"hulu-client-tealium-events-enabled"`
		HuluWebChromecastPlayerOptions struct {
			OverrideAdUnits string `json:"overrideAdUnits"`
			Touchstone      bool   `json:"touchstone"`
		} `json:"hulu-web-chromecast-player-options"`
		HuluClientEventPurchasePollingBaseInterval int    `json:"hulu-client-event-purchase-polling-base-interval"`
		HuluClientBrandedCollections               bool   `json:"hulu-client-branded-collections"`
		HuluClientSignupOnWebEnabled               bool   `json:"hulu-client-signup-on-web-enabled"`
		HuluClientCompassSitemapEnabled            bool   `json:"hulu-client-compass-sitemap-enabled"`
		HuluClientAdobeMetrics                     bool   `json:"hulu-client-adobe-metrics"`
		HuluClientGatewayAdLegalDisclaimer         string `json:"hulu-client-gateway-ad-legal-disclaimer"`
		HuluClientFeatureChannelFlipping           bool   `json:"hulu-client-feature-channel-flipping"`
		HuluClientFeaturePlaybackCdnSorting        bool   `json:"hulu-client-feature-playback-cdn-sorting"`
		HuluClientLoginMfa                         bool   `json:"hulu-client-login-mfa"`
		HuluClientLinksharingAppsflyer             bool   `json:"hulu-client-linksharing-appsflyer"`
		HuluClientFeatureHdr                       bool   `json:"hulu-client-feature-hdr"`
	} `json:"flags"`
	NoahProductTitle    string `json:"noah_product_title"`
	CriterionCollection int    `json:"criterion_collection"`
	KeyID               int    `json:"key_id"`
}
