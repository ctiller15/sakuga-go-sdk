package models

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

type RelatedTagResult [2]interface{} // First value is the tag - str, second value is the tag ID - int
