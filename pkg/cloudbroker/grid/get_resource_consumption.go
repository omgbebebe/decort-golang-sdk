package grid

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

type GetResourceConsumptionRequest struct {
	// ID of the grid
	// Required: true
	GridID uint64 `url:"gridId" json:"gridId" validate:"required"`
}

func (g Grid) GetResourceConsumption(ctx context.Context, req GetResourceConsumptionRequest) (*RecordResourcesConsumption, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/getResourceConsumption"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	result := RecordResourcesConsumption{}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
