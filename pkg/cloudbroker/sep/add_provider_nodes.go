package sep

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for add provider nodes
type AddProviderNodesRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id" validate:"required"`

	// List of node IDs
	// Required: true
	ProviderNIDs []uint64 `url:"provider_nids" json:"provider_nids" validate:"min=1"`
}

// AddProviderNodes add provider nodes to SEP parameters
func (s SEP) AddProviderNodes(ctx context.Context, req AddProviderNodesRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/sep/addProviderNodes"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
