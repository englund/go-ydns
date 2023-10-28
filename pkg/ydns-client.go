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

const YDNS_BASE_URL = "https://ydns.io/api/v1"

func NewYdnsClient(username *string, password *string) *YdnsClient {
	return &YdnsClient{username, password}
}

func (c *YdnsClient) Update(host *string) error {
	ip, err := getIp()
	if err != nil {
		return err
	}

	fmt.Println(*ip)
	return nil
}

func getIp() (*string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/ip", YDNS_BASE_URL))
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
