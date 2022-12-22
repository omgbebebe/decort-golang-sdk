package extnet

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

// Request struct for include list IPs
type IPsIncludeRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id"`

	// List of IPs for include to external network
	// Required: true
	IPs []string `url:"ips"`
}

func (erq IPsIncludeRequest) validate() error {
	if erq.NetID == 0 {
		return errors.New("validation-error: field NetID must be set")
	}
	if len(erq.IPs) == 0 {
		return errors.New("validation-error: field IPs must be set")
	}

	return nil
}

// IPsInclude include list IPs to external network pool
func (e ExtNet) IPsInclude(ctx context.Context, req IPsIncludeRequest) (bool, error) {
	err := req.validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/extnet/ipsInclude"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
