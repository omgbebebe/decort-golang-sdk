package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create VINS in resource group
type CreateInRGRequest struct {
	// VINS name
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Resource group ID
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Private network IP CIDR
	// Required: false
	IPCIDR string `url:"ipcidr,omitempty" json:"ipcidr,omitempty"`

	// External network ID
	// Required: false
	ExtNetID uint64 `url:"extNetId,omitempty" json:"extNetId,omitempty"`

	// External IP, related only for extNetId >= 0
	// Required: false
	ExtIP string `url:"extIp,omitempty" json:"extIp,omitempty"`

	// Description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Number of pre created reservations
	// Required: false
	PreReservationsNum uint `url:"preReservationsNum,omitempty" json:"preReservationsNum,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// CreateInRG creates VINS in resource group level
func (v VINS) CreateInRG(ctx context.Context, req CreateInRGRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/createInRG"

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
