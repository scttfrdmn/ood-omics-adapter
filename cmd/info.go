package cmd

import (
	"context"
	"encoding/json"
	"os"

	"github.com/scttfrdmn/ood-omics-adapter/internal/omics"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info <run-id>",
	Short: "Print full HealthOmics run details as JSON",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client, err := omics.New(ctx, region)
		if err != nil {
			return err
		}
		detail, err := client.GetRun(ctx, args[0])
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(detail)
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
