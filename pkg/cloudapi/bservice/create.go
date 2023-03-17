package bservice

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for BasicService
type CreateRequest struct {
	// Name of the service
	// Required: true
	Name string `url:"name" json:"name" validate:"required"`

	// ID of the Resource Group where this service will be placed
	// Required: true
	RGID uint64 `url:"rgId" json:"rgId" validate:"required"`

	// Name of the user to deploy SSH key for. Pass empty string if no SSH key deployment is required
	// Required: false
	SSHUser string `url:"sshUser,omitempty" json:"sshUser,omitempty"`

	// SSH key to deploy for the specified user. Same key will be deployed to all computes of the service
	// Required: false
	SSHKey string `url:"sshKey,omitempty" json:"sshKey,omitempty"`
}

// Create creates blank BasicService instance
func (b BService) Create(ctx context.Context, req CreateRequest) (uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/bservice/create"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseUint(string(res), 10, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
