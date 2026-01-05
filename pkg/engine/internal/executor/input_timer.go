package executor

import (
	"context"
	"time"

	"github.com/apache/arrow-go/v18/arrow"
)

// inputTimer tracks cumulative time spent reading from input pipelines.
// Use this to separate input wait time from local execution time.
type inputTimer struct {
	total time.Duration
}

// ReadFrom reads from the given pipeline while tracking elapsed time.
func (t *inputTimer) ReadFrom(ctx context.Context, input Pipeline) (arrow.RecordBatch, error) {
	start := time.Now()
	rec, err := input.Read(ctx)
	t.total += time.Since(start)
	return rec, err
}

// Duration returns the total time spent waiting on input reads.
func (t *inputTimer) Duration() time.Duration {
	return t.total
}

// Reset clears the accumulated duration.
func (t *inputTimer) Reset() {
	t.total = 0
}

