package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/scttfrdmn/ood-omics-adapter/internal/omics"
	"github.com/spf13/cobra"
)

// JobSpec is the HealthOmics-specific job submission payload.
type JobSpec struct {
	WorkflowID      string                 `json:"workflow_id"`
	WorkflowType    string                 `json:"workflow_type,omitempty"`
	RoleArn         string                 `json:"role_arn,omitempty"`
	OutputUri       string                 `json:"output_uri"`
	Parameters      map[string]interface{} `json:"parameters,omitempty"`
	StorageCapacity int32                  `json:"storage_capacity,omitempty"`
	JobName         string                 `json:"job_name,omitempty"`
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit an OOD job to AWS HealthOmics",
	Long:  "Reads a JSON job spec from stdin and submits it as a HealthOmics workflow run.",
	RunE: func(cmd *cobra.Command, args []string) error {
		var spec JobSpec
		if err := json.NewDecoder(os.Stdin).Decode(&spec); err != nil {
			return fmt.Errorf("decode job spec: %w", err)
		}

		if spec.WorkflowID == "" {
			return fmt.Errorf("job spec must include workflow_id")
		}
		if spec.OutputUri == "" {
			return fmt.Errorf("job spec must include output_uri")
		}

		effectiveRole := roleArn
		if effectiveRole == "" {
			effectiveRole = spec.RoleArn
		}
		if effectiveRole == "" {
			return fmt.Errorf("--role-arn is required (or set role_arn in job spec)")
		}

		ctx := context.Background()
		client, err := omics.New(ctx, region)
		if err != nil {
			return err
		}

		runID, err := client.StartRun(ctx, omics.RunSpec{
			WorkflowID:      spec.WorkflowID,
			WorkflowType:    spec.WorkflowType,
			RoleArn:         effectiveRole,
			OutputUri:       spec.OutputUri,
			Parameters:      spec.Parameters,
			StorageCapacity: spec.StorageCapacity,
			JobName:         spec.JobName,
		})
		if err != nil {
			return err
		}

		fmt.Println(runID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(submitCmd)
}
