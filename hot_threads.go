package main

import (
	"fmt"
)

var cmdHotThreads = &Command{
	Run:   runHotThreads,
	Usage: "hot_threads",
	Short: "display the hot threads",
	Long: `
	Display all hot threads on all nodes.
`,
}

func runHotThreads(c *Cluster, cmd *Command, args []string) {
	hotThreads := c.Stretch.GetHotThreads()
	fmt.Println(hotThreads)
}
