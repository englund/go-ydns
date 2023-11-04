package pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type YdnsClient struct {
	baseUrl  string
	username string
	password string
	client   *http.Client
}

func NewYdnsClient(baseUrl string, username string, password string) *YdnsClient {
	return &YdnsClient{
		baseUrl:  baseUrl,
		username: username,
		password: password,
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *YdnsClient) Update(host string, ip string) error {
	u, err := c.newRequestUrl("/update/", map[string]string{
		"host": host,
		"ip":   ip,
	})
	if err != nil {
		return fmt.Errorf("creating new request: %v", err)
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("creating new request: %v", err)
	}

	req.SetBasicAuth(c.username, c.password)
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d with body \"%s\"", resp.StatusCode, body)
	}

	result := string(body)

	if result != "ok" {
		return fmt.Errorf("unexpected response: %s", result)
	}

	return nil
}

func (c *YdnsClient) GetIp() (*string, error) {
	u, err := c.newRequestUrl("/ip", nil)
	if err != nil {
		return nil, fmt.Errorf("creating new request: %v", err)
	}

	resp, err := c.client.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	ip := string(body)
	return &ip, nil
}

func (c *YdnsClient) newRequestUrl(path string, params map[string]string) (*url.URL, error) {
	u, err := url.Parse(c.baseUrl + path)
	if err != nil {
		return nil, fmt.Errorf("parsing URL: %v", err)
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
	return u, nil
}
