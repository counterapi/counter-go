package cmd

import (
	"fmt"
	"log"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Down uses Counter to counts down.
func Down() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down",
		Short: "Count down for given name",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			name, _ := cmd.Flags().GetString("name")

			c := counter.NewCounter()

			resp, err := c.Down(namespace, name)
			if err != nil {
				return err
			}

			fmt.Println(prettyPrint(resp))

			return nil
		},
	}

	cmd.Flags().String("namespace", "", "Namespace")
	if err := cmd.MarkFlagRequired("namespace"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("name", "", "Name")
	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}
