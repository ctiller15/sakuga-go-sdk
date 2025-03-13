package sakugaapi

import (
	"encoding/json"

	"github.com/ctiller15/sakuga-go-sdk/models"
	"github.com/ctiller15/sakuga-go-sdk/utils"
)

type UsersAPI struct {
	URL string
}

func newUsersAPI(baseURL string) *UsersAPI {
	newAPI := UsersAPI{
		URL: baseURL + "/user",
	}
	return &newAPI
}

func mapUserSearchAPIResultToResponse(items []models.UserSearchAPIResultItem) ([]models.UserSearchAPIResponseItem, error) {
	userSearchAPIReponseResults := make([]models.UserSearchAPIResponseItem, 0)

	for _, item := range items {
		newItem := models.UserSearchAPIResponseItem(item)
		userSearchAPIReponseResults = append(userSearchAPIReponseResults, newItem)
	}

	return userSearchAPIReponseResults, nil
}

func (u *UsersAPI) Search(opts *models.UserSearchOptions) ([]models.UserSearchAPIResponseItem, error) {
	url, err := utils.CreateUserSearchUrl(u.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := utils.Fetch(url)
	if err != nil {
		return nil, err
	}

	userSearchAPIResult := make([]models.UserSearchAPIResultItem, 0)
	err = json.Unmarshal(body, &userSearchAPIResult)

	if err != nil {
		return nil, err
	}

	mappedResponse, err := mapUserSearchAPIResultToResponse(userSearchAPIResult)
	if err != nil {
		return nil, err
	}
	return mappedResponse, nil
}
