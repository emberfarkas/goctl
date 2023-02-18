package main

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestXxx(t *testing.T) {
	r, w := utf8.DecodeRuneInString("你好")
	t.Logf("%v, %v", r, w)
	if unicode.Is(unicode.Han, r) {
		t.Logf("汉字")
	}
}
