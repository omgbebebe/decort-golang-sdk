package sep

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create SEP object
type CreateRequest struct {
	// Grid ID
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`

	// SEP name
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Type of storage
	// Required: true
	SEPType string `url:"sep_type" json:"sep_type" validate:"required"`

	// Description
	// Required: false
	Description string `url:"description,omitempty" json:"description,omitempty"`

	// SEP config
	// Required: false
	Config string `url:"config,omitempty" json:"config,omitempty"`

	// List of provider node IDs
	// Required: false
	ProviderNIDs []uint64 `url:"provider_nids,omitempty" json:"provider_nids,omitempty"`

	// List of consumer node IDs
	// Required: false
	ConsumerNIDs []uint64 `url:"consumer_nids,omitempty" json:"consumer_nids,omitempty"`

	// Enable SEP after creation
	// Required: false
	Enable bool `url:"enable,omitempty" json:"enable,omitempty"`
}

// Create creates SEP object
func (s SEP) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/sep/create"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
