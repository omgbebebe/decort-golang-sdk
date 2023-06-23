package apiaccess

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for removing api from access group.
type APIsExcludeRequest struct {
	// APIAccess group ID
	// Required: true
	APIAccessID uint64 `url:"apiaccessId" json:"apiaccessId" validate:"required"`

	// APIs to remove from APIAccess group
	// Required: true
	APIs APIsEndpoints `url:"-" json:"-" validate:"required"`
}

type wrapperAPIsExcludeRequest struct {
	APIsExcludeRequest

	APIString string `url:"apis"`
}

// APIsExclude removes the listed API functions from the specified apiaccess group.
func (a APIAccess) APIsExclude(ctx context.Context, req APIsExcludeRequest) (*APIsEndpoints, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/apisExclude"

	info := APIsEndpoints{}

	apiJSON, err := json.Marshal(&req.APIs)
	if err != nil {
		return nil, err
	}

	apiJSONPretty, err := json.MarshalIndent(&req.APIs, "", "    ")
	if err != nil {
		return nil, err
	}
	fmt.Println(string(apiJSONPretty))

	reqWrapped := wrapperAPIsExcludeRequest{
		APIsExcludeRequest: req,
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
