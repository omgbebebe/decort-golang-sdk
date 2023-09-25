package decortsdk

import (
	"bytes"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/k8s"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker"

	"github.com/google/go-querystring/query"
	"repository.basistech.ru/BASIS/decort-golang-sdk/config"
)

// HTTP-client for platform
type DecortClient struct {
	decortURL  string
	client     *http.Client
	cfg        config.Config
	expiryTime time.Time
	mutex      *sync.Mutex
}

// Сlient builder
func New(cfg config.Config) *DecortClient {
	if cfg.Retries == 0 {
		cfg.Retries = 5
	}

	var expiryTime time.Time

	if cfg.Token != "" {
		expiryTime = time.Now().AddDate(0, 0, 1)
	}

	return &DecortClient{
		decortURL: cfg.DecortURL,
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					//nolint:gosec
					InsecureSkipVerify: cfg.SSLSkipVerify,
				},
			},
		},
		cfg:        cfg,
		expiryTime: expiryTime,
		mutex:      &sync.Mutex{},
	}
}

// CloudAPI builder
func (dc *DecortClient) CloudAPI() *cloudapi.CloudAPI {
	return cloudapi.New(dc)
}

// CloudBroker builder
func (dc *DecortClient) CloudBroker() *cloudbroker.CloudBroker {
	return cloudbroker.New(dc)
}

// DecortApiCall method for sending requests to the platform
func (dc *DecortClient) DecortApiCall(ctx context.Context, method, url string, params interface{}) ([]byte, error) {
	if k8sCreateReq, ok := params.(k8s.CreateRequest); ok {
		reqBody := &bytes.Buffer{}
		writer := multipart.NewWriter(reqBody)
		if k8sCreateReq.OidcCertificate != "" {
			part, _ := writer.CreateFormFile("oidcCertificate", "ca.crt")
			_, _ =io.Copy(part, strings.NewReader(k8sCreateReq.OidcCertificate))
		}

		_= writer.WriteField("name", k8sCreateReq.Name)
		_= writer.WriteField("rgId", strconv.FormatUint(k8sCreateReq.RGID, 10))
		_= writer.WriteField("k8ciId", strconv.FormatUint(k8sCreateReq.K8SCIID, 10))
		_= writer.WriteField("workerGroupName", k8sCreateReq.WorkerGroupName)
		_= writer.WriteField("networkPlugin", k8sCreateReq.NetworkPlugin)

		if k8sCreateReq.MasterSEPID != 0 {
			_= writer.WriteField("masterSepId", strconv.FormatUint(k8sCreateReq.MasterSEPID, 10))
		}
		if k8sCreateReq.MasterSEPPool != "" {
			_= writer.WriteField("masterSepPool", k8sCreateReq.MasterSEPPool)
		}
		if k8sCreateReq.WorkerSEPID != 0 {
			_= writer.WriteField("workerSepId", strconv.FormatUint(k8sCreateReq.WorkerSEPID, 10))
		}
		if k8sCreateReq.WorkerSEPPool != "" {
			_= writer.WriteField("workerSepPool", k8sCreateReq.WorkerSEPPool)
		}

		if k8sCreateReq.Labels != nil {
			for _, v := range k8sCreateReq.Labels {
				_= writer.WriteField("labels", v)
			}
		}
		if k8sCreateReq.Taints != nil {
			for _, v := range k8sCreateReq.Taints {
				_= writer.WriteField("taints", v)
			}
		}
		if k8sCreateReq.Annotations != nil {
			for _, v := range k8sCreateReq.Annotations {
				_= writer.WriteField("annotations", v)
			}
		}

		if k8sCreateReq.MasterCPU != 0 {
			_= writer.WriteField("masterCpu", strconv.FormatUint(uint64(k8sCreateReq.MasterCPU), 10))
		}
		if k8sCreateReq.MasterNum != 0 {
			_= writer.WriteField("masterNum", strconv.FormatUint(uint64(k8sCreateReq.MasterNum), 10))
		}
		if k8sCreateReq.MasterRAM != 0 {
			_= writer.WriteField("masterRam", strconv.FormatUint(uint64(k8sCreateReq.MasterRAM), 10))
		}
		if k8sCreateReq.MasterDisk != 0 {
			_= writer.WriteField("masterDisk", strconv.FormatUint(uint64(k8sCreateReq.MasterDisk), 10))
		}
		if k8sCreateReq.WorkerCPU != 0 {
			_= writer.WriteField("workerCpu", strconv.FormatUint(uint64(k8sCreateReq.WorkerCPU), 10))
		}
		if k8sCreateReq.WorkerNum != 0 {
			_= writer.WriteField("workerNum", strconv.FormatUint(uint64(k8sCreateReq.WorkerNum), 10))
		}
		if k8sCreateReq.WorkerRAM != 0 {
			_= writer.WriteField("workerRam", strconv.FormatUint(uint64(k8sCreateReq.WorkerRAM), 10))
		}
		if k8sCreateReq.WorkerDisk != 0 {
			_= writer.WriteField("workerDisk", strconv.FormatUint(uint64(k8sCreateReq.WorkerDisk), 10))
		}
		if k8sCreateReq.ExtNetID != 0 {
			_= writer.WriteField("extnetId", strconv.FormatUint(k8sCreateReq.ExtNetID, 10))
		}
		if k8sCreateReq.VinsId != 0 {
			_= writer.WriteField("vinsId", strconv.FormatUint(k8sCreateReq.VinsId, 10))
		}
		if !k8sCreateReq.WithLB {
			_= writer.WriteField("withLB", strconv.FormatBool(k8sCreateReq.WithLB))
		}

		_= writer.WriteField("highlyAvailable", strconv.FormatBool(k8sCreateReq.HighlyAvailable))

		if k8sCreateReq.AdditionalSANs != nil {
			for _, v := range k8sCreateReq.AdditionalSANs {
				_= writer.WriteField("additionalSANs", v)
			}
		}
		if k8sCreateReq.InitConfiguration != "" {
			_= writer.WriteField("initConfiguration", k8sCreateReq.InitConfiguration)
		}
		if k8sCreateReq.ClusterConfiguration != "" {
			_= writer.WriteField("clusterConfiguration", k8sCreateReq.ClusterConfiguration)
		}
		if k8sCreateReq.KubeletConfiguration != "" {
			_= writer.WriteField("kubeletConfiguration", k8sCreateReq.KubeletConfiguration)
		}
		if k8sCreateReq.KubeProxyConfiguration != "" {
			_= writer.WriteField("kubeProxyConfiguration", k8sCreateReq.KubeProxyConfiguration)
		}
		if k8sCreateReq.JoinConfiguration != "" {
			_= writer.WriteField("joinConfiguration", k8sCreateReq.JoinConfiguration)
		}
		if k8sCreateReq.Description != "" {
			_= writer.WriteField("desc", k8sCreateReq.Description)
		}
		if k8sCreateReq.UserData != "" {
			_= writer.WriteField("userData", k8sCreateReq.UserData)
		}

		_= writer.WriteField("extnetOnly", strconv.FormatBool(k8sCreateReq.ExtNetOnly))
		_= writer.FormDataContentType()

		ct := writer.FormDataContentType()

		writer.Close()
		req, err := http.NewRequestWithContext(ctx, method, dc.decortURL+"/restmachine"+url, reqBody)
		if err != nil {
			return nil, err
		}
		if err = dc.getToken(ctx); err != nil {
			return nil, err
		}

		resp, err := dc.domp(req, ct)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		respBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != 200 {
			return nil, errors.New(string(respBytes))
		}

		return respBytes, nil
	}

	values, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	body := strings.NewReader(values.Encode())
	req, err := http.NewRequestWithContext(ctx, method, dc.decortURL+"/restmachine"+url, body)
	if err != nil {
		return nil, err
	}

	if err = dc.getToken(ctx); err != nil {
		return nil, err
	}

	resp, err := dc.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(respBytes))
	}

	return respBytes, nil
}

