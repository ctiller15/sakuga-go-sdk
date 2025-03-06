package utils

import (
	"testing"

	"github.com/ctiller15/sakuga-go-sdk/sakugaapi"
	"github.com/stretchr/testify/assert"
)

func TestRelatedTagsURLCreation(t *testing.T) {
	t.Run("errors if no tag has been submitted", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := sakugaapi.TagRelatedOptions{}

		_, err := createRelatedTagsUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidTagOptions)
	})

	t.Run("errors if an invalid tag type has been provided", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := sakugaapi.TagRelatedOptions{
			Tags: []string{"animated"},
			Type: "invalid",
		}

		_, err := createRelatedTagsUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidTagType)
	})

	t.Run("creates a url with related tags", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := sakugaapi.TagRelatedOptions{
			Tags: []string{"animated"},
		}

		result, err := createRelatedTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		expected := "https://sakugabooru.com/tag/related.json?tags=animated"
		assert.Equal(t, expected, result)
	})

	t.Run("creates a url with multiple tags", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := sakugaapi.TagRelatedOptions{
			Tags: []string{"animated", "fighting"},
		}

		result, err := createRelatedTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		expected := "https://sakugabooru.com/tag/related.json?tags=animated fighting"
		assert.Equal(t, expected, result)
	})

	t.Run("creates a url with multiple tags and a type", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := sakugaapi.TagRelatedOptions{
			Tags: []string{"animated", "fighting"},
			Type: sakugaapi.TagTypeArtist,
		}

		result, err := createRelatedTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		expected := "https://sakugabooru.com/tag/related.json?tags=animated fighting&type=artist"
		assert.Equal(t, expected, result)
	})
}
