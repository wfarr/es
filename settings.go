package main

import (
	"fmt"

	"github.com/wfarr/termtable"
)

var cmdSettings = &Command{
	Run:   runSettings,
	Usage: "settings",
	Short: "display cluster settings",
	Long: `
	Show a tabular display of all persistent and transient cluster settings.
`,
}

func runSettings(c *Cluster, cmd *Command, args []string) error {
	settings, err := c.Stretch.GetSettings()

	if err != nil {
		return err
	}

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
