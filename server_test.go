package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"testing"

	"github.com/jeromelesaux/fizzbuzz/configuration"
	"github.com/jeromelesaux/fizzbuzz/model"
)

var serverIsLaunched = false

func TestServerStatsResponse200(t *testing.T) {
	// launch in background the server
	launchServer()

	// preparing the request
	r, err := http.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/stats", nil)
	if err != nil {
		t.Fatal(err)
	}
	c := http.DefaultClient

	// execute query
	resp, err := c.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	// check http server response
	if resp.StatusCode != 200 {
		t.Fatalf("expected http status code and gets %d\n", resp.StatusCode)
	}

	// parse the response body
	defer resp.Body.Close()
	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	buf := bytes.NewBuffer(msg)
	p := &model.Parameters{}

	// check json format
	if err := json.NewDecoder(buf).Decode(p); err != nil {
		t.Fatal(err)
	}

	// check the  objectcontent
	if p.Hits != 0 {
		t.Fatalf("expected 0 hits and gets %d\n", p.Hits)
	}
}

func TestSimpleFizbuzzResponse200(t *testing.T) {
	// launch in background the server
	launchServer()

	// preparing the request
	r, err := http.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/fizzbuzz?int1=2&int2=3&str1=fizz&str2=buzz&limit=10", nil)
	if err != nil {
		t.Fatal(err)
	}
	c := http.DefaultClient

	// execute query
	resp, err := c.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	// check http server response
	if resp.StatusCode != 202 {
		t.Fatalf("expected http status code and gets %d\n", resp.StatusCode)
	}
}

func launchServer() {
	if serverIsLaunched == true {
		return
	}
	os.Setenv("PORT", "3000")
	configuration.InitEnv()
	go main()
	serverIsLaunched = true
}

func benchServerResponses(i int) {
	int1 := rand.Intn(i + 1)
	int2 := rand.Intn(i + 1)
	limit := rand.Intn(1000)
	query := fmt.Sprintf("http://localhost:3000/api/v1/fizzbuzz?int1=%d&int2=%d&str1=fizz&str2=buzz&limit=%d", int1, int2, limit)
	r, _ := http.NewRequest(http.MethodGet, query, nil)
	r.Close = true
	c := http.DefaultClient
	// execute query
	res, _ := c.Do(r)
	res.Body.Close()
}

func BenchmarkFizzbuzz10(b *testing.B) {
	launchServer()
	for i := 0; i < b.N; i++ {
		benchServerResponses(i)
	}
}

func BenchmarkFizzbuzz20(b *testing.B) {
	launchServer()
	for i := 0; i < b.N; i++ {
		benchServerResponses(i)
	}
}

func BenchmarkFizzbuzz30(b *testing.B) {
	launchServer()
	for i := 0; i < b.N; i++ {
		benchServerResponses(i)
	}
}

func BenchmarkFizzbuzz40(b *testing.B) {
	launchServer()
	for i := 0; i < b.N; i++ {
		benchServerResponses(i)
	}
}

func BenchmarkFizzbuzz50(b *testing.B) {
	launchServer()
	for i := 0; i < b.N; i++ {
		benchServerResponses(i)
	}
}
