package main

import (
	"fmt"
	"strconv"

	"github.com/wfarr/termtable"
)

var cmdHealth = &Command{
	Run:   runHealth,
	Usage: "health",
	Short: "display the health of the cluster",
	Long: `
  Displays general cluster health information.
`,
}

func runHealth(c *Cluster, cmd *Command, args []string) {
	health := c.Stretch.GetHealth()

	t := termtable.NewTable(&termtable.TableOptions{Padding: 1, Header: []string{"CLUSTER HEALTH", ""}})

	t.AddRow([]string{"Name", health.ClusterName})
	t.AddRow([]string{"Status", health.Status})
	t.AddRow([]string{"Timed Out", strconv.FormatBool(health.TimedOut)})
	t.AddRow([]string{"Number of Nodes", strconv.Itoa(health.NumberOfNodes)})
	t.AddRow([]string{"Number of Data Nodes", strconv.Itoa(health.NumberOfDataNodes)})
	t.AddRow([]string{"Active Primary Shards", strconv.Itoa(health.ActivePrimaryShards)})
	t.AddRow([]string{"Active Shards", strconv.Itoa(health.ActiveShards)})
	t.AddRow([]string{"Relocating Shards", strconv.Itoa(health.RelocatingShards)})
	t.AddRow([]string{"Initializing Shards", strconv.Itoa(health.InitializingShards)})
	t.AddRow([]string{"Unassigned Shards", strconv.Itoa(health.UnassignedShards)})
	fmt.Println(t.Render())
}
