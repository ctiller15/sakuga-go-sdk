package sakugaapi

import (
	"encoding/json"
	"strings"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
)

type FavoritesAPI struct {
	URL string
}

func newFavoritesAPI(baseURL string) *FavoritesAPI {
	newAPI := FavoritesAPI{
		URL: baseURL + "/favorite",
	}

	return &newAPI
}

func mapFavoriteAPIResultToResponse(result models.FavoriteListUsersAPIResult) models.FavoriteListUsersResponse {
	parsedUsers := strings.Split(result.FavoritedUsers, ",")
	response := models.FavoriteListUsersResponse{
		FavoritedUsers: parsedUsers,
	}
	return response
}

func (f *FavoritesAPI) ListUsers(opts *models.FavoriteListUsersOptions) (models.FavoriteListUsersResponse, error) {
	url, err := utils.CreateFavoriteListUserURL(f.URL, opts)
	if err != nil {
		return models.FavoriteListUsersResponse{}, err
	}

	body, err := utils.Fetch(url)

	favoriteListAPIResult := models.FavoriteListUsersAPIResult{}
	err = json.Unmarshal(body, &favoriteListAPIResult)

	if err != nil {
		return models.FavoriteListUsersResponse{}, err
	}

	mappedResponse := mapFavoriteAPIResultToResponse(favoriteListAPIResult)

	return mappedResponse, nil
}
