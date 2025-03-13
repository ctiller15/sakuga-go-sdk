package sakugaapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/stretchr/testify/assert"
)

func TestNoteList(t *testing.T) {
	t.Run("base case, retrieve notes", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/note.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(noteDataListResponse)
			}
		}))

		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.NoteListOptions{}
		response, err := newAPI.Notes.List(&opts)
		assert.Nil(t, err)
		expected := []sakugamodels.NoteListResponseItem{
			{
				ID:        270,
				CreatedAt: time.Date(2024, time.February, 6, 22, 15, 57, 304000000, time.UTC),
				UpdatedAt: time.Date(2024, time.February, 6, 22, 15, 57, 304000000, time.UTC),
				CreatorID: 4456,
				X:         555,
				Y:         137,
				Width:     153,
				Height:    119,
				IsActive:  true,
				PostID:    248418,
				Body:      "object is always cel\nfull color trace paint",
				Version:   1,
			},
			{
				ID:        269,
				CreatedAt: time.Date(2024, time.February, 6, 22, 8, 57, 131000000, time.UTC),
				UpdatedAt: time.Date(2024, time.February, 6, 22, 8, 57, 131000000, time.UTC),
				CreatorID: 4456,
				X:         277,
				Y:         5,
				Width:     148,
				Height:    148,
				IsActive:  true,
				PostID:    248418,
				Body:      "Demonic Mirroring Ice Crystals ",
				Version:   1,
			},
		}
		assert.Equal(t, expected, response)
	})
}

func TestNoteSearch(t *testing.T) {
	t.Run("base case, retrieve notes", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/note/search.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(noteDataSearchResponse)
			}
		}))

		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.NoteSearchOptions{
			Query: "fight",
		}
		response, err := newAPI.Notes.Search(&opts)
		assert.Nil(t, err)
		expected := []sakugamodels.NoteListResponseItem{
			{
				ID: 4, CreatedAt: time.Date(2013, time.September, 3, 20, 29, 32, 200000000, time.UTC),
				UpdatedAt: time.Date(2013, time.September, 3, 20, 29, 32, 200000000, time.UTC),
				CreatorID: 55,
				X:         28,
				Y:         21,
				Width:     861,
				Height:    100,
				IsActive:  true,
				PostID:    1273,
				Body:      "        \"Well, even if I do stuff like this, there probably isn't anyone who understands what it means.\"\n         That was the first time I'd heard him make that complaint since the \"Moonrise Mannor\" screening.\n        I imagine those complaints of his, as he cheerfully streaked through a public street in the middle of winter, were his true feelings.\n        The life of Ikuhara Kunihiko already had no distinction between public and private.\n \n        Just what is this guy, who dares to dance in the nude, like in the other world, where he's really going for it?\n        The things he fights against are instinct and normalcy. I think those two are the big ones.\n        Basically, concern for others is part of it. In most cases, people are unconsciously dependent on those sorts of things.\n        But, Ikuhara Kunihiko just seems to intuitively understand that using those elements as a tools is an inherent part of being human as well. I think he already had the \"subjective goal\" necessary to do that as a student. In other words, a spire of determination.\n \n        So––\n       \n        Now, he seems to have rid himself of his discontent in not finding anyone who understands him. But, in no way do I mean to imply that was the result of finding a friend he'd always hoped for. It's really quite the opposite. It's more that, he's become acutely cognizant of the fact that everyone other than himself is an enemy.\n        Banished from paradise, man was made mortal.\n        I think, Ikuhara Kunihiko decides to \"be Ikuhara Kunihiko\" in film, TV and even in his New Year's cards, with that harsh reality firmly planted in his mind.\n        He was already well aware of the struggle with the truth that you are you and nothing more.\n        But, I desperately wish that I might be able to watch over what he does from now on, firm in the belief that there is some truth to that stance of his.\n        Peace.\n \n12/21/1994\n \n---\nNotes:\n \n[1] - I did some digging and I believe this story pertains to episode 31 of Sailor Moon, \"恋されて追われて!ルナの最悪の日.\" It won the Animage Anime Grand Prix for individual episode of a TV series in 1992. Ikuhara storyboarded it and it does involve cats and mice (it's that Luna-centric episode).\n \nReferences:\nhttp://ja.wikipedia.org/wiki/アニメグランプリ\nhttp://ja.wikipedia.org/wiki/美少女戦士セーラームーン_(アニメ)\n \n \n[2] - 製作進行: Apparently this Toei's unique nomenclature for 制作進行. \"Production Assistant\" or \"Production Runner\" seem to be the common English translations for the position. I've also seen \"Production Coordinator.\" According to Wikipedia, at least, the job is often a stepping stone to becoming a producer or director, so they do a variety of odd jobs to help out the main staff and keep things running smoothly.\n \nReferences:\nhttp://ja.wikipedia.org/wiki/%E5%88%B6%E4%BD%9C%E9%80%B2%E8%A1%8C#.E3.82.A2.E3.83.8B.E3.83.A1.E3.83.BC.E3.82.B7.E3.83.A7.E3.83.B3.E4.BD.9C.E5.93.81.E3.81.AE.E5.88.B6.E4.BD.9C.E9.80.B2.E8.A1.8C\n \n \n[3] - The last card in the set they copied for this book.",
				Version:   1,
			}, {
				ID:        70,
				CreatedAt: time.Date(2016, time.December, 12, 0, 1, 36, 576000000, time.UTC),
				UpdatedAt: time.Date(2016, time.December, 12, 0, 1, 36, 576000000, time.UTC),
				CreatorID: 933,
				X:         250,
				Y:         180,
				Width:     90,
				Height:    150,
				IsActive:  true,
				PostID:    28259,
				Body:      "Since then he's done the scene in Gall Force where monsters come out of the crumbling ship, and the scene at the end of Kikoukai Galient: Tetsu no Monshou where Galient cuts Jashinhei apart. In Machine Robo, he did Baikanfuu's henkei which became a bank cut, Baikanfuu and Mizuchi's fight scene in episode 6 etc.",
				Version:   1,
			}, {
				ID:        71,
				CreatedAt: time.Date(2016, time.December, 12, 0, 8, 26, 418000000, time.UTC),
				UpdatedAt: time.Date(2016, time.December, 12, 0, 8, 26, 418000000, time.UTC),
				CreatorID: 933,
				X:         202,
				Y:         176,
				Width:     47,
				Height:    159,
				IsActive:  true,
				PostID:    28259,
				Body:      "Recently he's done the fight scene with the monsters in Gakuen Tokusou Hikaruon and the baseball scene in Bats and Terry, and been involved in Bubblegum Crisis, among others.",
				Version:   1,
			},
		}
		assert.Equal(t, expected, response)
	})
}

