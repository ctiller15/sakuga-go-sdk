package models

import "time"

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
