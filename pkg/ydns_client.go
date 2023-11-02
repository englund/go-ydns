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
}

func NewYdnsClient(baseUrl string, username string, password string) *YdnsClient {
	return &YdnsClient{baseUrl, username, password}
}

func (c *YdnsClient) Update(host string, ip string) error {
	u, err := url.Parse(c.baseUrl)
	if err != nil {
		return fmt.Errorf("parsing base URL: %v", err)
	}

	u.Path = "/update/"
	q := u.Query()
	q.Set("host", host)
	q.Set("ip", ip)
	u.RawQuery = q.Encode()

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("creating new request: %v", err)
	}

	req.SetBasicAuth(c.username, c.password)
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %v", err)
	}

	result := string(body)

	if result != "ok" {
		return fmt.Errorf("unexpected response: %s", result)
	}

	fmt.Println(result)

	return nil
}

func (c *YdnsClient) GetIp() (*string, error) {
	u, err := url.Parse(c.baseUrl)
	if err != nil {
		return nil, fmt.Errorf("parsing base URL: %v", err)
	}

	u.Path = "/ip"

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	resp, err := client.Get(u.String())
	if err != nil {
		return nil, fmt.Errorf("making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %v", err)
	}

	ip := string(body)
	return &ip, nil
}
