package asset

type Status int64

const (
	StatusOffline Status = -1
	StatusOnline  Status = 0
)

type Code string

type Asset struct {
	id int64

	code Code

	name string

	status Status

	identity string
}

func NewAsset(id int64, code Code, name, identity string, status int64) *Asset {
	return &Asset{
		id:       id,
		code:     code,
		name:     name,
		status:   Status(status),
		identity: identity,
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

func (a *Asset) Identity() string {
	return a.identity
}
