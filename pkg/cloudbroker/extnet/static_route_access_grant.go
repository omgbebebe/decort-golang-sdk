package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for grant access to static route to Compute/ViNS
type StaticRouteAccessGrantRequest struct {
	// ExtNet ID to grant access
	// Required: true
	ExtNetID uint64 `url:"extnetId" json:"extnetId" validate:"required"`

	// Route ID to grant access, can be found in staticRouteList
	// Required: true
	RouteId uint64 `url:"routeId" json:"routeId" validate:"required"`

	// List of Compute IDs to grant access to this route
	// Required: false
	ComputeIds []uint64 `url:"computeIds,omitempty" json:"computeIds,omitempty"`
}

// Grant access to static route to Compute/ViNS
func (v ExtNet) StaticRouteAccessGrant(ctx context.Context, req StaticRouteAccessGrantRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/staticRouteAccessGrant"

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
