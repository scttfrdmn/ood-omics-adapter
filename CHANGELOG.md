# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Fixed
- Bumped `github.com/scttfrdmn/substrate` 0.45.2 → 0.65.0 and regenerated go.sum. The recorded 0.45.2 checksum no longer matched the module the proxy serves (upstream re-tag), which broke `go test -tags=integration` with a go.sum SECURITY ERROR. Integration tests now build and pass.

### Added
- Initial scaffold — OOD compute adapter for AWS HealthOmics Workflows, translating Open OnDemand job submissions to AWS HealthOmics API calls.
- CLI commands: `submit` (JSON job spec from stdin → HealthOmics workflow run), `status <run-id>` (OOD-normalized status), `delete <run-id>` (cancel a run), and `info <run-id>` (full run details as JSON).
- Unit tests for status state mapping.
- Substrate integration tests for the HealthOmics workflow run lifecycle.
- CI workflow with pinned action SHAs.
