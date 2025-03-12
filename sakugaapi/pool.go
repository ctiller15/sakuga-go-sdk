package sakugaapi

type PoolAPI struct {
	URL string
}

func newPoolAPI(baseURL string) *PoolAPI {
	newAPI := PoolAPI{
		URL: baseURL + "/pool",
	}
	return &newAPI
}
