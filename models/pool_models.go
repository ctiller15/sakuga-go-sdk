package models

import "time"

type PoolListOptions struct {
	Query string // The title
	Page  int    // The page
}

type PoolListAPIResultItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UserID      int    `json:"user_id"`
	IsPublic    bool   `json:"is_public"`
	PostCount   int    `json:"post_count"`
	Description string `json:"description"`
}

type PoolListResponseItem struct {
	ID          int
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      int
	IsPublic    bool
	PostCount   int
	Description string
}

type PoolShowOptions struct {
	ID   int // The pool ID number
	Page int // The Page
}

type PoolShowPostAPIResultItem struct {
	ID                  int           `json:"id"`
	Tags                string        `json:"tags"`
	CreatedAt           string        `json:"created_at"`
	CreatorID           int           `json:"creator_id"`
	Author              string        `json:"author"`
	Change              int           `json:"change"`
	Source              string        `json:"source"`
	Score               int           `json:"score"`
	MD5                 string        `json:"md5"`
	FileSize            int           `json:"file_size"`
	FileURL             string        `json:"file_url"`
	IsShownInIndex      bool          `json:"is_shown_in_index"`
	PreviewURL          string        `json:"preview_url"`
	PreviewWidth        int           `json:"preview_width"`
	PreviewHeight       int           `json:"preview_height"`
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
	HasChildren         bool          `json:"has_children"`
	ParentID            int           `json:"parent_id"`
	Status              string        `json:"status"`
	Width               int           `json:"width"`
	Height              int           `json:"height"`
	IsHeld              bool          `json:"is_held"`
	FramesPendingString string        `json:"frames_pending_string"`
	FramesPending       []interface{} `json:"frames_pending"` // TODO: Accurately define type.
	FramesString        string        `json:"frames_string"`
	Frames              []interface{} `json:"frames"` // TODO: Accurately define type.
}

type PoolShowPostAPIResult struct {
	ID          int                         `json:"id"`
	Name        string                      `json:"name"`
	CreatedAt   string                      `json:"created_at"`
	UpdatedAt   string                      `json:"updated_at"`
	UserID      int                         `json:"user_id"`
	IsPublic    bool                        `json:"is_public"`
	PostCount   int                         `json:"post_count"`
	Description string                      `json:"description"`
	Posts       []PoolShowPostAPIResultItem `json:"posts"`
}

type PoolShowPostResponseItem struct {
	ID                  int           `json:"id"`
	Tags                string        `json:"tags"`
	CreatedAt           time.Time     `json:"created_at"`
	CreatorID           int           `json:"creator_id"`
	Author              string        `json:"author"`
	Change              int           `json:"change"`
	Source              string        `json:"source"`
	Score               int           `json:"score"`
	MD5                 string        `json:"md5"`
	FileSize            int           `json:"file_size"`
	FileURL             string        `json:"file_url"`
	IsShownInIndex      bool          `json:"is_shown_in_index"`
	PreviewURL          string        `json:"preview_url"`
	PreviewWidth        int           `json:"preview_width"`
	PreviewHeight       int           `json:"preview_height"`
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
	HasChildren         bool          `json:"has_children"`
	ParentID            int           `json:"parent_id"`
	Status              string        `json:"status"`
	Width               int           `json:"width"`
	Height              int           `json:"height"`
	IsHeld              bool          `json:"is_held"`
	FramesPendingString string        `json:"frames_pending_string"`
	FramesPending       []interface{} `json:"frames_pending"` // TODO: Accurately define type.
	FramesString        string        `json:"frames_string"`
	Frames              []interface{} `json:"frames"` // TODO: Accurately define type.
}

type PoolShowPostResponse struct {
	ID          int
	Name        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      int
	IsPublic    bool
	PostCount   int
	Description string
	Posts       []PoolShowPostResponse
}
