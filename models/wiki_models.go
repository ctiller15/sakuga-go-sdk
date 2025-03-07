package models

import "time"

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
