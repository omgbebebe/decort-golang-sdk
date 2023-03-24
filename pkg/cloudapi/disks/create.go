package disks

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create disk
type CreateRequest struct {
	// ID of the account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// ID of the grid (platform)
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`

	// Name of disk
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Description of disk
	// Required: false
	Description string `url:"description,omitempty" json:"description,omitempty"`

	// Size in GB, default is 0
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`

	// Type of disk
	//	- B=Boot
	//	- D=Data
	//	- T=Temp
	// Required: true
	Type string `url:"type" json:"type" validate:"diskType"`

	// Size in GB default is 0
	// Required: false
	SSDSize uint64 `url:"ssdSize,omitempty" json:"ssdSize,omitempty"`

	// Max IOPS disk can perform defaults to 2000
	// Required: false
	IOPS uint64 `url:"iops,omitempty" json:"iops,omitempty"`

	// Storage endpoint provider ID to create disk
	// Required: false
	SEPID uint64 `url:"sep_id,omitempty" json:"sep_id,omitempty"`

	// Pool name to create disk
	// Required: false
	Pool string `url:"pool,omitempty" json:"pool,omitempty"`
}

// Create creates a disk
func (d Disks) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/create"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
