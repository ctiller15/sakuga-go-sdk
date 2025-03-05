package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAPIInitialization(t *testing.T) {
	newAPI := NewAPI()

	assert.NotNil(t, newAPI)
	assert.NotNil(t, newAPI.Posts)
}

func TestApiPostList(t *testing.T) {
	t.Run("Base case happy path", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create mock response
			if r.URL.Path == "/post.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(postDataResponse)
			}
		}))
		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)
		opts := PostsListOptions{}
		response, err := newAPI.Posts.List(&opts)
		assert.Nil(t, err)
		expected := []PostListResponseResult{
			{
				ID:                  216284,
				Tags:                "artist_unknown correction kubo-san_wa_mob_wo_yurusanai production_materials",
				CreatedAt:           time.Date(1970, time.January, 20, 4, 0, 42, 317000000, time.Local),
				UpdatedAt:           time.Date(1970, time.January, 20, 4, 0, 50, 468000000, time.Local),
				CreatorID:           20764,
				ApproverID:          4456,
				Author:              "JDMManga",
				Change:              1065989,
				Source:              "https://twitter.com/PINEJAM_info/status/1615627553791672321?s=20&t=rThR9qaXQaB6JtaaPczLoA",
				Score:               0,
				MD5:                 "94818e87adafcd4a8b857eca97417b69",
				FileSize:            2607352,
				FileExtension:       "",
				FileURL:             "https://www.sakugabooru.com/data/94818e87adafcd4a8b857eca97417b69.png",
				IsShownInIndex:      true,
				PreviewURL:          "https://www.sakugabooru.com/data/preview/94818e87adafcd4a8b857eca97417b69.jpg",
				PreviewWidth:        150,
				ActualPreviewWidth:  300,
				ActualPreviewHeight: 212,
				SampleURL:           "https://www.sakugabooru.com/data/sample/94818e87adafcd4a8b857eca97417b69.jpg",
				SampleWidth:         1500,
				SampleHeight:        1061,
				SampleFileSize:      237252,
				JpegURL:             "https://www.sakugabooru.com/data/94818e87adafcd4a8b857eca97417b69.png",
				JpegWidth:           7484,
				JpegHeight:          5292,
				JpegFileSize:        0,
				Rating:              "s",
				IsRatingLocked:      false,
				HasChildren:         false,
				ParentID:            0,
				Status:              "active",
				IsPending:           false,
				Width:               7484,
				Height:              5292,
				IsHeld:              false,
				FramesPendingString: "",
				FramesPending:       []interface{}{},
				FramesString:        "",
				Frames:              []interface{}{},
				IsNoteLocked:        false,
				LastNotedAt:         time.Date(1969, time.December, 31, 19, 0, 0, 0, time.Local),
				LastCommentedAt:     time.Date(1969, time.December, 31, 19, 0, 0, 0, time.Local),
			},
		}
		assert.Equal(t, expected, response)
	})
}
