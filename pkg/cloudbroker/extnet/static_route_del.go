package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for remove static route from ViNS
type StaticRouteDelRequest struct {
	// ExtNet ID to remove static route from
	// Required: true
	ExtNetID uint64 `url:"extnetId" json:"extnetId" validate:"required"`

	// Route ID to remove, can be found in staticRouteList
	// Required: true
	RouteId uint64 `url:"routeId" json:"routeId" validate:"required"`
}

// Remove static route from ViNS
func (v ExtNet) StaticRouteDel(ctx context.Context, req StaticRouteDelRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/staticRouteDel"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
