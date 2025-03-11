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

func TestTagsListURLCreation(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json")
	})

	t.Run("errors if limit is above 1000", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			Limit: 50000,
		}

		_, err := CreateListTagsUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidListTagOptions)
	})

	t.Run("errors if order is not date, count, or name", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			Order: "invalid",
		}

		_, err := CreateListTagsUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidListTagOptions)
	})

	t.Run("limit", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			Limit: 5,
		}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json?limit=5")
	})

	t.Run("page", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			Page: 5,
		}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json?page=5")
	})

	t.Run("order", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			Order: models.TagListOrderOptionDate,
		}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json?order=date")
	})

	t.Run("id", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			ID: 25325,
		}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json?id=25325")
	})

	t.Run("after_id", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			AfterID: 25000,
		}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json?after_id=25000")
	})

	t.Run("name", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			Name: "effects",
		}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json?name=effects")
	})

	t.Run("name_pattern", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			NamePattern: "abcd",
		}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json?name_pattern=abcd")
	})

	t.Run("combination of all options", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagListOptions{
			Page:        1,
			Limit:       5,
			Order:       models.TagListOrderOptionDate,
			ID:          26964,
			AfterID:     25000,
			Name:        "miyagawa",
			NamePattern: "o_m",
		}

		response, err := CreateListTagsUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, response, "https://sakugabooru.com/tag.json?limit=5&page=1&order=date&id=26964&after_id=25000&name=miyagawa&name_pattern=o_m")
	})
}

func TestRelatedTagsURLCreation(t *testing.T) {
	t.Run("errors if no tag has been submitted", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
		options := models.TagRelatedOptions{}

		_, err := CreateRelatedTagsUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidTagOptions)
	})

	t.Run("errors if an invalid tag type has been provided", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/tag"
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
func TestArtistsListURLCreation(t *testing.T) {
	t.Run("base case, generate url without options", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/artist"
		options := models.ArtistListOptions{}

		result, err := CreateArtistsListUrl(baseURL, &options)
		assert.Nil(t, err)
		expected := "https://sakugabooru.com/artist.json"
		assert.Equal(t, expected, result)
	})

	t.Run("errors if order is not date or name", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/artist"
		options := models.ArtistListOptions{
			Order: "invalid",
		}

		_, err := CreateArtistsListUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidArtistOptions)
	})

	t.Run("name", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/artist"
		options := models.ArtistListOptions{
			Name: "oda_eiichiro",
		}

		result, err := CreateArtistsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/artist.json?name=oda_eiichiro")
	})

	t.Run("order", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/artist"
		options := models.ArtistListOptions{
			Order: models.ArtistListOrderNameOption,
		}

		result, err := CreateArtistsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/artist.json?order=name")
	})

	t.Run("page", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/artist"
		options := models.ArtistListOptions{
			Page: 5,
		}

		result, err := CreateArtistsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/artist.json?page=5")

	})

	t.Run("name, order, and page", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/artist"
		options := models.ArtistListOptions{
			Name:  "a",
			Order: models.ArtistListOrderNameOption,
			Page:  2,
		}

		result, err := CreateArtistsListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/artist.json?name=a&page=2&order=name")
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

func TestWikiListURLCreation(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/wiki"
		options := models.WikiListOptions{}

		result, err := CreateWikiListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/wiki.json")
	})

	t.Run("order limited to title and date", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/wiki"
		options := models.WikiListOptions{
			Order: "invalid",
		}

		_, err := CreateWikiListUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidWikiOptions)
	})

	t.Run("order", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/wiki"
		options := models.WikiListOptions{
			Order: models.WikiListOrderOptionDate,
		}

		result, err := CreateWikiListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/wiki.json?order=date")
	})

	t.Run("limit", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/wiki"
		options := models.WikiListOptions{
			Limit: 50,
		}

		result, err := CreateWikiListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/wiki.json?limit=50")
	})

	t.Run("page", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/wiki"
		options := models.WikiListOptions{
			Page: 5,
		}

		result, err := CreateWikiListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/wiki.json?page=5")
	})

	t.Run("query", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/wiki"
		options := models.WikiListOptions{
			Query: "Background",
		}

		result, err := CreateWikiListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/wiki.json?query=Background")
	})

	t.Run("all parameters", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/wiki"
		options := models.WikiListOptions{
			Order: models.WikiListOrderOptionDate,
			Limit: 2,
			Page:  2,
			Query: "Background",
		}

		result, err := CreateWikiListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/wiki.json?order=date&limit=2&page=2&query=Background")
	})
}

func TestNotesListURLCreation(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteListOptions{}

		result, err := CreateNoteListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note.json")
	})

	t.Run("postID", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteListOptions{
			PostID: 251595,
		}

		result, err := CreateNoteListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note.json?post_id=251595")
	})
}

func TestNotesSearchURLCreation(t *testing.T) {
	t.Run("errors if no query is provided", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteSearchOptions{}

		_, err := CreateNoteSearchUrl(baseURL, &options)
		assert.NotNil(t, err)
		assert.ErrorIs(t, err, errInvalidNoteSearchOptions)
	})

	t.Run("successful case", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteSearchOptions{
			Query: "fight",
		}

		result, err := CreateNoteSearchUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note/search.json?query=fight")
	})
}

func TestNotesHistoryURLCreation(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteHistorySearchOptions{}

		result, err := CreateNoteHistoryUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note/history.json")
	})

	t.Run("limit", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteHistorySearchOptions{
			Limit: 5,
		}

		result, err := CreateNoteHistoryUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note/history.json?limit=5")
	})

	t.Run("page", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteHistorySearchOptions{
			Page: 1,
		}

		result, err := CreateNoteHistoryUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note/history.json?page=1")
	})

	t.Run("post_id", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteHistorySearchOptions{
			PostID: 1273,
		}

		result, err := CreateNoteHistoryUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note/history.json?post_id=1273")
	})

	t.Run("id", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteHistorySearchOptions{
			ID: 70,
		}

		result, err := CreateNoteHistoryUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note/history.json?id=70")
	})

	t.Run("all query params", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/note"
		options := models.NoteHistorySearchOptions{
			Limit:  5,
			ID:     4,
			Page:   1,
			PostID: 1273,
		}

		result, err := CreateNoteHistoryUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/note/history.json?limit=5&page=1&post_id=1273&id=4")
	})
}

func TestUserSearchURLCreation(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/user"
		options := models.UserSearchOptions{}

		result, err := CreateUserSearchUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/user.json")
	})

	t.Run("id", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/user"
		options := models.UserSearchOptions{
			ID: 10,
		}

		result, err := CreateUserSearchUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/user.json?id=10")
	})

	t.Run("name", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/user"
		options := models.UserSearchOptions{
			Name: "s",
		}

		result, err := CreateUserSearchUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/user.json?name=s")
	})
}

func TestForumListURLCreation(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/forum"
		options := models.ForumListOptions{}

		result, err := CreateForumListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/forum.json")
	})

	t.Run("parent_id", func(t *testing.T) {
		baseURL := "https://sakugabooru.com/forum"
		options := models.ForumListOptions{
			ParentID: 50,
		}

		result, err := CreateForumListUrl(baseURL, &options)
		assert.Nil(t, err)
		assert.Equal(t, result, "https://sakugabooru.com/forum.json?parent_id=50")
	})
}
