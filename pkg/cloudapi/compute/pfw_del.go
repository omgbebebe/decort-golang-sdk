package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete port forward rule
type PFWDelRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// ID of the rule to delete. If specified, all other arguments will be ignored
	// Required: false
	PFWID uint64 `url:"ruleId,omitempty" json:"ruleId,omitempty"`

	// External start port number for the rule
	// Required: false
	PublicPortStart uint64 `url:"publicPortStart,omitempty" json:"publicPortStart,omitempty"`

	// End port number (inclusive) for the ranged rule
	// Required: false
	PublicPortEnd uint64 `url:"publicPortEnd,omitempty" json:"publicPortEnd,omitempty"`

	// Internal base port number
	// Required: false
	LocalBasePort uint64 `url:"localBasePort,omitempty" json:"localBasePort,omitempty"`

	// Network protocol
	// either "tcp" or "udp"
	// Required: false
	Proto string `url:"proto,omitempty" json:"proto,omitempty"`
}

// PFWDel delete port forward rule
func (c Compute) PFWDel(ctx context.Context, req PFWDelRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/compute/pfwDel"

	res, err := c.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
