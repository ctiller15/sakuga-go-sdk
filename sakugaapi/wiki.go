package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
)

type WikiAPI struct {
	URL string
}

func newWikiAPI(baseURL string) *WikiAPI {
	newAPI := WikiAPI{
		URL: baseURL + "/wiki",
	}
	return &newAPI
}

func mapWikiListAPIItemsToResponse(items []models.WikiListAPIResultItem) ([]models.WikiListResponseItem, error) {
	response := make([]models.WikiListResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		mappedResponseItem := models.WikiListResponseItem{
			ID:        item.ID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			Title:     item.Title,
			Body:      item.Body,
			UpdaterID: item.UpdaterID,
			Locked:    item.Locked,
			Version:   item.Version,
		}

		response = append(response, mappedResponseItem)
	}

	return response, nil
}

func (w *WikiAPI) List(opts *models.WikiListOptions) ([]models.WikiListResponseItem, error) {
	url, err := utils.CreateWikiListUrl(w.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := utils.Fetch(url)

	wikiListItems := make([]models.WikiListAPIResultItem, 0)
	err = json.Unmarshal(body, &wikiListItems)

	if err != nil {
		return nil, err
	}

	mappedItems, err := mapWikiListAPIItemsToResponse(wikiListItems)
	if err != nil {
		return nil, err
	}

	return mappedItems, nil
}
