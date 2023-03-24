package sep

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get list of disk IDs
type DiskListRequest struct {
	// Storage endpoint provider ID
	// Required: true
	SEPID uint64 `url:"sep_id" json:"sep_id" validate:"required"`

	// Pool name
	// Required: false
	PoolName string `url:"pool_name,omitempty" json:"pool_name,omitempty"`
}

// DiskList get list of disk IDs, who use this SEP and pool (if provided)
func (s SEP) DiskList(ctx context.Context, req DiskListRequest) ([]uint64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/sep/diskList"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make([]uint64, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
