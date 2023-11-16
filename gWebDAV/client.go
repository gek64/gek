package gWebDAV

import "net/url"

type Client struct {
	URL           *url.URL
	Username      string
	Password      string
	AllowInsecure bool
}

func NewClient(host string, username string, password string, allowInsecure bool) (client *Client, err error) {
	parsedURL, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	return &Client{
		URL:           parsedURL,
		Username:      username,
		Password:      password,
		AllowInsecure: allowInsecure,
	}, nil
}
