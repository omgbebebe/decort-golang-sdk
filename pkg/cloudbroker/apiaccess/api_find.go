package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for finding apiaccess groups.
type APIFindRequest struct {
	// API function to find
	// Example: cloudbroker/k8s/create
	// Required: true
	APIName string `url:"apiName" json:"apiName" validate:"required"`
}

// APIFind outputs a list of apiaccess groups that mention the specified API function.
func (a APIAccess) APIFind(ctx context.Context, req APIFindRequest) ([]uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/apiFind"

	list := make([]uint64, 0)

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
