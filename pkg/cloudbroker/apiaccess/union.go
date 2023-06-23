package apiaccess

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for union.
type UnionRequest struct {
	// Recipient apiaccess group ID
	// Required: true
	RecipientID uint64 `url:"recipientId" json:"recipientId" validate:"required"`

	// Donor apiaccess group ID
	// Required: true
	DonorID uint64 `url:"donorId" json:"donorId" validate:"required"`
}

// Combines the API list of group #1 ("recipient") and group #2 ("donor"),
// writing the result to group #1 and avoiding duplicates in the list
func (a APIAccess) Union(ctx context.Context, req UnionRequest) (*APIsEndpoints, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/union"

	info := APIsEndpoints{}

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
