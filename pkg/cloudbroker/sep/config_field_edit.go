package sep

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for edit config fields
type ConfigFieldEditRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id" validate:"required"`

	// Field name
	// Required: true
	FieldName string `url:"field_name" json:"field_name" validate:"required"`

	// Field value
	// Required: true
	FieldValue string `url:"field_value" json:"field_value" validate:"required"`

	// Field type
	// Should be one of:
	//	- int
	//	- str
	//	- bool
	//	- list
	//	- dict
	// Required: true
	FieldType string `url:"field_type" json:"field_type" validate:"sepFieldType"`
}

// ConfigFieldEdit edit SEP config field value
func (s SEP) ConfigFieldEdit(ctx context.Context, req ConfigFieldEditRequest) (bool, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return false, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/sep/configFieldEdit"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
