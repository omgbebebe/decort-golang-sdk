package flipgroup

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for create FLIPGroup
type CreateRequest struct {
	// Account ID
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// FLIPGroup name
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// Network type
	// Should be one of:
	//	- EXTNET
	//	- VINS
	// Required: true
	NetType string `url:"netType" json:"netType" validate:"computeNetType"`

	// ID of external network or VINS
	// Required: true
	NetID uint64 `url:"netId" json:"netId" validate:"required"`

	// Type of client
	//	- 'compute'
	//	- 'vins' (will be later)
	// Required: true
	ClientType string `url:"clientType" json:"clientType" validate:"flipgroupClientType"`

	// IP address to associate with this group. If empty, the platform will autoselect IP address
	// Required: false
	IP string `url:"ip,omitempty" json:"ip,omitempty"`

	// Text description of this FLIPGorup instance
	// Required: false
	Description string `url:"desc,omitempty" json:"desc,omitempty"`
}

// Create method will create a new FLIPGorup in the specified Account
func (f FLIPGroup) Create(ctx context.Context, req CreateRequest) (*RecordFLIPGroup, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/flipgroup/create"

	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordFLIPGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
