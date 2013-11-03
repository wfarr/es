package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "net/http"
)

var ip string
var port string

func main() {
  flag.StringVar(&ip, "ip", "127.0.0.1", "The host elasticsearch is running on")
  flag.StringVar(&port, "port", "9200", "The port elasticsearch is running on")

  flag.Parse()

  cluster := new(Cluster)
  cluster.Ip = ip
  cluster.Port = port

  switch flag.Arg(0) {
    case "status":
      health, _, settings := cluster.GetHealth(), cluster.GetState(), cluster.GetSettings()
      var allocation string

      if (settings.Persistent.AllocationDisabled || settings.Transient.AllocationDisabled) {
        allocation = "disabled"
      } else {
        allocation = "enabled"
      }

      fmt.Printf("Cluster:\n  Name: %s\n  State: %s\n  Allocation: %s\n  Relocating Shards: %v\n  Initializing Shards: %v\n  Unassigned Shards: %v\n",
        health.ClusterName,
        health.Status,
        allocation,
        health.RelocatingShards,
        health.InitializingShards,
        health.UnassignedShards)
    default:
      usage()
      os.Exit(1)
  }

  os.Exit(0)
}

func usage() {
  if !flag.Parsed() {
    flag.Parse()
  }

  fmt.Fprintf(os.Stderr, "Usage: %s [flags] [command] <[subcommand]>\n\n", os.Args[0])
  fmt.Fprintf(os.Stderr, "FLAGS\n\n")
  flag.PrintDefaults()
  fmt.Fprintf(os.Stderr, "\n")
  fmt.Fprintf(os.Stderr, "COMMANDS\n\n")
  fmt.Fprintf(os.Stderr, "    status      display overall health\n\n")
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
