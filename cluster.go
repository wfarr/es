package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Cluster struct {
	URL string
}

type ClusterVersion struct {
	Number         string `json:"number"`
	BuildHash      string `json:"build_hash"`
	BuildTimestamp string `json:"build_timestamp"`
	BuildSnapshot  bool   `json:"build_snapshot"`
	LuceneVersion  string `json:"lucene_version"`
}

type ClusterInfo struct {
	Ok      bool           `json:"ok"`
	Status  int            `json:"status"`
	Name    string         `json:"name"`
	Version ClusterVersion `json:"version"`
	Tagline string         `json:"tagline"`
}

type ClusterHealth struct {
	Status              string `json:"status"`
	ClusterName         string `json:"cluster_name"`
	TimedOut            bool   `json:"timed_out"`
	NumberOfNodes       int    `json:"number_of_nodes"`
	NumberOfDataNodes   int    `json:"number_of_data_nodes"`
	ActivePrimaryShards int    `json:"active_primary_shards"`
	ActiveShards        int    `json:"active_shards"`
	RelocatingShards    int    `json:"relocating_shards"`
	InitializingShards  int    `json:"initializing_shards"`
	UnassignedShards    int    `json:"unassigned_shards"`
}

type ClusterState struct {
	ClusterName string `json:"cluster_name"`
	MasterNode  string `json:"master_node"`
}

type ClusterSettings struct {
	Persistent SettingsYo `json:"persistent"`
	Transient  SettingsYo `json:"transient"`
}

type SettingsYo struct {
	AllocationDisabled bool `json:"allocation_disabled"`
}

func (c *Cluster) GetVersion() (ClusterVersion) {
	ci := &ClusterInfo{}
	c.get("/", &ci)

	return ci.Version
}

func (c *Cluster) GetHealth() (data ClusterHealth) {
	c.get("/_cluster/health", &data)
	return
}

func (c *Cluster) GetState() (data ClusterState) {
	c.get("/_cluster/state", &data)
	return
}

func (c *Cluster) GetSettings() (data ClusterSettings) {
	c.get("/_cluster/settings", &data)
	return
}

func (c *Cluster) get(path string, buf interface{}) interface{} {
	resp, err := http.Get(c.URL + "/" + path)

	if err != nil || resp.StatusCode > 200 {
		fmt.Printf("Could not find cluster at `%s`!\n", c.URL)
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	err = json.Unmarshal(body, &buf)

	if err != nil {
		fmt.Println("That's not JSON buddy!")
		os.Exit(1)
	}

	return buf
}
