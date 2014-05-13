package main

import (
	"fmt"
	"strings"
)

var cmdHotThreads = &Command{
	Run:   runHotThreads,
	Usage: "hot_threads [<node name> [<node name 2> ...]]",
	Short: "display hot threads",
	Long: `
	Display all hot threads on all nodes, or just a single node given a node name.
`,
}

func runHotThreads(c *Cluster, cmd *Command, args []string) error {
	hotThreads := c.Stretch.GetHotThreads(args...)

	if strings.Trim(hotThreads, "\n") == "" {
		return fmt.Errorf("couldn't find any nodes for %v", args)
	}

	fmt.Println(hotThreads)
	return nil
}
