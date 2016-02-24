package slice

import (
	"reflect"
	"testing"
)

func TestShiftLeftInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ShiftLeftInt(slice, 3)
	if !reflect.DeepEqual(slice, []int{4, 5, 6, 7, 8, 9, 1, 2, 3}) {
		t.Error(slice)
	}
}

func TestShiftRightInt(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ShiftRightInt(slice, 3)
	if !reflect.DeepEqual(slice, []int{7, 8, 9, 1, 2, 3, 4, 5, 6}) {
		t.Error(slice)
	}
}
