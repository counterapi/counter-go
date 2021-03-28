package cmd

import (
	"fmt"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Get fetches counter information.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Fetches counter information",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")

			c := counter.NewCounter()

			resp, err := c.Get(name)
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
