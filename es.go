package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "os"
  "net/http"
)

func main() {
/*  ip := flag.String("ip", "127.0.0.1", "The host elasticsearch is running on")*/
/*  port := flag.String("port", "9200", "The port elasticsearch is running on")*/

  flag.Parse()

  switch flag.Arg(0) {
    case "top":
      health, state := getHealth(), getState()
      fmt.Printf("Cluster `%s` is %s\n", state.ClusterName, health.Status)
    default:
      flag.Usage()
      os.Exit(1)
  }

  os.Exit(0)
}

func getHealth() (data ClusterHealth) {
  get("/_cluster/health", &data)
  return
}

func getState() (data ClusterState) {
  get("/_cluster/state", &data)
  return
}

func get(path string, buf interface{}) (interface{}) {
  resp, err := http.Get("http://127.0.0.1:19200" + path)

  if err != nil || resp.StatusCode > 200 {
    fmt.Println("Could not find cluster at `127.0.0.1:19200`!")
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
