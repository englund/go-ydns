package pkg

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type YdnsClient struct {
	username *string
	password *string
}

const ydnsBaseUrl = "https://ydns.io/api/v1"

func NewYdnsClient(username *string, password *string) *YdnsClient {
	return &YdnsClient{username, password}
}

func (c *YdnsClient) Update(host *string, ip *string) error {
	url := fmt.Sprintf("%s/update/?host=%s&ip=%s", ydnsBaseUrl, *host, *ip)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(*c.username, *c.password)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	result := string(body)

	if !strings.Contains(result, "ok") {
		return errors.New(result)
	}

	fmt.Println(result)

	return nil
}

func (c *YdnsClient) GetIp() (*string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/ip", ydnsBaseUrl))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ip := string(body)
	return &ip, nil
}
