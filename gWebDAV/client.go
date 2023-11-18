package gWebDAV

import "net/url"

type Client struct {
	Endpoint      *url.URL
	Username      string
	Password      string
	AllowInsecure bool
}

func NewClient(endpoint string, username string, password string, allowInsecure bool) (client *Client, err error) {
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		Endpoint:      parsedURL,
		Username:      username,
		Password:      password,
		AllowInsecure: allowInsecure,
	}, nil
}
