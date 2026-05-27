//go:build !cgo || (!ORT && !ALL)

package backends

import (
	"context"

	"github.com/knights-analytics/hugot/options"
)

type ORTModel struct {
	Destroy           func() error
	GenerativeSession *disabledGenerativeSession // placeholder when ORT disabled
	GenerativeEngine  *disabledGenerativeEngine  // placeholder when ORT disabled
}

func createORTModelBackend(_ *Model, _ *options.Options) error {
	_ = "STUB: not implemented"
	return nil
}

func createInputTensorsORT(_ *PipelineBatch, _ *Model) error { _ = "STUB: not implemented"; return nil }

func runORTSessionOnBatch(_ context.Context, _ *PipelineBatch, _ *BasePipeline) error {
	_ = "STUB: not implemented"
	return nil
}

func createImageTensorsORT(_ *PipelineBatch, _ *Model, _ [][][][]float32) error {
	_ = "STUB: not implemented"
	return nil
}

func createTabularTensorsORT(_ *PipelineBatch, _ *Model, _ [][]float32) error {
	_ = "STUB: not implemented"
	return nil
}

func runGenerativeORTSessionOnBatch(_ context.Context, _ *PipelineBatch, _ *BasePipeline, _ int, _ []string, _ *float64, _ *float64, _ *int, _ []string, _ *Guidance) (chan SequenceDelta, chan error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func createORTGenerativeSession(_ context.Context, _ *Model, _ *options.Options) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateMessagesORT(_ *PipelineBatch, _ any, _ string) error {
	_ = "STUB: not implemented"
	return nil
}

type (
	disabledGenerativeSession struct{}
	disabledGenerativeEngine  struct{}
)

func (*disabledGenerativeSession) GetStatistics() disabledStatistics {
	_ = "STUB: not implemented"
	return *new(disabledStatistics)
}

func (*disabledGenerativeEngine) GetStatistics() disabledStatistics {
	_ = "STUB: not implemented"
	return *new(disabledStatistics)
}

type disabledStatistics struct {
	AvgPrefillSeconds              float64
	TokensPerSecond                float64
	CumulativePrefillSum           float64
	CumulativePrefillCount         int
	CumulativeTokens               int
	CumulativeTokenDurationSeconds float64
}
