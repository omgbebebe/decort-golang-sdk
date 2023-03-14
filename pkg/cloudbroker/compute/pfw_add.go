package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"repos.digitalenergy.online/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for add port forward rule
type PFWAddRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId"`

	// External start port number for the rule
	// Required: true
	PublicPortStart uint64 `url:"publicPortStart" json:"publicPortStart"`

	// End port number (inclusive) for the ranged rule
	// Required: false
	PublicPortEnd uint64 `url:"publicPortEnd,omitempty" json:"publicPortEnd,omitempty"`

	// Internal base port number
	// Required: true
	LocalBasePort uint64 `url:"localBasePort" json:"localBasePort"`

	// Network protocol
	// Should be one of:
	//	- tcp
	//	- udp
	// Required: true
	Proto string `url:"proto" json:"proto"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (crq PFWAddRequest) validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID must be set")
	}
	if crq.PublicPortStart == 0 {
		return errors.New("validation-error: field PublicPortStart must be set")
	}
	if crq.LocalBasePort == 0 {
		return errors.New("validation-error: field LocalBasePort must be set")
	}
	if crq.Proto == "" {
		return errors.New("validation-error: field Proto must be set")
	}
	validate := validators.StringInSlice(crq.Proto, []string{"tcp", "udp"})
	if !validate {
		return errors.New("validation-error: field Proto must be tcp or udp")
	}

	return nil
}

// PFWAdd add port forward rule
func (c Compute) PFWAdd(ctx context.Context, req PFWAddRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/compute/pfwAdd"

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
