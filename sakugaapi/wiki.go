package sakugaapi

type WikiAPI struct {
	URL string
}

func newWikiAPI(baseURL string) *WikiAPI {
	newAPI := WikiAPI{
		URL: baseURL + "/wiki",
	}
	return &newAPI
}
