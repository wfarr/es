package main

import (
	"fmt"
	"net/http"

	"net/http/httptest"

	// "testing"

	"github.com/wfarr/stretch-go"
)

func testServer(resp string) (ts *httptest.Server) {
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, resp)
	}))
	return
}

func makeClusterForTestServer(ts *httptest.Server) *Cluster {
	return &Cluster{&stretch.Cluster{&stretch.Client{URL: ts.URL}}}
}
