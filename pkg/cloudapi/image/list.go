package image

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListRequest struct to get list of available images
type ListRequest struct {
	// Find by storage endpoint provider ID
	// Required: false
	SEPID uint64 `url:"sepId,omitempty" json:"sepId,omitempty"`

	// Find by id
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by ID
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by status
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by architecture
	// Required: false
	Architecture string `url:"architecture,omitempty" json:"architecture,omitempty"`

	// Find by type
	// Required: false
	TypeImage string `url:"typeImage,omitempty" json:"typeImage,omitempty"`

	// Find by image size
	// Required: false
	ImageSize uint64 `url:"imageSize,omitempty" json:"imageSize,omitempty"`

	// Find by SEP name
	// Required: false
	SEPName string `url:"sepName,omitempty" json:"sepName,omitempty"`

	// Find by pool
	// Required: false
	Pool string `url:"pool,omitempty" json:"pool,omitempty"`

	// Find by public True or False
	// Required: false
	Public bool `url:"public,omitempty" json:"public,omitempty"`

	// Find by hot resize True or False
	// Required: false
	HotResize bool `url:"hotResize,omitempty" json:"hotResize,omitempty"`

	// Find by bootable True or False
	// Required: false
	Bootable bool `url:"bootable,omitempty" json:"bootable,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of available images as a ListImages struct, optionally filtering by account ID
func (i Image) List(ctx context.Context, req ListRequest) (*ListImages, error) {
	res, err := i.ListRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	list := ListImages{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// ListRaw gets list of available images as an array of bytes
func (i Image) ListRaw(ctx context.Context, req ListRequest) ([]byte, error) {
	url := "/cloudapi/image/list"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
