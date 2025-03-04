package sakugaapi

func NewPostsAPI() *PostsAPI {
	newAPI := PostsAPI{}
	return &newAPI
}

type PostsAPI struct {
}

// The SakugaAPI struct contains the structs for all of the methods
// as documented in the API.
type SakugaAPI struct {
	Posts *PostsAPI
}

// Initializes a new SakugaAPI and all of its child structs
func NewAPI() *SakugaAPI {
	newAPI := SakugaAPI{
		Posts: NewPostsAPI(),
	}

	return &newAPI
}
