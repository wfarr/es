package main

import (
	"strings"
)

type Cluster struct {
	Client *Client
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

type ClusterState struct {
	ClusterName string `json:"cluster_name"`
	MasterNode  string `json:"master_node"`
}

func (c *Cluster) Version() string {
	return c.GetVersion().Number
}

func (c *Cluster) GetVersion() ClusterVersion {
	ci := &ClusterInfo{}
	c.Client.Get(&ci, "/")

	return ci.Version
}

func (c *Cluster) GetState() (data ClusterState) {
	c.Client.Get(&data, "/_cluster/state")
	return
}

func (c *Cluster) OhNinety() bool {
	ver := cluster.GetVersion().Number
	segments := strings.Split(ver, ".")

	return segments[0] == "0" && segments[1] == "90"
}

func (c *Cluster) One() bool {
	ver := cluster.GetVersion().Number
	segments := strings.Split(ver, ".")

	return segments[0] == "1"
}
