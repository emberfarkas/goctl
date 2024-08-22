package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/bytedance/sonic"
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

type Data struct {
	Now time.Time
}

func TestInt1(t *testing.T) {
	d := Data{
		Now: time.Now(),
	}
	x, _ := sonic.Marshal(d)
	var xx Data
	err := sonic.Unmarshal([]byte(`{"Now":"2023-12-06 17:12:24"}`), &xx)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%v\n", d.Now)
	if err := json.Unmarshal(x, &xx); err != nil {
		return
	}
	fmt.Printf("%v\n", d.Now)
}
