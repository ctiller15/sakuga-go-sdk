package sakugaapi

import (
	"encoding/json"
	"strconv"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
)

type TagsAPI struct {
	URL string
}

func newTagsAPI(baseURL string) *TagsAPI {
	newAPI := TagsAPI{
		URL: baseURL + "/tag",
	}
	return &newAPI
}

func (t *TagsAPI) Related(opts *models.TagRelatedOptions) (map[string][]models.RelatedTagResponse, error) {
	url, err := utils.CreateRelatedTagsUrl(t.URL, opts)

	body, err := utils.Fetch(url)
	if err != nil {
		return nil, err
	}

	relatedTagAPIResponse := make(map[string][]models.RelatedTagResult)
	err = json.Unmarshal(body, &relatedTagAPIResponse)

	if err != nil {
		return nil, err
	}

	mappedResponse, err := mapTagRelatedAPIResultToResponse(relatedTagAPIResponse)
	if err != nil {
		return nil, err
	}
	return mappedResponse, nil
}

func (t *TagsAPI) List(opts *models.TagListOptions) ([]models.TagListResponseResult, error) {
	url := t.URL + ".json"

	body, err := utils.Fetch(url)

	if err != nil {
		return nil, err
	}

	tagListItems := make([]models.TagListAPIResponseItem, 0)
	err = json.Unmarshal(body, &tagListItems)

	if err != nil {
		return nil, err
	}

	return mapTagListAPIItemsToResponse(tagListItems), nil
	// work out options later.
}

func mapTagRelatedAPIResultToResponse(tagAPIResult map[string][]models.RelatedTagResult) (map[string][]models.RelatedTagResponse, error) {
	response := make(map[string][]models.RelatedTagResponse)

	for k, v := range tagAPIResult {
		newResult := make([]models.RelatedTagResponse, 0)

		for _, tagResult := range v {
			var ID int

			switch v := tagResult[1].(type) {
			case int:
				ID = v
			case string:
				convID, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}

				ID = convID
			}

			newTagResponse := models.RelatedTagResponse{
				Name: tagResult[0].(string),
				ID:   ID,
			}
			newResult = append(newResult, newTagResponse)
		}

		response[k] = newResult
	}

	return response, nil
}

func mapTagListAPIItemsToResponse(listItems []models.TagListAPIResponseItem) []models.TagListResponseResult {
	tagListResponseResults := make([]models.TagListResponseResult, 0)

	for _, item := range listItems {
		newItem := models.TagListResponseResult(item)
		tagListResponseResults = append(tagListResponseResults, newItem)
	}

	return tagListResponseResults
}
