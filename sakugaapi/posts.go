package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
)

// Struct that holds all the methods for posts
type PostsAPI struct {
	URL string
}

func newPostsAPI(baseURL string) *PostsAPI {
	newAPI := PostsAPI{
		URL: baseURL + "/post",
	}
	return &newAPI
}

// Can fetch by ID with the following: https://www.sakugabooru.com/post.json?tags=id:277788

// Fetches posts from sakugabooru
func (p *PostsAPI) List(opts *sakugamodels.PostsListOptions) ([]sakugamodels.PostListResponseResult, error) {
	url, err := sakugautils.CreatePostsListUrl(p.URL, opts)
	if err != nil {
		return nil, err
	}
	body, err := sakugautils.Fetch(url)
	if err != nil {
		return nil, err
	}

	postListItems := make([]sakugamodels.PostListAPIResponseItem, 0)
	err = json.Unmarshal(body, &postListItems)

	if err != nil {
		return nil, err
	}

	return mapPostListAPIItemsToResponse(postListItems), nil
	// work out options later.
}

func mapPostListAPIItemsToResponse(listItems []sakugamodels.PostListAPIResponseItem) []sakugamodels.PostListResponseResult {
	postListResponseResults := make([]sakugamodels.PostListResponseResult, 0)

	for _, item := range listItems {
		newResponseResult := sakugamodels.PostListResponseResult{
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
