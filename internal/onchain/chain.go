package onchain

import (
	"github.com/xlalon/golee/internal/onchain/conf"
)

type Chain struct {
	Code   Code              `json:"code"`
	Config *conf.ChainConfig `json:"config"`
}

func (*Chain) GetNodeInfo() (*Node, error) {
	return nil, nil
}

func (*Chain) GetBlockHash(height int64) (string, error) {
	_ = height
	return "", nil
}
