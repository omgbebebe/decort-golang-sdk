package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting list of all deleted apiaccess instances.
type ListDeletedRequest struct {
	// Page number.
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`

	// Page size.
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`
}

// ListDeleted gets list of all deleted apiaccess instances.
func (a APIAccess) ListDeleted(ctx context.Context, req ListDeletedRequest) (*ListAPIAccess, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/listDeleted"

	info := ListAPIAccess{}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
