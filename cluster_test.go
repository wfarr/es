package main

import (
	"fmt"
	"net/http"

	"net/http/httptest"

	"testing"
)

func testServer(resp string) (ts *httptest.Server) {
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, resp)
	}))
	return
}

func TestClusterVersion(t *testing.T) {
	ts := testServer(`{
  "ok" : true,
  "status" : 200,
  "name" : "boxen",
  "version" : {
    "number" : "0.90.5",
    "build_hash" : "c8714e8e0620b62638f660f6144831792b9dedee",
    "build_timestamp" : "2013-09-17T12:50:20Z",
    "build_snapshot" : false,
    "lucene_version" : "4.4"
  },
  "tagline" : "You Know, for Search"
	}`)

	defer ts.Close()

	cluster := &Cluster{&Client{URL: ts.URL}}
	v := cluster.GetVersion()

	if v.Number != "0.90.5" {
		t.Fail()
	}

	if v.BuildHash != "c8714e8e0620b62638f660f6144831792b9dedee" {
		t.Fail()
	}

	if v.BuildTimestamp != "2013-09-17T12:50:20Z" {
		t.Fail()
	}

	if v.BuildSnapshot != false {
		t.Fail()
	}

	if v.LuceneVersion != "4.4" {
		t.Fail()
	}
}
