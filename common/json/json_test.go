package json

import (
	"fmt"
	"testing"
)

func TestJGet(t *testing.T) {
	fmt.Println(JGet(`{"a": 1}`, "a"))
}
