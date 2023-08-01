package client

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

const (
	TOKEN_EXPIRE_PERIOD = 30
	TOKEN_RENEW_PERIOD  = 15
	ACTIVE_CHECK_PERIOD = 30
)

type CommvaultClient struct {
	target                  string
	token                   string
	username                string
	password                string
	insecure                bool
	httpclient              *http.Client
	tokenAutoRenewPeriod    time.Duration
	tokenAutoRenewTicker    *time.Ticker
	isActive                bool
	isActiveLock            sync.Mutex
	activeStatusCheckPeriod time.Duration
	activeStatusCheckTicker *time.Ticker

	Vm       *VmService
	Storage  *StorageService
	Bkp      *BkpService
	Database *DatabaseService
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
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	c := &CommvaultClient{target: target, username: username, password: password, insecure: insecure}
	c.httpclient = &http.Client{Transport: tr}
	c.Vm = &VmService{client: c}
	c.Storage = &StorageService{client: c}
	c.Bkp = &BkpService{client: c}
	c.Database = &DatabaseService{client: c}
	c.tokenAutoRenewPeriod = TOKEN_RENEW_PERIOD * time.Minute
	err := c.RenewToken()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CommvaultClient) NewRequest(method string, path string, params map[string]any, data *interface{}) (*http.Request, error) {

	baseUrl, err := url.Parse(c.target)
	if err != nil {
		return nil, err
	}
	baseUrl.Path = path
	if params != nil {
		ps := url.Values{}
		for k, v := range params {
			ps.Set(k, fmt.Sprintf("%v", v))
		}
		baseUrl.RawQuery = ps.Encode()
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
	req.Header.Add("Authtoken", c.token)

	return req, err
}

func (c *CommvaultClient) NewXmlRequest(method string, path string, params map[string]any, data *interface{}) (*http.Request, error) {

	req, err := c.NewRequest(method, path, params, data)
	req.Header.Add("Accept", "application/xml")

	return req, err
}

func (c *CommvaultClient) Do(req *http.Request, v interface{}) (*http.Response, error) {

	resp, err := c.httpclient.Do(req)
	if err != nil {
		fmt.Println("Do request failed")
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("[error] %s", err)
		}
	}()
	err = readJsonResponse(resp, v)

	return resp, err
}

func (c *CommvaultClient) DoXml(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpclient.Do(req)
	if err != nil {
		fmt.Println("Do request failed")
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("[error] %s", err)
		}
	}()

	// if err := validateResponse(resp); err != nil {
	// 	return resp, err
	// }
	err = readXmlResponse(resp, v)
	return resp, err
}

func readJsonResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := io.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	if c := r.StatusCode; c < 200 || c > 299 {
		return fmt.Errorf("[error] %s", bodyString)
	}

	err := json.Unmarshal(bodyBytes, &v)
	if err != nil {
		return fmt.Errorf("failed to parse body to %T\n%v", v, err)
	}
	return nil
}

func readXmlResponse(r *http.Response, v interface{}) error {
	if v == nil {
		return fmt.Errorf("nil interface provided to decodeResponse")
	}

	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)

	if c := r.StatusCode; c < 200 || c > 299 {
		return fmt.Errorf("[error] %s", bodyString)
	}

	fmt.Printf("== xml response ==\n%s", bodyString)

	err := xml.Unmarshal([]byte(bodyString), &v)
	return err
}

func (c *CommvaultClient) RenewToken() error {
	if c.GetActiveStatus() {
		token, err := GetCommvaultToken(c.target, c.username, c.password, c.insecure)
		if err != nil {
			return err
		}
		c.token = token
	}
	return nil
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
			err := c.RenewToken()
			if err != nil {
				fmt.Printf("[error]: failed to renew token\n%+v\n", err)
				continue
			}
			fmt.Printf("[info] Renew token\n")
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
			fmt.Printf("[error] %s", err)
		}
	}()

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		resp, err := ioutil.ReadAll(res.Body)
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
		resp, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf(string(resp))
	}
}

func (c *CommvaultClient) GetActiveStatus() bool {
	c.isActiveLock.Lock()
	defer c.isActiveLock.Unlock()
	return c.isActive
}

func (c *CommvaultClient) StartActiveCheck(ctx context.Context) {
	go c.checkActiveExecutor(ctx)
}

func (c *CommvaultClient) StopActiveCheck(ctx context.Context) {
	c.activeStatusCheckTicker.Stop()
}

func (c *CommvaultClient) checkActiveExecutor(ctx context.Context) {
	c.activeStatusCheckTicker = time.NewTicker(c.activeStatusCheckPeriod)
	for {
		select {
		case <-c.activeStatusCheckTicker.C:
			err := c.CheckActiveStatus()
			if err != nil {
				fmt.Printf("[error]: failed to check API endpoint status\n%+v\n", err)
				continue
			}
			fmt.Printf("[info] check API endpoint status\n")
		case <-ctx.Done():
			return
		}
	}
}

func (c *CommvaultClient) CheckActiveStatus() error {
	baseUrl, err := url.Parse(c.target)
	if err != nil {
		return err
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", baseUrl.Hostname(), baseUrl.Port()))
	if err != nil {
		c.isActiveLock.Lock()
		defer c.isActiveLock.Unlock()
		c.isActive = false
	} else {
		c.isActiveLock.Lock()
		defer c.isActiveLock.Unlock()
		c.isActive = true
	}

	defer func() {
		err := conn.Close()
		fmt.Printf("failed to close connection: %s\n", err)
	}()

	return nil
}
