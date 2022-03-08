package model

import (
	"net/http"
	"testing"
)

func TestParsingParameters(t *testing.T) {
	// init an empty request
	r, err := http.NewRequest(http.MethodGet, "http://localhost?int1=10&int2=101&limit=5&str1=aaaa&str2=bbb", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// parsing the http  parameters
	p, err := ParseParameters(r)
	// check parsing
	if err != nil {
		t.Fatalf(err.Error())
	}

	// check values
	if p.Int1 != 10 {
		t.Fatalf("error for int1 parameter expected 10 and gets %d\n", p.Int1)
	}
	if p.Int2 != 101 {
		t.Fatalf("error for int2 parameter expected 10 and gets %d\n", p.Int2)
	}
	if p.Limit != 5 {
		t.Fatalf("error for int1 parameter expected 5 and gets %d\n", p.Limit)
	}
	if p.Str1 != "aaaa" {
		t.Fatalf("error for str1 parameter expected 5 and gets %s\n", p.Str1)
	}
	if p.Str2 != "bbb" {
		t.Fatalf("error for str2 parameter expected 5 and gets %s\n", p.Str2)
	}
}

func TestEmptyParameters(t *testing.T) {
	r, err := http.NewRequest(http.MethodGet, "http://localhost?int1=10&int2=101&limit=5&str1=aaaa&str2=bbb", nil)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// parsing the http  parameters
	_, err = ParseParameters(r)
	// check parsing
	if err != nil {
		t.Fatalf(err.Error())
	}
}
