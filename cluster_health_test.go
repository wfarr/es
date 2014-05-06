package main

import (
  "testing"
)

func TestClusterHealth(t *testing.T) {
  ts := testServer(`{
    "status": "tangerine",
    "cluster_name": "foobar",
    "timed_out" : false,
    "number_of_nodes" : 1,
    "number_of_data_nodes" : 1,
    "active_primary_shards" : 10,
    "active_shards" : 20,
    "relocating_shards" : 2,
    "initializing_shards" : 0,
    "unassigned_shards" : 0
  }`)

  defer ts.Close()

  cluster := &Cluster{&Client{URL: ts.URL}}
  health := cluster.GetHealth()

  if health.Status != "tangerine" {
    t.Fail()
  }
  if health.ClusterName != "foobar" {
    t.Fail()
  }
  if health.TimedOut != false {
    t.Fail()
  }
  if health.NumberOfNodes != 1 {
    t.Fail()
  }
  if health.NumberOfDataNodes != 1 {
    t.Fail()
  }
  if health.ActivePrimaryShards != 10 {
    t.Fail()
  }
  if health.ActiveShards != 20 {
    t.Fail()
  }
  if health.RelocatingShards != 2 {
    t.Fail()
  }
  if health.InitializingShards != 0 {
    t.Fail()
  }
  if health.UnassignedShards != 0 {
    t.Fail()
  }
}
