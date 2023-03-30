package deposit

import (
	"fmt"
	"testing"

	"github.com/xlalon/golee/pkg/math/decimal"
)

func TestAmountVO_ToAmount(t *testing.T) {
	amountVO := NewAmountVO("uband", decimal.NewFromInt(63000000), 6, decimal.NewFromInt(0))
	fmt.Println("amount:", amountVO.ToAmount())
}

func TestAmountVO_ToAmountRaw(t *testing.T) {
	amountVO := NewAmountVO("uband", decimal.NewFromInt(0), 6, decimal.NewFromInt(63))
	fmt.Println("amount raw:", amountVO.ToAmountRaw())
}
