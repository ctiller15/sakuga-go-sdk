package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/stretchr/testify/assert"
)

func TestPoolList(t *testing.T) {
	t.Run("base case, happy path", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/pool.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(poolDataListResponse)
			}
		}))
		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.PoolListOptions{}
		response, err := newAPI.Pools.List(&opts)
		assert.Nil(t, err)
		expected := []sakugamodels.PoolListResponseItem{
			{
				ID: 257, Name: "PPURI", CreatedAt: time.Date(2024, time.October, 19, 17, 58, 29, 325000000, time.UTC), UpdatedAt: time.Date(2025, time.January, 5, 1, 31, 19, 717000000, time.UTC), UserID: 1028, IsPublic: true, PostCount: 300, Description: "All the animation works from the Korean studio."},
			{
				ID: 249, Name: "Stair-Climbing", CreatedAt: time.Date(2024, time.June, 14, 17, 48, 9, 391000000, time.UTC), UpdatedAt: time.Date(2025, time.February, 1, 1, 40, 41, 690000000, time.UTC), UserID: 13069, IsPublic: true, PostCount: 37, Description: "Specialized *character acting* featuring stair-climbing, is generally considered to be one of the most difficult types of character acting to draw well. Preferably contains challenging/interesting layouts and movements (of the character acting) that capitalizes on the stair setting. Would prefer to avoid sequences where the movement is handled similarly to a standard run cycle."},
			{
				ID: 246, Name: "Botanical_Sakuga", CreatedAt: time.Date(2024, time.April, 5, 16, 50, 37, 446000000, time.UTC), UpdatedAt: time.Date(2025, time.February, 20, 15, 7, 19, 790000000, time.UTC), UserID: 18138, IsPublic: true, PostCount: 111, Description: "Well animated scenes of plants!"},
			{
				ID: 237, Name: "Kyoani_playing_cards", CreatedAt: time.Date(2024, time.March, 7, 16, 29, 11, 632000000, time.UTC), UpdatedAt: time.Date(2025, time.January, 23, 20, 53, 32, 593000000, time.UTC), UserID: 21916, IsPublic: true, PostCount: 31, Description: "パラパラトランプ"},
			{
				ID: 232, Name: "Sign_Language", CreatedAt: time.Date(2024, time.January, 2, 11, 50, 25, 13000000, time.UTC), UpdatedAt: time.Date(2024, time.November, 19, 22, 51, 11, 781000000, time.UTC), UserID: 20764, IsPublic: true, PostCount: 47, Description: ""}}

		assert.Equal(t, expected, response)
	})
}

