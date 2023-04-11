package client

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type RestyClient struct {
	*resty.Client
}

func NewRestyClient(baseURL string, opts ...Option) *RestyClient {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}
	_client := resty.New()
	_client.SetBaseURL(baseURL)
	_client.SetHeaders(option.headers)

	return &RestyClient{_client}
}

func (c *RestyClient) Get(ctx context.Context, url string, params map[string]string, opts ...Option) (string, error) {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}

	req := c.NewRequest().SetContext(ctx).SetQueryParams(params)
	if option.headers != nil {
		req = req.SetHeaders(option.headers)
	}

	resp, err := req.Get(url)
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

func (c *RestyClient) Post(ctx context.Context, url string, data map[string]interface{}, opts ...Option) (string, error) {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}

	req := c.NewRequest().SetContext(ctx).SetBody(data)
	if option.headers != nil {
		req = req.SetHeaders(option.headers)
	}

	resp, err := req.Post(url)
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
