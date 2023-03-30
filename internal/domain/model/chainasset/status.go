package chainasset

type Status string

const (
	ChainStatusOnline  Status = "ONLINE"
	ChainStatusOffline        = "OFFLINE"

	AssetStatusOnline  Status = "ONLINE"
	AssetStatusOffline        = "OFFLINE"
)
