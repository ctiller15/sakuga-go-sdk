package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
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

func mapWikiListAPIItemsToResponse(items []sakugamodels.WikiListAPIResultItem) ([]sakugamodels.WikiListResponseItem, error) {
	response := make([]sakugamodels.WikiListResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		mappedResponseItem := sakugamodels.WikiListResponseItem{
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

func (w *WikiAPI) List(opts *sakugamodels.WikiListOptions) ([]sakugamodels.WikiListResponseItem, error) {
	url, err := sakugautils.CreateWikiListUrl(w.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := sakugautils.Fetch(url)

	wikiListItems := make([]sakugamodels.WikiListAPIResultItem, 0)
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
