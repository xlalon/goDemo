package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	xurl "net/url"
	"path"
	"strings"
)

type DefaultClient struct {
	BaseUrl string
	Headers map[string]string
	client  *http.Client
}

func NewRestClient(baseURL string, opts ...Option) *DefaultClient {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}

	return &DefaultClient{
		BaseUrl: baseURL,
		Headers: option.headers,
		// TODO
		client: http.DefaultClient,
	}
}

func (r *DefaultClient) Get(ctx context.Context, url string, params map[string]string, opts ...Option) (string, error) {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}
	if option.headers != nil {
		r.Headers = option.headers
	}

	_fullUrl := fullUrl(r.BaseUrl, url, params)
	if _fullUrl == "" {
		return "", errors.New("wrong url")
	}
	req, err := r.request(ctx, http.MethodGet, _fullUrl, nil)
	if err != nil {
		return "", errors.New("bad request")
	}

	return r.doRequest(req)
}

func (r *DefaultClient) Post(ctx context.Context, url string, data map[string]interface{}, opts ...Option) (string, error) {
	option := &Options{}
	for _, opt := range opts {
		opt(option)
	}
	if option.headers != nil {
		r.Headers = option.headers
	}

	_fullUrl := fullUrl(r.BaseUrl, url, nil)
	if _fullUrl == "" {
		return "", errors.New("wrong url")
	}
	var d []byte
	if data != nil {
		var err error
		d, err = json.Marshal(data)
		if err != nil {
			return "", errors.New("wrong body")
		}
	}
	req, err := r.request(ctx, http.MethodPost, _fullUrl, strings.NewReader(string(d)))
	if err != nil {
		return "", errors.New("bad request")
	}

	return r.doRequest(req)
}

func (r *DefaultClient) request(ctx context.Context, method, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, urlStr, body)
	if err != nil {
		return nil, err
	}
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	return req, err
}

func (r *DefaultClient) doRequest(request *http.Request) (string, error) {
	response, err := r.client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)

	return string(b), err
}

func fullUrl(baseURL, url string, params map[string]string) string {

	_fullURL, err := xurl.Parse(strings.TrimSuffix(baseURL, "/"))
	if err != nil {
		return ""
	}
	urlUrl, err := xurl.Parse(url)
	if err != nil {
		return ""
	}

	_fullURL.Path = path.Join(_fullURL.Path, urlUrl.Path)

	query := _fullURL.Query()
	for k, vs := range urlUrl.Query() {
		for _, v := range vs {
			query.Set(k, v)
		}
	}
	for k, v := range params {
		query.Set(k, v)
	}
	_fullURL.RawQuery = query.Encode()

	fullUrlStr := _fullURL.String()
	fmt.Println("Full Request URL:", fullUrlStr)

	return fullUrlStr
}
