package apiaccess

import (
	"encoding/json"
	"net/http"

	"context"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for subtracting.
type SubtractRequest struct {
	// ID of the API access group to subtract from. This group will contain the difference.
	MinuendID uint64 `url:"minuendId" json:"minuendId" validate:"required"`

	// ID of the API access group which is subtracted. This group is unchanged.
	SubtrahendID uint64 `url:"subtrahendId" json:"subtrahendId" validate:"required"`
}

// Subtract removes such APIs from MinuendID that match APIs from SubtrahendID.
func (a APIAccess) Subtruct(ctx context.Context, req SubtractRequest) (*APIsEndpoints, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/apiaccess/subtract"

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
