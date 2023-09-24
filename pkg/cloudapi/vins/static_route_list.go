package vins

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for static route list
type StaticRouteListRequest struct {
	// ViNS ID to show list of static routes
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`
}

// Show list of static routes for ViNS
func (v VINS) StaticRouteList(ctx context.Context, req StaticRouteListRequest) (*ListStaticRoutes, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/vins/staticRouteList"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListStaticRoutes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
