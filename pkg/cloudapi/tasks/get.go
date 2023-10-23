package tasks

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get information about task
type GetRequest struct {
	// ID of audit
	// Required: true
	AuditID string `url:"auditId" json:"auditId" validate:"required"`
}

// Get gets background API task status and result as a RecordAsyncTask struct
func (t Tasks) Get(ctx context.Context, req GetRequest) (*RecordAsyncTask, error) {
	res, err := t.GetRaw(ctx, req)
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

// GetRaw gets background API task status and result as an array of bytes
func (t Tasks) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/tasks/get"

	res, err := t.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
