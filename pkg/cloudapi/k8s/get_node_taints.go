package k8s

import (
	"context"
	"errors"
	"net/http"
)

type GetNodeTaintsRequest struct {
	K8SID  uint64 `url:"k8sId"`
	NodeID uint64 `url:"nodeId"`
}

func (krq GetNodeTaintsRequest) Validate() error {
	if krq.K8SID == 0 {
		return errors.New("validation-error: field K8SID can not be empty or equal to 0")
	}
	if krq.NodeID == 0 {
		return errors.New("validation-error: field NodeID can not be empty or equal to 0")
	}

	return nil
}

func (k8s K8S) GetNodeTaints(ctx context.Context, req GetNodeTaintsRequest) (string, error) {
	err := req.Validate()
	if err != nil {
		return "", err
	}

	url := "/cloudapi/k8s/getNodeTaints"

	res, err := k8s.client.DecortApiCall(ctx, http.MethodPost, url, req)
	if err != nil {
		return "", err
	}

	return string(res), nil
}
