package sakugaapi

import (
	"encoding/json"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
)

type ArtistsAPI struct {
	URL string
}

func newArtistsAPI(baseURL string) *ArtistsAPI {
	newAPI := ArtistsAPI{
		URL: baseURL + "/artist",
	}
	return &newAPI
}

func (a *ArtistsAPI) List(opts *sakugamodels.ArtistListOptions) ([]sakugamodels.ArtistListAPIResponseItem, error) {
	url := a.URL + ".json"

	body, err := sakugautils.Fetch(url)

	if err != nil {
		return nil, err
	}

	artistListItems := make([]sakugamodels.ArtistListAPIResultItem, 0)
	err = json.Unmarshal(body, &artistListItems)

	if err != nil {
		return nil, err
	}

	return mapArtistListAPIItemsToResponse(artistListItems), nil
}

func mapArtistListAPIItemsToResponse(listItems []sakugamodels.ArtistListAPIResultItem) []sakugamodels.ArtistListAPIResponseItem {
	artistListResponseResults := make([]sakugamodels.ArtistListAPIResponseItem, 0)

	for _, item := range listItems {
		newItem := sakugamodels.ArtistListAPIResponseItem(item)
		artistListResponseResults = append(artistListResponseResults, newItem)
	}

	return artistListResponseResults
}
