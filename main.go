package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/wfarr/stretch-go"
)

type Cluster struct {
	Stretch *stretch.Cluster
}

var cluster Cluster

type Command struct {
	// args does not include the command name
	Run  func(cluster *Cluster, cmd *Command, args []string)
	Flag flag.FlagSet

	Usage string // first word is the command name
	Short string // `es help` output
	Long  string // `es help cmd` output
}

func (c *Command) printUsage() {
	if c.Runnable() {
		fmt.Printf("Usage: es %s\n\n", c.Usage)
	}
	fmt.Println(strings.Trim(c.Long, "\n"))
}

func (c *Command) Name() string {
	name := c.Usage
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

func (c *Command) Runnable() bool {
	return c.Run != nil
}

const extra = " (extra)"

func (c *Command) List() bool {
	return c.Short != "" && !strings.HasSuffix(c.Short, extra)
}

func (c *Command) ListAsExtra() bool {
	return c.Short != "" && strings.HasSuffix(c.Short, extra)
}

func (c *Command) ShortExtra() string {
	return c.Short[:len(c.Short)-len(extra)]
}

// Running `es help` will list commands in this order.
var commands = []*Command{
	cmdHelp,
	cmdHealth,
	cmdAllocation,
}

var (
	flagApp  string
	flagLong bool
)

func main() {
	cluster := &Cluster{&stretch.Cluster{&stretch.Client{URL: "http://127.0.0.1:9200"}}}

	args := os.Args[1:]

	if len(args) < 1 {
		usage()
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Run != nil {
			cmd.Flag.Usage = func() {
				cmd.printUsage()
			}

			if err := cmd.Flag.Parse(args[1:]); err != nil {
				os.Exit(2)
			}

			cmd.Run(cluster, cmd, cmd.Flag.Args())
			return
		}
	}
}
