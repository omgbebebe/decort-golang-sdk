package grid

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for change grid settings
type ChangeSettingsRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"id" json:"id" validate:"required"`

	// Json data of the new settings will override old data
	// Required: true
	Settings string `url:"settings" json:"settings" validate:"required"`
}

// ChangeSettings changes grid settings
func (g Grid) ChangeSettings(ctx context.Context, req ChangeSettingsRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/changeSettings"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
