package disks

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"repos.digitalenergy.online/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create disk
type CreateRequest struct {
	// ID of the account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId"`

	// ID of the grid (platform)
	// Required: true
	GID uint64 `url:"gid" json:"gid"`

	// Name of disk
	// Required: true
	Name string `url:"name" json:"name"`

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
	Type string `url:"type" json:"type"`

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

func (drq CreateRequest) validate() error {
	if drq.AccountID == 0 {
		return errors.New("validation-error: field AccountID can not be empty or equal to 0")
	}
	if drq.GID == 0 {
		return errors.New("validation-error: field GID can not be empty or equal to 0")
	}
	if drq.Name == "" {
		return errors.New("validation-error: field Name can not be empty")
	}
	validType := validators.StringInSlice(drq.Type, []string{"B", "D", "T"})
	if !validType {
		return errors.New("validation-error: field Type must be set as B, D or T")
	}

	return nil
}

// Create creates a disk
func (d Disks) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := req.validate()
	if err != nil {
		return 0, err
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
