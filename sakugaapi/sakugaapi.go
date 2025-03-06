package sakugaapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const (
	defaultApiURL = "https://sakugabooru.com"
)

type TagListAPIResponseItem struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Count     int    `json:"count"`
	Type      int    `json:"type"`
	Ambiguous bool   `json:"ambiguous"`
}

type TagListResponseResult struct {
	ID        int
	Name      string
	Count     int
	Type      int
	Ambiguous bool
}

type PostListAPIResponseItem struct {
	ID                  int           `json:"id"`
	Tags                string        `json:"tags"`
	CreatedAt           int64         `json:"created_at"`
	UpdatedAt           int64         `json:"updated_at"`
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
	LastNotedAt         int64         `json:"last_noted_at"`
	LastCommentedAt     int64         `json:"last_commented_at"`
}

type PostListResponseResult struct {
	ID                  int
	Tags                string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	CreatorID           int
	ApproverID          int
	Author              string
	Change              int
	Source              string
	Score               int
	MD5                 string
	FileSize            int
	FileExtension       string
	FileURL             string
	IsShownInIndex      bool
	PreviewURL          string
	PreviewWidth        int
	ActualPreviewWidth  int
	ActualPreviewHeight int
	SampleURL           string
	SampleWidth         int
	SampleHeight        int
	SampleFileSize      int
	JpegURL             string
	JpegWidth           int
	JpegHeight          int
	JpegFileSize        int
	Rating              string
	IsRatingLocked      bool
	HasChildren         bool
	ParentID            int
	Status              string
	IsPending           bool
	Width               int
	Height              int
	IsHeld              bool
	FramesPendingString string
	FramesPending       []interface{}
	FramesString        string
	Frames              []interface{}
	IsNoteLocked        bool
	LastNotedAt         time.Time
	LastCommentedAt     time.Time
}

func mapPostListAPIItemsToResponse(listItems []PostListAPIResponseItem) []PostListResponseResult {
	postListResponseResults := make([]PostListResponseResult, 0)

	for _, item := range listItems {
		newResponseResult := PostListResponseResult{
			ID:                  item.ID,
			Tags:                item.Tags,
			CreatedAt:           time.UnixMilli(item.CreatedAt),
			UpdatedAt:           time.UnixMilli(item.UpdatedAt),
			CreatorID:           item.CreatorID,
			ApproverID:          item.ApproverID,
			Author:              item.Author,
			Change:              item.Change,
			Source:              item.Source,
			MD5:                 item.MD5,
			FileSize:            item.FileSize,
			FileExtension:       item.FileExtension,
			FileURL:             item.FileURL,
			IsShownInIndex:      item.IsShownInIndex,
			PreviewURL:          item.PreviewURL,
			PreviewWidth:        item.PreviewWidth,
			ActualPreviewWidth:  item.ActualPreviewWidth,
			ActualPreviewHeight: item.ActualPreviewHeight,
			SampleURL:           item.SampleURL,
			SampleWidth:         item.SampleWidth,
			SampleHeight:        item.SampleHeight,
			SampleFileSize:      item.SampleFileSize,
			JpegURL:             item.JpegURL,
			JpegWidth:           item.JpegWidth,
			JpegHeight:          item.JpegHeight,
			JpegFileSize:        item.JpegFileSize,
			Rating:              item.Rating,
			IsRatingLocked:      item.IsRatingLocked,
			HasChildren:         item.HasChildren,
			ParentID:            item.ParentID,
			Status:              item.Status,
			IsPending:           item.IsPending,
			Width:               item.Width,
			Height:              item.Height,
			IsHeld:              item.IsHeld,
			FramesPendingString: item.FramesPendingString,
			FramesPending:       item.FramesPending,
			FramesString:        item.FramesString,
			Frames:              item.Frames,
			IsNoteLocked:        item.IsNoteLocked,
			LastNotedAt:         time.UnixMilli(item.LastNotedAt),
			LastCommentedAt:     time.UnixMilli(item.LastCommentedAt),
		}

		postListResponseResults = append(postListResponseResults, newResponseResult)
	}

	return postListResponseResults
}

