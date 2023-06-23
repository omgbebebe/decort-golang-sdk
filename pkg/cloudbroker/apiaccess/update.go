package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for updating apis of apiaccess group.
type UpdateRequest struct {
	// APIAccess group ID
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`

	// APIs to remove from APIAccess group
	// Required: false
	APIs APIsEndpoints `url:"-" json:"-"`
}

type wrapperUpdateRequest struct {
	UpdateRequest

	APIString string `url:"apis"`
}

// Update updates apis of apiaccess group.
func (a APIAccess) Update(ctx context.Context, req UpdateRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/update"

	reqWrapped := wrapperUpdateRequest{
		UpdateRequest: req,
	}

	apiJSON, err := json.Marshal(&req.APIs)
	if err != nil {
		return false, err
	}

	reqWrapped.APIString = string(apiJSON)

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
