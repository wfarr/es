package main

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

func (c *Cluster) GetVersion() (ClusterVersion) {
	ci := &ClusterInfo{}
	c.Client.Get(&ci, "/")

	return ci.Version
}

func (c *Cluster) GetHealth() (data ClusterHealth) {
	c.Client.Get(&data, "/_cluster/health")
	return
}

func (c *Cluster) GetState() (data ClusterState) {
	c.Client.Get(&data, "/_cluster/state")
	return
}
