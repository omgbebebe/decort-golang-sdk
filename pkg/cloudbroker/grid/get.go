package grid

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get grid details
type GetRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gridId" json:"gridId" validate:"required"`
}

// Get gets information about grid by ID as a RecordGrid struct
func (g Grid) Get(ctx context.Context, req GetRequest) (*RecordGrid, error) {
	res, err := g.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordGrid{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets information about grid by ID as an array of bytes
func (g Grid) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/get"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
