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
	errInvalidTagOptions     = errors.New("invalid options - requires at least one tag argument")
	errInvalidTagType        = fmt.Errorf("invalid tag type - must be one of %v", constants.VALID_TAG_TYPES)
	errInvalidCommentOptions = errors.New("invalid options - id must be set")
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
