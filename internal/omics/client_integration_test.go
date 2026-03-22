//go:build integration

package omics_test

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	substrate "github.com/scttfrdmn/substrate"

	. "github.com/scttfrdmn/ood-omics-adapter/internal/omics"
)

// omicsRoundTripper strips the "workflows-" prefix that the HealthOmics SDK
// prepends to the host for workflow run operations. When using a local test
// server the rewritten host (e.g. "workflows-127.0.0.1:PORT") does not resolve,
// so we undo the rewrite before the connection is attempted.
type omicsRoundTripper struct{ base http.RoundTripper }

func (t *omicsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.HasPrefix(req.URL.Host, "workflows-") {
		req = req.Clone(req.Context())
		req.URL.Host = strings.TrimPrefix(req.URL.Host, "workflows-")
		req.Host = req.URL.Host
	}
	return t.base.RoundTrip(req)
}

// withTestClient injects the substrate-compatible HTTP client into the HealthOmics
// config so that the workflows-host rewrite does not break local test servers.
func withTestClient() func(*config.LoadOptions) error {
	return config.WithHTTPClient(&http.Client{
		Transport: &omicsRoundTripper{base: http.DefaultTransport},
	})
}

// TestStartGetCancelRun_Substrate exercises the full HealthOmics workflow run
// lifecycle (StartRun → GetRun → CancelRun) against the substrate emulator.
func TestStartGetCancelRun_Substrate(t *testing.T) {
	ts := substrate.StartTestServer(t)
	t.Setenv("AWS_ENDPOINT_URL", ts.URL)
	t.Setenv("AWS_ACCESS_KEY_ID", "test")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "test")

	ctx := context.Background()
	client, err := New(ctx, "us-east-1",
		withTestClient(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
	)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	spec := RunSpec{
		WorkflowID:      "1234567",
		WorkflowType:    "PRIVATE",
		RoleArn:         "arn:aws:iam::123456789012:role/OmicsWorkflowRole",
		OutputUri:       "s3://my-bucket/omics-outputs/",
		Parameters:      map[string]interface{}{"ref_fasta": "s3://my-bucket/ref/hg38.fa"},
		StorageCapacity: 100,
		JobName:         "ood-integration-test",
	}

	// StartRun
	runID, err := client.StartRun(ctx, spec)
	if err != nil {
		t.Fatalf("StartRun: %v", err)
	}
	if runID == "" {
		t.Fatal("expected non-empty run ID")
	}
	t.Logf("started omics run: %s", runID)

	// GetRun
	detail, err := client.GetRun(ctx, runID)
	if err != nil {
		t.Fatalf("GetRun: %v", err)
	}
	if detail == nil {
		t.Fatal("GetRun: got nil output")
	}
	t.Logf("run status: %s", detail.Status)

	// CancelRun
	err = client.CancelRun(ctx, runID)
	if err != nil {
		t.Fatalf("CancelRun: %v", err)
	}
	t.Log("run cancelled successfully")
}

// TestGetRun_NotFound_Substrate verifies that GetRun returns an error
// for a run ID that was never created.
func TestGetRun_NotFound_Substrate(t *testing.T) {
	ts := substrate.StartTestServer(t)
	t.Setenv("AWS_ENDPOINT_URL", ts.URL)
	t.Setenv("AWS_ACCESS_KEY_ID", "test")
	t.Setenv("AWS_SECRET_ACCESS_KEY", "test")

	ctx := context.Background()
	client, err := New(ctx, "us-east-1",
		withTestClient(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
	)
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	_, err = client.GetRun(ctx, "0000000000")
	if err == nil {
		t.Fatal("expected error for non-existent run, got nil")
	}
	if !strings.Contains(err.Error(), "omics") {
		t.Logf("error (acceptable): %v", err)
	}
}
