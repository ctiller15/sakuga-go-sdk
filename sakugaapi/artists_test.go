package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/stretchr/testify/assert"
)

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
