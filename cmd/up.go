package cmd

import (
	"fmt"
	"log"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Up uses Counter to counts up.
func Up() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "Count up for given name",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			name, _ := cmd.Flags().GetString("name")

			c := counter.NewCounter()

			resp, err := c.Up(namespace, name)
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
