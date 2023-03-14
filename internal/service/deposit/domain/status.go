package domain

type Status string

const (
	DepositStatusPending   Status = "pending"
	DepositStatusFinished         = "finished"
	DepositStatusCancelled        = "cancelled"
	DepositStatusSwapped          = "swapped"
)
