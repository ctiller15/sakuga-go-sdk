package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/stretchr/testify/assert"
)

func TestWikiList(t *testing.T) {
	t.Run("base case, happy path", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/wiki.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(wikiDataListResponse)
			}
		}))
		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.WikiListOptions{}
		response, err := newAPI.Wiki.List(&opts)
		assert.Nil(t, err)
		expected := []sakugamodels.WikiListResponseItem{
			{
				ID:        42,
				CreatedAt: time.Date(2016, time.October, 5, 13, 30, 31, 967000000, time.UTC),
				UpdatedAt: time.Date(2016, time.October, 5, 13, 30, 31, 967000000, time.UTC),
				Title:     "3d_background",
				Body:      "Background is made with 3dcg. Implicates [[cgi]].",
				UpdaterID: 933,
				Locked:    false,
				Version:   1,
			},
			{
				ID: 312, CreatedAt: time.Date(2024, time.November, 1, 5, 17, 8, 294000000, time.UTC),
				UpdatedAt: time.Date(2024, time.November, 1, 5, 17, 8, 294000000, time.UTC),
				Title:     "aimman_ibrahim",
				Body:      "Brother of \"Jack-Amin Ibrahim\":https://www.sakugabooru.com/post?tags=jack-amin_ibrahim",
				UpdaterID: 34147,
				Locked:    false,
				Version:   1,
			},
			{
				ID:        58,
				CreatedAt: time.Date(2017, time.August, 30, 11, 10, 14, 286000000, time.UTC),
				UpdatedAt: time.Date(2017, time.August, 30, 11, 43, 21, 687000000, time.UTC),
				Title:     "ai_monogatari:_9_love_stories",
				Body:      "h2. 「HERO」\r\n\r\n監督・キャラクターデザイン \r\n森本晃司 Kouji Morimoto\r\n\r\nキャラクターデザイン・作画監督\r\n清山滋崇 Shigetaka Kiyoyama\r\n\r\n原画\r\n誟橋伸司 Shinji Morohashi\r\n後藤孝宏 Takahiro Gotou\r\n西村\u3000智 Satoshi Nishimura\r\n安藤正浩 Masahiro Andou (the other one)\r\n田辺\u3000修 Osamu Tanabe\r\n結城明宏 Akihiro Yuuki\r\n\r\n\r\n\r\n\r\nh2. 「夜をぶっとばせ」\r\n\r\n監督\r\n浜津\u3000守 Mamoru Hamatsu\r\n\r\nキャラクターデザイン・作画監督・原画\r\n村瀬修功 Shuukou Murase\r\n\r\n\r\n\r\n\r\nh2. 「ライオンとペリカン」\r\n\r\n監督\r\n澤井幸次 Kouji Sawai\r\n\r\nキャラクタデザイン・作画監督\r\n戸倉紀元 Norimoto Tokura\r\n\r\n原画\r\n戸倉紀元 Norimoto Tokura\r\n藤川\u3000太 Futoshi Fujikawa\r\n\r\n",
				UpdaterID: 933,
				Locked:    false,
				Version:   4,
			},
			{
				ID:        211,
				CreatedAt: time.Date(2022, time.November, 12, 12, 51, 55, 235000000, time.UTC),
				UpdatedAt: time.Date(2022, time.November, 12, 12, 51, 55, 235000000, time.UTC),
				Title:     "aito_ohashi",
				Body:      "大橋 藍人",
				UpdaterID: 6259,
				Locked:    false,
				Version:   1,
			},
			{
				ID:        202,
				CreatedAt: time.Date(2022, time.October, 3, 14, 31, 32, 798000000, time.UTC),
				UpdatedAt: time.Date(2022, time.October, 3, 14, 31, 32, 798000000, time.UTC),
				Title:     "akari_ranzaki",
				Body:      "藍崎\u3000灯",
				UpdaterID: 6259,
				Locked:    false,
				Version:   1,
			},
		}
		assert.Equal(t, expected, response)
	})
}
