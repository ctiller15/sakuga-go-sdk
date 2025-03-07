package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
)

type CommentsAPI struct {
	URL string
}

func newCommentsAPI(baseURL string) *CommentsAPI {
	newAPI := CommentsAPI{
		URL: baseURL + "/comment",
	}
	return &newAPI
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