func TestNoteHistory(t *testing.T) {
	t.Run("base case, retrieve notes", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/note/history.json" {
				w.WriteHeader(http.StatusOK)
				w.Write(noteDataHistoryResponse)
			}
		}))

		defer server.Close()

		newAPI := NewAPI()
		newAPI.SetHomeURL(server.URL)

		opts := sakugamodels.NoteHistorySearchOptions{}
		response, err := newAPI.Notes.History(&opts)
		assert.Nil(t, err)
		expected := []sakugamodels.NoteHistoryResponseItem{
			{
				CreatedAt: time.Date(2021, time.September, 23, 22, 46, 0, 458000000, time.UTC),
				UpdatedAt: time.Date(2025, time.January, 14, 12, 8, 35, 705000000, time.UTC),
				CreatorID: 34147,
				X:         66,
				Y:         230,
				Width:     148,
				Height:    148,
				IsActive:  true,
				PostID:    123229,
				Body:      "sp",
				Version:   6,
			},
			{
				CreatedAt: time.Date(2021, time.September, 23, 22, 46, 0, 458000000, time.UTC),
				UpdatedAt: time.Date(2024, time.December, 26, 7, 3, 0, 633000000, time.UTC),
				CreatorID: 15559,
				X:         66,
				Y:         230,
				Width:     148,
				Height:    148,
				IsActive:  true,
				PostID:    123229,
				Body:      "sp",
				Version:   5,
			},
			{
				CreatedAt: time.Date(2021, time.September, 23, 22, 46, 0, 458000000, time.UTC),
				UpdatedAt: time.Date(2024, time.December, 17, 2, 39, 27, 997000000, time.UTC),
				CreatorID: 15559,
				X:         66,
				Y:         230,
				Width:     148,
				Height:    148,
				IsActive:  true,
				PostID:    123229,
				Body:      "sp",
				Version:   4,
			},
			{
				CreatedAt: time.Date(2021, time.September, 23, 22, 46, 0, 458000000, time.UTC),
				UpdatedAt: time.Date(2024, time.December, 17, 2, 37, 57, 75000000, time.UTC),
				CreatorID: 15559,
				X:         66,
				Y:         230,
				Width:     148,
				Height:    148,
				IsActive:  true,
				PostID:    123229,
				Body:      "sp",
				Version:   3,
			},
			{
				CreatedAt: time.Date(2021, time.September, 23, 22, 46, 0, 458000000, time.UTC),
				UpdatedAt: time.Date(2024, time.August, 16, 0, 54, 13, 780000000, time.UTC),
				CreatorID: 8283,
				X:         66,
				Y:         230,
				Width:     148,
				Height:    148,
				IsActive:  true,
				PostID:    123229,
				Body:      "sp",
				Version:   2,
			},
			{
				CreatedAt: time.Date(2021, time.September, 23, 22, 46, 0, 458000000, time.UTC),
				UpdatedAt: time.Date(2021, time.September, 23, 22, 46, 0, 458000000, time.UTC),
				CreatorID: 19493,
				X:         66,
				Y:         230,
				Width:     148,
				Height:    148,
				IsActive:  true,
				PostID:    123229,
				Body:      "sp",
				Version:   1,
			},
		}
		assert.Equal(t, expected, response)
	})
}
