package k8ci

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get information about K8CI
type GetRequest struct {
	// ID of the K8 catalog item to get
	// Required: true
	K8CIID uint64 `url:"k8ciId" json:"k8ciId" validate:"required"`
}

// Get gets details of the specified K8 catalog item
func (k K8CI) Get(ctx context.Context, req GetRequest) (*RecordK8CI, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/k8ci/get"

	res, err := k.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	info := RecordK8CI{}
	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
