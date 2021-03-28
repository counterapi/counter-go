package cmd

import (
	"fmt"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Up uses Counter to counts up.
func Up() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "up",
		Short: "Count up for given name",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")

			c := counter.NewCounter()

			resp, err := c.Up(name)
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
