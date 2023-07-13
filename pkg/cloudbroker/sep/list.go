package sep

import (
	"context"
	"encoding/json"
	"net/http"
)

// Request struct for get list of SEPs
type ListRequest struct {
	// Find by ID
	// Required: false
	ByID uint64 `url:"by_id,omitempty" json:"by_id,omitempty"`

	// Find by name
	// Required: false
	Name string `url:"name,omitempty" json:"name,omitempty"`

	// Find by gId
	// Required: false
	GID uint64 `url:"gId,omitempty" json:"gId,omitempty"`

	// Find by sep type
	// Required: false
	Type string `url:"type,omitempty" json:"type,omitempty"`

	// Find by provided physical node id
	// Required: false
	ProvidedBy uint64 `url:"providedBy,omitempty" json:"providedBy,omitempty"`

	// Find by techStatus
	// Required: false
	TechStatus string `url:"techStatus,omitempty" json:"techStatus,omitempty"`

	// Find by consumed physical node id
	// Required: false
	ConsumedBy uint64 `url:"consumedBy,omitempty" json:"consumedBy,omitempty"`

	// Page size
	// Required: false
	Size uint64 `url:"size,omitempty" json:"size,omitempty"`

	// Page number
	// Required: false
	Page uint64 `url:"page,omitempty" json:"page,omitempty"`
}

// List gets list of SEPs
func (s SEP) List(ctx context.Context, req ListRequest) (*ListSEP, error) {
	url := "/cloudbroker/sep/list"

	res, err := s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return nil, err
	}

	list := ListSEP{}

	err = json.Unmarshal(res, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}
