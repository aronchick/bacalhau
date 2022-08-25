package noop

import (
	"context"

	"github.com/filecoin-project/bacalhau/pkg/executor"
	"github.com/filecoin-project/bacalhau/pkg/job"
	"github.com/filecoin-project/bacalhau/pkg/system"
	"github.com/filecoin-project/bacalhau/pkg/verifier"
	"github.com/filecoin-project/bacalhau/pkg/verifier/results"
	"go.opentelemetry.io/otel/trace"
)

type NoopVerifier struct {
	stateResolver *job.StateResolver
	results       *results.Results
}

func NewNoopVerifier(
	cm *system.CleanupManager,
	resolver *job.StateResolver,
) (*NoopVerifier, error) {
	results, err := results.NewResults()
	if err != nil {
		return nil, err
	}
	return &NoopVerifier{
		stateResolver: resolver,
		results:       results,
	}, nil
}

func (noopVerifier *NoopVerifier) IsInstalled(ctx context.Context) (bool, error) {
	return true, nil
}

func (noopVerifier *NoopVerifier) GetShardResultPath(
	ctx context.Context,
	jobID string,
	shardIndex int,
) (string, error) {
	return noopVerifier.results.EnsureShardResultsDir(jobID, shardIndex)
}

func (noopVerifier *NoopVerifier) GetShardProposal(
	ctx context.Context,
	jobID string,
	shardIndex int,
	shardResultPath string,
) ([]byte, error) {
	return []byte{}, nil
}

// each shard must have >= concurrency states
// and they must be either JobStateError or JobStateVerifying
func (noopVerifier *NoopVerifier) IsExecutionComplete(
	ctx context.Context,
	jobID string,
) (bool, error) {
	return noopVerifier.stateResolver.CheckShardStates(ctx, jobID, func(
		shardStates []executor.JobShardState,
		concurrency int,
	) (bool, error) {
		return noopVerifier.results.CheckShardStates(shardStates, concurrency)
	})
}

func (noopVerifier *NoopVerifier) VerifyJob(
	ctx context.Context,
	jobID string,
) ([]verifier.VerifierResult, error) {
	ctx, span := newSpan(ctx, "VerifyJob")
	defer span.End()
	results := []verifier.VerifierResult{}
	jobState, err := noopVerifier.stateResolver.GetJobState(ctx, jobID)
	if err != nil {
		return results, err
	}
	for _, shardState := range job.FlattenShardStates(jobState) { //nolint:gocritic
		if shardState.State != executor.JobStateVerifying {
			continue
		}
		results = append(results, verifier.VerifierResult{
			JobID:      jobID,
			NodeID:     shardState.NodeID,
			ShardIndex: shardState.ShardIndex,
			Verified:   true,
		})
	}
	return results, nil
}

func newSpan(ctx context.Context, apiName string) (context.Context, trace.Span) {
	return system.Span(ctx, "verifier/noop", apiName)
}

// Compile-time check that NoopVerifier implements the correct interface:
var _ verifier.Verifier = (*NoopVerifier)(nil)
