package pkg

import "fmt"

type YdnsClient struct {
	username *string
	password *string
}

func NewYdnsClient(username *string, password *string) *YdnsClient {
	return &YdnsClient{username, password}
}

func (c *YdnsClient) Update(host *string) {
	fmt.Println(*host)
	fmt.Println(*c.username)
	fmt.Println(*c.password)
}
