package cmd

import (
	"fmt"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Set sets counter.
func Set() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Sets counter",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			count, _ := cmd.Flags().GetString("count")

			c := counter.NewCounter()

			resp, err := c.Set(name, count)
			if err != nil {
				return err
			}

			fmt.Println(prettyPrint(resp))

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name")
	_ = cmd.MarkFlagRequired("name")

	cmd.Flags().String("count", "", "Count value for the counter")
	_ = cmd.MarkFlagRequired("count")

	return cmd
}
