package main

import (
	"strings"
)

func (c *Cluster) versionNumber() string {
	return c.Stretch.GetInfo().Version.Number
}

func (c *Cluster) ohNinety() bool {
	ver := c.versionNumber()
	segments := strings.Split(ver, ".")

	return segments[0] == "0" && segments[1] == "90"
}

func (c *Cluster) one() bool {
	ver := c.versionNumber()
	segments := strings.Split(ver, ".")

	return segments[0] == "1"
}
