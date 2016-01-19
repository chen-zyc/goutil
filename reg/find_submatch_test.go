package reg

import (
	"reflect"
	"regexp"
	"testing"
	"strconv"
)

func TestFindStringSubmatchMap(t *testing.T) {
	reg, err := regexp.Compile(`abc_(?P<n1>\d{4}), def_(?P<n2>\d{4}), ghi_(?P<n3>\d{4})`)
	if err != nil {
		t.Error(err)
		return
	}
	text := "abc_1234, def_1235, ghi_1236"
	match := FindStringSubmatchMap(reg, text)
	if len(match) == 0 {
		t.Error("cann't match " + text)
		return
	}
	if len(match) != 3 {
		t.Errorf("the length of submatch should be 3, but " + strconv.Itoa(len(match)))
		return
	}
	wantMath := map[string]string{
		"n1": "1234",
		"n2": "1235",
		"n3": "1236",
	}
	if !reflect.DeepEqual(match, wantMath) {
		t.Errorf("got match[%v], but want[%v]", match, wantMath)
	}
}
