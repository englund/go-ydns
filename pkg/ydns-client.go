package pkg

import (
	"fmt"
	"io"
	"net/http"
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
	fmt.Println(*ip)
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
