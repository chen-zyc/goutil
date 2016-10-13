package gostring

import (
	"reflect"
	"testing"
)

func TestSplitTextToWords(t *testing.T) {
	examples := []struct {
		text string
		want []string
	}{
		{"中国", []string{"中", "国"}},
		{"abc", []string{"abc"}},
		{"abc汗", []string{"abc", "汗"}},
		{"abc 汗", []string{"abc", " ", "汗"}},
	}

	for _, e := range examples {
		words := SplitTextToWords([]byte(e.text))
		act := make([]string, len(words))
		for i, word := range words {
			act[i] = string(word)
		}
		if len(words) != len(e.want) || !reflect.DeepEqual(act, e.want) {
			t.Errorf("SplitTextToWords(%s) got %v, want %v", e.text, act, e.want)
		}
	}
}
