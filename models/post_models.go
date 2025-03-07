package models

import "time"

type PostsListOptions struct {
	Limit int      // How many posts you want to retrieve. Hard limit of 100 per request
	Page  int      // The page number. It starts at 1.
	Tags  []string // Tags to search for. Any tag combination will work, including meta-tags
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
