package sakugaapi

import (
	"encoding/json"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
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

func (a *ArtistsAPI) List(opts *models.ArtistListOptions) ([]models.ArtistListAPIResponseItem, error) {
	url := a.URL + ".json"

	body, err := utils.Fetch(url)

	if err != nil {
		return nil, err
	}

	artistListItems := make([]models.ArtistListAPIResultItem, 0)
	err = json.Unmarshal(body, &artistListItems)

	if err != nil {
		return nil, err
	}

	return mapArtistListAPIItemsToResponse(artistListItems), nil
}

func mapArtistListAPIItemsToResponse(listItems []models.ArtistListAPIResultItem) []models.ArtistListAPIResponseItem {
	artistListResponseResults := make([]models.ArtistListAPIResponseItem, 0)

	for _, item := range listItems {
		newItem := models.ArtistListAPIResponseItem(item)
		artistListResponseResults = append(artistListResponseResults, newItem)
	}

	return artistListResponseResults
}
