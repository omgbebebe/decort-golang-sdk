package vins

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for IP relese
type IPReleaseRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId"`

	// IP address
	// Required: false
	IPAddr string `url:"ipAddr,omitempty"`

	// MAC address
	// Required: false
	MAC string `url:"mac,omitempty"`
}

func (vrq IPReleaseRequest) validate() error {
	if vrq.VINSID == 0 {
		return errors.New("validation-error: field VINSID must be set")
	}

	return nil
}

// IPRelese delete IP reservation matched by specified IP & MAC address combination.
// If both IP and MAC address are empty strings, all IP reservations will be deleted.
func (v VINS) IPRelease(ctx context.Context, req IPReleaseRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/vins/ipRelease"

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
