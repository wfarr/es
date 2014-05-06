package main

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


func (c *Cluster) GetHealth() (data ClusterHealth) {
  c.Client.Get(&data, "/_cluster/health")
  return
}
