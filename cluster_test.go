package main

import (
	"fmt"
	"net/http"

	"net/http/httptest"
)

import "testing"

func MockServer(resp string) (ts *httptest.Server) {
	ts = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, resp)
	}))
	return
}

func TestGetHealth(t *testing.T) {
	ts := MockServer(`{"status": "tangerine"}`)
	defer ts.Close()

	cluster := &Cluster{
		URL: ts.URL,
	}

	status := cluster.GetHealth().Status

	if status != "tangerine" {
		t.Log("Expected status to be `tangerine`, got", status)
		t.Fail()
	}
}
