package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/wfarr/stretch-go"
)

// Cluster wrapper around stretch.Cluster to hide details of that type
type Cluster struct {
	Stretch *stretch.Cluster
}

var cluster Cluster

// Command object for our CLI sub-commands
type Command struct {
	// args does not include the command name
	Run  func(cluster *Cluster, cmd *Command, args []string) error
	Flag flag.FlagSet

	Usage string // first word is the command name
	Short string // `es help` output
	Long  string // `es help cmd` output
}

func (c *Command) renderUsage() (output string) {
	if c.Runnable() {
		output += fmt.Sprintf("Usage: es %s\n\n", c.Usage)
	}
	output += fmt.Sprintf(strings.Trim(c.Long, "\n"))

	return
}

func (c *Command) printUsage() {
	fmt.Println(c.renderUsage())
}

// Name returns the name of the command as a string
func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// Runnable returns true if there is a func assigned to c.Run
func (c *Command) Runnable() bool {
	return c.Run != nil
}

const extra = " (extra)"

// List returns true if the command should be listed (ie. is not an extra command)
func (c *Command) List() bool {
	return c.Short != "" && !strings.HasSuffix(c.Short, extra)
}

// ListAsExtra returns true if the command should be listed as an extra instead.
func (c *Command) ListAsExtra() bool {
	return c.Short != "" && strings.HasSuffix(c.Short, extra)
}

// ShortExtra returns the short usage of the extra section, without the extra text
func (c *Command) ShortExtra() string {
	return c.Short[:len(c.Short)-len(extra)]
}

// Running `es help` will list commands in this order.
var commands = []*Command{
	cmdHelp,
	cmdVersion,
	cmdHealth,
	cmdSettings,
	cmdAllocation,
	cmdNodes,
	cmdHotThreads,
}

var (
	flagApp  string
	flagLong bool
)

func main() {
	cluster := &Cluster{&stretch.Cluster{&stretch.Client{URL: "http://127.0.0.1:9200"}}}

	args := os.Args[1:]

	if len(args) < 1 {
		printUsage()
		os.Exit(2)
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() {
				cmd.printUsage()
			}

			if err := cmd.Flag.Parse(args[1:]); err != nil {
				os.Exit(2)
			}

			err := cmd.Run(cluster, cmd, cmd.Flag.Args())

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			return
		}
	}

	printUsage()
	os.Exit(2)
}
