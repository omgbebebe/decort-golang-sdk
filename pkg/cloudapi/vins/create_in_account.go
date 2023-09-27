package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create VINS in account
type CreateInAccountRequest struct {
	// VINS name
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// ID of account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Grid ID
	// Required: false
	GID uint64 `url:"gid,omitempty" json:"gid,omitempty"`

	// Private network IP CIDR
	// Required: false
	IPCIDR string `url:"ipcidr,omitempty" json:"ipcidr,omitempty"`

	// Description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Number of pre created reservations
	// Required: false
	PreReservationsNum uint64 `url:"preReservationsNum,omitempty" json:"preReservationsNum,omitempty"`
}

// CreateInAccount creates VINS in account level
func (v VINS) CreateInAccount(ctx context.Context, req CreateInAccountRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/vins/createInAccount"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
