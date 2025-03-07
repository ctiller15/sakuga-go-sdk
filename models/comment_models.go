package models

import "time"

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
