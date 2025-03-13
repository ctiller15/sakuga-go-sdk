package sakugaapi

import (
	"encoding/json"
	"strings"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
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

func mapFavoriteAPIResultToResponse(result sakugamodels.FavoriteListUsersAPIResult) sakugamodels.FavoriteListUsersResponse {
	parsedUsers := strings.Split(result.FavoritedUsers, ",")
	response := sakugamodels.FavoriteListUsersResponse{
		FavoritedUsers: parsedUsers,
	}
	return response
}

func (f *FavoritesAPI) ListUsers(opts *sakugamodels.FavoriteListUsersOptions) (sakugamodels.FavoriteListUsersResponse, error) {
	url, err := sakugautils.CreateFavoriteListUserURL(f.URL, opts)
	if err != nil {
		return sakugamodels.FavoriteListUsersResponse{}, err
	}

	body, err := sakugautils.Fetch(url)
	if err != nil {
		return sakugamodels.FavoriteListUsersResponse{}, err
	}

	favoriteListAPIResult := sakugamodels.FavoriteListUsersAPIResult{}
	err = json.Unmarshal(body, &favoriteListAPIResult)

	if err != nil {
		return sakugamodels.FavoriteListUsersResponse{}, err
	}

	mappedResponse := mapFavoriteAPIResultToResponse(favoriteListAPIResult)

	return mappedResponse, nil
}
