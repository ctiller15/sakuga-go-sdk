package sakugaapi

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
)

const (
	defaultApiURL = "https://sakugabooru.com"
)

func mapTagRelatedAPIResultToResponse(tagAPIResult map[string][]relatedTagResult) (map[string][]models.RelatedTagResponse, error) {
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

func mapPostListAPIItemsToResponse(listItems []models.PostListAPIResponseItem) []models.PostListResponseResult {
	postListResponseResults := make([]models.PostListResponseResult, 0)

	for _, item := range listItems {
		newResponseResult := models.PostListResponseResult{
			ID:                  item.ID,
			Tags:                item.Tags,
			CreatedAt:           time.UnixMilli(item.CreatedAt),
			UpdatedAt:           time.UnixMilli(item.UpdatedAt),
			CreatorID:           item.CreatorID,
			ApproverID:          item.ApproverID,
			Author:              item.Author,
			Change:              item.Change,
			Source:              item.Source,
			MD5:                 item.MD5,
			FileSize:            item.FileSize,
			FileExtension:       item.FileExtension,
			FileURL:             item.FileURL,
			IsShownInIndex:      item.IsShownInIndex,
			PreviewURL:          item.PreviewURL,
			PreviewWidth:        item.PreviewWidth,
			ActualPreviewWidth:  item.ActualPreviewWidth,
			ActualPreviewHeight: item.ActualPreviewHeight,
			SampleURL:           item.SampleURL,
			SampleWidth:         item.SampleWidth,
			SampleHeight:        item.SampleHeight,
			SampleFileSize:      item.SampleFileSize,
			JpegURL:             item.JpegURL,
			JpegWidth:           item.JpegWidth,
			JpegHeight:          item.JpegHeight,
			JpegFileSize:        item.JpegFileSize,
			Rating:              item.Rating,
			IsRatingLocked:      item.IsRatingLocked,
			HasChildren:         item.HasChildren,
			ParentID:            item.ParentID,
			Status:              item.Status,
			IsPending:           item.IsPending,
			Width:               item.Width,
			Height:              item.Height,
			IsHeld:              item.IsHeld,
			FramesPendingString: item.FramesPendingString,
			FramesPending:       item.FramesPending,
			FramesString:        item.FramesString,
			Frames:              item.Frames,
			IsNoteLocked:        item.IsNoteLocked,
			LastNotedAt:         time.UnixMilli(item.LastNotedAt),
			LastCommentedAt:     time.UnixMilli(item.LastCommentedAt),
		}

		postListResponseResults = append(postListResponseResults, newResponseResult)
	}

	return postListResponseResults
}

func mapTagListAPIItemsToResponse(listItems []models.TagListAPIResponseItem) []models.TagListResponseResult {
	tagListResponseResults := make([]models.TagListResponseResult, 0)

	for _, item := range listItems {
		newItem := models.TagListResponseResult(item)
		tagListResponseResults = append(tagListResponseResults, newItem)
	}

	return tagListResponseResults
}

func (p *PostsAPI) List(opts *models.PostsListOptions) ([]models.PostListResponseResult, error) {
	var resBody []byte

	res, err := http.Get(p.URL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	postListItems := make([]models.PostListAPIResponseItem, 0)
	err = json.Unmarshal(resBody, &postListItems)

	if err != nil {
		return nil, err
	}

	return mapPostListAPIItemsToResponse(postListItems), nil
	// work out options later.
}

// Struct that holds all the methods for posts
type PostsAPI struct {
	URL string
}

type TagsAPI struct {
	URL string
}

func newPostsAPI(baseURL string) *PostsAPI {
	newAPI := PostsAPI{
		URL: baseURL + "/post.json",
	}
	return &newAPI
}

type relatedTagResult [2]interface{} // First value is the tag - str, second value is the tag ID - int

func (t *TagsAPI) Related(opts *models.TagRelatedOptions) (map[string][]models.RelatedTagResponse, error) {
	url, err := utils.CreateRelatedTagsUrl(t.URL, opts)

	var resBody []byte

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	relatedTagAPIResponse := make(map[string][]relatedTagResult)
	err = json.Unmarshal(resBody, &relatedTagAPIResponse)

	if err != nil {
		return nil, err
	}

	// Incorrect. Map to final response body
	mappedResponse, err := mapTagRelatedAPIResultToResponse(relatedTagAPIResponse)
	if err != nil {
		return nil, err
	}
	return mappedResponse, nil
}

func (t *TagsAPI) List(opts *models.TagListOptions) ([]models.TagListResponseResult, error) {
	var resBody []byte

	url := t.URL + ".json"

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	tagListItems := make([]models.TagListAPIResponseItem, 0)
	err = json.Unmarshal(resBody, &tagListItems)

	if err != nil {
		return nil, err
	}

	return mapTagListAPIItemsToResponse(tagListItems), nil
	// work out options later.
}

func newTagsAPI(baseURL string) *TagsAPI {
	newAPI := TagsAPI{
		URL: baseURL + "/tag",
	}
	return &newAPI
}

// The SakugaAPI struct contains the structs for all of the methods
// as documented in the API.
type SakugaAPI struct {
	Posts *PostsAPI
	Tags  *TagsAPI
	URL   string
}

// Initializes a new SakugaAPI and all of its child structs
func NewAPI() *SakugaAPI {
	newAPI := SakugaAPI{
		Posts: newPostsAPI(defaultApiURL),
		Tags:  newTagsAPI(defaultApiURL),
		URL:   defaultApiURL,
	}

	return &newAPI
}

func (s *SakugaAPI) SetHomeURL(newURL string) {
	s.Posts = newPostsAPI(newURL)
	s.Tags = newTagsAPI(newURL)
	s.URL = newURL
}
