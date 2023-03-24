package image

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for sync create image
type SyncCreateRequest struct {
	// Name of the rescue disk
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// URL where to download media from
	// Required: true
	URL string `url:"url" json:"url" validate:"required"`

	// Grid (platform) ID where this template should be create in
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`

	// Boot type of image
	// Should be one of:
	//	- bios
	//	- UEFI
	// Required: true
	BootType string `url:"boottype" json:"boottype" validate:"imageBootType"`

	// Image type
	// Should be one of:
	//	- linux
	//	- windows
	//	- or other
	// Required: true
	ImageType string `url:"imagetype" json:"imagetype" validate:"imageType"`

	// Does this machine supports hot resize
	// Required: false
	HotResize bool `url:"hotresize,omitempty" json:"hotresize,omitempty"`

	// Optional username for the image
	// Required: false
	Username string `url:"username,omitempty" json:"username,omitempty"`

	// Optional password for the image
	// Required: false
	Password string `url:"password,omitempty" json:"password,omitempty"`

	// Account ID to make the image exclusive
	// Required: false
	AccountID uint64 `url:"accountId,omitempty" json:"accountId,omitempty"`

	// Username for upload binary media
	// Required: false
	UsernameDL string `url:"usernameDL,omitempty" json:"usernameDL,omitempty"`

	// Password for upload binary media
	// Required: false
	PasswordDL string `url:"passwordDL,omitempty" json:"passwordDL,omitempty"`

	// Storage endpoint provider ID
	// Required: false
	SEPID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Pool for image create
	// Required: false
	PoolName string `url:"poolName,omitempty" json:"poolName,omitempty"`

	// Binary architecture of this image
	// Should be one of:
	//	- X86_64
	//	- PPC64_LE
	// Required: false
	Architecture string `url:"architecture,omitempty" json:"architecture,omitempty"`

	// List of types of compute suitable for image
	// Example: [ "KVM_X86" ]
	// Required: true
	Drivers []string `url:"drivers" json:"drivers" validate:"min=1,max=2,imageDrivers"`

	// Bootable image or not
	// Required: false
	Bootable bool `url:"bootable,omitempty" json:"bootable,omitempty"`
}

// SyncCreate creates image from a media identified by URL (in synchronous mode)
func (i Image) SyncCreate(ctx context.Context, req SyncCreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/image/syncCreateImage"

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
