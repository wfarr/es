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

  switch flag.Arg(0) {
    case "top":
      health, state := getHealth(), getState()
      fmt.Printf("Cluster `%s` is %s\n", state.ClusterName, health.Status)
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

  fmt.Fprintf(os.Stderr, "%s [flags] [command] <[subcommand]>\n\n", os.Args[0])
  fmt.Fprintf(os.Stderr, "FLAGS\n\n")
  flag.PrintDefaults()
  fmt.Fprintf(os.Stderr, "\n")
  fmt.Fprintf(os.Stderr, "COMMANDS\n\n")
  fmt.Fprintf(os.Stderr, "    top      display overall health\n\n")
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
