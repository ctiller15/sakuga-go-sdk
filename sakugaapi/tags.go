package sakugaapi

import (
	"encoding/json"
	"strconv"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
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

func (t *TagsAPI) Related(opts *sakugamodels.TagRelatedOptions) (map[string][]sakugamodels.RelatedTagResponse, error) {
	url, err := sakugautils.CreateRelatedTagsUrl(t.URL, opts)

	body, err := sakugautils.Fetch(url)
	if err != nil {
		return nil, err
	}

	relatedTagAPIResponse := make(map[string][]sakugamodels.RelatedTagResult)
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

func (t *TagsAPI) List(opts *sakugamodels.TagListOptions) ([]sakugamodels.TagListResponseResult, error) {
	url := t.URL + ".json"

	body, err := sakugautils.Fetch(url)

	if err != nil {
		return nil, err
	}

	tagListItems := make([]sakugamodels.TagListAPIResponseItem, 0)
	err = json.Unmarshal(body, &tagListItems)

	if err != nil {
		return nil, err
	}

	return mapTagListAPIItemsToResponse(tagListItems), nil
	// work out options later.
}

func mapTagRelatedAPIResultToResponse(tagAPIResult map[string][]sakugamodels.RelatedTagResult) (map[string][]sakugamodels.RelatedTagResponse, error) {
	response := make(map[string][]sakugamodels.RelatedTagResponse)

	for k, v := range tagAPIResult {
		newResult := make([]sakugamodels.RelatedTagResponse, 0)

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

			newTagResponse := sakugamodels.RelatedTagResponse{
				Name: tagResult[0].(string),
				ID:   ID,
			}
			newResult = append(newResult, newTagResponse)
		}

		response[k] = newResult
	}

	return response, nil
}

func mapTagListAPIItemsToResponse(listItems []sakugamodels.TagListAPIResponseItem) []sakugamodels.TagListResponseResult {
	tagListResponseResults := make([]sakugamodels.TagListResponseResult, 0)

	for _, item := range listItems {
		newItem := sakugamodels.TagListResponseResult(item)
		tagListResponseResults = append(tagListResponseResults, newItem)
	}

	return tagListResponseResults
}
