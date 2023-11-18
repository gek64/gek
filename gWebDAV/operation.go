package gWebDAV

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"net/url"
)

func (c *Client) SendWebDavRequest(method string, path string, data []byte) (resp *http.Response, err error) {
	folderUrl, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// build request
	request, err := http.NewRequest(method, c.Endpoint.ResolveReference(folderUrl).String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	request.SetBasicAuth(c.Username, c.Password)

	// allow insecure cert in all clients
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// allow insecure cert in a client
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.AllowInsecure},
	}
	client := http.Client{Transport: tr}

	// do request
	return client.Do(request)
}

func (c *Client) Mkdir(path string) (resp *http.Response, err error) {
	return c.SendWebDavRequest("MKCOL", path, nil)
}

func (c *Client) Upload(path string, data []byte) (resp *http.Response, err error) {
	return c.SendWebDavRequest("PUT", path, data)
}

func (c *Client) Download(path string) (resp *http.Response, err error) {
	return c.SendWebDavRequest("GET", path, nil)
}

func (c *Client) Delete(path string) (resp *http.Response, err error) {
	return c.SendWebDavRequest("DELETE", path, nil)
}

func (c *Client) Exist(path string) (resp *http.Response, err error) {
	return c.SendWebDavRequest("PROPFIND", path, nil)
}
