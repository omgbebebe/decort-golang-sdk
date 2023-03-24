package grid

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create system space
type CreateSystemSpaceRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"id" json:"id" validate:"required"`

	// Name of the account/cloudspace to be created for the system
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// ID of the specific image
	// Required: true
	ImageID uint64 `url:"imageId" json:"imageId" validate:"required"`

	// Size of base volume
	// Required: true
	BootSize uint64 `url:"bootsize" json:"bootsize" validate:"required"`

	// Data disk size in gigabytes
	// Required: true
	DataDiskSize uint64 `url:"dataDiskSize" json:"dataDiskSize" validate:"required"`

	// ID of the specific size
	// Required: false
	SizeID uint64 `url:"sizeId,omitempty" json:"sizeId,omitempty"`

	// Number of vcpus to provide
	// Required: false
	VCPUS uint64 `url:"vcpus,omitempty" json:"vcpus,omitempty"`

	// Amount of memory to provide
	// Required: false
	Memory uint64 `url:"memory,omitempty" json:"memory,omitempty"`
}

// CreateSystemSpace creates system space
func (g Grid) CreateSystemSpace(ctx context.Context, req CreateSystemSpaceRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/createSystemSpace"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
