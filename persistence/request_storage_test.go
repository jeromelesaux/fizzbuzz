package persistence

import (
	"sort"
	"testing"

	"github.com/jeromelesaux/fizzbuzz/model"
)

func TestSortPairs(t *testing.T) {
	pair0 := Pair{P: model.Parameters{
		Int1:  1,
		Int2:  2,
		Limit: 10,
		Str1:  "aa",
		Str2:  "bb",
	},
		V: 10}
	pair1 := Pair{P: model.Parameters{
		Int1:  1,
		Int2:  3,
		Limit: 10,
		Str1:  "aa",
		Str2:  "bb",
	},
		V: 5}

	pairList := make(Pairlist, 0)
	pairList = append(pairList, pair1)
	pairList = append(pairList, pair0)

	sort.Sort(sort.Reverse(pairList))
	if pairList[0].V != 10 {
		t.Fatalf("expected 10 occurence and gets %d\n", pairList[0].V)
	}

}

func TestAdd(t *testing.T) {
	param := model.Parameters{
		Int1:  1,
		Int2:  2,
		Limit: 10,
		Str1:  "aa",
		Str2:  "bb",
	}
	emptyParam := model.Parameters{}
	for i := 0; i < 5; i++ {
		Add(param)
	}
	for i := 0; i < 10; i++ {
		Add(emptyParam)
	}
	if RequestsStored[emptyParam] != 10 {
		t.Fatalf("expected 10 occurence for empty param and gets %d\n", RequestsStored[emptyParam])
	}
	if RequestsStored[param] != 5 {
		t.Fatalf("expected 5 occurence for  param and gets %d\n", RequestsStored[param])
	}
}
