package chain

import "github.com/xlalon/golee/core/model/asset"

type Status int64

const (
	StatusOffline Status = -1
	StatusOnline  Status = 0
)

type Code string

type Chain struct {
	id int64

	code Code

	name string

	status Status

	assets []*asset.Asset
}

func NewChain(id int64, code Code, name string, status int64) *Chain {
	return &Chain{
		id:     id,
		code:   code,
		name:   name,
		status: Status(status),
	}
}

func (c *Chain) Id() int64 {
	return c.id
}

func (c *Chain) Code() Code {
	return c.code
}

func (c *Chain) Name() string {
	return c.name
}

func (c *Chain) Status() Status {
	return c.status
}

func (c *Chain) Offline() bool {
	if c.status != StatusOffline {
		c.status = StatusOffline
		return true
	}
	return false
}

func (c *Chain) Online() bool {
	if c.status != StatusOnline {
		c.status = StatusOnline
		return true
	}
	return false
}

func (c *Chain) Assets() []*asset.Asset {
	return c.assets
}
