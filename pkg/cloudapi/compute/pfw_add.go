package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for add port forward rule
type PFWAddRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// External start port number for the rule
	// Required: true
	PublicPortStart uint64 `url:"publicPortStart" json:"publicPortStart" validate:"required"`

	// End port number (inclusive) for the ranged rule
	// Default value: -1
	// Required: false
	PublicPortEnd int64 `url:"publicPortEnd,omitempty" json:"publicPortEnd,omitempty"`

	// Internal base port number
	// Required: true
	LocalBasePort uint64 `url:"localBasePort" json:"localBasePort" validate:"required"`

	// Network protocol
	// either "tcp" or "udp"
	// Required: true
	Proto string `url:"proto" json:"proto" validate:"proto"`
}

// PFWAdd add port forward rule
func (c Compute) PFWAdd(ctx context.Context, req PFWAddRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/pfwAdd"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
