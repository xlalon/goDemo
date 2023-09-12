package mysql

import (
	"fmt"
	"testing"
)

func TestIdGenerator_NextID(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Println(NextID())
	}
}
