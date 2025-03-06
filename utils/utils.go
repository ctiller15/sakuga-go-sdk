package utils

import (
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strings"

	"github.com/ctiller15/sakuga-go-sdk/sakugaapi"
)

var (
	errInvalidTagOptions = errors.New("invalid options - requires at least one tag argument")
	errInvalidTagType    = fmt.Errorf("invalid tag type - must be one of %v", sakugaapi.VALID_TAG_TYPES)
)

func createRelatedTagsUrl(baseURL string, opts *sakugaapi.TagRelatedOptions) (string, error) {
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
		if !slices.Contains(sakugaapi.VALID_TAG_TYPES, opts.Type) {
			return "", errInvalidTagType
		}

		finalURL += "&type=" + opts.Type
	}

	return finalURL, nil
}
