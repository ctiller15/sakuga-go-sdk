package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
)

type PoolAPI struct {
	URL string
}

func newPoolAPI(baseURL string) *PoolAPI {
	newAPI := PoolAPI{
		URL: baseURL + "/pool",
	}
	return &newAPI
}

func mapPoolListAPIItemsToResponse(items []models.PoolListAPIResultItem) ([]models.PoolListResponseItem, error) {
	response := make([]models.PoolListResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		mappedResponseItem := models.PoolListResponseItem{
			ID:          item.ID,
			Name:        item.Name,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			UserID:      item.UserID,
			IsPublic:    item.IsPublic,
			PostCount:   item.PostCount,
			Description: item.Description,
		}

		response = append(response, mappedResponseItem)
	}

	return response, nil
}

func mapPoolShowAPIItemsToResponse(items []models.PoolShowPostAPIResultItem) ([]models.PoolShowPostResponseItem, error) {
	response := make([]models.PoolShowPostResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		newItem := models.PoolShowPostResponseItem{
			ID:                  item.ID,
			Tags:                item.Tags,
			CreatedAt:           createdAt,
			CreatorID:           item.CreatorID,
			Author:              item.Author,
			Change:              item.Change,
			Source:              item.Source,
			Score:               item.Score,
			MD5:                 item.MD5,
			FileSize:            item.FileSize,
			FileURL:             item.FileURL,
			IsShownInIndex:      item.IsShownInIndex,
			PreviewURL:          item.PreviewURL,
			PreviewWidth:        item.PreviewWidth,
			PreviewHeight:       item.PreviewHeight,
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
			HasChildren:         item.HasChildren,
			ParentID:            item.ParentID,
			Status:              item.Status,
			Width:               item.Width,
			Height:              item.Height,
			IsHeld:              item.IsHeld,
			FramesPendingString: item.FramesPendingString,
			FramesPending:       item.FramesPending,
			FramesString:        item.FramesString,
			Frames:              item.Frames,
		}

		response = append(response, newItem)
	}

	return response, nil
}

func mapPoolShowAPIItemToResponse(item models.PoolShowPostAPIResult) (models.PoolShowPostResponse, error) {
	createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
	if err != nil {
		return models.PoolShowPostResponse{}, err
	}

	updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
	if err != nil {
		return models.PoolShowPostResponse{}, err
	}

	posts, err := mapPoolShowAPIItemsToResponse(item.Posts)

	response := models.PoolShowPostResponse{
		ID:          item.ID,
		Name:        item.Name,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		UserID:      item.UserID,
		IsPublic:    item.IsPublic,
		PostCount:   item.PostCount,
		Description: item.Description,
		Posts:       posts,
	}

	return response, nil
}

func (p *PoolAPI) List(opts *models.PoolListOptions) ([]models.PoolListResponseItem, error) {
	url, err := utils.CreatePoolListUrl(p.URL, opts)

	if err != nil {
		return nil, err
	}

	body, err := utils.Fetch(url)

	if err != nil {
		return nil, err
	}

	poolListItems := make([]models.PoolListAPIResultItem, 0)
	err = json.Unmarshal(body, &poolListItems)

	if err != nil {
		return nil, err
	}

	mappedResponse, err := mapPoolListAPIItemsToResponse(poolListItems)
	if err != nil {
		return nil, err
	}

	return mappedResponse, nil
}

func (p *PoolAPI) Show(opts *models.PoolShowOptions) (models.PoolShowPostResponse, error) {
	url, err := utils.CreatePoolShowUrl(p.URL, opts)

	if err != nil {
		return models.PoolShowPostResponse{}, err
	}

	body, err := utils.Fetch(url)

	if err != nil {
		return models.PoolShowPostResponse{}, err
	}

	poolShowItem := models.PoolShowPostAPIResult{}
	err = json.Unmarshal(body, &poolShowItem)

	if err != nil {
		return models.PoolShowPostResponse{}, err
	}

	mappedResponse, err := mapPoolShowAPIItemToResponse(poolShowItem)
	if err != nil {
		return models.PoolShowPostResponse{}, err
	}

	return mappedResponse, nil
}
