package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
)

type NotesAPI struct {
	URL string
}

func newNotesAPI(baseURL string) *NotesAPI {
	newAPI := NotesAPI{
		URL: baseURL + "/note",
	}
	return &newAPI
}

func mapNotesListAPIItemsToResponse(items []sakugamodels.NoteListAPIResultItem) ([]sakugamodels.NoteListResponseItem, error) {
	response := make([]sakugamodels.NoteListResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		mappedResponseItem := sakugamodels.NoteListResponseItem{
			ID:        item.ID,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			CreatorID: item.CreatorID,
			X:         item.X,
			Y:         item.Y,
			Width:     item.Width,
			Height:    item.Height,
			IsActive:  item.IsActive,
			PostID:    item.PostID,
			Body:      item.Body,
			Version:   item.Version,
		}

		response = append(response, mappedResponseItem)
	}

	return response, nil
}

func mapNotesHistoryAPIItemsToResponse(items []sakugamodels.NoteHistoryAPIResultItem) ([]sakugamodels.NoteHistoryResponseItem, error) {
	response := make([]sakugamodels.NoteHistoryResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		mappedResponseItem := sakugamodels.NoteHistoryResponseItem{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			CreatorID: item.CreatorID,
			X:         item.X,
			Y:         item.Y,
			Width:     item.Width,
			Height:    item.Height,
			IsActive:  item.IsActive,
			PostID:    item.PostID,
			Body:      item.Body,
			Version:   item.Version,
		}

		response = append(response, mappedResponseItem)
	}

	return response, nil
}

func (n *NotesAPI) List(opts *sakugamodels.NoteListOptions) ([]sakugamodels.NoteListResponseItem, error) {
	url, err := sakugautils.CreateNoteListUrl(n.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := sakugautils.Fetch(url)
	if err != nil {
		return nil, err
	}

	notesListItems := make([]sakugamodels.NoteListAPIResultItem, 0)
	err = json.Unmarshal(body, &notesListItems)

	if err != nil {
		return nil, err
	}

	mappedItems, err := mapNotesListAPIItemsToResponse(notesListItems)
	if err != nil {
		return nil, err
	}

	return mappedItems, nil
}

// Allow to fuzzy search on notes
func (n *NotesAPI) Search(opts *sakugamodels.NoteSearchOptions) ([]sakugamodels.NoteListResponseItem, error) {
	url, err := sakugautils.CreateNoteSearchUrl(n.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := sakugautils.Fetch(url)
	if err != nil {
		return nil, err
	}

	notesListItems := make([]sakugamodels.NoteListAPIResultItem, 0)
	err = json.Unmarshal(body, &notesListItems)

	if err != nil {
		return nil, err
	}

	mappedItems, err := mapNotesListAPIItemsToResponse(notesListItems)
	if err != nil {
		return nil, err
	}

	return mappedItems, nil
}

func (n *NotesAPI) History(opts *sakugamodels.NoteHistorySearchOptions) ([]sakugamodels.NoteHistoryResponseItem, error) {
	url, err := sakugautils.CreateNoteHistoryUrl(n.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := sakugautils.Fetch(url)
	if err != nil {
		return nil, err
	}

	notesListItems := make([]sakugamodels.NoteHistoryAPIResultItem, 0)
	err = json.Unmarshal(body, &notesListItems)

	if err != nil {
		return nil, err
	}

	mappedItems, err := mapNotesHistoryAPIItemsToResponse(notesListItems)
	if err != nil {
		return nil, err
	}

	return mappedItems, nil
}
