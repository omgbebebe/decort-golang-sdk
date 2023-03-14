package rg

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"repos.digitalenergy.online/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for set default network
type SetDefNetRequest struct {
	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId"`

	// Network type
	// Should be one of:
	//	- "PUBLIC"
	//	- "PRIVATE"
	// Required: true
	NetType string `url:"netType" json:"netType"`

	// Network ID
	// Required: false
	NetID uint64 `url:"netId" json:"netId"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

func (rgrq SetDefNetRequest) validate() error {
	if rgrq.RGID == 0 {
		return errors.New("validation-error: field RGID must be set")
	}
	validate := validators.StringInSlice(rgrq.NetType, []string{"PUBLIC", "PRIVATE"})
	if !validate {
		return errors.New("validation-error: field NetType must be one of PRIVATE or PUBLIC")
	}

	return nil
}

// SetDefNet sets default network for attach associated virtual machines
func (r RG) SetDefNet(ctx context.Context, req SetDefNetRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
	}

	url := "/cloudbroker/rg/setDefNet"

	res, err := r.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
