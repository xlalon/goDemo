package wallet

type AccountDTO struct {
	Id       int64      `json:"id"`
	Chain    string     `json:"chain"`
	Address  string     `json:"address"`
	Label    string     `json:"label"`
	Memo     string     `json:"memo"`
	Status   string     `json:"status"`
	Sequence int64      `json:"sequence"`
	Balances []*Balance `json:"balances"`
}
