package sakugaapi

import (
	"encoding/json"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
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

func mapNotesListAPIItemsToResponse(items []models.NoteListAPIResultItem) ([]models.NoteListResponseItem, error) {
	response := make([]models.NoteListResponseItem, 0)

	for _, item := range items {
		createdAt, err := time.Parse(time.RFC3339, item.CreatedAt)
		if err != nil {
			return nil, err
		}

		updatedAt, err := time.Parse(time.RFC3339, item.UpdatedAt)
		if err != nil {
			return nil, err
		}

		mappedResponseItem := models.NoteListResponseItem{
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

func (n *NotesAPI) List(opts *models.NoteListOptions) ([]models.NoteListResponseItem, error) {
	url, err := utils.CreateNoteListUrl(n.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := utils.Fetch(url)

	notesListItems := make([]models.NoteListAPIResultItem, 0)
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
func (n *NotesAPI) Search(opts *models.NoteSearchOptions) ([]models.NoteListResponseItem, error) {
	url, err := utils.CreateNoteSearchUrl(n.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := utils.Fetch(url)

	notesListItems := make([]models.NoteListAPIResultItem, 0)
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
