package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

type transportLegacy struct {
	base       http.RoundTripper
	username   string
	password   string
	retries    uint64
	token      string
	decortURL  string
	mutex      *sync.Mutex
	expiryTime time.Time
}

func (t *transportLegacy) RoundTrip(request *http.Request) (*http.Response, error) {
	if t.token == "" || time.Now().After(t.expiryTime) {
		body := fmt.Sprintf("username=%s&password=%s", t.username, t.password)
		bodyReader := strings.NewReader(body)

		req, _ := http.NewRequestWithContext(request.Context(), "POST", t.decortURL+"/restmachine/cloudapi/user/authenticate", bodyReader)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, err := t.base.RoundTrip(req)
		if err != nil {
			return nil, fmt.Errorf("unable to get token: %w", err)
		}

		tokenBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode != 200 {
			return nil, fmt.Errorf("unable to get token: %s", tokenBytes)
		}

		token := string(tokenBytes)
		t.token = token
		t.expiryTime = time.Now().AddDate(0, 0, 1)
	}

	tokenValue := fmt.Sprintf("&authkey=%s", t.token)
	tokenReader := strings.NewReader(tokenValue)

	newBody := io.MultiReader(request.Body, tokenReader)

	req, _ := http.NewRequestWithContext(request.Context(), request.Method, request.URL.String(), newBody)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	var resp *http.Response
	var err error
	for i := uint64(0); i < t.retries; i++ {
		t.mutex.Lock()
		resp, err = t.base.RoundTrip(req)
		t.mutex.Unlock()
		if err == nil {
			if resp.StatusCode == 200 {
				return resp, nil
			}
			respBytes, _ := io.ReadAll(resp.Body)
			err = fmt.Errorf("%s", respBytes)
			resp.Body.Close()
		}
		if err != nil {
			return nil, fmt.Errorf("could not execute request: %w", err)
		}
		time.Sleep(time.Second * 5)
	}
	return nil, fmt.Errorf("could not execute request: %w", err)
}