func TestPoolPostList(t *testing.T) {
	t.Run("base case, happy path", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.RequestURI() == "/pool/show.json?id=215" {
				w.WriteHeader(http.StatusOK)
				w.Write(poolShowDataResponse)
			}
		}))
		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.PoolShowOptions{
			ID: 215,
		}
		response, err := newAPI.Pools.Show(&opts)
		assert.Nil(t, err)
		expected := sakugamodels.PoolShowPostResponse{
			ID:          215,
			Name:        "Stick_Figure_Animations",
			CreatedAt:   time.Date(2023, time.June, 21, 19, 45, 51, 333000000, time.UTC),
			UpdatedAt:   time.Date(2024, time.August, 8, 21, 7, 28, 95000000, time.UTC),
			UserID:      18138,
			IsPublic:    false,
			PostCount:   20,
			Description: "Animations containing Stick Figures. Whether it'd fighting, dancing or mucking about, in general.",
			Posts: []sakugamodels.PoolShowPostResponseItem{
				{
					ID:                  22733,
					Tags:                "animated fighting philips_lacanlale smears web western",
					CreatedAt:           time.Date(2016, time.May, 6, 1, 23, 30, 476000000, time.UTC),
					CreatorID:           995,
					Author:              "kintuin",
					Change:              1156350,
					Source:              "",
					Score:               124,
					MD5:                 "92da968affd0ab4a7a31e4fb7e2755f6",
					FileSize:            884515,
					FileURL:             "https://www.sakugabooru.com/data/92da968affd0ab4a7a31e4fb7e2755f6.gif",
					IsShownInIndex:      true,
					PreviewURL:          "https://www.sakugabooru.com/data/preview/92da968affd0ab4a7a31e4fb7e2755f6.jpg",
					PreviewWidth:        150,
					PreviewHeight:       105,
					ActualPreviewWidth:  300,
					ActualPreviewHeight: 210,
					SampleURL:           "https://www.sakugabooru.com/data/92da968affd0ab4a7a31e4fb7e2755f6.gif",
					SampleWidth:         500,
					SampleHeight:        350,
					SampleFileSize:      0,
					JpegURL:             "https://www.sakugabooru.com/data/92da968affd0ab4a7a31e4fb7e2755f6.gif",
					JpegWidth:           500,
					JpegHeight:          350,
					JpegFileSize:        0,
					Rating:              "s",
					HasChildren:         false,
					ParentID:            0,
					Status:              "active",
					Width:               500,
					Height:              350,
					IsHeld:              false,
					FramesPendingString: "",
					FramesPending:       []interface{}{},
					FramesString:        "",
					Frames:              []interface{}{},
				}, {
					ID:                  29421,
					Tags:                "animated beams debris effects fighting smears smoke tommy web",
					CreatedAt:           time.Date(2017, time.January, 5, 10, 11, 41, 217000000, time.UTC),
					CreatorID:           508,
					Author:              "Ashita",
					Change:              1161941,
					Source:              "",
					Score:               42,
					MD5:                 "7c4ea5f75f6901e2ec5bb4c8289a9e93",
					FileSize:            1179880,
					FileURL:             "https://www.sakugabooru.com/data/7c4ea5f75f6901e2ec5bb4c8289a9e93.mp4",
					IsShownInIndex:      true,
					PreviewURL:          "https://www.sakugabooru.com/data/preview/7c4ea5f75f6901e2ec5bb4c8289a9e93.jpg",
					PreviewWidth:        150,
					PreviewHeight:       113,
					ActualPreviewWidth:  256,
					ActualPreviewHeight: 192,
					SampleURL:           "https://www.sakugabooru.com/data/7c4ea5f75f6901e2ec5bb4c8289a9e93.mp4",
					SampleWidth:         256,
					SampleHeight:        192,
					SampleFileSize:      0,
					JpegURL:             "https://www.sakugabooru.com/data/7c4ea5f75f6901e2ec5bb4c8289a9e93.mp4",
					JpegWidth:           256,
					JpegHeight:          192,
					JpegFileSize:        0,
					Rating:              "s",
					HasChildren:         false,
					ParentID:            0,
					Status:              "active",
					Width:               256,
					Height:              192,
					IsHeld:              false,
					FramesPendingString: "",
					FramesPending:       []interface{}{},
					FramesString:        "",
					Frames:              []interface{}{},
				}, {
					ID:                  135811,
					Tags:                "animated animator_vs._animation_8:_the_showdown animator_vs._animation_series artist_unknown beams effects explosions fighting fire flying guzzu impact_frames liquid shadowqrow smears smoke web western",
					CreatedAt:           time.Date(2020, time.October, 26, 2, 32, 47, 778000000, time.UTC),
					CreatorID:           5084,
					Author:              "Armando",
					Change:              1323223,
					Source:              "https://www.youtube.com/watch?v=0a1r0JaONS4",
					Score:               115,
					MD5:                 "277c0d99ef7aa2c9a5c443aa7580f0d0",
					FileSize:            20868373,
					FileURL:             "https://www.sakugabooru.com/data/277c0d99ef7aa2c9a5c443aa7580f0d0.mp4",
					IsShownInIndex:      true,
					PreviewURL:          "https://www.sakugabooru.com/data/preview/277c0d99ef7aa2c9a5c443aa7580f0d0.jpg",
					PreviewWidth:        150,
					PreviewHeight:       84,
					ActualPreviewWidth:  300,
					ActualPreviewHeight: 169,
					SampleURL:           "https://www.sakugabooru.com/data/277c0d99ef7aa2c9a5c443aa7580f0d0.mp4",
					SampleWidth:         854,
					SampleHeight:        480,
					SampleFileSize:      0,
					JpegURL:             "https://www.sakugabooru.com/data/277c0d99ef7aa2c9a5c443aa7580f0d0.mp4",
					JpegWidth:           854,
					JpegHeight:          480,
					JpegFileSize:        0,
					Rating:              "s",
					HasChildren:         false,
					ParentID:            0,
					Status:              "active",
					Width:               854,
					Height:              480,
					IsHeld:              false,
					FramesPendingString: "",
					FramesPending:       []interface{}{},
					FramesString:        "",
					Frames:              []interface{}{},
				}, {
					ID:                  135820,
					Tags:                "animated animator_vs._animation_8:_the_showdown animator_vs._animation_series artist_unknown beams debris effects explosions falling fighting fire flying guzzu liquid rotation shadowqrow smears smoke web western",
					CreatedAt:           time.Date(2020, time.October, 26, 4, 38, 51, 279000000, time.UTC),
					CreatorID:           5084,
					Author:              "Armando",
					Change:              1323224,
					Source:              "https://www.youtube.com/watch?v=0a1r0JaONS4",
					Score:               168,
					MD5:                 "f4078f9ef536d04bb6d3e038beb94547",
					FileSize:            19062220,
					FileURL:             "https://www.sakugabooru.com/data/f4078f9ef536d04bb6d3e038beb94547.mp4",
					IsShownInIndex:      true,
					PreviewURL:          "https://www.sakugabooru.com/data/preview/f4078f9ef536d04bb6d3e038beb94547.jpg",
					PreviewWidth:        150,
					PreviewHeight:       84,
					ActualPreviewWidth:  300,
					ActualPreviewHeight: 169,
					SampleURL:           "https://www.sakugabooru.com/data/f4078f9ef536d04bb6d3e038beb94547.mp4",
					SampleWidth:         854,
					SampleHeight:        480,
					SampleFileSize:      0,
					JpegURL:             "https://www.sakugabooru.com/data/f4078f9ef536d04bb6d3e038beb94547.mp4",
					JpegWidth:           854,
					JpegHeight:          480,
					JpegFileSize:        0,
					Rating:              "s",
					HasChildren:         true,
					ParentID:            0,
					Status:              "active",
					Width:               854,
					Height:              480,
					IsHeld:              false,
					FramesPendingString: "",
					FramesPending:       []interface{}{},
					FramesString:        "",
					Frames:              []interface{}{},
				}, {
					ID:                  135822,
					Tags:                "animated animator_vs._animation_8:_the_showdown animator_vs._animation_series artist_unknown effects fighting fire flying guzzu lightning liquid shadowqrow smears smoke sparks web western wind",
					CreatedAt:           time.Date(2020, time.October, 26, 4, 41, 5, 45000000, time.UTC),
					CreatorID:           5084,
					Author:              "Armando",
					Change:              1323225,
					Source:              "https://www.youtube.com/watch?v=0a1r0JaONS4",
					Score:               58,
					MD5:                 "1f9314adcc612ede1a25692b37b63647",
					FileSize:            13446784,
					FileURL:             "https://www.sakugabooru.com/data/1f9314adcc612ede1a25692b37b63647.mp4",
					IsShownInIndex:      true,
					PreviewURL:          "https://www.sakugabooru.com/data/preview/1f9314adcc612ede1a25692b37b63647.jpg",
					PreviewWidth:        150,
					PreviewHeight:       84,
					ActualPreviewWidth:  300,
					ActualPreviewHeight: 169,
					SampleURL:           "https://www.sakugabooru.com/data/1f9314adcc612ede1a25692b37b63647.mp4",
					SampleWidth:         854,
					SampleHeight:        480,
					SampleFileSize:      0,
					JpegURL:             "https://www.sakugabooru.com/data/1f9314adcc612ede1a25692b37b63647.mp4",
					JpegWidth:           854,
					JpegHeight:          480,
					JpegFileSize:        0,
					Rating:              "s",
					HasChildren:         false,
					ParentID:            0,
					Status:              "active",
					Width:               854,
					Height:              480,
					IsHeld:              false,
					FramesPendingString: "",
					FramesPending:       []interface{}{},
					FramesString:        "",
					Frames:              []interface{}{}}}}
		assert.Equal(t, expected, response)
	})
}
