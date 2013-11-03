package main

import (
  "encoding/json"
  "io/ioutil"
  "net/http"
  "fmt"
  "os"
)

type Cluster struct {
  Ip string
  Port string
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

func (c *Cluster) GetHealth() (data ClusterHealth) {
  get("/_cluster/health", &data)
  return
}

func (c *Cluster) GetState() (data ClusterState) {
  get("/_cluster/state", &data)
  return
}

func (c *Cluster) GetSettings() (data ClusterSettings) {
  get("/_cluster/settings", &data)
  return
}

func get(path string, buf interface{}) (interface{}) {
  resp, err := http.Get("http://" + ip + ":" + port + path)

  if err != nil || resp.StatusCode > 200 {
    fmt.Printf("Could not find cluster at `%s:%s`!\n", ip, port)
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
