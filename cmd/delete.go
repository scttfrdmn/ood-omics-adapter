package cmd

import (
	"context"
	"fmt"

	"github.com/scttfrdmn/ood-omics-adapter/internal/omics"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <run-id>",
	Short: "Cancel a HealthOmics run",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client, err := omics.New(ctx, region)
		if err != nil {
			return err
		}
		if err := client.CancelRun(ctx, args[0]); err != nil {
			return err
		}
		fmt.Printf("Run %s cancelled\n", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
