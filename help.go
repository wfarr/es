package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var helpEnviron = &Command{
	Usage: "environ",
	Short: "environment variables used by es",
	Long: `
Several environment variables affect es's behavior.

ELASTICSEARCH_HOST

  Its default value is 127.0.0.1.

ELASTICSEARCH_PORT

  Its default value is 9200.

`,
}

var cmdVersion = &Command{
	Run:   runVersion,
	Usage: "version",
	Short: "show es version",
	Long:  `Version shows the es client version string.`,
}

func runVersion(cluster *Cluster, cmd *Command, args []string) error {
	fmt.Println("0.1.1")
	return nil
}

var cmdHelp = &Command{
	Usage: "help [topic]",
	Long:  `Help shows usage for a command or other topic.`,
}

func init() {
	cmdHelp.Run = runHelp // break init loop
}

func runHelp(cluster *Cluster, cmd *Command, args []string) error {
	if len(args) == 0 {
		printUsage()
		return nil // not os.Exit(2); success
	}
	if len(args) != 1 {
		log.Fatal("too many arguments")
	}

	for _, cmd := range commands {
		if cmd.name() == args[0] {
			cmd.printUsage()
			return nil
		}
	}

	return fmt.Errorf("unknown help topic: %q. Run 'es help'.\n", args[0])
}

var usageTemplate = template.Must(template.New("usage").Parse(`
Usage: es [command] [options] [arguments]


Commands:
{{range .Commands}}{{if .Runnable}}{{if .List}}
    {{.Name | printf "%-8s"}}  {{.Short}}{{end}}{{end}}{{end}}

Run 'es help [command]' for details.

`[1:]))

func printUsage() {
	usageTemplate.Execute(os.Stdout, struct {
		Commands []*Command
		Dev      bool
	}{
		commands,
		false,
	})
}
