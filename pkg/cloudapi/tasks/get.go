package tasks

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get infromation about task
type GetRequest struct {
	// ID of audit
	// Required: true
	AuditID string `url:"auditId" json:"auditId" validate:"required"`
}

// Get gets background API task status and result
func (t Tasks) Get(ctx context.Context, req GetRequest) (*RecordAsyncTask, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/tasks/get"

	res, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordAsyncTask{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil

}
