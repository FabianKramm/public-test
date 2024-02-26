package public_test

import "fmt"

func NewClient() *Client {
	return &Client{}
}

type Client struct{}

func (c *Client) IsPro() {
	fmt.Println("IS NOT PRO")
}
