package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for unshare data disk
type UnshareRequest struct {
	// ID of the disk to unshare
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`
}

// Unshare unshares data disk
func (d Disks) Unshare(ctx context.Context, req UnshareRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/disks/unshare"

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
