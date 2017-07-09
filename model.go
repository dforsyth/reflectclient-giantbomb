package giantbomb

type ApiKey struct {
	ApiKey *string `json:"api_key"`
}

type Meta struct {
	NumberOfTotalResults int    `json:"number_of_total_results"`
	Error                string `json:"error"`
	StatusCode           int    `json:"status_code"`
}

type Image struct {
	IconImg   string `json:"icon_img"`
	TinyUrl   string `json:"tiny_url"`
	SmallUrl  string `json:"small_url"`
	MediumUrl string `json:"medium_url"`
	ThumbUrl  string `json:"thumb_url"`
	SuperUrl  string `json:"super_url"`
	ScreenUrl string `json:"screen_url"`
}

type Video struct {
	ApiDetailUrl  string `json:"api_detail_url"`
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Image         *Image `json:"image"`
	HighUrl       string `json:"high_url"`
	LowUrl        string `json:"low_url"`
	HdUrl         string `json:"hd_url"`
	LengthSeconds int    `json:"length_seconds"`
	User          string `json:"user"`
	PublishDate   string `json:"publish_date"`
	SiteDetailUrl string `json:"site_detail_url"`
	YoutubeId     string `json:"youtube_id"`
	VideoType     string `json:"video_type"`
	Deck          string `json:"deck"`
}

type VideoType struct {
	ApiDetailUrl  string `json:"api_detail_url"`
	Deck          string `json:"deck"`
	Id            int    `json:"id"`
	Name          string `json:"name"`
	SiteDetailUrl string `json:"site_detail_url"`
}

type SearchResult struct {
	Id           int64  `json:"id"`
	ResourceType string `json:"resource_type"`
	Name         string `json:"name"`
	Image        *Image `json:"image"`
}
