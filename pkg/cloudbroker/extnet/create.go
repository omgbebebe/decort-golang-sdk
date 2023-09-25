package extnet

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

	// List of static routes, each item must have destination, netmask, and gateway fields
	// Required: false
	Routes []Route `url:"-" json:"routes,omitempty" validate:"omitempty,dive"`
}

type wrapperCreateRequest struct {
	CreateRequest
	Routes []string `url:"routes,omitempty"`
}

// Create creates new external network into platform
func (e ExtNet) Create(ctx context.Context, req CreateRequest) (uint64, error) {
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

	reqWrapped := wrapperCreateRequest{
		CreateRequest: req,
		Routes:        routes,
	}

	url := "/cloudbroker/extnet/create"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
