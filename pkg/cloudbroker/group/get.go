package group

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Reuqest struct for getting details of the specified group.
type GetRequest struct {
	// Group ID
	// Required: true
	GroupID string `url:"groupId" json:"groupId" validate:"required"`
}

// Get gets details of the specified group.
func (g Group) Get(ctx context.Context, req GetRequest) (*ItemGroup, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/group/get"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := ItemGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
