# ood-omics-adapter

[![CI](https://github.com/scttfrdmn/ood-omics-adapter/actions/workflows/ci.yml/badge.svg)](https://github.com/scttfrdmn/ood-omics-adapter/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/scttfrdmn/ood-omics-adapter)](https://goreportcard.com/report/github.com/scttfrdmn/ood-omics-adapter)
[![codecov](https://codecov.io/gh/scttfrdmn/ood-omics-adapter/branch/main/graph/badge.svg)](https://codecov.io/gh/scttfrdmn/ood-omics-adapter)
[![Go Reference](https://pkg.go.dev/badge/github.com/scttfrdmn/ood-omics-adapter.svg)](https://pkg.go.dev/github.com/scttfrdmn/ood-omics-adapter)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

OOD compute adapter for AWS HealthOmics Workflows. Translates Open OnDemand job submissions to AWS HealthOmics API calls.

## Commands

| Command | Description |
|---------|-------------|
| `submit` | Submit an OOD job spec (JSON from stdin) as a HealthOmics workflow run |
| `status <run-id>` | Get OOD-normalized status of a HealthOmics run |
| `delete <run-id>` | Cancel a HealthOmics run |
| `info <run-id>` | Print full run details as JSON |

## Usage

```bash
echo '{"workflow_id":"wf-1234","output_uri":"s3://my-bucket/results/","role_arn":"arn:aws:iam::123456789012:role/omics-exec"}' | \
  ood-omics-adapter submit --region us-east-1

ood-omics-adapter status <run-id>
ood-omics-adapter delete <run-id>
```

## Infrastructure

Terraform in `aws-openondemand` with `adapters_enabled = ["omics"]` provisions:
- IAM policy on the OOD instance role (omics:StartRun, GetRun, CancelRun, ListRuns)
- iam:PassRole scoped to omics.amazonaws.com
