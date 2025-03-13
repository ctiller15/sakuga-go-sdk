package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/stretchr/testify/assert"
)

func TestUsersSearch(t *testing.T) {
	t.Run("base case, happy path", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/user.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(userDataSearchResponse)
			}
		}))

		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.UserSearchOptions{}
		response, err := newAPI.Users.Search(&opts)
		assert.Nil(t, err)
		expected := []sakugamodels.UserSearchAPIResponseItem{
			{
				Name: "sakurga56",
				ID:   39722,
			},
			{
				Name: "Kanegane",
				ID:   39721,
			},
			{
				Name: "bonofyah",
				ID:   39720,
			},
			{
				Name: "antoines_workshop",
				ID:   39719,
			},
			{
				Name: "Mekkies",
				ID:   39718,
			},
			{
				Name: "Phoslos",
				ID:   39717,
			},
			{
				Name: "gugull",
				ID:   39716,
			},
			{
				Name: "ForestElephant",
				ID:   39715,
			},
			{
				Name: "suibo",
				ID:   39714,
			},
			{
				Name: "qJvECnDaBJcH",
				ID:   39713,
			},
			{
				Name: "Etozy",
				ID:   39712,
			},
			{
				Name: "奥利奥",
				ID:   39711,
			},
			{
				Name: "alazoro",
				ID:   39710,
			},
			{
				Name: "nikomi_shio",
				ID:   39709,
			},
			{
				Name: "Barusu",
				ID:   39708,
			},
			{
				Name: "GEEanimations",
				ID:   39707,
			},
			{
				Name: "GT",
				ID:   39706,
			},
			{
				Name: "MARL_",
				ID:   39705,
			},
			{
				Name: "Aley",
				ID:   39704,
			},
			{
				Name: "anzu",
				ID:   39703,
			},
		}
		assert.Equal(t, expected, response)
	})
}
