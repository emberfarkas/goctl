package main

import (
	"fmt"
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

func TestInt(t *testing.T) {
	n := fmt.Sprintf("%.4f", float64(11000)/float64(10000))
	t.Errorf("%v", n) // 1.1000
}

func TestInt1(t *testing.T) {

}
