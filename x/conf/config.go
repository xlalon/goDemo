package conf

import "time"

type Config struct {
	NodeUrl string

	BlockTime         time.Duration
	IrreversibleBlock int64
}

type ChainConfig struct {
	Band *Config
}
