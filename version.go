package main

import (
	"fmt"
	"runtime"
)

var cmdVersion = &Command{
	Run:   runVersion,
	Usage: "version",
	Short: "show es version",
	Long:  `Version shows the es client version string.`,
}

func runVersion(cluster *Cluster, cmd *Command, args []string) error {
	fmt.Printf("es 0.2.1 (built with %v)\n", runtime.Version())
	return nil
}
