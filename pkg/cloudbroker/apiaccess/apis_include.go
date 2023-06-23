package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for adding api to access group.
type APIsIncludeRequest struct {
	// APIAccess group ID.
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`

	// APIs to add to APIAccess group.
	// Required: true
	APIs APIsEndpoints `url:"-" json:"-" validate:"required"`
}

type wrapperAPIsIncludeRequest struct {
	APIsIncludeRequest

	APIString string `url:"apis"`
}

// APIsInclude adds the listed API functions to the specified apiaccess group.
func (a APIAccess) APIsInclude(ctx context.Context, req APIsIncludeRequest) (*APIsEndpoints, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/apisInclude"

	info := APIsEndpoints{}

	apiJSON, err := json.Marshal(&req.APIs)
	if err != nil {
		return nil, err
	}

	reqWrapped := wrapperAPIsIncludeRequest{
		APIsIncludeRequest: req,
		APIString:          string(apiJSON),
	}

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, reqWrapped)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
