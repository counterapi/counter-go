package cmd

import (
	"fmt"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Down uses Counter to counts down.
func Down() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down",
		Short: "Count down for given name",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")

			c := counter.NewCounter()

			resp, err := c.Down(name)
			if err != nil {
				return err
			}

			fmt.Println(prettyPrint(resp))

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}
