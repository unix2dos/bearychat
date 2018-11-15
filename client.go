// A simple package for interacting with bearychat's API.
package bearychat

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Hook string
}

func (c Client) Send(message io.Reader) (*http.Response, error) {
	req := &http.Client{
		Timeout: time.Second * 30,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return req.Post(c.Hook, "application/json", message)
}

func (c *Client) SetHook(hook string) *Client {
	c.Hook = hook
	return c
}

func NewClient(hook string) *Client {
	return &Client{hook}
}

// Exported client.
var DefaultClient = NewClient("")

func SetHook(hook string) *Client {
	return DefaultClient.SetHook(hook)
}

func Send(message io.Reader) (*http.Response, error) {
	return DefaultClient.Send(message)
}
