package extnet

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for static route list
type StaticRouteListRequest struct {
	// ExtNet ID to show list of static routes
	// Required: true
	ExtNetID uint64 `url:"extnetId" json:"extnetId" validate:"required"`
}

// Show list of static routes for ViNS
func (v ExtNet) StaticRouteList(ctx context.Context, req StaticRouteListRequest) (*ListStaticRoutes, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/staticRouteList"

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
