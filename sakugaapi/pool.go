package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
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

func mapPoolListAPIItemsToResponse(items []sakugamodels.PoolListAPIResultItem) ([]sakugamodels.PoolListResponseItem, error) {
	response := make([]sakugamodels.PoolListResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
		if err != nil {
			return nil, err
		}
		mappedResponseItem := sakugamodels.PoolListResponseItem{
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

func mapPoolShowAPIItemsToResponse(items []sakugamodels.PoolShowPostAPIResultItem) ([]sakugamodels.PoolShowPostResponseItem, error) {
	response := make([]sakugamodels.PoolShowPostResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		newItem := sakugamodels.PoolShowPostResponseItem{
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

func mapPoolShowAPIItemToResponse(item sakugamodels.PoolShowPostAPIResult) (sakugamodels.PoolShowPostResponse, error) {
	createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
	if err != nil {
		return sakugamodels.PoolShowPostResponse{}, err
	}

	updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
	if err != nil {
		return sakugamodels.PoolShowPostResponse{}, err
	}

	posts, err := mapPoolShowAPIItemsToResponse(item.Posts)

	response := sakugamodels.PoolShowPostResponse{
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

func (p *PoolAPI) List(opts *sakugamodels.PoolListOptions) ([]sakugamodels.PoolListResponseItem, error) {
	url, err := sakugautils.CreatePoolListUrl(p.URL, opts)

	if err != nil {
		return nil, err
	}

	body, err := sakugautils.Fetch(url)

	if err != nil {
		return nil, err
	}

	poolListItems := make([]sakugamodels.PoolListAPIResultItem, 0)
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

func (p *PoolAPI) Show(opts *sakugamodels.PoolShowOptions) (sakugamodels.PoolShowPostResponse, error) {
	url, err := sakugautils.CreatePoolShowUrl(p.URL, opts)

	if err != nil {
		return sakugamodels.PoolShowPostResponse{}, err
	}

	body, err := sakugautils.Fetch(url)

	if err != nil {
		return sakugamodels.PoolShowPostResponse{}, err
	}

	poolShowItem := sakugamodels.PoolShowPostAPIResult{}
	err = json.Unmarshal(body, &poolShowItem)

	if err != nil {
		return sakugamodels.PoolShowPostResponse{}, err
	}

	mappedResponse, err := mapPoolShowAPIItemToResponse(poolShowItem)
	if err != nil {
		return sakugamodels.PoolShowPostResponse{}, err
	}

	return mappedResponse, nil
}
