package one

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const authHeaderName = "X-ZT1-Auth"

var (
	APIVersion     = "0.1.0"
	UserAgent      = fmt.Sprintf("zt-one/%s", APIVersion)
	DefaultBaseURL = "http://localhost:9993"
)

type Client struct {
	baseURL *url.URL
	apiKey  string
	client  *http.Client
}

func NewClient(apiKey string) *Client {
	u, _ := url.ParseRequestURI(DefaultBaseURL)
	return &Client{
		client:  &http.Client{},
		apiKey:  apiKey,
		baseURL: u,
	}
}

func (c *Client) SetBaseURL(base string) error {
	u, err := url.ParseRequestURI(base)
	if err != nil {
		return err
	}

	c.baseURL = u
	return nil
}

func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("User-Agent", UserAgent)
	req.Header.Add(authHeaderName, c.apiKey)
	return c.client.Do(req)
}

func (c *Client) wrapJSON(path string, obj interface{}) error {
	u, err := url.ParseRequestURI(path)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", c.baseURL.ResolveReference(u).String(), nil)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Response status was not 200, was %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(obj)
}
