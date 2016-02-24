package slice

import (
	"reflect"
	"testing"
)

func TestReserve(t *testing.T) {
	slice := []int{3, 9, 0, -3, -5}
	ReverseInt(slice)
	if !reflect.DeepEqual(slice, []int{-5, -3, 0, 9, 3}) {
		t.Error(slice)
	}
}
