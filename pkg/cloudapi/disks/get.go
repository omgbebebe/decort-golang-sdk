package disks

import (
	"context"
	"encoding/json"
	"net/http"

	"repository.basistech.ru/BASIS/decort-golang-sdk/internal/validators"
)

// GetRequest struct to get information about disk
type GetRequest struct {
	// ID of the disk
	// Required: true
	DiskID uint64 `url:"diskId" json:"diskId" validate:"required"`
}

// Get gets disk details as a RecordDisk struct
// Notice: the devicename field is the name as it is passed to the kernel (kname in linux) for unattached disks this field has no relevant value
func (d Disks) Get(ctx context.Context, req GetRequest) (*RecordDisk, error) {
	res, err := d.GetRaw(ctx, req)
	if err != nil {
		return nil, err
	}

	info := RecordDisk{}

	err = json.Unmarshal(res, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetRaw gets disk details as an array of bytes
// Notice: the devicename field is the name as it is passed to the kernel (kname in linux) for unattached disks this field has no relevant value
func (d Disks) GetRaw(ctx context.Context, req GetRequest) ([]byte, error) {
	err := validators.ValidateRequest(req)
	if err != nil {
		for _, validationError := range validators.GetErrors(err) {
			return nil, validators.ValidationError(validationError)
		}
	}

	url := "/cloudapi/disks/get"

	res, err := d.client.DecortApiCall(ctx, http.MethodPost, url, req)
	return res, err
}
