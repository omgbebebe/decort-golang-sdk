package vins

import (
	"context"
	"net/http"
	"strings"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for IP reserve
type IPReserveRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Type of the reservation
	// Should be one of:
	//	- DHCP
	//	- VIP
	//	- EXCLUDE
	// Required: true
	Type string `url:"type" json:"type" validate:"vinsType"`

	// IP address to use. Non-empty string is required for type "EXCLUDE".
	// Ignored for types "DHCP" and "VIP".
	// Required: false
	IPAddr string `url:"ipAddr,omitempty" json:"ipAddr,omitempty"`

	// MAC address to associate with IP reservation.
	// Ignored for type "EXCLUDE",
	// non-empty string is required for "DHCP" and "VIP"
	// Required: false
	MAC string `url:"mac,omitempty" json:"mac,omitempty"`

	// ID of the compute, associated with this reservation of type "DHCP".
	// Ignored for other types
	// Required: false
	ComputeID uint64 `url:"computeId,omitempty" json:"computeId,omitempty"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// IPReserve creates reservation on ViNS DHCP
func (v VINS) IPReserve(ctx context.Context, req IPReserveRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/ipReserve"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(res), "\"", "")

	return result, nil
}
