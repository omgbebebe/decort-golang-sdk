package flipgroup

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get information about FLIPGroup
type GetRequest struct {
	// FLIPGroup ID
	// Required: true
	FLIPGroupID uint64 `url:"flipgroupId" json:"flipgroupId" validate:"required"`
}

// Get gets details of the specified Floating IP group as a RecordFLIPGroup struct
func (f FLIPGroup) Get(ctx context.Context, req GetRequest) (*RecordFLIPGroup, error) {
	res, err := f.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordFLIPGroup{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets details of the specified Floating IP group as an array of bytes
func (f FLIPGroup) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/flipgroup/get"

	res, err := f.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
