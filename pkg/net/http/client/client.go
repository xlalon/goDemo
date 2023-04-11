package client

import (
	"context"
)

type Client interface {
	Get(ctx context.Context, url string, params map[string]string, opts ...Option) (string, error)
	Post(ctx context.Context, url string, data map[string]interface{}, opts ...Option) (string, error)
}

type Options struct {
	headers map[string]string
}

type Option func(options *Options)

func WithHeaders(headers map[string]string) Option {
	return func(options *Options) {
		options.headers = headers
	}
}
