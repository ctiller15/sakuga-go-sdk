package utils

import (
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/ctiller15/sakuga-go-sdk/constants"
	"github.com/ctiller15/sakuga-go-sdk/models"
)

var (
	errInvalidTagOptions = errors.New("invalid options - requires at least one tag argument")
	errInvalidTagType    = fmt.Errorf("invalid tag type - must be one of %v", constants.VALID_TAG_TYPES)
)

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
