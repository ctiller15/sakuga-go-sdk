package sakugaapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

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
		fmt.Println(response)
		t.Errorf("Finish the test!")
	})
}
