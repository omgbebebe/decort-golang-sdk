package compute

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for attach network
type NetAttachRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// Network type
	// 'EXTNET' for connect to external network directly
	// and 'VINS' for connect to ViNS
	// Required: true
	NetType string `url:"netType" json:"netType" validate:"computeNetType"`

	// Network ID for connect to
	// For EXTNET - external network ID
	// For VINS - VINS ID
	// Required: true
	NetID uint64 `url:"netId" json:"netId" validate:"required"`

	// Directly required IP address for new network interface
	// Required: false
	IPAddr string `url:"ipAddr,omitempty" json:"ipAddr,omitempty"`
}

// NetAttach attach network to compute and gets info about network
func (c Compute) NetAttach(ctx context.Context, req NetAttachRequest) (*RecordNetAttach, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/netAttach"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordNetAttach{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
