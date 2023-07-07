package extnet

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create external network
type CreateRequest struct {
	// External network name
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`

	// IP network CIDR
	// For example 192.168.0.0/24
	// Required: true
	IPCIDR string `url:"ipcidr" json:"ipcidr" validate:"required"`

	// VLAN ID
	// Required: true
	VLANID uint64 `url:"vlanId" json:"vlanId" validate:"required"`

	// External network gateway IP address
	// Required: false
	Gateway string `url:"gateway,omitempty" json:"gateway,omitempty"`

	// List of DNS addresses
	// Required: false
	DNS []string `url:"dns,omitempty" json:"dns,omitempty"`

	// List of NTP addresses
	// Required: false
	NTP []string `url:"ntp,omitempty" json:"ntp,omitempty"`

	// IPs to check network availability
	// Required: false
	CheckIPs []string `url:"checkIps,omitempty" json:"checkIps,omitempty"`

	// If true - platform DHCP server will not be created
	// Required: false
	Virtual bool `url:"virtual,omitempty" json:"virtual,omitempty"`

	// Optional description
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`

	// Start of IP range to be explicitly included
	// Required: false
	StartIP string `url:"startIP,omitempty" json:"startIP,omitempty"`

	// End of IP range to be explicitly included
	// Required: false
	EndIP string `url:"endIP,omitempty" json:"endIP,omitempty"`

	// IP to create VNFDev with
	// Required: false
	VNFDevIP string `url:"vnfdevIP,omitempty" json:"vnfdevIP,omitempty"`

	// Number of pre created reservations
	// Required: false
	PreReservationsNum uint64 `url:"preReservationsNum,omitempty" json:"preReservationsNum,omitempty"`

	// OpenvSwith bridge name for ExtNet connection
	// Required: false
	OVSBridge string `url:"ovsBridge,omitempty" json:"ovsBridge,omitempty"`
}

// Create creates new external network into platform
func (e ExtNet) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/create"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
