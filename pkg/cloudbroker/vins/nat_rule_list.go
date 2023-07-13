package vins

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list of NAT rules
type NATRuleListRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// NATRuleList gets list of NAT (port forwarding) rules
func (v VINS) NATRuleList(ctx context.Context, req NATRuleListRequest) (*ListNATRules, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/natRuleList"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListNATRules{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
