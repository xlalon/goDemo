package conf

import "time"

type Config struct {
	Band *ChainConfig `yaml:"band"`
	Waxp *ChainConfig `yaml:"waxp"`
}

type ChainConfig struct {
	NodeUrl         string `yaml:"node_url"`
	ExternalNodeUrl string `yaml:"external_node_url"`

	BlockTime         time.Duration `yaml:"block_time"`
	IrreversibleBlock int64         `yaml:"irreversible_block"`

	WalletDepositUrl string `yaml:"wallet_deposit_url"`
	WalletHotUrl     string `yaml:"wallet_hot_url"`

	SupportMemo    bool   `yaml:"support_memo"`
	DepositAddress string `yaml:"deposit_address"`
	HotAddress     string `yaml:"hot_address"`
}
