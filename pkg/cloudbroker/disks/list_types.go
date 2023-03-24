package disks

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list types of disks
type ListTypesRequest struct {
	// Show detailed disk types by seps
	// Required: true
	Detailed bool `url:"detailed" json:"detailed" validate:"required"`
}

// ListTypes gets list defined disk types
func (d Disks) ListTypes(ctx context.Context, req ListTypesRequest) ([]interface{}, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/disks/listTypes"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make([]interface{}, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil

}
