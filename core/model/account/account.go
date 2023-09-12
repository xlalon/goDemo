package account

type Status string

const (
	StatusDeleted    Status = "DELETED"
	StatusActive     Status = "ACTIVE"
	StatusDeprecated Status = "DEPRECATED"
)

type Account struct {
	id int64

	address string
	memo    string

	status Status
}

func NewAccount(id int64, address, memo, status string) *Account {
	return &Account{
		id:      id,
		address: address,
		memo:    memo,
		status:  Status(status),
	}
}

func (a *Account) Id() int64 {
	return a.id
}

func (a *Account) Address() string {
	return a.address
}

func (a *Account) Memo() string {
	return a.memo
}

func (a *Account) Status() Status {
	return a.status
}

func (a *Account) Active() bool {
	if a.status != StatusActive {
		a.status = StatusActive
		return true
	}
	return false
}

func (a *Account) Deprecated() bool {
	if a.status != StatusDeprecated {
		a.status = StatusDeprecated
		return true
	}
	return false
}

func (a *Account) Deleted() bool {
	if a.status != StatusDeleted {
		a.status = StatusDeleted
		return true
	}
	return false
}
