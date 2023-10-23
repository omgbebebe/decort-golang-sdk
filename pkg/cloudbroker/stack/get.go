package stack

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get list of stacks
type GetRequest struct {
	// Find by ID
	// Required: true
	StackId uint64 `url:"stackId" json:"stackId" validate:"required"`
}

// Get gets stack details by ID as an InfoStack struct
func (i Stack) Get(ctx context.Context, req GetRequest) (*InfoStack, error) {
	res, err := i.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := InfoStack{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets stack details by ID as an array of bytes
func (i Stack) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/stack/get"

	res, err := i.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
