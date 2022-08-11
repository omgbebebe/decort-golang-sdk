package lb

import (
	"context"
	"encoding/json"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type ListRequest struct {
	IncludeDeleted bool   `url:"includedeleted"`
	Page           uint64 `url:"page"`
	Size           uint64 `url:"size"`
}

func (l LB) List(ctx context.Context, req ListRequest, options ...opts.DecortOpts) (LBList, error) {
	url := "/lb/list"
	prefix := "/cloudapi"

	option := opts.New(options)

	if option != nil {
		if option.IsAdmin {
			prefix = "/" + option.AdminValue
		}
	}
	url = prefix + url
	lbListRaw, err := l.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return nil, err
	}

	lbList := LBList{}
	err = json.Unmarshal(lbListRaw, &lbList)
	if err != nil {
		return nil, err
	}

	return lbList, nil

}
