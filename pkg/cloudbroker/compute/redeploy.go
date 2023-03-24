package compute

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for redeploy
type RedeployRequest struct {
	// ID of compute instance
	// Required: true
	ComputeID uint64 `url:"computeId" json:"computeId" validate:"required"`

	// ID of the new OS image, if image change is required
	// Required: false
	ImageID uint64 `url:"imageId,omitempty" json:"imageId,omitempty"`

	// New size for the boot disk in GB, if boot disk size change is required
	// Required: false
	DiskSize uint64 `url:"diskSize,omitempty" json:"diskSize,omitempty"`

	// How to handle data disks connected to this compute instance
	// Should be one of:
	//	- KEEP
	//	- DETACH
	//	- DESTROY
	// Required: false
	DataDisks string `url:"dataDisks,omitempty" json:"dataDisks,omitempty" validate:"omitempty,computeDataDisks"`

	// Should the compute be restarted upon successful redeploy
	// Required: false
	AutoStart bool `url:"autoStart,omitempty" json:"autoStart,omitempty"`

	// Set this flag to True to force stop running compute instance and redeploy next
	// Required: false
	ForceStop bool `url:"forceStop,omitempty" json:"forceStop,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// Redeploy redeploy compute
func (c Compute) Redeploy(ctx context.Context, req RedeployRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/compute/redeploy"

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
