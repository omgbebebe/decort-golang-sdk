package vins

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

type Route struct {
	// Destination network
	Destination string `url:"destination" json:"destination" validate:"required"`

	//Destination network mask in 255.255.255.255 format
	Netmask string `url:"netmask" json:"netmask" validate:"required"`

	//Next hop host, IP address from ViNS ID free IP pool
	Gateway string `url:"gateway" json:"gateway" validate:"required"`
}

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

	// List of static routes, each item must have destination, netmask, and gateway fields
	// Required: false
	Routes []Route `url:"-" json:"routes,omitempty" validate:"omitempty,dive"`
}

type wrapperCreateRequestInAcc struct {
	CreateInAccountRequest
	Routes []string `url:"routes,omitempty"`
}

// CreateInAccount creates VINS in account level
func (v VINS) CreateInAccount(ctx context.Context, req CreateInAccountRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	var routes []string

	if len(req.Routes) != 0 {
		routes = make([]string, 0, len(req.Routes))

		for r := range req.Routes {
			b, err := json.Marshal(req.Routes[r])
			if err != nil {
				return 0, err
			}

			routes = append(routes, string(b))
		}
	} else {
		routes = []string{}
	}

	reqWrapped := wrapperCreateRequestInAcc{
		CreateInAccountRequest: req,
		Routes:                 routes,
	}

	url := "/cloudapi/vins/createInAccount"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