func (dc *DecortClient) getToken(ctx context.Context) error {
	dc.mutex.Lock()
	defer dc.mutex.Unlock()

	if dc.cfg.Token == "" || time.Now().After(dc.expiryTime) {
		body := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&response_type=id_token", dc.cfg.AppID, dc.cfg.AppSecret)
		bodyReader := strings.NewReader(body)

		dc.cfg.SSOURL = strings.TrimSuffix(dc.cfg.SSOURL, "/")

		req, _ := http.NewRequestWithContext(ctx, "POST", dc.cfg.SSOURL+"/v1/oauth/access_token", bodyReader)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := dc.client.Do(req)
		if err != nil {
			return fmt.Errorf("cannot get token: %w", err)
		}

		tokenBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != 200 {
			return fmt.Errorf("cannot get token: %s", tokenBytes)
		}

		token := string(tokenBytes)

		dc.cfg.Token = token
		dc.expiryTime = time.Now().AddDate(0, 0, 1)
	}

	return nil
}

func (dc *DecortClient) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "bearer "+dc.cfg.Token)
	req.Header.Set("Accept", "application/json")

	var resp *http.Response
	var err error
	buf, _ := io.ReadAll(req.Body)
	req = req.Clone(req.Context())

	for i := uint64(0); i < dc.cfg.Retries; i++ {
		req.Body = io.NopCloser(bytes.NewBuffer(buf))
		resp, err = dc.client.Do(req)

		// if err == nil {
		if resp.StatusCode == 200 {
			return resp, err
		}
		respBytes, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("%s", respBytes)
		resp.Body.Close()
		// }
	}

	return nil, fmt.Errorf("could not execute request: %w", err)
}

func (dc *DecortClient) domp(req *http.Request, ctype string) (*http.Response, error) {
	req.Header.Add("Content-Type", ctype)
	req.Header.Add("Authorization", "bearer "+dc.cfg.Token)
	req.Header.Set("Accept", "application/json")

	var resp *http.Response
	var err error
	buf, _ := io.ReadAll(req.Body)
	req = req.Clone(req.Context())

	for i := uint64(0); i < dc.cfg.Retries; i++ {
		req.Body = io.NopCloser(bytes.NewBuffer(buf))
		resp, err = dc.client.Do(req)

		// if err == nil {
		if resp.StatusCode == 200 {
			return resp, err
		}
		respBytes, _ := io.ReadAll(resp.Body)
		err = fmt.Errorf("%s", respBytes)
		resp.Body.Close()
		// }
	}

	return nil, fmt.Errorf("could not execute request: %w", err)
}
