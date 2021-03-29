package main

import (
	"github.com/counterapi/counter/cmd"

	"github.com/spf13/cobra"
)

// RootCommand is the main entry point of this application.
func RootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:          "counter",
		Short:        "Counter API CLI",
		Long:         "CLI command to counter things",
		SilenceUsage: true,
	}

	root.AddCommand(cmd.Version())
	root.AddCommand(cmd.Up())
	root.AddCommand(cmd.Down())
	root.AddCommand(cmd.Get())
	root.AddCommand(cmd.Counts())
	root.AddCommand(cmd.Set())

	return root
}

func main() {
	root := RootCommand()
	_ = root.Execute()
}
