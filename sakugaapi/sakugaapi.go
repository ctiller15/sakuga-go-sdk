package sakugaapi

const (
	defaultApiURL = "https://sakugabooru.com"
)

// The SakugaAPI struct contains the structs for all of the methods
// as documented in the API.
type SakugaAPI struct {
	Posts    *PostsAPI
	Tags     *TagsAPI
	Artists  *ArtistsAPI
	Comments *CommentsAPI
	URL      string
}

// Initializes a new SakugaAPI and all of its child structs
func NewAPI() *SakugaAPI {
	newAPI := SakugaAPI{
		Posts:    newPostsAPI(defaultApiURL),
		Tags:     newTagsAPI(defaultApiURL),
		Artists:  newArtistsAPI(defaultApiURL),
		Comments: newCommentsAPI(defaultApiURL),
		URL:      defaultApiURL,
	}

	return &newAPI
}

func (s *SakugaAPI) SetHomeURL(newURL string) {
	s.Posts = newPostsAPI(newURL)
	s.Tags = newTagsAPI(newURL)
	s.Artists = newArtistsAPI(newURL)
	s.Comments = newCommentsAPI(newURL)
	s.URL = newURL
}
