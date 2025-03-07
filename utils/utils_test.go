package utils

import (
	"testing"

	"github.com/ctiller15/sakuga-go-sdk/constants"
	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestRelatedTagsURLCreation(t *testing.T) {
	t.Run("errors if no tag has been submitted", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := models.TagRelatedOptions{}

		_, err := CreateRelatedTagsUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidTagOptions)
	})

	t.Run("errors if an invalid tag type has been provided", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := models.TagRelatedOptions{
			Tags: []string{"animated"},
			Type: "invalid",
		}

		_, err := CreateRelatedTagsUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidTagType)
	})

	t.Run("creates a url with related tags", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := models.TagRelatedOptions{
			Tags: []string{"animated"},
		}

		result, err := CreateRelatedTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		expected := "https://sakugabooru.com/tag/related.json?tags=animated"
		assert.Equal(t, expected, result)
	})

	t.Run("creates a url with multiple tags", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := models.TagRelatedOptions{
			Tags: []string{"animated", "fighting"},
		}

		result, err := CreateRelatedTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		expected := "https://sakugabooru.com/tag/related.json?tags=animated fighting"
		assert.Equal(t, expected, result)
	})

	t.Run("creates a url with multiple tags and a type", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag/"
		options := models.TagRelatedOptions{
			Tags: []string{"animated", "fighting"},
			Type: constants.TagTypeArtist,
		}

		result, err := CreateRelatedTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		expected := "https://sakugabooru.com/tag/related.json?tags=animated fighting&type=artist"
		assert.Equal(t, expected, result)
	})
}

func TestCommentShowURLCreation(t *testing.T) {
	t.Run("Fails to create url if ID is not provided", func(t *testing.T) {
		baseURL := "https://sakurabooru.com/comment/"

		options := models.CommentShowOptions{}

		_, err := CreateCommentShowUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidCommentOptions)
	})

	t.Run("Successfully creates url if ID is provided", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/comment/"

		options := models.CommentShowOptions{
			ID: 90003,
		}

		result, err := CreateCommentShowUrl(baseURL, &options)
		assert.Nil(t, err)
		expected := "https://sakugabooru.com/comment/show.json?id=90003"
		assert.Equal(t, expected, result)
	})
}
