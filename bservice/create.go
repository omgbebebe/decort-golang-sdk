package bservice

import (
	"context"
	"errors"
	"strconv"

	"github.com/rudecs/decort-sdk/opts"
	"github.com/rudecs/decort-sdk/typed"
)

type CreateRequest struct {
	Name    string `url:"name"`
	RGID    uint64 `url:"rgId"`
	SSHUser string `url:"sshUser,omitempty"`
	SSHKey  string `url:"sshKey,omitempty"`
}

func (bsrq CreateRequest) Validate() error {
	if bsrq.Name == "" {
		return errors.New("field Name can not be empty")
	}

	if bsrq.RGID == 0 {
		return errors.New("field RGID can not be empty or equal to 0")
	}

	return nil
}

func (b BService) Create(ctx context.Context, req CreateRequest, options ...opts.DecortOpts) (uint64, error) {
	if err := req.Validate(); err != nil {
		return 0, err
	}

	url := "/cloudapi/bservice/create"
	res, err := b.client.DecortApiCall(ctx, typed.POST, url, req)
	if err != nil {
		return 0, err
	}

	return strconv.ParseUint(string(res), 10, 64)
}
