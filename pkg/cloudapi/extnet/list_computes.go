package extnet

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list computes
type ListComputesRequest struct {
	// Filter by account ID
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`
}

// ListComputes gets computes from account with extnets
func (e ExtNet) ListComputes(ctx context.Context, req ListComputesRequest) (ListExtNetComputes, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/extnet/listComputes"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListExtNetComputes{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
