package lb

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get detailed information about load balancer
type GetRequest struct {
	// ID of the load balancer to get details for
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`
}

// Get gets detailed information about load balancer
func (l LB) Get(ctx context.Context, req GetRequest) (*RecordLB, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/lb/get"

	res, err := l.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordLB{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
