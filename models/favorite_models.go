package models

// Valid example that doesn't give a 404. (If the ID has no results it 404's)
// https://sakugabooru.com/favorite/list_users.json?id=3

type FavoriteListUsersOptions struct {
	ID int // The post ID
}

type FavoriteListUsersAPIResult struct {
	FavoritedUsers string `json:"favorited_users"`
}

type FavoriteListUsersAPIResponse struct {
	FavoritedUsers []string
}
