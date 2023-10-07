package cmd

import (
	"fmt"
	"log"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Set sets counter.
func Set() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Sets counter",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			name, _ := cmd.Flags().GetString("name")
			count, _ := cmd.Flags().GetString("count")

			c := counter.NewCounter()

			resp, err := c.Set(namespace, name, count)
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

	cmd.Flags().String("count", "", "Count value for the counter")
	if err := cmd.MarkFlagRequired("count"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}
