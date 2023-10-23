package group

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get details of the specified group.
type GetRequest struct {
	// Group ID
	// Required: true
	GroupID string `url:"groupId" json:"groupId" validate:"required"`
}

// Get gets details of the specified group as an ItemGroup struct
func (g Group) Get(ctx context.Context, req GetRequest) (*ItemGroup, error) {
	res, err := g.GetRaw(ctx, req)
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

// GetRaw gets details of the specified group as an array of bytes
func (g Group) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/group/get"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
