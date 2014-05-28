package main

import (
	"fmt"

	"github.com/wfarr/termtable"
)

var cmdNodes = &Command{
	Run:   runNodes,
	Usage: "nodes",
	Short: "display a list of nodes in the cluster",
	Long: `
	Display a list of all nodes in the cluster.
`,
}

func runNodes(c *Cluster, cmd *Command, args []string) error {
	nodes, err := c.Stretch.GetNodes()

	if err != nil {
		return err
	}

	t := termtable.NewTable(&termtable.TableOptions{Padding: 1, Header: []string{"NAME", "HOSTNAME", "VERSION", "HTTP ADDRESS", "ATTRIBUTES"}})

	for _, node := range nodes.Nodes {
		t.AddRow([]string{node.Name, node.Hostname + node.Host, node.Version, node.HTTPAddress, fmt.Sprintf("%v", node.Attributes)})
	}

	fmt.Println(t.Render())
	return nil
}
