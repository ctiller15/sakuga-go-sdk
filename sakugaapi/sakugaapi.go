package sakugaapi

import (
	"encoding/json"
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

func mapArtistListAPIItemsToResponse(listItems []models.ArtistListAPIResultItem) []models.ArtistListAPIResponseItem {
	artistListResponseResults := make([]models.ArtistListAPIResponseItem, 0)

	for _, item := range listItems {
		newItem := models.ArtistListAPIResponseItem(item)
		artistListResponseResults = append(artistListResponseResults, newItem)
	}

	return artistListResponseResults
}

func mapCommentShowAPIItemToResponse(item models.CommentShowAPIResultItem) (models.CommentShowResponseItem, error) {
	// Apparently the format is ISO 8601, but this is working for now so I'm happy.
	//https://stackoverflow.com/questions/522251/whats-the-difference-between-iso-8601-and-rfc-3339-date-formats
	parsedTime, err := time.Parse(time.RFC3339, item.CreatedAt)
	if err != nil {
		return models.CommentShowResponseItem{}, err
	}

	return models.CommentShowResponseItem{
		ID:        item.ID,
		CreatedAt: parsedTime,
		PostID:    item.PostID,
		Creator:   item.Creator,
		CreatorID: item.CreatorID,
		Body:      item.Body,
	}, nil
}

func (p *PostsAPI) List(opts *models.PostsListOptions) ([]models.PostListResponseResult, error) {
	url := p.URL + ".json"
	body, err := utils.Fetch(url)
	if err != nil {
		return nil, err
	}

	postListItems := make([]models.PostListAPIResponseItem, 0)
	err = json.Unmarshal(body, &postListItems)

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

type ArtistsAPI struct {
	URL string
}

type CommentsAPI struct {
	URL string
}

func newPostsAPI(baseURL string) *PostsAPI {
	newAPI := PostsAPI{
		URL: baseURL + "/post",
	}
	return &newAPI
}

type relatedTagResult [2]interface{} // First value is the tag - str, second value is the tag ID - int

func (t *TagsAPI) Related(opts *models.TagRelatedOptions) (map[string][]models.RelatedTagResponse, error) {
	url, err := utils.CreateRelatedTagsUrl(t.URL, opts)

	body, err := utils.Fetch(url)
	if err != nil {
		return nil, err
	}

	relatedTagAPIResponse := make(map[string][]relatedTagResult)
	err = json.Unmarshal(body, &relatedTagAPIResponse)

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

func (c *CommentsAPI) Show(opts *models.CommentShowOptions) (models.CommentShowResponseItem, error) {
	url, err := utils.CreateCommentShowUrl(c.URL, opts)
	if err != nil {
		return models.CommentShowResponseItem{}, err
	}

	body, err := utils.Fetch(url)

	if err != nil {
		return models.CommentShowResponseItem{}, err
	}

	commentShowItem := models.CommentShowAPIResultItem{}
	err = json.Unmarshal(body, &commentShowItem)

	if err != nil {
		return models.CommentShowResponseItem{}, err
	}

	commentShowResponse, err := mapCommentShowAPIItemToResponse(commentShowItem)
	if err != nil {
		return models.CommentShowResponseItem{}, err
	}

	return commentShowResponse, nil
}

func newTagsAPI(baseURL string) *TagsAPI {
	newAPI := TagsAPI{
		URL: baseURL + "/tag",
	}
	return &newAPI
}

func newArtistsAPI(baseURL string) *ArtistsAPI {
	newAPI := ArtistsAPI{
		URL: baseURL + "/artist",
	}
	return &newAPI
}

func newCommentsAPI(baseURL string) *CommentsAPI {
	newAPI := CommentsAPI{
		URL: baseURL + "/comment",
	}
	return &newAPI
}

// The SakugaAPI struct contains the structs for all of the methods
// as documented in the API.
type SakugaAPI struct {
	Posts    *PostsAPI
	Tags     *TagsAPI
	Artists  *ArtistsAPI
	Comments *CommentsAPI
	URL      string
}

// Initializes a new SakugaAPI and all of its child structs
func NewAPI() *SakugaAPI {
	newAPI := SakugaAPI{
		Posts:    newPostsAPI(defaultApiURL),
		Tags:     newTagsAPI(defaultApiURL),
		Artists:  newArtistsAPI(defaultApiURL),
		Comments: newCommentsAPI(defaultApiURL),
		URL:      defaultApiURL,
	}

	return &newAPI
}

func (s *SakugaAPI) SetHomeURL(newURL string) {
	s.Posts = newPostsAPI(newURL)
	s.Tags = newTagsAPI(newURL)
	s.Artists = newArtistsAPI(newURL)
	s.Comments = newCommentsAPI(newURL)
	s.URL = newURL
}
