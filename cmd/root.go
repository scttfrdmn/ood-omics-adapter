package cmd

import (
	"github.com/spf13/cobra"
)

var (
	region  string
	roleArn string
)

var version = "dev" // overridden at release time via -ldflags -X .../cmd.version

var rootCmd = &cobra.Command{
	Version: version,
	Use:   "ood-omics-adapter",
	Short: "OOD compute adapter for AWS HealthOmics",
	Long:  "Translates Open OnDemand job submissions to AWS HealthOmics Workflows API calls.",
}

// Execute runs the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&region, "region", "us-east-1", "AWS region")
	rootCmd.PersistentFlags().StringVar(&roleArn, "role-arn", "", "IAM role ARN for HealthOmics run execution")
}
