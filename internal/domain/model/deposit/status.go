package deposit

type Status string

const (
	DepositStatusPending   Status = "PENDING"
	DepositStatusFinished         = "FINISHED"
	DepositStatusCancelled        = "CANCELLED"
	DepositStatusSwapped          = "SWAPPED"
)
