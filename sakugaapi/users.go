package sakugaapi

type UsersAPI struct {
	URL string
}

func newUsersAPI(baseURL string) *UsersAPI {
	newAPI := UsersAPI{
		URL: baseURL + "/user",
	}
	return &newAPI
}
