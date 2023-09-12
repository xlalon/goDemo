package band

import (
	"context"

	"github.com/xlalon/golee/x/chain"
	"github.com/xlalon/golee/x/conf"
)

type Band struct {
	chain.Chain
}

func NewBand(conf *conf.Config) *Band {
	return &Band{Chain: chain.Chain{
		Code:   "BAND",
		Config: conf,
	}}
}

func (band *Band) Version(ctx context.Context) (string, error) {
	return "1.1.1", nil
}
