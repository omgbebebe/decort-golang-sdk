package sep

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for exclude consumer nodes
type DelConsumerNodesRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id" validate:"required"`

	// List of consumer node IDs
	// Required: true
	ConsumerNIDs []uint64 `url:"consumer_nids" json:"consumer_nids" validate:"min=1"`
}

// DelConsumerNodes exclude consumer nodes from SEP parameters
func (s SEP) DelConsumerNodes(ctx context.Context, req DelConsumerNodesRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/sep/delConsumerNodes"

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
