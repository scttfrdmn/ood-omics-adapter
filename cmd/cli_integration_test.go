//go:build integration

package cmd

// CLI-level integration tests for ood-omics-adapter are not feasible with the
// substrate emulator because the AWS HealthOmics SDK rewrites the request host
// for workflow run operations (StartRun, GetRun, CancelRun) by prepending
// "workflows-" to whatever hostname is configured.  For example, a substrate
// server running at "127.0.0.1:PORT" becomes "workflows-127.0.0.1:PORT",
// which fails DNS resolution and causes every request to error before it
// reaches the emulator.
//
// The internal package integration tests in
// internal/omics/client_integration_test.go work around this by injecting a
// custom HTTP transport (omicsRoundTripper) that strips the prefix.  The CLI
// commands construct their own SDK clients and have no hook for injecting a
// custom transport without code changes.
//
// Full lifecycle coverage at the CLI level should be added once the omics
// client constructor exposes a transport/options injection point, or once
// substrate gains a DNS alias that absorbs the "workflows-" prefix.
