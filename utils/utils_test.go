package utils

import (
	"testing"

	"github.com/ctiller15/sakuga-go-sdk/constants"
	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/stretchr/testify/assert"
)

func TestPostListURLCreation(t *testing.T) {
	baseURL := "https://sakugabooru.com/post"

	t.Run("empty case", func(t *testing.T) {
		options := models.PostsListOptions{}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)

		assert.Equal(t, result, "https://sakugabooru.com/post.json")
	})

	t.Run("errors if limit is higher than 100", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: 500,
		}

		_, err := CreatePostsListUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidPostOptions)
	})

	t.Run("errors if limit is lower than 1", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: -10,
		}

		_, err := CreatePostsListUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidPostOptions)
	})

	t.Run("successful query - limit only", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: 50,
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)

		assert.Equal(t, result, "https://sakugabooru.com/post.json?limit=50")

	})

	t.Run("errors if page is lower than 1", func(t *testing.T) {
		options := models.PostsListOptions{
			Page: -5,
		}

		_, err := CreatePostsListUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidPostOptions)
	})

	t.Run("successful query - page only", func(t *testing.T) {
		options := models.PostsListOptions{
			Page: 1,
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?page=1")
	})

	t.Run("successful query - single tag only", func(t *testing.T) {
		options := models.PostsListOptions{
			Tags: []string{"effects"},
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?tags=effects")
	})

	t.Run("successful query - multiple tags", func(t *testing.T) {
		options := models.PostsListOptions{
			Tags: []string{"effects", "fighting"},
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?tags=effects fighting")
	})

	t.Run("successful query, limit and page", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: 10,
			Page:  5,
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?limit=10&page=5")
	})

	t.Run("successful query, limit and tags - single", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: 10,
			Tags:  []string{"smoke"},
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?limit=10&tags=smoke")
	})

	t.Run("successful query, limit and tags - multiple", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: 10,
			Tags:  []string{"smoke", "effects"},
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?limit=10&tags=smoke effects")
	})

	t.Run("page and tags - single", func(t *testing.T) {
		options := models.PostsListOptions{
			Page: 10,
			Tags: []string{"smoke"},
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?page=10&tags=smoke")
	})

	t.Run("page and tags - multiple", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: 10,
			Tags:  []string{"smoke", "effects"},
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?limit=10&tags=smoke effects")
	})

	t.Run("limit, page, and tags - single", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: 10,
			Page:  5,
			Tags:  []string{"smoke"},
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?limit=10&page=5&tags=smoke")
	})

	t.Run("limit, page, and tags - multiple", func(t *testing.T) {
		options := models.PostsListOptions{
			Limit: 10,
			Page:  5,
			Tags:  []string{"smoke", "effects"},
		}

		result, err := CreatePostsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/post.json?limit=10&page=5&tags=smoke effects")
	})
}

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
