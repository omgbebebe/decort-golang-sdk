package user

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting a list of the matching usernames for a given string.
type GetMatchingUsernamesRequest struct {
	// Regex of the usernames to searched for.
	// Required: true
	UsernameRegex string `url:"usernameregex" json:"usernameregex" validate:"required"`

	// The number of usernames to return.
	// Required: true
	Limit uint64 `url:"limit" json:"limit" validate:"required"`
}

// GetMatchingUsernames gets a list of the matching usernames for a given string.
func (u User) GetMatchingUsernames(ctx context.Context, req GetMatchingUsernamesRequest) (ListMatchingUsernames, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/getMatchingUsernames"

	list := ListMatchingUsernames{}

	res, err := u.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &list)

	if err != nil {
		return nil, err
	}

	return list, nil
}
