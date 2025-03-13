package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
)

type ForumAPI struct {
	URL string
}

func newForumAPI(baseURL string) *ForumAPI {
	newAPI := ForumAPI{
		URL: baseURL + "/forum",
	}
	return &newAPI
}

func mapForumListAPIResultToResponse(items []sakugamodels.ForumListAPIResultItem) ([]sakugamodels.ForumListResponseItem, error) {
	forumListAPIResponseResults := make([]sakugamodels.ForumListResponseItem, 0)

	for _, item := range items {
		updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		newItem := sakugamodels.ForumListResponseItem{
			Body:      item.Body,
			Creator:   item.Creator,
			CreatorID: item.CreatorID,
			ID:        item.ID,
			ParentID:  item.ParentID,
			Title:     item.Title,
			UpdatedAt: updatedAt,
			Pages:     item.Pages,
		}

		forumListAPIResponseResults = append(forumListAPIResponseResults, newItem)
	}

	return forumListAPIResponseResults, nil
}

func (f *ForumAPI) List(opts *sakugamodels.ForumListOptions) ([]sakugamodels.ForumListResponseItem, error) {
	url, err := sakugautils.CreateForumListUrl(f.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := sakugautils.Fetch(url)
	if err != nil {
		return nil, err
	}

	forumListAPIResult := make([]sakugamodels.ForumListAPIResultItem, 0)
	err = json.Unmarshal(body, &forumListAPIResult)

	if err != nil {
		return nil, err
	}

	mappedResponse, err := mapForumListAPIResultToResponse(forumListAPIResult)
	if err != nil {
		return nil, err
	}
	return mappedResponse, nil
}