func mapTagListAPIItemsToResponse(listItems []TagListAPIResponseItem) []TagListResponseResult {
	tagListResponseResults := make([]TagListResponseResult, 0)

	for _, item := range listItems {
		newItem := TagListResponseResult(item)
		tagListResponseResults = append(tagListResponseResults, newItem)
	}

	return tagListResponseResults
}

const (
	TagTypeGeneral   string = "general"
	TagTypeArtist    string = "artist"
	TagTypeCopyright string = "copyright"
	TagTypeCharacter string = "character"
)

var VALID_TAG_TYPES = []string{TagTypeGeneral, TagTypeArtist, TagTypeCopyright, TagTypeCharacter}

type TagRelatedOptions struct {
	Tags []string // The tag names to query
	Type string   // Restrict results to this tag type. Can be "general", "artist", "copyright", or "character"
}

type TagListOptions struct {
	Limit       int    // How many tags to retrieve. Setting to 0 returns every tag.
	Page        int    // The page number
	Order       string // can be date, count, or name
	ID          int    // The ID number of the tag
	AfterID     int    // Return all tags with an id number greater than this.
	Name        string // The exact name of the tag
	NamePattern string // Search for any tag that has this parameter in its name
}

type PostsListOptions struct {
	Limit int      // How many posts you want to retrieve. Hard limit of 100 per request
	Page  int      // The page number. It starts at 1.
	Tags  []string // Tags to search for. Any tag combination will work, including meta-tags
}

func (p *PostsAPI) List(opts *PostsListOptions) ([]PostListResponseResult, error) {
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

	return mapPostListAPIItemsToResponse(postListItems), nil
	// work out options later.
}

// Struct that holds all the methods for posts
type PostsAPI struct {
	URL string
}

type TagsAPI struct {
	URL string
}

func newPostsAPI(baseURL string) *PostsAPI {
	newAPI := PostsAPI{
		URL: baseURL + "/post.json",
	}
	return &newAPI
}

type relatedTagResult [2]string // First value is the tag, second value is the tag ID

func (t *TagsAPI) Related(opts *TagRelatedOptions) (map[string][]relatedTagResult, error) {
	var resBody []byte

	res, err := http.Get(t.URL)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	relatedTagAPIResponse := make(map[string][]relatedTagResult)
	err = json.Unmarshal(resBody, &relatedTagAPIResponse)

	if err != nil {
		return nil, err
	}

	// Incorrect. Map to final response body
	return relatedTagAPIResponse, nil
}

func (t *TagsAPI) List(opts *TagListOptions) ([]TagListResponseResult, error) {
	var resBody []byte

	res, err := http.Get(t.URL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	tagListItems := make([]TagListAPIResponseItem, 0)
	err = json.Unmarshal(resBody, &tagListItems)

	if err != nil {
		return nil, err
	}

	return mapTagListAPIItemsToResponse(tagListItems), nil
	// work out options later.
}

func newTagsAPI(baseURL string) *TagsAPI {
	newAPI := TagsAPI{
		URL: baseURL + "/tag.json",
	}
	return &newAPI
}

// The SakugaAPI struct contains the structs for all of the methods
// as documented in the API.
type SakugaAPI struct {
	Posts *PostsAPI
	Tags  *TagsAPI
	URL   string
}

// Initializes a new SakugaAPI and all of its child structs
func NewAPI() *SakugaAPI {
	newAPI := SakugaAPI{
		Posts: newPostsAPI(defaultApiURL),
		Tags:  newTagsAPI(defaultApiURL),
		URL:   defaultApiURL,
	}

	return &newAPI
}

func (s *SakugaAPI) SetHomeURL(newURL string) {
	s.Posts = newPostsAPI(newURL)
	s.Tags = newTagsAPI(newURL)
	s.URL = newURL
}
