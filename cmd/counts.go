package cmd

import (
	"fmt"
	"log"

	"github.com/counterapi/counter/pkg/counter"

	"github.com/spf13/cobra"
)

// Counts fetches counts of counter.
func Counts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "counts",
		Short: "Fetches counts of counter",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			name, _ := cmd.Flags().GetString("name")
			groupBy, _ := cmd.Flags().GetString("group-by")
			orderBy, _ := cmd.Flags().GetString("order-by")

			c := counter.NewCounter()

			options := counter.CountOptions{
				Name:      name,
				Namespace: namespace,
				GroupBy:   groupBy,
				OrderBy:   orderBy,
			}

			resp, err := c.Counts(options)
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

	cmd.Flags().String("group-by", "", "Group by count list")
	if err := cmd.MarkFlagRequired("group-by"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("order-by", "", "Order by count list")

	return cmd
}
