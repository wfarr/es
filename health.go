package main

import (
	"fmt"
	"strconv"

	"github.com/wfarr/stretch-go"
	"github.com/wfarr/termtable"
)

var cmdHealth = &Command{
	Run:   runHealth,
	Usage: "health [index]",
	Short: "display the health of the cluster",
	Long: `
	Displays general cluster health information.

	If the argument 'index' is given, displays health by-index.
`,
}

func runHealth(c *Cluster, cmd *Command, args []string) error {
	health := c.Stretch.GetHealth()

	if len(args) > 0 && args[0] == "index" {
		return renderIndexHealth(health)
	}

	return renderClusterHealth(health)
}

func renderIndexHealth(health stretch.ClusterHealth) error {
	t := termtable.NewTable(&termtable.TableOptions{
		Padding: 1,
		Header: []string{
			"INDEX",
			"STATUS",
			"SHARDS",
			"REPLICAS",
			"ACT. PRIM. SHARDS",
			"ACTIVE SHARDS",
			"RELOCATING",
			"INITIALIZING",
			"UNASSIGNED",
		},
	})

	for indexName, indexHealth := range health.Indices {
		t.AddRow([]string{
			indexName,
			indexHealth.Status,
			strconv.Itoa(indexHealth.NumberOfShards),
			strconv.Itoa(indexHealth.NumberOfReplicas),
			strconv.Itoa(indexHealth.ActivePrimaryShards),
			strconv.Itoa(indexHealth.ActiveShards),
			strconv.Itoa(indexHealth.RelocatingShards),
			strconv.Itoa(indexHealth.InitializingShards),
			strconv.Itoa(indexHealth.UnassignedShards),
		})
	}

	fmt.Println(t.Render())

	return nil
}

func renderClusterHealth(health stretch.ClusterHealth) error {
	t := termtable.NewTable(&termtable.TableOptions{
		Padding: 1,
		Header:  []string{"CLUSTER HEALTH", ""},
	})

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

	return nil
}
