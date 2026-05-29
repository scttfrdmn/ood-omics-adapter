# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial scaffold — OOD compute adapter for AWS HealthOmics Workflows, translating Open OnDemand job submissions to AWS HealthOmics API calls.
- CLI commands: `submit` (JSON job spec from stdin → HealthOmics workflow run), `status <run-id>` (OOD-normalized status), `delete <run-id>` (cancel a run), and `info <run-id>` (full run details as JSON).
- Unit tests for status state mapping.
- Substrate integration tests for the HealthOmics workflow run lifecycle.
- CI workflow with pinned action SHAs.
