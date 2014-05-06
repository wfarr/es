package main

type FullClusterSettings struct {
  Persistent ClusterSettings `json:"persistent"`
  Transient  ClusterSettings `json:"transient"`
}

type ClusterSettings struct {
  // 0.90 allocation settings
  ClusterRoutingAllocationDisableAllocation        bool   `json:"cluster.routing.allocation.disable_allocation,omitempty"`
  ClusterRoutingAllocationDisableReplicaAllocation bool   `json:"cluster.routing.allocation.disable_replica_allocation,omitempty"`

  // 1.0 allocation settings
  ClusterRoutingAllocationEnable                   string `json:"cluster.routing.allocation.enable"`
}

func (c *Cluster) GetSettings() (data FullClusterSettings) {
  c.Client.Get(&data, "/_cluster/settings")
  return
}

func (c *Cluster) SetSettings(settings interface{}) (err error) {
  err = c.Client.Put(nil, "/_cluster/settings", settings)
  return
}
