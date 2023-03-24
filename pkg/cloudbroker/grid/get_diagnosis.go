package grid

import (
	"context"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for get platform snapshot with additional diagnosis
type GetDiagnosisRequest struct {
	// Grid (platform) ID
	// Required: true
	GID uint64 `url:"gid" json:"gid" validate:"required"`
}

// GetDiagnosis get platform snapshot with additional diagnosis info like a logs, etc
func (g Grid) GetDiagnosis(ctx context.Context, req GetDiagnosisRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/getDiagnosis"

	res, err := g.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}

// GetDiagnosisGET get platform snapshot with additional diagnosis info like a logs, etc
func (g Grid) GetDiagnosisGET(ctx context.Context, req GetDiagnosisRequest) (string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return "", validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/grid/getDiagnosis"

	res, err := g.client.DecortApiCall(ctx, http.MethodGet, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
