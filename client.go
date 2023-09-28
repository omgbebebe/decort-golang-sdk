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
	k8s_ca "repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudapi/k8s"
	"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker"
	k8s_cb"repository.basistech.ru/BASIS/decort-golang-sdk/pkg/cloudbroker/k8s"
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
	k8sCaCreateReq, okCa := params.(k8s_ca.CreateRequest)
	k8sCbCreateReq, okCb := params.(k8s_cb.CreateRequest)

	if okCa {
		reqBody := &bytes.Buffer{}
		writer := multipart.NewWriter(reqBody)
		if k8sCaCreateReq.OidcCertificate != "" {
			part, _ := writer.CreateFormFile("oidcCertificate", "ca.crt")
			_, _ = io.Copy(part, strings.NewReader(k8sCaCreateReq.OidcCertificate))
		}

		_ = writer.WriteField("name", k8sCaCreateReq.Name)
		_ = writer.WriteField("rgId", strconv.FormatUint(k8sCaCreateReq.RGID, 10))
		_ = writer.WriteField("k8ciId", strconv.FormatUint(k8sCaCreateReq.K8SCIID, 10))
		_ = writer.WriteField("workerGroupName", k8sCaCreateReq.WorkerGroupName)
		_ = writer.WriteField("networkPlugin", k8sCaCreateReq.NetworkPlugin)

		if k8sCaCreateReq.MasterSEPID != 0 {
			_ = writer.WriteField("masterSepId", strconv.FormatUint(k8sCaCreateReq.MasterSEPID, 10))
		}
		if k8sCaCreateReq.MasterSEPPool != "" {
			_ = writer.WriteField("masterSepPool", k8sCaCreateReq.MasterSEPPool)
		}
		if k8sCaCreateReq.WorkerSEPID != 0 {
			_ = writer.WriteField("workerSepId", strconv.FormatUint(k8sCaCreateReq.WorkerSEPID, 10))
		}
		if k8sCaCreateReq.WorkerSEPPool != "" {
			_ = writer.WriteField("workerSepPool", k8sCaCreateReq.WorkerSEPPool)
		}

		if k8sCaCreateReq.Labels != nil {
			for _, v := range k8sCaCreateReq.Labels {
				_ = writer.WriteField("labels", v)
			}
		}
		if k8sCaCreateReq.Taints != nil {
			for _, v := range k8sCaCreateReq.Taints {
				_ = writer.WriteField("taints", v)
			}
		}
		if k8sCaCreateReq.Annotations != nil {
			for _, v := range k8sCaCreateReq.Annotations {
				_ = writer.WriteField("annotations", v)
			}
		}

		if k8sCaCreateReq.MasterCPU != 0 {
			_ = writer.WriteField("masterCpu", strconv.FormatUint(uint64(k8sCaCreateReq.MasterCPU), 10))
		}
		if k8sCaCreateReq.MasterNum != 0 {
			_ = writer.WriteField("masterNum", strconv.FormatUint(uint64(k8sCaCreateReq.MasterNum), 10))
		}
		if k8sCaCreateReq.MasterRAM != 0 {
			_ = writer.WriteField("masterRam", strconv.FormatUint(uint64(k8sCaCreateReq.MasterRAM), 10))
		}
		if k8sCaCreateReq.MasterDisk != 0 {
			_ = writer.WriteField("masterDisk", strconv.FormatUint(uint64(k8sCaCreateReq.MasterDisk), 10))
		}
		if k8sCaCreateReq.WorkerCPU != 0 {
			_ = writer.WriteField("workerCpu", strconv.FormatUint(uint64(k8sCaCreateReq.WorkerCPU), 10))
		}
		if k8sCaCreateReq.WorkerNum != 0 {
			_ = writer.WriteField("workerNum", strconv.FormatUint(uint64(k8sCaCreateReq.WorkerNum), 10))
		}
		if k8sCaCreateReq.WorkerRAM != 0 {
			_ = writer.WriteField("workerRam", strconv.FormatUint(uint64(k8sCaCreateReq.WorkerRAM), 10))
		}
		if k8sCaCreateReq.WorkerDisk != 0 {
			_ = writer.WriteField("workerDisk", strconv.FormatUint(uint64(k8sCaCreateReq.WorkerDisk), 10))
		}
		if k8sCaCreateReq.ExtNetID != 0 {
			_ = writer.WriteField("extnetId", strconv.FormatUint(k8sCaCreateReq.ExtNetID, 10))
		}
		if k8sCaCreateReq.VinsId != 0 {
			_ = writer.WriteField("vinsId", strconv.FormatUint(k8sCaCreateReq.VinsId, 10))
		}
		if !k8sCaCreateReq.WithLB {
			_ = writer.WriteField("withLB", strconv.FormatBool(k8sCaCreateReq.WithLB))
		}

		_ = writer.WriteField("highlyAvailableLB", strconv.FormatBool(k8sCaCreateReq.HighlyAvailable))

		if k8sCaCreateReq.AdditionalSANs != nil {
			for _, v := range k8sCaCreateReq.AdditionalSANs {
				_ = writer.WriteField("additionalSANs", v)
			}
		}
		if k8sCaCreateReq.InitConfiguration != "" {
			_ = writer.WriteField("initConfiguration", k8sCaCreateReq.InitConfiguration)
		}
		if k8sCaCreateReq.ClusterConfiguration != "" {
			_ = writer.WriteField("clusterConfiguration", k8sCaCreateReq.ClusterConfiguration)
		}
		if k8sCaCreateReq.KubeletConfiguration != "" {
			_ = writer.WriteField("kubeletConfiguration", k8sCaCreateReq.KubeletConfiguration)
		}
		if k8sCaCreateReq.KubeProxyConfiguration != "" {
			_ = writer.WriteField("kubeProxyConfiguration", k8sCaCreateReq.KubeProxyConfiguration)
		}
		if k8sCaCreateReq.JoinConfiguration != "" {
			_ = writer.WriteField("joinConfiguration", k8sCaCreateReq.JoinConfiguration)
		}
		if k8sCaCreateReq.Description != "" {
			_ = writer.WriteField("desc", k8sCaCreateReq.Description)
		}
		if k8sCaCreateReq.UserData != "" {
			_ = writer.WriteField("userData", k8sCaCreateReq.UserData)
		}

		_ = writer.WriteField("extnetOnly", strconv.FormatBool(k8sCaCreateReq.ExtNetOnly))
		_ = writer.FormDataContentType()

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
	} else if okCb {
		reqBody := &bytes.Buffer{}
		writer := multipart.NewWriter(reqBody)
		if k8sCbCreateReq.OidcCertificate != "" {
			part, _ := writer.CreateFormFile("oidcCertificate", "ca.crt")
			_, _ = io.Copy(part, strings.NewReader(k8sCbCreateReq.OidcCertificate))
		}

		_ = writer.WriteField("name", k8sCbCreateReq.Name)
		_ = writer.WriteField("rgId", strconv.FormatUint(k8sCbCreateReq.RGID, 10))
		_ = writer.WriteField("k8ciId", strconv.FormatUint(k8sCbCreateReq.K8CIID, 10))
		_ = writer.WriteField("workerGroupName", k8sCbCreateReq.WorkerGroupName)
		_ = writer.WriteField("networkPlugin", k8sCbCreateReq.NetworkPlugin)

		if k8sCbCreateReq.MasterSEPID != 0 {
			_ = writer.WriteField("masterSepId", strconv.FormatUint(k8sCbCreateReq.MasterSEPID, 10))
		}
		if k8sCbCreateReq.MasterSEPPool != "" {
			_ = writer.WriteField("masterSepPool", k8sCbCreateReq.MasterSEPPool)
		}
		if k8sCbCreateReq.WorkerSEPID != 0 {
			_ = writer.WriteField("workerSepId", strconv.FormatUint(k8sCbCreateReq.WorkerSEPID, 10))
		}
		if k8sCbCreateReq.WorkerSEPPool != "" {
			_ = writer.WriteField("workerSepPool", k8sCbCreateReq.WorkerSEPPool)
		}

		if k8sCbCreateReq.Labels != nil {
			for _, v := range k8sCbCreateReq.Labels {
				_ = writer.WriteField("labels", v)
			}
		}
		if k8sCbCreateReq.Taints != nil {
			for _, v := range k8sCbCreateReq.Taints {
				_ = writer.WriteField("taints", v)
			}
		}
		if k8sCbCreateReq.Annotations != nil {
			for _, v := range k8sCbCreateReq.Annotations {
				_ = writer.WriteField("annotations", v)
			}
		}

		if k8sCbCreateReq.MasterCPU != 0 {
			_ = writer.WriteField("masterCpu", strconv.FormatUint(k8sCbCreateReq.MasterCPU, 10))
		}
		if k8sCbCreateReq.MasterNum != 0 {
			_ = writer.WriteField("masterNum", strconv.FormatUint(k8sCbCreateReq.MasterNum, 10))
		}
		if k8sCbCreateReq.MasterRAM != 0 {
			_ = writer.WriteField("masterRam", strconv.FormatUint(k8sCbCreateReq.MasterRAM, 10))
		}
		if k8sCbCreateReq.MasterDisk != 0 {
			_ = writer.WriteField("masterDisk", strconv.FormatUint(k8sCbCreateReq.MasterDisk, 10))
		}
		if k8sCbCreateReq.WorkerCPU != 0 {
			_ = writer.WriteField("workerCpu", strconv.FormatUint(k8sCbCreateReq.WorkerCPU, 10))
		}
		if k8sCbCreateReq.WorkerNum != 0 {
			_ = writer.WriteField("workerNum", strconv.FormatUint(k8sCbCreateReq.WorkerNum, 10))
		}
		if k8sCbCreateReq.WorkerRAM != 0 {
			_ = writer.WriteField("workerRam", strconv.FormatUint(k8sCbCreateReq.WorkerRAM, 10))
		}
		if k8sCbCreateReq.WorkerDisk != 0 {
			_ = writer.WriteField("workerDisk", strconv.FormatUint(k8sCbCreateReq.WorkerDisk, 10))
		}
		if k8sCbCreateReq.ExtNetID != 0 {
			_ = writer.WriteField("extnetId", strconv.FormatUint(k8sCbCreateReq.ExtNetID, 10))
		}
		if k8sCbCreateReq.VinsId != 0 {
			_ = writer.WriteField("vinsId", strconv.FormatUint(k8sCbCreateReq.VinsId, 10))
		}
		if !k8sCbCreateReq.WithLB {
			_ = writer.WriteField("withLB", strconv.FormatBool(k8sCbCreateReq.WithLB))
		}

		_ = writer.WriteField("highlyAvailableLB", strconv.FormatBool(k8sCbCreateReq.HighlyAvailable))

		if k8sCbCreateReq.AdditionalSANs != nil {
			for _, v := range k8sCbCreateReq.AdditionalSANs {
				_ = writer.WriteField("additionalSANs", v)
			}
		}
		if k8sCbCreateReq.InitConfiguration != "" {
			_ = writer.WriteField("initConfiguration", k8sCbCreateReq.InitConfiguration)
		}
		if k8sCbCreateReq.ClusterConfiguration != "" {
			_ = writer.WriteField("clusterConfiguration", k8sCbCreateReq.ClusterConfiguration)
		}
		if k8sCbCreateReq.KubeletConfiguration != "" {
			_ = writer.WriteField("kubeletConfiguration", k8sCbCreateReq.KubeletConfiguration)
		}
		if k8sCbCreateReq.KubeProxyConfiguration != "" {
			_ = writer.WriteField("kubeProxyConfiguration", k8sCbCreateReq.KubeProxyConfiguration)
		}
		if k8sCbCreateReq.JoinConfiguration != "" {
			_ = writer.WriteField("joinConfiguration", k8sCbCreateReq.JoinConfiguration)
		}
		if k8sCbCreateReq.Description != "" {
			_ = writer.WriteField("desc", k8sCbCreateReq.Description)
		}
		if k8sCbCreateReq.UserData != "" {
			_ = writer.WriteField("userData", k8sCbCreateReq.UserData)
		}

		_ = writer.WriteField("extnetOnly", strconv.FormatBool(k8sCbCreateReq.ExtNetOnly))
		_ = writer.FormDataContentType()

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

	// var resp *http.Response
	// var err error
	buf, _ := io.ReadAll(req.Body)
	// req = req.Clone(req.Context())

	// for i := uint64(0); i < dc.cfg.Retries; i++ {
	req.Body = io.NopCloser(bytes.NewBuffer(buf))
	resp, err := dc.client.Do(req)

	// if err == nil {
	if resp.StatusCode == 200 {
		return resp, err
	}
	respBytes, _ := io.ReadAll(resp.Body)
	err = fmt.Errorf("%s", respBytes)
	resp.Body.Close()
	// }
	// }

	return nil, fmt.Errorf("could not execute request: %w", err)
}

func (dc *DecortClient) domp(req *http.Request, ctype string) (*http.Response, error) {
	req.Header.Add("Content-Type", ctype)
	req.Header.Add("Authorization", "bearer "+dc.cfg.Token)
	req.Header.Set("Accept", "application/json")

	// var resp *http.Response
	// var err error
	buf, _ := io.ReadAll(req.Body)
	// req = req.Clone(req.Context())

	// for i := uint64(0); i < dc.cfg.Retries; i++ {
	req.Body = io.NopCloser(bytes.NewBuffer(buf))
	resp, err := dc.client.Do(req)

	// if err == nil {
	if resp.StatusCode == 200 {
		return resp, err
	}
	respBytes, _ := io.ReadAll(resp.Body)
	err = fmt.Errorf("%s", respBytes)
	resp.Body.Close()
	// }
	// }

	return nil, fmt.Errorf("could not execute request: %w", err)
}
