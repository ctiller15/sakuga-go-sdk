package models

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
