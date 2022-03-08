package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jeromelesaux/fizzbuzz/configuration"
	"github.com/jeromelesaux/fizzbuzz/model"
)

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
	os.Setenv("PORT", "3000")
	configuration.InitEnv()
	go main()
}
