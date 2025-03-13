package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/stretchr/testify/assert"
)

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
		opts := sakugamodels.CommentShowOptions{
			ID: 90003,
		}
		response, err := newAPI.Comments.Show(&opts)
		assert.Nil(t, err)
		expected := sakugamodels.CommentShowResponseItem{
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
