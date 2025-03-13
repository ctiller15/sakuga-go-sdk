package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
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

func (c *CommentsAPI) Show(opts *sakugamodels.CommentShowOptions) (sakugamodels.CommentShowResponseItem, error) {
	url, err := sakugautils.CreateCommentShowUrl(c.URL, opts)
	if err != nil {
		return sakugamodels.CommentShowResponseItem{}, err
	}

	body, err := sakugautils.Fetch(url)

	if err != nil {
		return sakugamodels.CommentShowResponseItem{}, err
	}

	commentShowItem := sakugamodels.CommentShowAPIResultItem{}
	err = json.Unmarshal(body, &commentShowItem)

	if err != nil {
		return sakugamodels.CommentShowResponseItem{}, err
	}

	commentShowResponse, err := mapCommentShowAPIItemToResponse(commentShowItem)
	if err != nil {
		return sakugamodels.CommentShowResponseItem{}, err
	}

	return commentShowResponse, nil
}

func mapCommentShowAPIItemToResponse(item sakugamodels.CommentShowAPIResultItem) (sakugamodels.CommentShowResponseItem, error) {
	// Apparently the format is ISO 8601, but this is working for now so I'm happy.
	//https://stackoverflow.com/questions/522251/whats-the-difference-between-iso-8601-and-rfc-3339-date-formats
	parsedTime, err := time.Parse(time.RFC3339, item.CreatedAt)
	if err != nil {
		return sakugamodels.CommentShowResponseItem{}, err
	}

	return sakugamodels.CommentShowResponseItem{
		ID:        item.ID,
		CreatedAt: parsedTime,
		PostID:    item.PostID,
		Creator:   item.Creator,
		CreatorID: item.CreatorID,
		Body:      item.Body,
	}, nil
}
