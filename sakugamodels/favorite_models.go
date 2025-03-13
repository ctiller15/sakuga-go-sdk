package sakugamodels

type FavoriteListUsersOptions struct {
	ID int // The post ID
}

type FavoriteListUsersAPIResult struct {
	FavoritedUsers string `json:"favorited_users"`
}

type FavoriteListUsersResponse struct {
	FavoritedUsers []string
}
