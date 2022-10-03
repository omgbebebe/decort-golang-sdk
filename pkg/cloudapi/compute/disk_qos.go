package compute

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DiskQOSRequest struct {
	ComputeID uint64 `url:"computeId"`
	DiskID    uint64 `url:"diskId"`
	Limits    string `url:"limits"`
}

func (crq DiskQOSRequest) Validate() error {
	if crq.ComputeID == 0 {
		return errors.New("validation-error: field ComputeID can not be empty or equal to 0")
	}

	if crq.DiskID == 0 {
		return errors.New("validation-error: field DiskID can not be empty or equal to 0")
	}

	if crq.Limits == "" {
		return errors.New("validation-error: field Limits can not be empty")
	}

	return nil
}

func (c Compute) DiskQOS(ctx context.Context, req DiskQOSRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudapi/compute/diskQos"

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
