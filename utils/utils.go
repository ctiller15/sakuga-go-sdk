package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"
	"strconv"
	"strings"

	"github.com/ctiller15/sakuga-go-sdk/constants"
	"github.com/ctiller15/sakuga-go-sdk/models"
)

var (
	errInvalidPostOptions       = errors.New("invalid options - request limit must be between 1 and 100, page must be at least 1")
	errInvalidTagOptions        = errors.New("invalid options - requires at least one tag argument")
	errInvalidListTagOptions    = errors.New("invalid tag options")
	errInvalidTagType           = fmt.Errorf("invalid tag type - must be one of %v", constants.VALID_TAG_TYPES)
	errInvalidCommentOptions    = errors.New("invalid options - id must be set")
	errInvalidArtistOptions     = errors.New("invalid options for artist request")
	errInvalidWikiOptions       = errors.New("invalid options for wiki")
	errInvalidNoteSearchOptions = errors.New("invalid options for note search")
)

func Fetch(url string) ([]byte, error) {
	var body []byte
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)

	// Handle status codes.
	// 	Status Code 	Meaning
	// 200 OK 	Request was successful
	// 403 Forbidden 	Access denied
	// 404 Not Found 	Not found
	// 420 Invalid Record 	Record could not be saved
	// 421 User Throttled 	User is throttled, try again later
	// 422 Locked 	The resource is locked and cannot be modified
	// 423 Already Exists 	Resource already exists
	// 424 Invalid Parameters 	The given parameters were invalid
	// 500 Internal Server Error 	Some unknown error occurred on the server
	// 503 Service Unavailable 	Server cannot currently handle the request, try again later

	if err != nil {
		return nil, err
	}

	return body, nil
}

func CreatePostsListUrl(baseURL string, opts *models.PostsListOptions) (string, error) {
	// 0 is the default value. If we see that we ignore that query entirely.
	if opts.Limit != 0 && (opts.Limit > 100 || opts.Limit < 1) {
		return "", errInvalidPostOptions
	}

	if opts.Page != 0 && opts.Page < 1 {
		return "", errInvalidPostOptions
	}

	finalURL := baseURL + ".json"

	queryStringParams := make([]string, 0)

	if opts.Limit != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("limit=%d", opts.Limit))
	}

	if opts.Page != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("page=%d", opts.Page))
	}

	if len(opts.Tags) > 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("tags=%s", strings.Join(opts.Tags, " ")))
	}

	if len(queryStringParams) > 0 {
		finalURL += "?" + strings.Join(queryStringParams, "&")
	}

	return finalURL, nil
}

func CreateListTagsUrl(baseURL string, opts *models.TagListOptions) (string, error) {
	finalURL := baseURL + ".json"

	if opts.Limit != 0 && opts.Limit > 1000 {
		return "", errInvalidListTagOptions
	}

	// TODO: investigate if random is supported for tag list
	// TODO: add support for descending variations
	if opts.Order != "" && !slices.Contains(models.TagListOrderOptions, opts.Order) {
		return "", errInvalidListTagOptions
	}

	queryStringParams := make([]string, 0)

	// It may be worth building a generic builder for these params.
	if opts.Limit != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("limit=%d", opts.Limit))
	}

	if opts.Page != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("page=%d", opts.Page))
	}

	if opts.Order != "" {
		queryStringParams = append(queryStringParams, fmt.Sprintf("order=%s", opts.Order))
	}

	if opts.ID != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("id=%d", opts.ID))
	}

	if opts.AfterID != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("after_id=%d", opts.AfterID))
	}

	if opts.Name != "" {
		queryStringParams = append(queryStringParams, fmt.Sprintf("name=%s", opts.Name))
	}

	if opts.NamePattern != "" {
		queryStringParams = append(queryStringParams, fmt.Sprintf("name_pattern=%s", opts.NamePattern))
	}

	if len(queryStringParams) > 0 {
		finalURL += "?" + strings.Join(queryStringParams, "&")
	}

	return finalURL, nil
}

