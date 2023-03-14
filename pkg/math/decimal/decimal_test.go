package decimal

import (
	"fmt"
	"testing"
)

func TestDecimal_Pow(t *testing.T) {
	precession := NewFromInt(6)

	amount := NewFromInt(10)
	amountRaw := NewFromInt(10000000)

	fmt.Println("amountRaw to amount", amount.Mul(NewFromInt(10).Pow(precession)))
	fmt.Println("amount to amountRaw", amountRaw.Div(NewFromInt(10).Pow(precession)))
}
