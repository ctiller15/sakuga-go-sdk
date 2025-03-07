package models

import "time"

type TagRelatedOptions struct {
	Tags []string // The tag names to query
	Type string   // Restrict results to this tag type. Can be "general", "artist", "copyright", or "character"
}

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

type RelatedTagResponse struct {
	Name string
	ID   int
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

type ArtistListOptions struct {
	Name  string // The name or name fragment of the artist
	Order string // Can be date or name
	Page  int    // The Page number
}

type ArtistListAPIResultItem struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	AliasID int      `json:"alias_id"`
	GroupID int      `json:"group_id"`
	Urls    []string `json:"urls"`
}

// TODO: figure out how to handle nullable values.
// For integers/ids we can likely return a -1 value as an indicator

type ArtistListAPIResponseItem struct {
	ID      int
	Name    string
	AliasID int
	GroupID int
	Urls    []string
}

type CommentShowOptions struct {
	ID int // The ID of the comment to show. Required.
}

type CommentShowAPIResultItem struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	PostID    int    `json:"post_id"`
	Creator   string `json:"creator"`
	CreatorID int    `json:"creator_id"`
	Body      string `json:"body"`
}

type CommentShowResponseItem struct {
	ID        int
	CreatedAt time.Time
	PostID    int
	Creator   string
	CreatorID int
	Body      string
}

type WikiListOptions struct {
	Order string // How you want pages to be ordered. Can be `title` or `date`.
	Limit int    // The number of pages to retrieve
	Page  int    // The page number
	Query string // A word or phrase to search for
}

type WikiListAPIResultItem struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UpdaterID int    `json:"updater_id"`
	Locked    bool   `json:"locked"`
	Version   int    `json:"version"`
}

type WikiListResponseItem struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	Body      string
	UpdaterID int
	Locked    bool
	Version   int
}

type NoteListOptions struct {
	PostID int // The post id number to retrieve notes for
}

type NoteListAPIResultItem struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatorID int    `json:"creator_id"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsActive  bool   `json:"is_active"`
	PostID    int    `json:"post_id"`
	Body      string `json:"body"`
	Version   int    `json:"version"`
}

type NoteListResponseItem struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatorID int
	X         int
	Y         int
	Width     int
	Height    int
	IsActive  bool
	PostID    int
	Body      string
	Version   int
}

type NoteSearchOptions struct {
	Query string // A word or phrase to search for
}

// Same result body as the List endpoint.

type NoteHistorySearchOptions struct {
	Limit  int // How many versions to retrieve
	Page   int // The page offset
	PostID int // The post ID number to retrieve note versions for
	ID     int // The note id number to retrieve versions for
}

type NoteHistoryAPIResultItem struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatorID int    `json:"creator_id"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsActive  bool   `json:"is_active"`
	PostID    int    `json:"post_id"`
	Body      string `json:"body"`
	Version   int    `json:"version"`
}

type NoteHistoryResponseItem struct {
	CreatedAt string
	UpdatedAt string
	CreatorID int
	X         int
	Y         int
	Width     int
	Height    int
	IsActive  bool
	PostID    int
	Body      string
	Version   int
}

type UserSearchOptions struct {
	ID   int    // The id number of the user
	Name string // The name of the user
}

type UserSearchAPIResultItem struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

type UserSearchAPIResponseItem struct {
	Name string
	ID   int
}

type ForumListOptions struct {
	ParentID int // The parent ID number. Returns all responses to that post.
}

type ForumListAPIResultItem struct {
	Body      string `json:"body"`
	Creator   string `json:"creator"`
	CreatorID int    `json:"creator_id"`
	ID        int    `json:"id"`
	ParentID  int    `json:"parent_id"`
	Title     string `json:"title"`
	UpdatedAt string `json:"updated_at"`
	Pages     int    `json:"pages"`
}

type ForumListResponseItem struct {
	Body      string
	Creator   string
	CreatorID int
	ID        int
	ParentID  int
	Title     string
	UpdatedAt time.Time
	Pages     int
}

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

// Valid example that doesn't give a 404. (If the ID has no results it 404's)
// https://sakugabooru.com/favorite/list_users.json?id=3

type FavoriteListUsersOptions struct {
	ID int // The post ID
}

type FavoriteListUsersAPIResult struct {
	FavoritedUsers string `json:"favorited_users"`
}

type FavoriteListUsersAPIResponse struct {
	FavoritedUsers []string
}

type RelatedTagResult [2]interface{} // First value is the tag - str, second value is the tag ID - int
