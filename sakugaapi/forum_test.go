package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/stretchr/testify/assert"
)

func TestForumList(t *testing.T) {
	t.Run("base case, happy path", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/forum.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(forumDataListResponse)
			}
		}))

		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.ForumListOptions{}
		response, err := newAPI.Forum.List(&opts)
		assert.Nil(t, err)
		expected := []sakugamodels.ForumListResponseItem{
			{
				Body:      "If the website breaks or does something unexpected, post about it here.\r\n\r\nYou can also contact the booru Admin team via the following channels:\r\n\r\naers - http://sakuga.yshi.org/user/show/1 / https://twitter.com/aers00\r\nkvin - http://sakuga.yshi.org/user/show/2 / https://twitter.com/Yuyucow\r\nme - http://sakuga.yshi.org/user/show/4 / https://twitter.com/Kraker2k\r\n\r\nYou can also visit #sakuga on irc.rizon.net and leave a message there.\r\n\r\n",
				Creator:   "Kraker2k",
				CreatorID: 4,
				ID:        12,
				ParentID:  0,
				Title:     "Post errors and admin questions here",
				UpdatedAt: time.Date(2024, time.August, 29, 21, 13, 33, 244000000, time.UTC),
				Pages:     2,
			},
			{
				Body:      "is there a way to search by how long or how big it is?\r\nlike longer than/shorter than  seconds\r\nor by size bigger/smaller\r\nalso is there a way to order the anime by how many posts in their tags there is or something",
				Creator:   "amirdrama",
				CreatorID: 24087,
				ID:        1580,
				ParentID:  0,
				Title:     "is there a way to search by how long or how big it is?",
				UpdatedAt: time.Date(2024, time.April, 14, 19, 32, 1, 866000000, time.UTC),
				Pages:     1,
			},
			{
				Body:      " Anyone know of a community or place to get feedback and talk about creating animation? Like a serious one with good activity, not full of people that look like they drew stuff in MS paint (no offense to them). Like a discord or website.",
				Creator:   "dracobuster",
				CreatorID: 24948,
				ID:        1352,
				ParentID:  0,
				Title:     "Practicing Animation",
				UpdatedAt: time.Date(2024, time.April, 1, 16, 51, 22, 984000000, time.UTC),
				Pages:     1,
			},
			{
				Body:      "Hello, several months ago I created a pool named \"Regeneration\". Where is it now? Is it violated some rules or something like?",
				Creator:   "Nojtovsky",
				CreatorID: 19519,
				ID:        1563,
				ParentID:  0,
				Title:     "Where is my pool?",
				UpdatedAt: time.Date(2024, time.April, 1, 15, 26, 9, 392000000, time.UTC),
				Pages:     1,
			},
			{
				Body:      "is there a place to archive anime official arts?  \r\nlike anime Blu-ray covers/arts that are for magazine etc...",
				Creator:   "amirdrama",
				CreatorID: 24087,
				ID:        1473,
				ParentID:  0,
				Title:     "is there a place to archive anime official arts?",
				UpdatedAt: time.Date(2024, time.March, 22, 14, 44, 51, 107000000, time.UTC),
				Pages:     1,
			},
			{
				Body:      "I really liked this cut in AoT final season but CG police took it down for a very unjustified reason imo. I personally don't know how best to record/reupload this\r\nhttps://www.sakugabooru.com/post/show/241415",
				Creator:   "ItsAMemeMario",
				CreatorID: 23275,
				ID:        1565,
				ParentID:  0,
				Title:     "Any way to bring a deleted post back?",
				UpdatedAt: time.Date(2024, time.March, 6, 12, 2, 40, 117000000, time.UTC),
				Pages:     1,
			},
		}
		assert.Equal(t, expected, response)
	})
}
