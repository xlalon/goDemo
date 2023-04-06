package ecode

var (
	DepositAssetInvalid           = New(3000001)
	DepositAssetChange            = New(3000002)
	DepositAmountInvalid          = New(3000003)
	DepositAmountChange           = New(3000004)
	DepositCoinValueInvalid       = New(3000005)
	DepositCoinValueChange        = New(3000006)
	DepositItemInvalid            = New(3000007)
	DepositStatusInvalid          = New(3000008)
	DepositReceiverAddressInvalid = New(3000009)
	DepositReceiverAddressChange  = New(30000010)
	DepositTxIdChange             = New(30000011)
	DepositTxIdChainInvalid       = New(30000012)
	DepositTxIdHashInvalid        = New(30000013)
	DepositTxIdVOutInvalid        = New(30000014)

	CursorInvalid = New(30010001)
	CursorChange  = New(30010002)
)
