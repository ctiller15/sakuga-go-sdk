package sakugaapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/models"
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
			if r.URL.Path == "/post.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(postDataResponse)
			}
		}))
		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)
		opts := models.PostsListOptions{}
		response, err := newAPI.Posts.List(&opts)
		assert.Nil(t, err)
		expected := []models.PostListResponseResult{
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

func TestAPITagList(t *testing.T) {
	t.Run("Happy path - retrieve tags", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/tag.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(tagDataResponse)
			}
		}))
		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)
		opts := models.TagListOptions{}
		response, err := newAPI.Tags.List(&opts)
		assert.Nil(t, err)
		expected := []models.TagListResponseResult{
			{
				ID:        14391,
				Name:      "",
				Count:     0,
				Type:      0,
				Ambiguous: false,
			},
			{
				ID:        25325,
				Name:      "(???)",
				Count:     0,
				Type:      0,
				Ambiguous: false,
			},
			{
				ID:        25384,
				Name:      "*",
				Count:     0,
				Type:      0,
				Ambiguous: false,
			},
			{
				ID:        22702,
				Name:      "?",
				Count:     0,
				Type:      0,
				Ambiguous: false,
			},
			{
				ID:        24506,
				Name:      "\\",
				Count:     0,
				Type:      0,
				Ambiguous: false,
			},
			{
				ID:        26578,
				Name:      "]",
				Count:     0,
				Type:      0,
				Ambiguous: false,
			},
			{
				ID:        25768,
				Name:      "_",
				Count:     0,
				Type:      0,
				Ambiguous: false,
			},
			{
				ID:        25547,
				Name:      "0:00/0:46",
				Count:     0,
				Type:      0,
				Ambiguous: false,
			},
			{
				ID:        5620,
				Name:      "00:08",
				Count:     5,
				Type:      3,
				Ambiguous: false,
			},
			{
				ID:        11330,
				Name:      "001",
				Count:     1,
				Type:      3,
				Ambiguous: false,
			},
		}
		assert.Equal(t, expected, response)
	})
}

func TestApiTagRelated(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.URL.RequestURI()
		fmt.Println(val)
		if r.URL.RequestURI() == "/tag/related.json?tags=animated" {
			w.WriteHeader(http.StatusOK)
			w.Write(tagDataRelatedResponse)
		}
	}))
	defer server.Close()

	t.Run("Happy path - retrieve related tags", func(t *testing.T) {

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)
		opts := models.TagRelatedOptions{
			Tags: []string{"animated"},
		}
		response, err := newAPI.Tags.Related(&opts)
		assert.Nil(t, err)
		expected := map[string][]models.RelatedTagResponse{
			"animated": {
				{
					Name: "animated",
					ID:   155193,
				},
				{
					Name: "effects",
					ID:   88897,
				},
				{
					Name: "artist_unknown",
					ID:   81238,
				},
				{
					Name: "character_acting",
					ID:   60788,
				},
				{
					Name: "smoke",
					ID:   42239,
				},
				{
					Name: "smears",
					ID:   41842,
				},
				{
					Name: "fighting",
					ID:   38855,
				},
				{
					Name: "creatures",
					ID:   26170,
				},
				{
					Name: "presumed",
					ID:   23921,
				},
				{
					Name: "liquid",
					ID:   20841,
				},
				{
					Name: "debris",
					ID:   18986,
				},
				{
					Name: "explosions",
					ID:   16904,
				},
				{
					Name: "hair",
					ID:   16480,
				},
				{
					Name: "running",
					ID:   15126,
				},
				{
					Name: "fabric",
					ID:   15021,
				},
				{
					Name: "western",
					ID:   14087,
				},
				{
					Name: "background_animation",
					ID:   12432,
				},
				{
					Name: "mecha",
					ID:   12013,
				},
				{
					Name: "fire",
					ID:   11356,
				},
				{
					Name: "sparks",
					ID:   10068,
				},
				{
					Name: "animals",
					ID:   9591,
				},
				{
					Name: "impact_frames",
					ID:   9347,
				},
				{
					Name: "lightning",
					ID:   9244,
				},
				{
					Name: "cgi",
					ID:   8381,
				},
				{
					Name: "production_materials",
					ID:   7688,
				},
			},
		}
		assert.Equal(t, expected, response)
	})
}

