package lb

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get detailed information about load balancer
type GetRequest struct {
	// ID of the load balancer to get details for
	// Required: true
	LBID uint64 `url:"lbId" json:"lbId" validate:"required"`
}

// Get gets detailed information about load balancer as a RecordLB struct
func (lb LB) Get(ctx context.Context, req GetRequest) (*RecordLB, error) {
	res, err := lb.GetRaw(ctx, req)
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

// GetRaw gets detailed information about load balancer as an array of bytes
func (lb LB) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/lb/get"

	res, err := lb.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
