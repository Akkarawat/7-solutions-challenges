package baconipsum

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	HTTPClient *http.Client
	BaseUrl string
}

func NewBaconipsumClient(baseUrl string) *Client {
	return &Client{
		HTTPClient: &http.Client{},
		BaseUrl: baseUrl,
	}
}

func (c *Client) fetch(contentType ContentType, paragraphCount int) (io.ReadCloser, error) {

	baseURL, err := url.Parse(c.BaseUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	params := url.Values{}
	params.Add("type", string(contentType))
	params.Add("paras", fmt.Sprintf("%d", paragraphCount))
	params.Add("format", "text")
	baseURL.RawQuery = params.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (c *Client) GetMeatText(hasFiller bool) (io.ReadCloser, error) {
	var contentType ContentType
	if hasFiller {
		contentType = CONTENT_TYPE_MEAT_AND_FILLER
	} else {
		contentType = CONTENT_TYPE_MEAT
	}
	return c.fetch(contentType, 99)
}