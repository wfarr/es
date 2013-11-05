package main

import (
	"fmt"
	"net/http"

	"net/http/httptest"
)

import "testing"

func SetupTestServer(handler func(w http.ResponseWriter, r *http.Request)) (ts *httptest.Server) {
	ts = httptest.NewServer(http.HandlerFunc(handler))
	return
}

func TestGetHealth(t *testing.T) {
	ts := SetupTestServer(func(w http.ResponseWriter, r *http.Request) {
		status := `{"status": "tangerine"}`

		fmt.Fprintf(w, status)
	})

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
