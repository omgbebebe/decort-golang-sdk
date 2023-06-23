package user

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting user's audits.
type GetAuditRequest struct {
	// Name of user (get audits for current user if set to empty).
	// Required: false
	Username string `url:"username,omitempty" json:"username,omitempty"`

	// Page number.
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size, maximum - 100.
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// GetAudits gets user's audits.
func (u User) GetAudit(ctx context.Context, req GetAuditRequest) (ListAudits, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/user/getAudit"

	res, err := u.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListAudits{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
