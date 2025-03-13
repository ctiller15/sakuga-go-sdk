package sakugamodels

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

const (
	TagListOrderOptionDate  = "date"
	TagListOrderOptionCount = "count"
	TagListOrderOptionName  = "name"
)

var TagListOrderOptions = []string{
	TagListOrderOptionDate,
	TagListOrderOptionCount,
	TagListOrderOptionName,
}

type TagListOptions struct {
	Limit       int    // How many tags to retrieve. Setting to 0 returns every tag.
	Page        int    // The page number
	Order       string // can be date, count, or name
	ID          int    // The ID number of the tag. Cannot be used with other options.
	AfterID     int    // Return all tags with an id number greater than this.
	Name        string // The exact name of the tag
	NamePattern string // Search for any tag that has this parameter in its name
}

// * I legitimately have no clue what NamePattern does after testing, and there doesn't seem to be much documentation on it.
// Name seems to fuzzy match strings, but NamePattern seems to just get everything.
// If anyone knows and wants to put up a PR, I would greatly appreciate it.

type RelatedTagResult [2]interface{} // First value is the tag - str, second value is the tag ID - int
