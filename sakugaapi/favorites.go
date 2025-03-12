package sakugaapi

type FavoritesAPI struct {
	URL string
}

func newFavoritesAPI(baseURL string) *FavoritesAPI {
	newAPI := FavoritesAPI{
		URL: baseURL + "/favorite",
	}

	return &newAPI
}
