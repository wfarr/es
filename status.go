package main

import (
	"fmt"
)

var cmdStatus = &Command{
	Run:   runStatus,
	Usage: "status",
	Short: "display the high-level status of the cluster",
	Long: `
  Displays general cluster health information, such as the cluster state,
  the master node, and the current shard counts.
`,
}

func runStatus(cluster *Cluster, cmd *Command, args []string) {
	health, state, settings := cluster.GetHealth(), cluster.GetState(), cluster.GetSettings()
	var allocation string

	if settings.Persistent.AllocationDisabled || settings.Transient.AllocationDisabled {
		allocation = "disabled"
	} else {
		allocation = "enabled"
	}

	fmt.Printf("Cluster:\n  Name: %s\n  State: %s\n  Master: %s\n  Allocation: %s\n  Relocating Shards: %v\n  Initializing Shards: %v\n  Unassigned Shards: %v\n",
		health.ClusterName,
		health.Status,
		state.MasterNode,
		allocation,
		health.RelocatingShards,
		health.InitializingShards,
		health.UnassignedShards,
	)
}
