package tasks

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get background API task status and result
type GetRequest struct {
	// ID of audit GUID
	// Required: true
	AuditID string `url:"auditId" json:"auditId" validate:"required"`
}

// Get gets background API task status and result
func (t Tasks) Get(ctx context.Context, req GetRequest) (*RecordTask, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/tasks/get"

	res, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	item := RecordTask{}

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}
