package main

import (
	"errors"
	"fmt"

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

		t := termtable.NewTable(&termtable.TableOptions{Padding: 1, MaxColWidth: 90, Header: []string{"SETTING NAME", "VALUE"}})

		t.AddRow([]string{"PERSISTENT SETTINGS", ""})

		for key, value := range settings.Persistent {
			t.AddRow([]string{key, value})
		}

		t.AddRow([]string{"", ""})

		t.AddRow([]string{"TRANSIENT SETTINGS", ""})

		for key, value := range settings.Transient {
			t.AddRow([]string{key, value})
		}

		fmt.Println(t.Render())

		return nil
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
