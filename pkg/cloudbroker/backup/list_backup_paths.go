package backup

import (
	"context"
	"encoding/json"
	"net/http"
	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// Request struct for getting list of backup paths
type ListBackupPathsRequest struct {
	// Grid ID
	GID uint64 `url:"gridId" json:"gridId" validate:"required"`
}

// ListBackupPaths gets list of backup paths
func (b Backup) ListBackupPaths(ctx context.Context, req ListBackupPathsRequest) ([]string, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudbroker/backup/listBackupPaths"

	res, err := b.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := make([]string, 0)

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return list, nil
}
