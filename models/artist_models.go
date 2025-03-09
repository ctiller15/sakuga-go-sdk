package models

const (
	ArtistListOrderDateOption = "date"
	ArtistListOrderNameOption = "name"
)

var ValidArtistListOrderOptions = []string{ArtistListOrderDateOption, ArtistListOrderNameOption}

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
