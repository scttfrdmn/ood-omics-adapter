package cmd

import (
	"testing"

	omicstypes "github.com/aws/aws-sdk-go-v2/service/omics/types"
)

func TestOmicsStateToOod(t *testing.T) {
	tests := []struct {
		state    omicstypes.RunStatus
		expected string
	}{
		{omicstypes.RunStatusPending, "queued"},
		{omicstypes.RunStatusStarting, "queued"},
		{omicstypes.RunStatusRunning, "running"},
		{omicstypes.RunStatusStopping, "running"},
		{omicstypes.RunStatusCompleted, "completed"},
		{omicstypes.RunStatusFailed, "failed"},
		{omicstypes.RunStatusDeleted, "failed"},
		{omicstypes.RunStatusCancelled, "cancelled"},
		{omicstypes.RunStatus("UNKNOWN_STATE"), "undetermined"},
	}

	for _, tt := range tests {
		t.Run(string(tt.state), func(t *testing.T) {
			got := omicsStateToOod(tt.state)
			if got != tt.expected {
				t.Errorf("omicsStateToOod(%q) = %q, want %q", tt.state, got, tt.expected)
			}
		})
	}
}
