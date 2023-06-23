package user

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for showing list of dicts with information about
// apiaccess groups contains to the user.
type APIAccessListRequest struct {
	// ID of the user to list API access groups for.
	// Required: true
	UserID string `url:"userId" json:"userId" validate:"required"`
}

// APIAccessList shows list of dicts with information about apiaccess groups contains to the user.
func (u User) APIAccessList(ctx context.Context, req APIAccessListRequest) (ListAPIAccess, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/apiaccessList"

	list := ListAPIAccess{}

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
