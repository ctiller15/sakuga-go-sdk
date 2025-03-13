package sakugamodels

// The options for the User.Search method
type UserSearchOptions struct {
	ID   int    // The id number of the user
	Name string // The name of the user
}

type UserSearchAPIResultItem struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// The response for the User.Search method
type UserSearchAPIResponseItem struct {
	Name string
	ID   int
}
