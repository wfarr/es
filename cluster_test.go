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
	ts := MockServer(`{
		"status": "tangerine",
		"cluster_name": "foobar"
	}`)
	defer ts.Close()

	cluster := &Cluster{
		URL: ts.URL,
	}
	health := cluster.GetHealth()

	if health.Status != "tangerine" {
		t.Log("Expected status to be `tangerine`, got", health.Status)
		t.Fail()
	}

	if health.ClusterName != "foobar" {
		t.Log("Expected cluster name to be `foobar`, got", health.ClusterName)
		t.Fail()
	}
}
