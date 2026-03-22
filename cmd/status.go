package cmd

import (
	"context"
	"encoding/json"
	"os"

	omicstypes "github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/scttfrdmn/ood-omics-adapter/internal/omics"
	internalood "github.com/scttfrdmn/ood-omics-adapter/internal/ood"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status <run-id>",
	Short: "Get the status of a HealthOmics run",
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

		js := internalood.JobStatus{
			ID:     args[0],
			Status: omicsStateToOod(detail.Status),
		}
		if detail.StatusMessage != nil {
			js.Message = *detail.StatusMessage
		}

		return json.NewEncoder(os.Stdout).Encode(js)
	},
}

func omicsStateToOod(s omicstypes.RunStatus) string {
	switch s {
	case omicstypes.RunStatusPending, omicstypes.RunStatusStarting:
		return internalood.StatusQueued
	case omicstypes.RunStatusRunning, omicstypes.RunStatusStopping:
		return internalood.StatusRunning
	case omicstypes.RunStatusCompleted:
		return internalood.StatusCompleted
	case omicstypes.RunStatusFailed, omicstypes.RunStatusDeleted:
		return internalood.StatusFailed
	case omicstypes.RunStatusCancelled:
		return internalood.StatusCancelled
	default:
		return internalood.StatusUnknown
	}
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
