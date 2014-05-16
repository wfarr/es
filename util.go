package main

import (
	"strings"
)

func (c *Cluster) versionNumber() (string, error) {
	info, err := c.Stretch.GetInfo()

	if err != nil {
		return "", err
	}

	return info.Version.Number, nil
}

func (c *Cluster) ohNinety() (bool, error) {
	ver, err := c.versionNumber()

	if err != nil {
		return false, err
	}

	segments := strings.Split(ver, ".")

	return segments[0] == "0" && segments[1] == "90", nil
}

func (c *Cluster) one() (bool, error) {
	ver, err := c.versionNumber()

	if err != nil {
		return false, err
	}

	segments := strings.Split(ver, ".")

	return segments[0] == "1", nil
}
