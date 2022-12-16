package persistence

import (
	"testing"

	"github.com/jeromelesaux/fizzbuzz/model"
)

func TestInitialiseDB(t *testing.T) {
	initialiseDB()
}

func TestInsertNewParameters(t *testing.T) {
	p := model.Parameters{Int1: 1, Int2: 5, Str1: "fizz", Str2: "buzz", Limit: 10}
	for i := 0; i < 10; i++ {
		if err := AddInDB(p); err != nil {
			t.Fatalf(err.Error())
		}
	}

	err := DeleteStatsDB()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMostFrequentHitsInDb(t *testing.T) {
	err := DeleteStatsDB()
	if err != nil {
		t.Fatal(err)
	}
	p := model.Parameters{Int1: 1, Int2: 5, Str1: "fizz", Str2: "buzz", Limit: 100}
	for i := 0; i < 10; i++ {
		if err := AddInDB(p); err != nil {
			t.Fatalf(err.Error())
		}
	}

	p, err = GetMostFrequentDB()
	if err != nil {
		t.Fatalf(err.Error())
	}
	if p.Hits != 10 {
		t.Fatalf("expected 10 hits and gets %d hits in db\n", p.Hits)
	}
	t.Log(p)

}
