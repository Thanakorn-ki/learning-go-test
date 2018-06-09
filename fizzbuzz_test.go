package fizzbuzz

import (
	"testing"
)

func TestTextC(t *testing.T) {
	f := Fizzbuzz("2")
	expected := "2"
	if f != expected {
		t.Errorf("\n%#v\nis expected but got\n%#v\n", expected, f)
	}
}

func TestTextCC(t *testing.T) {
	f := Fizzbuzz("2")
	expected := "2"
	if f != expected {
		t.Errorf("\n%#v\nis expected but got\n%#v\n", expected, f)
	}
}
