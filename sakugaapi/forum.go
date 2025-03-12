package sakugaapi

type ForumAPI struct {
	URL string
}

func newForumAPI(baseURL string) *ForumAPI {
	newAPI := ForumAPI{
		URL: baseURL + "/forum",
	}
	return &newAPI
}
