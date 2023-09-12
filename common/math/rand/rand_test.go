package rand

import (
	"fmt"
	"testing"
)

func TestUIntNBetween(t *testing.T) {
	fmt.Println(UIntNBetween(100, 1000))
}

func TestDigitalMemo(t *testing.T) {
	fmt.Println(DigitalMemo())
}
