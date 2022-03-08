package engine

import (
	"testing"

	"github.com/jeromelesaux/fizzbuzz/model"
)

func TestFizzbuzzEngine(t *testing.T) {
	p := &model.Parameters{
		Limit: 30,
		Int1:  2,
		Int2:  3,
		Str1:  "FIZZ",
		Str2:  "BUZZ",
	}
	result := DoFizzbuzz(p)
	if len(result) != int(p.Limit) {
		t.Fatalf("expected the size of %d elements and gets %d\n", p.Limit, len(result))
	}
	t.Log(result)
}
