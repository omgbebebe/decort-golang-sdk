package account

import (
	"context"
	"net/http"
	"strconv"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for calculate the currently consumed cloud units of the specified type for all cloudspaces and resource groups in the account
type GetConsumedCloudUnitsByTypeRequest struct {
	// ID an account
	// Required: true
	AccountID uint64 `url:"accountId" json:"accountId" validate:"required"`

	// Cloud unit resource type
	// Required: true
	CUType string `url:"cutype" json:"cutype" validate:"required,accountCUType"`
}

// GetConsumedCloudUnitsByType calculates the currently consumed cloud units of the specified type for all cloudspaces
// and resource groups in the account.
// Possible types of cloud units are include:
//
//   - CU_M: returns consumed memory in MB
//   - CU_C: returns number of virtual cpu cores
//   - CU_D: returns consumed virtual disk storage in GB
//   - CU_S: returns consumed primary storage (NAS) in TB
//   - CU_A: returns consumed secondary storage (Archive) in TB
//   - CU_NO: returns sent/received network transfer in operator in GB
//   - CU_NP: returns sent/received network transfer peering in GB
//   - CU_I: returns number of public IPs
func (a Account) GetConsumedCloudUnitsByType(ctx context.Context, req GetConsumedCloudUnitsByTypeRequest) (float64, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return 0, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/account/getConsumedCloudUnitsByType"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return 0, err
	}

	result, err := strconv.ParseFloat(string(res), 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
