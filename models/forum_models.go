package models

import "time"

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