func CreateRelatedTagsUrl(baseURL string, opts *models.TagRelatedOptions) (string, error) {
	if len(opts.Tags) == 0 {
		return "", errInvalidTagOptions
	}

	finalURL, err := url.JoinPath(baseURL, "related.json")
	if err != nil {
		return "", err
	}

	tagsString := strings.Join(opts.Tags, " ")

	finalURL += "?tags=" + tagsString

	if opts.Type != "" {
		if !slices.Contains(constants.VALID_TAG_TYPES, opts.Type) {
			return "", errInvalidTagType
		}

		finalURL += "&type=" + opts.Type
	}

	return finalURL, nil
}

func CreateArtistsListUrl(baseURL string, opts *models.ArtistListOptions) (string, error) {
	if opts.Order != "" && !slices.Contains(models.ValidArtistListOrderOptions, opts.Order) {
		return "", errInvalidArtistOptions
	}

	finalURL := baseURL + ".json"

	queryStringParams := make([]string, 0)

	if opts.Name != "" {
		queryStringParams = append(queryStringParams, fmt.Sprintf("name=%s", opts.Name))
	}

	if opts.Page != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("page=%d", opts.Page))
	}

	if opts.Order != "" {
		queryStringParams = append(queryStringParams, fmt.Sprintf("order=%s", opts.Order))
	}

	if len(queryStringParams) > 0 {
		finalURL += "?" + strings.Join(queryStringParams, "&")
	}

	return finalURL, nil
}

func CreateCommentShowUrl(baseURL string, opts *models.CommentShowOptions) (string, error) {
	if opts.ID == 0 {
		return "", errInvalidCommentOptions
	}

	finalURL, err := url.JoinPath(baseURL, "show.json")

	if err != nil {
		return "", err
	}

	finalURL += "?id=" + strconv.Itoa(opts.ID)

	return finalURL, nil
}

func CreateWikiListUrl(baseURL string, opts *models.WikiListOptions) (string, error) {
	finalURL := baseURL + ".json"

	if opts.Order != "" && !slices.Contains(models.WikiListOrderOptions, opts.Order) {
		return "", errInvalidWikiOptions
	}

	queryStringParams := make([]string, 0)

	if opts.Order != "" {
		queryStringParams = append(queryStringParams, fmt.Sprintf("order=%s", opts.Order))
	}

	if opts.Limit != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("limit=%d", opts.Limit))
	}

	if opts.Page != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("page=%d", opts.Page))
	}

	if opts.Query != "" {
		queryStringParams = append(queryStringParams, fmt.Sprintf("query=%s", opts.Query))

	}

	if len(queryStringParams) > 0 {
		finalURL += "?" + strings.Join(queryStringParams, "&")
	}

	return finalURL, nil
}

func CreateNoteListUrl(baseURL string, opts *models.NoteListOptions) (string, error) {
	finalURL := baseURL + ".json"

	if opts.PostID != 0 {
		finalURL += fmt.Sprintf("?post_id=%d", opts.PostID)
	}
	return finalURL, nil
}

func CreateNoteSearchUrl(baseURL string, opts *models.NoteSearchOptions) (string, error) {
	finalURL, err := url.JoinPath(baseURL, "search.json")
	if err != nil {
		return "", err
	}

	if opts.Query == "" {
		return "", errInvalidNoteSearchOptions
	}

	finalURL += "?query=" + opts.Query

	return finalURL, nil
}

func CreateNoteHistoryUrl(baseURL string, opts *models.NoteHistorySearchOptions) (string, error) {
	finalURL, err := url.JoinPath(baseURL, "history.json")
	if err != nil {
		return "", err
	}

	queryStringParams := make([]string, 0)

	if opts.Limit != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("limit=%d", opts.Limit))
	}

	if opts.Page != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("page=%d", opts.Page))
	}

	if opts.PostID != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("post_id=%d", opts.PostID))
	}

	if opts.ID != 0 {
		queryStringParams = append(queryStringParams, fmt.Sprintf("id=%d", opts.ID))
	}

	if len(queryStringParams) > 0 {
		finalURL += "?" + strings.Join(queryStringParams, "&")
	}

	return finalURL, nil
}
