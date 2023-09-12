package x

import (
	"context"
	"fmt"
	"testing"

	"github.com/xlalon/golee/common/json"
	"github.com/xlalon/golee/x/conf"
)

var (
	testConf = &conf.ChainConfig{
		Band: &conf.Config{
			NodeUrl:           "https://laozi1.bandchain.org/api",
			BlockTime:         7,
			IrreversibleBlock: 10,
		}}
)

func TestGetChain(t *testing.T) {
	Init(testConf)
	c, err := GetChain("BAND")
	if err != nil {
		t.Fatal(err)
	}
	json.JPrint("TestGetChain", c)
	fmt.Println(c.Version(context.Background()))
}
