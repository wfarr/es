package main

import (
  "testing"
)

func TestClusterSettingsOhNinety(t *testing.T) {
  ts := testServer(`{
    "persistent": {
      "cluster.routing.allocation.disable_allocation": false,
      "cluster.routing.allocation.disable_replica_allocation": false
    },
    "transient": {
      "cluster.routing.allocation.disable_allocation": true,
      "cluster.routing.allocation.disable_replica_allocation": true
    }
  }`)

  defer ts.Close()

  cluster := &Cluster{URL: ts.URL}
  settings := cluster.GetSettings()

  if (settings.Persistent.ClusterRoutingAllocationDisableAllocation != false) {
    t.Fail()
  }

  if (settings.Persistent.ClusterRoutingAllocationDisableReplicaAllocation != false) {
    t.Fail()
  }

  if (!settings.Transient.ClusterRoutingAllocationDisableAllocation) {
    t.Fail()
  }

  if (!settings.Transient.ClusterRoutingAllocationDisableReplicaAllocation) {
    t.Fail()
  }
}

func TestClusterSettingsOneOh(t *testing.T) {
  ts := testServer(`{
    "persistent": {
      "cluster.routing.allocation.enable": "all"
    },
    "transient": {
      "cluster.routing.allocation.enable": "new_primaries"
    }
  }`)

  defer ts.Close()

  cluster := &Cluster{URL: ts.URL}
  settings := cluster.GetSettings()

  if (settings.Persistent.ClusterRoutingAllocationEnable != "all") {
    t.Fail()
  }

  if (settings.Transient.ClusterRoutingAllocationEnable != "new_primaries") {
    t.Fail()
  }
}
