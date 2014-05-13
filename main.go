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
	if c.runnable() {
		output += fmt.Sprintf("Usage: es %s\n\n", c.Usage)
	}
	output += fmt.Sprintf(strings.Trim(c.Long, "\n"))

	return
}

func (c *Command) printUsage() {
	fmt.Println(c.renderUsage())
}

func (c *Command) name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) runnable() bool {
	return c.Run != nil
}

const extra = " (extra)"

func (c *Command) list() bool {
	return c.Short != "" && !strings.HasSuffix(c.Short, extra)
}

func (c *Command) listAsExtra() bool {
	return c.Short != "" && strings.HasSuffix(c.Short, extra)
}

func (c *Command) shortExtra() string {
	return c.Short[:len(c.Short)-len(extra)]
}

// Running `es help` will list commands in this order.
var commands = []*Command{
	cmdHelp,
	cmdHealth,
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
		if cmd.name() == args[0] && cmd.Run != nil {
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
}
