package one

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const authHeaderName = "X-ZT1-Auth"

var (
	// APIVersion is the version of this library
	APIVersion = "0.2.0"
	// UserAgent is the entire HTTP User-Agent sent as header.
	UserAgent = fmt.Sprintf("zt-one/%s", APIVersion)
	// DefaultBaseURL is the default URL to contact, can be overridden with the (*Client).SetBaseURL() call.
	DefaultBaseURL = "http://localhost:9993"
)

// Client is a http client tailored to talk to ZeroTier One.
type Client struct {
	baseURL *url.URL
	apiKey  string
	client  *http.Client
}

// NewClient creates a new *Client. An API key must be passed at this time.
func NewClient(apiKey string) *Client {
	u, _ := url.ParseRequestURI(DefaultBaseURL)
	return &Client{
		client:  &http.Client{},
		apiKey:  apiKey,
		baseURL: u,
	}
}

// SetBaseURL sets the base URL to the value.
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

func (c *Client) makeBaseReq(method, path string, body io.Reader) (*http.Request, error) {
	u, err := url.ParseRequestURI(path)
	if err != nil {
		return nil, err
	}

	return http.NewRequest(method, c.baseURL.ResolveReference(u).String(), body)
}

func (c *Client) wrapJSON(path string, obj interface{}) error {
	req, err := c.makeBaseReq("GET", path, nil)
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
