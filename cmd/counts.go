package cmd

import (
	"fmt"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Counts fetches counts of counter.
func Counts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "counts",
		Short: "Fetches counts of counter",
		RunE: func(cmd *cobra.Command, args []string) error {
			name, _ := cmd.Flags().GetString("name")
			groupBy, _ := cmd.Flags().GetString("group-by")
			orderBy, _ := cmd.Flags().GetString("order-by")

			c := counter.NewCounter()

			options := counter.CountOptions{
				Name:    name,
				GroupBy: groupBy,
				OrderBy: orderBy,
			}

			resp, err := c.Counts(options)
			if err != nil {
				return err
			}

			fmt.Println(prettyPrint(resp))

			return nil
		},
	}

	cmd.Flags().String("name", "", "Name")
	_ = cmd.MarkFlagRequired("name")

	cmd.Flags().String("group-by", "", "Group by count list")
	_ = cmd.MarkFlagRequired("group-by")

	cmd.Flags().String("order-by", "", "Order by count list")

	return cmd
}
