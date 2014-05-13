package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/wfarr/termtable"
)

var cmdAllocation = &Command{
	Run:   runAllocation,
	Usage: "allocation [<setting>]",
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

	If no settings is given, display the current cluster allocation settings.
`,
}

func runAllocation(c *Cluster, cmd *Command, args []string) error {
	var settingName string
	var settingValue interface{}
	var foundValidValue bool

	if len(args) > 1 {
		return errors.New(cmd.renderUsage())
	}

	if len(args) == 0 {
		settings := c.Stretch.GetSettings()

		t := termtable.NewTable(&termtable.TableOptions{Padding: 1, Header: []string{"SETTING TYPE", "SETTING NAME", "VALUE"}})

		t.AddRow([]string{"persistent", "cluster.routing.allocation.disable_allocation", strconv.FormatBool(settings.Persistent.ClusterRoutingAllocationDisableAllocation)})

		if settings.Persistent.ClusterRoutingAllocationEnable != "" {
			t.AddRow([]string{"persistent", "cluster.routing.allocation.enable", settings.Persistent.ClusterRoutingAllocationEnable})
		}

		t.AddRow([]string{"transient", "cluster.routing.allocation.disable_allocation", strconv.FormatBool(settings.Transient.ClusterRoutingAllocationDisableAllocation)})

		if settings.Persistent.ClusterRoutingAllocationEnable != "" {
			t.AddRow([]string{"persistent", "cluster.routing.allocation.enable", settings.Transient.ClusterRoutingAllocationEnable})
		}

		fmt.Println(t.Render())

		return nil
	}

	if c.One() {
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
			return errors.New(fmt.Sprintf("Received an invalid setting for cluster version %v: %v\n\n", c.VersionNumber(), args[0]))
		}

	} else if c.OhNinety() {
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
			return errors.New(fmt.Sprintf("Received an invalid setting for cluster version %v: %v\n\n", c.VersionNumber(), args[0]))
		}
	} else {
		return errors.New("Don't know anything about this cluster version!")
	}

	newSettings := make(map[string]map[string]interface{})
	newSettings["persistent"] = make(map[string]interface{})
	newSettings["persistent"][settingName] = settingValue
	err := c.Stretch.SetSettings(newSettings)

	if err != nil {
		return errors.New(fmt.Sprintf("Failed to update settings!\n\n%v", err))
	}

	fmt.Printf("Successfully set %v=%v\n", settingName, settingValue)

	return nil
}
