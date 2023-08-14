package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	TOKEN_EXPIRE_PERIOD     = 30
	TOKEN_RENEW_PERIOD      = 15
	ACTIVE_CHECK_PERIOD_SEC = 10
)

type CommvaultClient struct {
	target               string
	token                string
	username             string
	password             string
	insecure             bool
	httpclient           *http.Client
	tokenLock            sync.Mutex
	tokenAutoRenewPeriod time.Duration
	tokenAutoRenewTicker *time.Ticker
	tokenManualTrigger   chan bool
	tokenFlushTrigger    chan bool

	Vm       *VmService
	Storage  *StorageService
	Bkp      *BkpService
	Database *DatabaseService
	Status   *ClientStatusService
}

func NewClient(target string, username string, password string, insecure bool) (*CommvaultClient, error) {

	target = strings.TrimRight(target, "/")
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		target = fmt.Sprintf("https://%s", target)
	}

	if username == "" || password == "" {
		return nil, fmt.Errorf("[error] Must specify username and password")
	}

	// cookieJar, _ := cookiejar.New(nil)
	// #nosec G402
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: insecure},
		TLSHandshakeTimeout: 3 * time.Second,
	}
	c := &CommvaultClient{target: target, username: username, password: password, insecure: insecure}
	c.tokenLock = sync.Mutex{}
	c.tokenManualTrigger = make(chan bool)
	c.tokenFlushTrigger = make(chan bool)
	c.httpclient = &http.Client{
		Transport: tr,
	}
	c.Vm = &VmService{client: c}
	c.Storage = &StorageService{client: c}
	c.Bkp = &BkpService{client: c}
	c.Database = &DatabaseService{client: c}
	c.Status = NewClientStatusService(c, ACTIVE_CHECK_PERIOD_SEC*time.Second)
	c.tokenAutoRenewPeriod = TOKEN_RENEW_PERIOD * time.Minute

	err := c.Status.CheckStatus()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CommvaultClient) Url(path *string, params *map[string]any) (*url.URL, error) {
	baseUrl, err := url.Parse(c.target)
	if err != nil {
		return nil, err
	}
	if path != nil {
		baseUrl.Path = *path
	}
	if params != nil {
		ps := url.Values{}
		for k, v := range *params {
			ps.Set(k, fmt.Sprintf("%v", v))
		}
		baseUrl.RawQuery = ps.Encode()
	}
	return baseUrl, nil
}

func (c *CommvaultClient) NewRequest(method string, path string, params map[string]any, data *interface{}) (*http.Request, error) {

	baseUrl, err := c.Url(&path, &params)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	if data != nil {
		jsonString, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, baseUrl.String(), bytes.NewBuffer(jsonString))
		if err != nil {
			return nil, err
		}
	}

	if data != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authtoken", c.GetToken())

	return req, err
}

func (c *CommvaultClient) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.httpclient.Do(req)
	if err != nil {
		fmt.Print("[error] Do request failed\n")
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("[error] %s\n", err)
		}
	}()
	err = readJsonResponse(resp, v)

	return resp, err
}

func (c *CommvaultClient) IsActive() bool {
	return c.Status != nil && c.Status.GetIsActive() && c.GetToken() != ""
}

func readJsonResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := io.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	if c := r.StatusCode; c < 200 || c > 299 {
		return fmt.Errorf("%s", bodyString)
	}

	err := json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return fmt.Errorf("failed to parse body to %T\n%v", v, err)
	}
	return nil
}

func (c *CommvaultClient) RenewToken() error {
	if c.Status != nil && c.Status.GetIsActive() {
		token, err := GetCommvaultToken(c.target, c.username, c.password, c.insecure)
		if err != nil {
			return err
		}
		c.tokenLock.Lock()
		defer c.tokenLock.Unlock()
		c.token = token
	}
	return nil
}

func (c *CommvaultClient) FlushToken() {
	c.tokenLock.Lock()
	defer c.tokenLock.Unlock()
	c.token = ""
}

func (c *CommvaultClient) GetToken() string {
	c.tokenLock.Lock()
	defer c.tokenLock.Unlock()
	return c.token
}

func (c *CommvaultClient) StartTokenAutoRenew(ctx context.Context) {
	go c.renewTokenExecutor(ctx)
}

func (c *CommvaultClient) StopTokenAutoRenew(ctx context.Context) {
	c.tokenAutoRenewTicker.Stop()
}

func (c *CommvaultClient) renewTokenExecutor(ctx context.Context) {
	c.tokenAutoRenewTicker = time.NewTicker(c.tokenAutoRenewPeriod)
	for {
		select {
		case <-c.tokenAutoRenewTicker.C:
			fmt.Printf("[info] Start renew token (auto trigger)\n")
			err := c.RenewToken()
			if err != nil {
				fmt.Printf("[error]: failed to renew token\n%+v\n", err)
				continue
			}
		case <-c.tokenManualTrigger:
			fmt.Printf("[info] Start renew token (manual trigger)\n")
			err := c.RenewToken()
			if err != nil {
				fmt.Printf("[error]: failed to renew token\n%+v\n", err)
				continue
			}
		case <-c.tokenFlushTrigger:
			c.FlushToken()
			fmt.Printf("[info] Flush token\n")
		case <-ctx.Done():
			return
		}
	}
}

func GetCommvaultToken(apiEndpoint string, username string, password string, insecure bool) (string, error) {
	// #nosec G402
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		},
		Timeout: 3 * time.Second,
	}
	url := fmt.Sprintf("%s/webconsole/api/Login", apiEndpoint)

	body, err := json.Marshal(map[string]any{
		"username": username,
		"password": base64.StdEncoding.EncodeToString([]byte(password)),
		"timeout":  TOKEN_EXPIRE_PERIOD,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create body: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Printf("[error] %s\n", err)
		}
	}()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		resp, err := io.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		uResp := LoginResp{}
		err = json.Unmarshal(resp, &uResp)
		if err != nil {
			return "", err
		}
		if uResp.Token != "" {
			return uResp.Token, nil
		} else if len(uResp.ErrList) > 0 {
			errorMsgs := []string{}
			for _, e := range uResp.ErrList {
				errorMsgs = append(errorMsgs, fmt.Sprintf("[error] %s (code %d)", e.ErrLogMessage, e.ErrorCode))
			}
			return "", fmt.Errorf(strings.Join(errorMsgs, "\n"))
		}
		return "", fmt.Errorf("[error] failed to get a token")
	} else {
		resp, err := io.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf(string(resp))
	}
}
