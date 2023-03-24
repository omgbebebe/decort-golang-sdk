package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for rename disk
type RenameRequest struct {
	// ID of the disk to rename
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`

	// New name of disk
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`
}

// Rename rename disk
func (d Disks) Rename(ctx context.Context, req RenameRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/rename"

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
