package vins

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for delete NAT rule
type NATRuleDelRequest struct {
	// VINS ID
	// Required: true
	VINSID uint64 `url:"vinsId" json:"vinsId" validate:"required"`

	// ID of the rule to delete.
	// Pass -1 to clear all rules at once
	// Required: true
	RuleID int64 `url:"ruleId" json:"ruleId" validate:"required"`

	// Reason for action
	// Required: false
	Reason string `url:"reason,omitempty" json:"reason,omitempty"`
}

// NATRuleDel delete NAT (port forwarding) rule on VINS
func (v VINS) NATRuleDel(ctx context.Context, req NATRuleDelRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/vins/natRuleDel"

	res, err := v.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
