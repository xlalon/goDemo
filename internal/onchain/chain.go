package onchain

import (
	"context"
	"github.com/xlalon/golee/internal/onchain/conf"
)

type Chain struct {
	Code   Code              `json:"code"`
	Config *conf.ChainConfig `json:"config"`
}

func (*Chain) GetNodeInfo(ctx context.Context) (*Node, error) {
	_ = ctx
	return nil, nil
}

func (*Chain) GetBlockHash(ctx context.Context, height int64) (string, error) {
	_ = ctx
	_ = height
	return "", nil
}
