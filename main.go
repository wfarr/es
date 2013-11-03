package main

import (
  "flag"
  "fmt"
  "os"
)

var ip string
var port string

func main() {
  flag.StringVar(&ip, "ip", "127.0.0.1", "The host elasticsearch is running on")
  flag.StringVar(&port, "port", "9200", "The port elasticsearch is running on")

  flag.Parse()

  cluster := &Cluster{
    Ip: ip,
    Port: port,
  }

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
