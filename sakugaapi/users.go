package sakugaapi

import (
	"encoding/json"

	"github.com/ctiller15/sakuga-go-sdk/sakugamodels"
	"github.com/ctiller15/sakuga-go-sdk/sakugautils"
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

func mapUserSearchAPIResultToResponse(items []sakugamodels.UserSearchAPIResultItem) ([]sakugamodels.UserSearchAPIResponseItem, error) {
	userSearchAPIReponseResults := make([]sakugamodels.UserSearchAPIResponseItem, 0)

	for _, item := range items {
		newItem := sakugamodels.UserSearchAPIResponseItem(item)
		userSearchAPIReponseResults = append(userSearchAPIReponseResults, newItem)
	}

	return userSearchAPIReponseResults, nil
}

func (u *UsersAPI) Search(opts *sakugamodels.UserSearchOptions) ([]sakugamodels.UserSearchAPIResponseItem, error) {
	url, err := sakugautils.CreateUserSearchUrl(u.URL, opts)
	if err != nil {
		return nil, err
	}

	body, err := sakugautils.Fetch(url)
	if err != nil {
		return nil, err
	}

	userSearchAPIResult := make([]sakugamodels.UserSearchAPIResultItem, 0)
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
