package sakugaapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/stretchr/testify/assert"
)

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
		opts := sakugamodels.TagListOptions{}
		response, err := newAPI.Tags.List(&opts)
		assert.Nil(t, err)
		expected := []sakugamodels.TagListResponseResult{
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
		opts := sakugamodels.TagRelatedOptions{
			Tags: []string{"animated"},
		}
		response, err := newAPI.Tags.Related(&opts)
		assert.Nil(t, err)
		expected := map[string][]sakugamodels.RelatedTagResponse{
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
