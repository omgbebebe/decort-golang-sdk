package sep

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request  struct for grant access to SEP
type AccessGrantRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id" validate:"required"`

	// Account ID to grant access to the specified SEP. If 0,
	// the SEP will be available for all accounts with no exceptions
	// Required: true
	AccountID uint64 `url:"account_id" json:"account_id" validate:"required"`
}

// AccessGrant grant access to SEP
func (s SEP) AccessGrant(ctx context.Context, req AccessGrantRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/sep/accessGrant"

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
