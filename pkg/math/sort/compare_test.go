package sort

import "testing"

func TestMax(t *testing.T) {
	values := []int64{-10, -999, 1, 3, 10, 23, 1, 2, 999}
	maxValue := Max(values)
	if maxValue != 999 {
		t.Fatal("Max Fail")
	}
}

func TestMin(t *testing.T) {
	values := []int64{-10, -999, 1, 3, 10, 23, 1, 2, 999}
	minValue := Min(values)
	if minValue != -999 {
		t.Fatal("Max Fail")
	}
}