func TestArtistAPIList(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RequestURI() == "/artist.json" {
			w.WriteHeader(http.StatusOK)
			w.Write(artistDataListResponse)
		}
	}))
	defer server.Close()

	t.Run("happy path, retrieve artists", func(t *testing.T) {
		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)
		opts := models.ArtistListOptions{}
		response, err := newAPI.Artists.List(&opts)
		assert.Nil(t, err)
		expected := []models.ArtistListAPIResponseItem([]models.ArtistListAPIResponseItem{
			{
				ID:      1172,
				Name:    "심땅",
				AliasID: 1171,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      1166,
				Name:    "장석민",
				AliasID: 1165,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      1170,
				Name:    "파랑상자",
				AliasID: 1169,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      647,
				Name:    ".",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      672,
				Name:    ">",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      593,
				Name:    "0nepeach",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      1053,
				Name:    "666ban",
				AliasID: 1052,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      1192,
				Name:    "abegen",
				AliasID: 1191,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      1068,
				Name:    "adam_zheng",
				AliasID: 1066,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID:      651,
				Name:    "aida_saakian",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 531, Name: "aimkid",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{"https://twitter.com/aimkidblast"},
			},
			{
				ID: 997, Name: "aimman_ibrahim",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 1027, Name: "aiowaruru",
				AliasID: 616,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 799, Name: "aito_ohashi",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 1189, Name: "akapape",
				AliasID: 811,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 782, Name: "akari_ranzaki",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 871, Name: "aki_deguchi",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 559, Name: "akihiro_ota",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 970, Name: "akiko_kudo",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 505, Name: "akio_watanabe",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 811, Name: "akira_hamaguchi",
				AliasID: 0,
				GroupID: 0,
				Urls: []string{
					"https://x.com/oyasirazumonai",
					"https://x.com/akapape",
				},
			},
			{
				ID: 753, Name: "akira_kikuchi",
				AliasID: 0,
				GroupID: 0,
				Urls: []string{
					"https://www.animenewsnetwork.com/encyclopedia/people.php?id=19539",
					"https://w.atwiki.jp/sakuga/pages/1138.amp",
					"https://w.atwiki.jp/anime_wiki/pages/6287.amp",
				},
			},
			{
				ID: 136, Name: "akira_takata",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 911, Name: "akira_tanaka",
				AliasID: 910,
				GroupID: 0,
				Urls:    []string{},
			},
			{
				ID: 975, Name: "ako_no_akio",
				AliasID: 0,
				GroupID: 0,
				Urls:    []string{},
			},
		},
		)
		assert.Equal(t, expected, response)
	})
}

func TestCommentAPIShow(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RequestURI() == "/comment/show.json?id=90003" {
			w.WriteHeader(http.StatusOK)
			w.Write(commentDataShowResponse)
		}
	}))
	defer server.Close()

	t.Run("happy path, retrieve a comment", func(t *testing.T) {
		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)
		opts := models.CommentShowOptions{
			ID: 90003,
		}
		response, err := newAPI.Comments.Show(&opts)
		assert.Nil(t, err)
		expected := models.CommentShowResponseItem{
			ID:        90003,
			CreatedAt: time.Date(2023, time.December, 26, 20, 52, 33, 419000000, time.UTC),
			PostID:    107257,
			Creator:   "Ivorybacon",
			CreatorID: 12904,
			Body:      "I'm tempted to say that Otsuka is here given some of the smears on dispay. 0:05 and 0:11 for instance appear to have those very jagged smear lines that she seems to do often. (post #217312) and (post #215778) for comparison.\r\n\r\nShe's listed as an AD on this episode so there's a chance that it might just be a correction on her part.",
		}
		assert.Equal(t, expected, response)
	})
}
