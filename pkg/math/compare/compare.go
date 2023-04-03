package compare

import "math"

func Max(values []int64) int64 {
	var maxValue int64
	for _, value := range values {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func Min(values []int64) int64 {
	var minValue int64 = math.MaxInt64
	for _, value := range values {
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}
