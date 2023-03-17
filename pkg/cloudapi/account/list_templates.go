package account

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list templates
type ListTemplatesRequest struct {
	// ID an account
	// Required: true
    AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Include deleted images
	// Required: false
	IncludeDeleted bool `url:"includedeleted" json:"includedeleted"`
}

// ListTemplates gets list templates which can be managed by this account
func (a Account) ListTemplates(ctx context.Context, req ListTemplatesRequest) (ListTemplates, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/account/listTemplates"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListTemplates{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
