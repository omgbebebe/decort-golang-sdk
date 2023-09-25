package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for add static route
type StaticRouteAddRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Destination network
	// Required: true
	Destination string `url:"destination" json:"destination" validate:"required"`

	// Destination network mask in 255.255.255.255 format
	// Required: true
	Netmask string `url:"netmask" json:"netmask" validate:"required"`

	// Next hop host, IP address from ViNS ID free IP pool
	// Required: true
	Gateway string `url:"gateway" json:"gateway" validate:"required"`

	// List of Compute IDs which have access to this route
	// Required: false
	ComputeIds []uint64 `url:"computeIds,omitempty" json:"computeIds,omitempty"`
}

// StaticRouteAdd add new static route to ViNS
func (v VINS) StaticRouteAdd(ctx context.Context, req StaticRouteAddRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/staticRouteAdd"

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
