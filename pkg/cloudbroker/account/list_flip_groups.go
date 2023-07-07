package account

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list FLIPGroups
type ListFLIPGroupsRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by vinsId
	// Required: false
	VINSID uint64 `url:"vinsId,omitempty" json:"vinsId,omitempty"`

	// Find by VINS name
	// Required: false
	VINSName string `url:"vinsName,omitempty" json:"vinsName,omitempty"`

	// Find by external network id
	// Required: false
	ExtNetID uint64 `url:"extnetId,omitempty" json:"extnetId,omitempty"`

	// Find by IP
	// Required: false
	ByIP string `url:"byIp,omitempty" json:"byIp,omitempty"`

	// Find by flipGroup Id
	// Required: false
	FLIPGroupID uint64 `url:"flipGroupId,omitempty" json:"flipGroupId,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListFLIPGroups gets list all FLIPGroups under specified account, accessible by the user
func (a Account) ListFLIPGroups(ctx context.Context, req ListFLIPGroupsRequest) (*ListFLIPGroups, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/account/listFlipGroups"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListFLIPGroups{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
