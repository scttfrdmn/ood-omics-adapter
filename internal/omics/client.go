// Package omics wraps the AWS HealthOmics API for the OOD adapter.
package omics

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/omics"
	"github.com/aws/aws-sdk-go-v2/service/omics/document"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
)

// Client wraps the AWS HealthOmics client.
type Client struct {
	svc    *omics.Client
	region string
}

// New creates an HealthOmics client using the default AWS credential chain.
func New(ctx context.Context, region string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("load AWS config: %w", err)
	}
	return &Client{svc: omics.NewFromConfig(cfg), region: region}, nil
}

// RunSpec holds the parameters for a HealthOmics workflow run.
type RunSpec struct {
	WorkflowID      string
	WorkflowType    string
	RoleArn         string
	OutputUri       string
	Parameters      map[string]interface{}
	StorageCapacity int32
	JobName         string
}


// StartRun submits a HealthOmics workflow run and returns the run ID.
func (c *Client) StartRun(ctx context.Context, spec RunSpec) (string, error) {
	input := &omics.StartRunInput{
		WorkflowId: aws.String(spec.WorkflowID),
		OutputUri:  aws.String(spec.OutputUri),
		RoleArn:    aws.String(spec.RoleArn),
	}
	if spec.WorkflowType != "" {
		input.WorkflowType = types.WorkflowType(spec.WorkflowType)
	}
	if spec.JobName != "" {
		input.Name = aws.String(spec.JobName)
	}
	if spec.StorageCapacity > 0 {
		input.StorageCapacity = aws.Int32(spec.StorageCapacity)
	}
	if len(spec.Parameters) > 0 {
		input.Parameters = document.NewLazyDocument(spec.Parameters)
	}

	out, err := c.svc.StartRun(ctx, input)
	if err != nil {
		return "", fmt.Errorf("omics StartRun: %w", err)
	}
	return aws.ToString(out.Id), nil
}

// GetRun returns the current detail of a HealthOmics run.
func (c *Client) GetRun(ctx context.Context, runID string) (*omics.GetRunOutput, error) {
	out, err := c.svc.GetRun(ctx, &omics.GetRunInput{Id: aws.String(runID)})
	if err != nil {
		return nil, fmt.Errorf("omics GetRun: %w", err)
	}
	return out, nil
}

// CancelRun cancels a HealthOmics workflow run.
func (c *Client) CancelRun(ctx context.Context, runID string) error {
	_, err := c.svc.CancelRun(ctx, &omics.CancelRunInput{Id: aws.String(runID)})
	if err != nil {
		return fmt.Errorf("omics CancelRun: %w", err)
	}
	return nil
}
