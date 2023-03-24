package sep

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for grant access to pool SEP
type AccessGrantToPoolRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id" validate:"required"`

	// Pool name
	// Required: true
	PoolName string `url:"pool_name" json:"pool_name" validate:"required"`

	// Account ID to grant access to the specified pool SEP
	// Required: false
	AccountID uint64 `url:"account_id,omitempty" json:"account_id,omitempty"`

	// Resource group to grant access to the specified pool SEP
	// Required: false
	RGID uint64 `url:"resgroup_id,omitempty" json:"resgroup_id,omitempty"`
}

// AccessGrantToPool grant access to pool SEP
func (s SEP) AccessGrantToPool(ctx context.Context, req AccessGrantToPoolRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/sep/accessGrantToPool"

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
