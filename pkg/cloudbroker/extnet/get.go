package extnet

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get information about external network
type GetRequest struct {
	// ID of external network
	// Required: true
	NetID uint64 `url:"net_id" json:"net_id" validate:"required"`
}

// Get gets external network details as a RecordExtNet struct
func (e ExtNet) Get(ctx context.Context, req GetRequest) (*RecordExtNet, error) {
	res, err := e.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordExtNet{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets external network details as an array of bytes
func (e ExtNet) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/extnet/get"

	res, err := e.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
