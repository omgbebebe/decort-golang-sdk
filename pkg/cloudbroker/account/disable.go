package account

import (
	"context"
	"errors"
	"net/http"
	"strconv"
)

type DisableRequest struct {
	AccountID uint64 `url:"accountId"`
	Reason    string `url:"reason"`
}

func (arq DisableRequest) Validate() error {
	if arq.AccountID == 0 {
		return errors.New("validation-error: field AccountID must be set")
	}
	if arq.Reason == "" {
		return errors.New("validation-error: field Reason must be set")
	}
	return nil
}

func (a Account) Disable(ctx context.Context, req DisableRequest) (bool, error) {
	err := req.Validate()
	if err != nil {
		return false, err
	}

	url := "/cloudbroker/account/disable"

	res, err := a.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return false, err
	}

	result, err := strconv.ParseBool(string(res))
	if err != nil {
		return false, err
	}

	return result, nil
}
