package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/rudecs/decort-sdk/internal/validators"
)

type PFWAddRequest struct {
	ComputeID       uint64 `url:"computeId"`
	PublicPortStart uint64 `url:"publicPortStart"`
	PublicPortEnd   uint64 `url:"publicPortEnd,omitempty"`
	LocalBasePort   uint64 `url:"localBasePort"`
	Proto           string `url:"proto"`
}

func (crq PFWAddRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}
	if crq.PublicPortStart == 0 {
		return errors.New("validation-error: field PublicPortStart can not be empty or equal to 0")
	}
	if crq.LocalBasePort == 0 {
		return errors.New("validation-error: field LocalBasePort can not be empty or equal to 0")
	}
	if crq.Proto == "" {
		return errors.New("validation-error: field Proto can not be empty")
	}
	validator := validators.StringInSlice(crq.Proto, []string{"tcp", "udp"})
	if !validator {
		return errors.New("validation-error: field Proto can be only tcp or udp")
	}

	return nil
}

func (c Compute) PFWAdd(ctx context.Context, req PFWAddRequest) (uint64, error) {
	err := req.Validate()
	if err != nil {
		return 0, err
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
