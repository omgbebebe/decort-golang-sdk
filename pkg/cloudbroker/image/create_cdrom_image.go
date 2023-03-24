package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create CD-ROM image
type CreateCDROMImageRequest struct {
	// Name of the rescue disk
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// URL where to download ISO from
	// Required: true
	URL string `url:"url" json:"url" validate:"required,url"`

	// Grid (platform) ID where this CD-ROM image should be create in
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`

	// Account ID to make the image exclusive
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Storage endpoint provider ID for place rescue CD
	// Required: false
	SEPID uint64 `url:"sep_id,omitempty" json:"sep_id,omitempty"`

	// Pool for place rescue CD
	// Required: false
	PoolName string `url:"pool_name,omitempty" json:"pool_name,omitempty"`

	// Username for remote media download
	// Required: false
	UsernameDL string `url:"usernameDL,omitempty" json:"usernameDL,omitempty"`

	// Password for remote media download
	// Required: false
	PasswordDl string `url:"passwordDL,omitempty" json:"passwordDL,omitempty"`

	// Binary architecture of this image
	// Should be one of:
	//	- X86_64
	//	- PPC64_LE
	// Required: false
	Architecture string `url:"architecture,omitempty" json:"architecture,omitempty"`

	// List of types of compute suitable for image.
	// Example: [ "KVM_X86" ]
	// Required: true
	Drivers []string `url:"drivers" json:"drivers" validate:"min=1,max=2,imageDrivers"`
}

// CreateCDROMImage creates CD-ROM image from an ISO identified by URL
func (i Image) CreateCDROMImage(ctx context.Context, req CreateCDROMImageRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/createCDROMImage"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
