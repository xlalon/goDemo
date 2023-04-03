package compare

import "testing"

func TestMax(t *testing.T) {
	values := []int64{1, 3, 10, 23, 1, 2}
	maxValue := Max(values)
	if maxValue != 23 {
		t.Fatal("Max Fail")
	}
}

func TestMin(t *testing.T) {
	values := []int64{1, 3, 10, 23, 1, 2, 0}
	minValue := Min(values)
	if minValue != 0 {
		t.Fatal("Max Fail")
	}
}
