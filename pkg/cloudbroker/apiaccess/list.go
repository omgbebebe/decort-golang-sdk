package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting list of all non deleted apiaccess instances.
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name apiaccess
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by status apiaccess
	// Required: false
	Status string `url:"status,omitempty" json:"status,omitempty"`

	// Find by created actor
	// Required: false
	CreatedBy string `url:"createdBy,omitempty" json:"createdBy,omitempty"`

	// Find by created after time (unix timestamp)
	// Required: false
	CreatedAfter uint64 `url:"createdAfter,omitempty" json:"createdAfter,omitempty"`

	// Find by created before time (unix timestamp)
	// Required: false
	CreatedBefore uint64 `url:"createdBefore,omitempty" json:"createdBefore,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size, maximum - 100
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// List gets list of all non deleted apiaccess instances.
func (a APIAccess) List(ctx context.Context, req ListRequest) (*ListAPIAccess, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/list"

	info := ListAPIAccess{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
