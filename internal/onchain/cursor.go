package onchain

type Direction string

const (
	CursorDirectionAsc  Direction = "ASC"
	CursorDirectionDesc           = "DESC"
)

type Cursor struct {
	Chain Code
	// scan by height range
	Height int64
	// scan by account
	Account   *Account
	TxHash    string
	Direction Direction //DESC
	Index     int64
}

func NewCursor(chain Code, height int64, address string, label Label, txHash string, direction Direction, index int64) *Cursor {

	if label != AccountDeposit && label != AccountHot {
		return nil
	}

	if direction != CursorDirectionAsc && direction != CursorDirectionDesc {
		return nil
	}
	return &Cursor{
		Chain:  chain,
		Height: height,
		Account: &Account{
			Chain:   chain,
			Address: address,
			Label:   label,
		},
		TxHash:    txHash,
		Direction: direction,
		Index:     index,
	}
}
