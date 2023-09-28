package lab_5_test

import (
	"testing"

	"lab_5"
)

func TestFindMinMax(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	min, max := lab_5.FindMinMax(arr)

	if min != 1 {
		t.Errorf("Expected minimum value to be 1 but got %d", min)
	}
	if max != 5 {
		t.Errorf("Expected maximum value to be 5 but got %d", max)
	}
}
