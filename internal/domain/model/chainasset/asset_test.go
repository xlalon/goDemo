package chainasset

import (
	"github.com/xlalon/golee/pkg/math/decimal"
	"testing"
)

func TestAsset_CalculateAmount(t *testing.T) {
	asset := NewAsset(&AssetDTO{
		Id:         1,
		Code:       "BAND",
		Name:       "BAND",
		Chain:      "BAND",
		Identity:   "uband",
		Precession: 6,
		Status:     string(AssetStatusOnline),
		Setting:    nil,
	})
	result, _ := decimal.NewFromString("1")
	if !asset.CalculateAmount(decimal.NewFromInt(1000000)).Equal(result) {
		t.Fatal("CalculateAmount Fail")
	}
}
