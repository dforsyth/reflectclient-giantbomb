package giantbomb

import (
	"github.com/dforsyth/reflectclient"
	"net/http"
)

// Stores a token for the API.
// Implement this in the host language.
type TokenProvider interface {
	GetToken() string
}

//go:generate wrap -service=GiantBombService
type GiantBombService struct {
	Videos     func(*VideosParams) (*VideoListResponse, error)  `rc_method:"GET" rc_path:"videos/"`
	Video      func(*VideoParams) (*VideoResponse, error)       `rc_method:"GET" rc_path:"video/{video_id}/"`
	VideoTypes func() (*VideoTypesResponse, error)              `rc_method:"GET" rc_path:"video_types/"`
	Validate   func(*ValidateParams) (*ApiKeyResponse, error)   `rc_method:"GET" rc_path:"validate/"`
	Search     func(*SearchParams) (*SearchListResponse, error) `rc_method:"GET" rc_path:"search/"`
}

type VideosParams struct {
	Offset    int    `rc_name:"offset" rc_feature:"query"`
	Limit     int    `rc_name:"limit" rc_feature:"query" rc_options:"omitempty"`
	VideoType string `rc_name:"video_type" rc_feature:"query" rc_options:"omitempty"`
}

type VideoParams struct {
	VideoId int `rc_name:"video_id" rc_feature:"path"`
}

type ValidateParams struct {
	LinkCode string `rc_name:"link_code" rc_feature:"query"`
}

type SearchParams struct {
	Query     string `rc_name:"query" rc_feature:"query"`
	Resources string `rc_name="resources" rc_feature:"query"`
}

func CreateGiantBombServiceDefault(baseUrl string, tokenProvider TokenProvider, useragent string) (*GiantBombService, error) {
	return CreateGiantBombService(baseUrl, tokenProvider, useragent, http.DefaultClient)
}

func CreateGiantBombService(baseUrl string, tokenProvider TokenProvider, userAgent string, httpClient *http.Client) (*GiantBombService, error) {
	client, err := reflectclient.NewBuilder().
		SetHttpClient(httpClient).
		BaseUrl(baseUrl).
		SetUnmarshaler(&reflectclient.JsonUnmarshaler{}).
		AddRequestTransformer(func(r *http.Request) *http.Request {
			q := r.URL.Query()
			q.Set("format", "json")
			r.URL.RawQuery = q.Encode()
			// If a token is available, add it as a query param.
			if token := tokenProvider.GetToken(); token != "" {
				q := r.URL.Query()
				q.Set("api_key", token)
				r.URL.RawQuery = q.Encode()
			}
			return r
		}).
		AddRequestTransformer(func(r *http.Request) *http.Request {
			r.Header.Set("User-Agent", userAgent)
			return r
		}).
		Build()
	if err != nil {
		return nil, err
	}

	service := new(GiantBombService)
	if err := client.Init(service); err != nil {
		return nil, err
	}

	return service, nil
}

type VideoResponse struct {
	*Meta
	Results *Video `json:"results"`
}

type VideoListResponse struct {
	*Meta
	Results []*Video `json:"results"`
}

func (r *VideoListResponse) Length() int {
	return len(r.Results)
}

func (r *VideoListResponse) ItemAt(pos int) *Video {
	return r.Results[pos]
}

type VideoTypesResponse struct {
	*Meta
	Results []*VideoType `json:"results"`
}

func (r *VideoTypesResponse) Length() int {
	return len(r.Results)
}

func (r *VideoTypesResponse) ItemAt(pos int) *VideoType {
	return r.Results[pos]
}

type ApiKeyResponse struct {
	ApiKey string `json:"api_key"`
}

type SearchListResponse struct {
	*Meta
	Results []*SearchResult `json:"results"`
}

func (r *SearchListResponse) Length() int {
	return len(r.Results)
}

func (r *SearchListResponse) ItemAt(pos int) *SearchResult {
	return r.Results[pos]
}
