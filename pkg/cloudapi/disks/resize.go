package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for resize disk
type ResizeRequest struct {
	// ID of the disk to resize
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// New size of the disk in GB
	// Required: true
	Size uint64 `url:"size" json:"size" validate:"required"`
}

// Resize resize disk
// Returns 200 if disk is resized online, else will return 202,
// in that case please stop and start your machine after changing the disk size, for your changes to be reflected.
// This method will not be used for disks, assigned to computes. Only unassigned disks and disks, assigned with "old" virtual machines.
func (d Disks) Resize(ctx context.Context, req ResizeRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/resize"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}

// Resize2 resize disk
// Returns 200 if disk is resized online, else will return 202,
// in that case please stop and start your machine after changing the disk size, for your changes to be reflected.
// This method will not be used for disks, assigned to "old" virtual machines. Only unassigned disks and disks, assigned with computes.
func (d Disks) Resize2(ctx context.Context, req ResizeRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/resize2"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
