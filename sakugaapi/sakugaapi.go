package sakugaapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	defaultApiURL = "https://sakugabooru.com"
)

type PostListAPIResponseItem struct {
	ID                  int           `json:"id"`
	Tags                string        `json:"tags"`
	CreatedAt           int           `json:"created_at"`
	UpdatedAt           int           `json:"updated_at"`
	CreatorID           int           `json:"creator_id"`
	ApproverID          int           `json:"approver_id"`
	Author              string        `json:"author"`
	Change              int           `json:"change"`
	Source              string        `json:"source"`
	Score               int           `json:"score"`
	MD5                 string        `json:"md5"`
	FileSize            int           `json:"file_size"`
	FileExtension       string        `json:"file_extension"`
	FileURL             string        `json:"file_url"`
	IsShownInIndex      bool          `json:"is_shown_in_index"`
	PreviewURL          string        `json:"preview_url"`
	PreviewWidth        int           `json:"preview_width"`
	ActualPreviewWidth  int           `json:"actual_preview_width"`
	ActualPreviewHeight int           `json:"actual_preview_height"`
	SampleURL           string        `json:"sample_url"`
	SampleWidth         int           `json:"sample_width"`
	SampleHeight        int           `json:"sample_height"`
	SampleFileSize      int           `json:"sample_file_size"`
	JpegURL             string        `json:"jpeg_url"`
	JpegWidth           int           `json:"jpeg_width"`
	JpegHeight          int           `json:"jpeg_height"`
	JpegFileSize        int           `json:"jpeg_file_size"`
	Rating              string        `json:"rating"`
	IsRatingLocked      bool          `json:"is_rating_locked"`
	HasChildren         bool          `json:"has_children"`
	ParentID            int           `json:"parent_id"`
	Status              string        `json:"status"`
	IsPending           bool          `json:"is_pending"`
	Width               int           `json:"width"`
	Height              int           `json:"height"`
	IsHeld              bool          `json:"is_held"`
	FramesPendingString string        `json:"frames_pending_string"`
	FramesPending       []interface{} `json:"frames_pending"` // TODO: Accurately define type.
	FramesString        string        `json:"frames_string"`
	Frames              []interface{} `json:"frames"` // TODO: Accurately define type.
	IsNoteLocked        bool          `json:"is_note_locked"`
	LastNotedAt         int           `json:"last_noted_at"`
	LastCommentedAt     int           `json:"last_commented_at"`
}

type PostsListOptions struct {
	Limit int      // How many posts you want to retrieve. Hard limit of 100 per request
	Page  int      // The page number. It starts at 1.
	Tags  []string // Tags to search for. Any tag combination will work, including meta-tags
}

func (p *PostsAPI) List(opts *PostsListOptions) ([]PostListAPIResponseItem, error) {
	var resBody []byte

	res, err := http.Get(p.URL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	postListItems := make([]PostListAPIResponseItem, 0)
	err = json.Unmarshal(resBody, &postListItems)

	if err != nil {
		return nil, err
	}

	return postListItems, nil
	// work out options later.
	// TODO: Map to a more user-friendly data structure
}

// Struct that holds all the methods for posts
type PostsAPI struct {
	URL string
}

func newPostsAPI(baseURL string) *PostsAPI {
	newAPI := PostsAPI{
		URL: baseURL + "/post.json",
	}
	return &newAPI
}

// The SakugaAPI struct contains the structs for all of the methods
// as documented in the API.
type SakugaAPI struct {
	Posts *PostsAPI
	URL   string
}

// Initializes a new SakugaAPI and all of its child structs
func NewAPI() *SakugaAPI {
	newAPI := SakugaAPI{
		Posts: newPostsAPI(defaultApiURL),
		URL:   defaultApiURL,
	}

	return &newAPI
}

func (s *SakugaAPI) SetHomeURL(newURL string) {
	s.Posts = newPostsAPI(newURL)
	s.URL = newURL
}
