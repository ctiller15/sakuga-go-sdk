package sakugaapi

const (
	defaultApiURL = "https://sakugabooru.com"
)

// The SakugaAPI struct contains the structs for all of the methods
// as documented in the API.
type SakugaAPI struct {
	Posts     *PostsAPI
	Tags      *TagsAPI
	Artists   *ArtistsAPI
	Comments  *CommentsAPI
	Wiki      *WikiAPI
	Notes     *NotesAPI
	Users     *UsersAPI
	Forum     *ForumAPI
	Pools     *PoolAPI
	Favorites *FavoritesAPI
	URL       string
}

// Initializes a new SakugaAPI and all of its child structs
func NewAPI() *SakugaAPI {
	newAPI := SakugaAPI{
		Posts:     newPostsAPI(defaultApiURL),
		Tags:      newTagsAPI(defaultApiURL),
		Artists:   newArtistsAPI(defaultApiURL),
		Comments:  newCommentsAPI(defaultApiURL),
		Wiki:      newWikiAPI(defaultApiURL),
		Notes:     newNotesAPI(defaultApiURL),
		Users:     newUsersAPI(defaultApiURL),
		Forum:     newForumAPI(defaultApiURL),
		Pools:     newPoolAPI(defaultApiURL),
		Favorites: newFavoritesAPI(defaultApiURL),
		URL:       defaultApiURL,
	}

	return &newAPI
}

func (s *SakugaAPI) SetHomeURL(newURL string) {
	s.Posts = newPostsAPI(newURL)
	s.Tags = newTagsAPI(newURL)
	s.Artists = newArtistsAPI(newURL)
	s.Comments = newCommentsAPI(newURL)
	s.Wiki = newWikiAPI(newURL)
	s.Notes = newNotesAPI(newURL)
	s.Users = newUsersAPI(newURL)
	s.Forum = newForumAPI(newURL)
	s.Pools = newPoolAPI(newURL)
	s.Favorites = newFavoritesAPI(newURL)
	s.URL = newURL
}
