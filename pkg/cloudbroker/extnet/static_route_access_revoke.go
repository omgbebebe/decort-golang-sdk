package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for revoke access to static route to Compute/ViNS
type StaticRouteAccessRevokeRequest struct {
	// ExtNet ID to revoke access
	// Required: true
	ExtNetID uint64 `url:"extnetId" json:"extnetId" validate:"required"`

	// Route ID to revoke access, can be found in staticRouteList
	// Required: true
	RouteId uint64 `url:"routeId" json:"routeId" validate:"required"`

	// List of Compute IDs to revoke access to this route
	// Required: false
	ComputeIds []uint64 `url:"computeIds,omitempty" json:"computeIds,omitempty"`
}

// Revoke access to static route to Compute/ViNS
func (v ExtNet) StaticRouteAccessRevoke(ctx context.Context, req StaticRouteAccessRevokeRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/staticRouteAccessRevoke"

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
