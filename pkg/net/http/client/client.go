package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	xurl "net/url"
	"path"
	"strings"
	"time"
)

type Config struct {
	BaseUrl string        `yaml:"base_url"`
	Headers xurl.Values   `yaml:"headers"`
	Timeout time.Duration `yaml:"timeout"`
}

type RestfulClient struct {
	BaseUrl string
	Headers xurl.Values
	client  *http.Client
}

func NewRestfulClient(config *Config) *RestfulClient {
	return &RestfulClient{
		BaseUrl: config.BaseUrl,
		Headers: config.Headers,
		// TODO
		client: http.DefaultClient,
	}
}

func (r *RestfulClient) Get(url string, params xurl.Values) (string, error) {
	fullUrl := fullUrl(r.BaseUrl, url, params)
	if fullUrl == "" {
		return "", errors.New("wrong url")
	}
	req, err := r.request(http.MethodGet, fullUrl, nil)
	if err != nil {
		return "", errors.New("bad request")
	}
	return r.doRequest(req)
}

func (r *RestfulClient) Post(url string, data map[string]interface{}) (string, error) {
	fullUrl := fullUrl(r.BaseUrl, url, nil)
	if fullUrl == "" {
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
	req, err := r.request(http.MethodPost, fullUrl, strings.NewReader(string(d)))
	if err != nil {
		return "", errors.New("bad request")
	}
	return r.doRequest(req)
}

func (r *RestfulClient) request(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, vs := range r.Headers {
		for _, v := range vs {
			req.Header.Set(k, v)
		}
	}
	return req, err
}

func (r *RestfulClient) doRequest(request *http.Request) (string, error) {
	response, err := r.client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	b, err := ioutil.ReadAll(response.Body)

	return string(b), err
}

func fullUrl(baseUrl, url string, params xurl.Values) string {

	fullUrl, err := xurl.Parse(strings.TrimSuffix(baseUrl, "/"))
	if err != nil {
		return ""
	}
	urlUrl, err := xurl.Parse(url)
	if err != nil {
		return ""
	}

	fullUrl.Path = path.Join(fullUrl.Path, urlUrl.Path)

	query := fullUrl.Query()
	for k, vs := range urlUrl.Query() {
		for _, v := range vs {
			query.Set(k, v)
		}
	}
	for k, vs := range params {
		for _, v := range vs {
			query.Set(k, v)
		}
	}
	fullUrl.RawQuery = query.Encode()

	fullUrlStr := fullUrl.String()
	fmt.Println("Full Request URL:", fullUrlStr)

	return fullUrlStr
}
