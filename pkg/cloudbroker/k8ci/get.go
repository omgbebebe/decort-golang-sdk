package k8ci

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get information about K8CI
type GetRequest struct {
	// ID of the K8 catalog item to get
	// Required: true
	K8CIID uint64 `url:"k8ciId" json:"k8ciId" validate:"required"`
}

// Get gets details of the specified K8 catalog item as a RecordK8CI struct
func (k K8CI) Get(ctx context.Context, req GetRequest) (*RecordK8CI, error) {
	res, err := k.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	item := RecordK8CI{}

	err = json.Unmarshal(res, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

// GetRaw gets details of the specified K8 catalog item as an array of bytes
func (k K8CI) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/k8ci/get"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
