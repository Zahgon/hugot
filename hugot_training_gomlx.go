package hugot

import (
	"context"

	"github.com/gomlx/gomlx/pkg/ml/train/losses"
	"github.com/gomlx/gomlx/pkg/ml/train/optimizers"

	"github.com/knights-analytics/hugot/backends"
)

type GOMLXTrainingOptions struct {
	Optimizer optimizers.Interface
	Loss      losses.LossFn
}

type stoppingError struct{}

func (e stoppingError) Error() string { _ = "STUB: not implemented"; return "" }

func NewGoTrainingSession[T backends.Pipeline](ctx context.Context, config TrainingConfig) (*TrainingSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewXLATrainingSession[T backends.Pipeline](ctx context.Context, config TrainingConfig) (*TrainingSession, error) {
	_ = "STUB: not implemented"
	// Disabled for now until we have auto installs globally
	return nil, nil
}

func newGoMLXTrainingSession(s *TrainingSession) (*TrainingSession, error) {
	_ = "STUB: not implemented"
	// set defaults
	return nil, nil
}

func TrainGoMLX(s *TrainingSession) error { _ = "STUB: not implemented"; return nil }

// freeze the layers if requested

// identify the layer number in the variable name

// inputIDs, attentionMask, tokenTypeIDs if present

// we mean pool the results if needed e.g. if dimensions are [batch, seq, hidden]

// Loop for given number of epochs.

// trigger stopping

// we rely on try catch because an error is returned if there is an initialization error but
// a panic will be thrown if e.g. dataset reset fails.
