package account

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for updaate account
type UpdateRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Name of the account
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Max size of memory in MB
	// Required: false
	MaxMemoryCapacity int64 `url:"maxMemoryCapacity,omitempty" json:"maxMemoryCapacity,omitempty"`

	// Max size of aggregated vdisks in GB
	// Required: false
	MaxVDiskCapacity int64 `url:"maxVDiskCapacity,omitempty" json:"maxVDiskCapacity,omitempty"`

	// Max number of CPU cores
	// Required: false
	MaxCPUCapacity int64 `url:"maxCPUCapacity,omitempty" json:"maxCPUCapacity,omitempty"`

	// Max sent/received network transfer peering
	// Required: false
	MaxNetworkPeerTransfer int64 `url:"maxNetworkPeerTransfer,omitempty" json:"maxNetworkPeerTransfer,omitempty"`

	// Max number of assigned public IPs
	// Required: false
	MaxNumPublicIP int64 `url:"maxNumPublicIP,omitempty" json:"maxNumPublicIP,omitempty"`

	// If true send emails when a user is granted access to resources
	// Required: false
	SendAccessEmails bool `url:"sendAccessEmails,omitempty" json:"sendAccessEmails,omitempty"`

	// Limit (positive) or disable (0) GPU resources
	// Required: false
	GPUUnits int64 `url:"gpu_units,omitempty" json:"gpu_units,omitempty"`
}

// Update updates an account name and resource types and limits
func (a Account) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/account/update"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
