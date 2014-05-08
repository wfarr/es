package main

import (
	"strings"
)

func (c *Cluster) VersionNumber() string {
	return c.Stretch.GetInfo().Version.Number
}

func (c *Cluster) OhNinety() bool {
	ver := c.VersionNumber()
	segments := strings.Split(ver, ".")

	return segments[0] == "0" && segments[1] == "90"
}

func (c *Cluster) One() bool {
	ver := c.VersionNumber()
	segments := strings.Split(ver, ".")

	return segments[0] == "1"
}
