package decimal

import (
	"fmt"
	"testing"
)

func TestDecimal_Div(t *testing.T) {
	amountZero := NewFromInt(0)
	amountTen := NewFromInt(10)
	fmt.Println("0 / 10", amountZero.Div(amountTen))
	//fmt.Println("10 / 0", amountTen.Div(amountZero))
}

func TestDecimal_Pow(t *testing.T) {
	precession := NewFromInt(6)

	amount := NewFromInt(10)
	amountRaw := NewFromInt(10000000)

	fmt.Println("amountRaw to amount", amount.Mul(NewFromInt(10).Pow(precession)))
	fmt.Println("amount to amountRaw", amountRaw.Div(NewFromInt(10).Pow(precession)))
}
