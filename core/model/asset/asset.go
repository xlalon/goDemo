package asset

import "github.com/xlalon/golee/core/model/chain"

type Status string

const (
	StatusOffline Status = "OFFLINE"
	StatusOnline  Status = "ONLINE"
	StatusSuspend Status = "SUSPEND"
)

type Code string

type Asset struct {
	id int64

	code      Code
	name      string
	status    Status
	identity  string
	chainCode chain.Code
}

func NewAsset(id int64, code Code, name, identity, status string, chainCode chain.Code) *Asset {
	return &Asset{
		id:        id,
		code:      code,
		name:      name,
		status:    Status(status),
		identity:  identity,
		chainCode: chainCode,
	}
}

func (a *Asset) Id() int64 {
	return a.id
}

func (a *Asset) Code() Code {
	return a.code
}

func (a *Asset) Name() string {
	return a.name
}

func (a *Asset) Status() Status {
	return a.status
}

func (a *Asset) Offline() bool {
	if a.status != StatusOffline {
		a.status = StatusOffline
		return true
	}
	return false
}

func (a *Asset) Online() bool {
	if a.status != StatusOnline {
		a.status = StatusOnline
		return true
	}
	return false
}

func (a *Asset) Suspend() bool {
	if a.status != StatusSuspend {
		a.status = StatusSuspend
		return true
	}
	return false
}

func (a *Asset) Identity() string {
	return a.identity
}

func (a *Asset) Chain() chain.Code {
	return a.chainCode
}
