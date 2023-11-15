package spotify

type SpotifyPlayer struct {
	metadata Track
	paused bool
	trackTime int64
	uri string
	contextUri string
}

type PlaybackState struct {
	Current   string `json:"current"`
	TrackTime int    `json:"trackTime"`
	Track     Track  `json:"track"`
}

type Artist struct {
	Gid  string `json:"gid"`
	Name string `json:"name"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type Image struct {
	FileID string `json:"fileId"`
	Size   string `json:"size"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type CoverGroup struct {
	Image []Image `json:"image"`
}

type Album struct {
	Gid        string     `json:"gid"`
	Name       string     `json:"name"`
	Artist     []Artist   `json:"artist"`
	Label      string     `json:"label"`
	Date       Date       `json:"date"`
	CoverGroup CoverGroup `json:"coverGroup"`
}

type ExternalID struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

type File struct {
	FileID string `json:"fileId"`
	Format string `json:"format,omitempty"`
}

type Preview struct {
	FileID string `json:"fileId"`
	Format string `json:"format"`
}

type Licensor struct {
	UUID string `json:"uuid"`
}

type Track struct {
	Gid                   string       `json:"gid"`
	Name                  string       `json:"name"`
	Album                 Album        `json:"album"`
	Artist                []Artist     `json:"artist"`
	Number                int          `json:"number"`
	DiscNumber            int          `json:"discNumber"`
	Duration              int          `json:"duration"`
	Popularity            int          `json:"popularity"`
	ExternalID            []ExternalID `json:"externalId"`
	File                  []File       `json:"file"`
	Preview               []Preview    `json:"preview"`
	EarliestLiveTimestamp int          `json:"earliestLiveTimestamp"`
	HasLyrics             bool         `json:"hasLyrics"`
	Licensor              Licensor     `json:"licensor"`
}

type TracksState struct {
	Current struct {
		URI      string `json:"uri"`
		UID      string `json:"uid"`
		Gid      string `json:"gid"`
		Metadata struct {
			ArtistURI string `json:"artist_uri"`
			AlbumURI  string `json:"album_uri"`
			Duration  string `json:"duration"`
		} `json:"metadata"`
	} `json:"current"`
	Next []struct {
		URI      string `json:"uri"`
		UID      string `json:"uid"`
		Gid      string `json:"gid"`
		Metadata struct {
			AlbumArtistName        string `json:"album_artist_name"`
			AlbumTitle             string `json:"album_title"`
			AlbumURI               string `json:"album_uri"`
			ArtistName             string `json:"artist_name"`
			ArtistURI              string `json:"artist_uri"`
			CollectionCanAdd       string `json:"collection.can_add"`
			CollectionCanBan       string `json:"collection.can_ban"`
			CollectionInCollection string `json:"collection.in_collection"`
			CollectionIsBanned     string `json:"collection.is_banned"`
			Duration               string `json:"duration"`
			HasLyrics              string `json:"has_lyrics"`
			ImageSmallURL          string `json:"image_small_url"`
			ImageURL               string `json:"image_url"`
			MarkedForDownload      string `json:"marked_for_download"`
			Title                  string `json:"title"`
		} `json:"metadata,omitempty"`
	} `json:"next"`
	Prev []any `json:"prev"`
}

type SearchResult struct {
	Results struct {
		Tracks struct {
			Hits []struct {
				Name    string `json:"name"`
				URI     string `json:"uri"`
				Image   string `json:"image"`
				Artists []struct {
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"artists"`
				Album struct {
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"album"`
				Duration    int     `json:"duration"`
				Mogef19     bool    `json:"mogef19"`
				Popularity  float64 `json:"popularity"`
				LyricsMatch bool    `json:"lyricsMatch"`
			} `json:"hits"`
			Total int `json:"total"`
		} `json:"tracks"`
		Albums struct {
			Hits []struct {
				Name    string `json:"name"`
				URI     string `json:"uri"`
				Image   string `json:"image"`
				Artists []struct {
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"artists"`
			} `json:"hits"`
			Total int `json:"total"`
		} `json:"albums"`
		Artists struct {
			Hits []struct {
				Name     string `json:"name"`
				URI      string `json:"uri"`
				Image    string `json:"image"`
				Verified bool   `json:"verified"`
			} `json:"hits"`
			Total int `json:"total"`
		} `json:"artists"`
		Playlists struct {
			Hits []struct {
				Name           string `json:"name"`
				URI            string `json:"uri"`
				Image          string `json:"image"`
				FollowersCount int    `json:"followersCount"`
				Author         string `json:"author"`
				Personalized   bool   `json:"personalized"`
			} `json:"hits"`
			Total int `json:"total"`
		} `json:"playlists"`
		Profiles struct {
			Hits  []interface{} `json:"hits"`
			Total int           `json:"total"`
		} `json:"profiles"`
		Genres struct {
			Hits  []interface{} `json:"hits"`
			Total int           `json:"total"`
		} `json:"genres"`
		TopHit struct {
			Hits []struct {
				Name    string `json:"name"`
				URI     string `json:"uri"`
				Image   string `json:"image"`
				Artists []struct {
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"artists"`
				Album struct {
					Name string `json:"name"`
					URI  string `json:"uri"`
				} `json:"album"`
				Duration    int     `json:"duration"`
				Mogef19     bool    `json:"mogef19"`
				Popularity  float64 `json:"popularity"`
				LyricsMatch bool    `json:"lyricsMatch"`
			} `json:"hits"`
			Total int `json:"total"`
		} `json:"topHit"`
		Shows struct {
			Hits []struct {
				Name         string `json:"name"`
				URI          string `json:"uri"`
				Image        string `json:"image"`
				ShowType     string `json:"showType"`
				MusicAndTalk bool   `json:"musicAndTalk"`
			} `json:"hits"`
			Total int `json:"total"`
		} `json:"shows"`
		Audioepisodes struct {
			Hits []struct {
				Name         string `json:"name"`
				URI          string `json:"uri"`
				Image        string `json:"image"`
				Explicit     bool   `json:"explicit"`
				Duration     int    `json:"duration"`
				MusicAndTalk bool   `json:"musicAndTalk"`
			} `json:"hits"`
			Total int `json:"total"`
		} `json:"audioepisodes"`
	} `json:"results"`
	RequestID       string   `json:"requestId"`
	CategoriesOrder []string `json:"categoriesOrder"`
}

type InstanceData struct {
	DeviceID        string `json:"device_id"`
	DeviceName      string `json:"device_name"`
	DeviceType      string `json:"device_type"`
	CountryCode     string `json:"country_code"`
	PreferredLocale string `json:"preferred_locale"`
}