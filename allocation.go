package main

import (
	"errors"
	"fmt"
)

var cmdAllocation = &Command{
	Run:   runAllocation,
	Usage: "allocation <setting>",
	Short: "control cluster allocation settings",
	Long: `
	Manage cluster allocation settings.

	For Elasticsearch clusters running 0.90.x, valid options are:
		* enable
		* disable

	For Elasticsearch clusters running 1.x, valid options are:
		* all (alias: enable)
		* primaries
		* new_primaries
		* none (alias: disable)
`,
}

func runAllocation(c *Cluster, cmd *Command, args []string) error {
	var settingName string
	var settingValue interface{}
	var foundValidValue bool

	if len(args) != 1 {
		return errors.New(cmd.renderUsage())
	}

	if c.one() {
		settingName = "cluster.routing.allocation.enable"
		validValues := [...]string{"all", "primaries", "new_primaries", "none", "enable", "disable"}

		for _, v := range validValues {
			if args[0] == v {
				foundValidValue = true

				switch v {
				case "enable":
					settingValue = "all"
				case "disable":
					settingValue = "none"
				default:
					settingValue = v
				}

			}
		}

		if !foundValidValue {
			return fmt.Errorf("received an invalid setting for cluster version %v: %v\n\n", c.versionNumber(), args[0])
		}

	} else if c.ohNinety() {
		settingName = "cluster.routing.allocation.disable_allocation"
		validValues := [...]string{"enable", "disable"}

		for _, v := range validValues {
			if args[0] == v {
				foundValidValue = true

				if v == "enable" {
					settingValue = false
				} else {
					settingValue = true
				}
			}
		}

		if !foundValidValue {
			return fmt.Errorf("received an invalid setting for cluster version %v: %v\n\n", c.versionNumber(), args[0])
		}
	} else {
		return errors.New("don't know anything about this cluster version")
	}

	newSettings := make(map[string]map[string]interface{})
	newSettings["persistent"] = make(map[string]interface{})
	newSettings["persistent"][settingName] = settingValue
	err := c.Stretch.SetSettings(newSettings)

	if err != nil {
		return fmt.Errorf("failed to update settings!\n\n%v", err)
	}

	fmt.Printf("Successfully set %v=%v\n", settingName, settingValue)

	return nil
}
