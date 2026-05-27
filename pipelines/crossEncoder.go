package pipelines

import (
	"context"
	"time"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
)

type CrossEncoderPipeline struct {
	*backends.BasePipeline
	scoreThreshold *float32
	statistics     CrossEncoderStatistics
	batchSize      int
	sortResults    bool
}
type CrossEncoderStatistics struct {
	TotalQueries     uint64
	TotalDocuments   uint64
	AverageLatency   time.Duration
	AverageBatchSize float64
	FilteredResults  uint64
}
type CrossEncoderResult struct {
	Document string
	Score    float32
	Index    int
}
type CrossEncoderOutput struct {
	Results []CrossEncoderResult
}

func WithBatchSize(size int) backends.PipelineOption[*CrossEncoderPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithSortResults() backends.PipelineOption[*CrossEncoderPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithScoreThreshold(threshold float32) backends.PipelineOption[*CrossEncoderPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func (t *CrossEncoderOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

func NewCrossEncoderPipeline(sessionContext context.Context, config backends.PipelineConfig[*CrossEncoderPipeline], s *options.Options, model *backends.Model) (*CrossEncoderPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// INTERFACE IMPLEMENTATIONS.

func (p *CrossEncoderPipeline) IsGenerative() bool { _ = "STUB: not implemented"; return false }

func (p *CrossEncoderPipeline) GetModel() *backends.Model { _ = "STUB: not implemented"; return nil }

func (p *CrossEncoderPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

func (p *CrossEncoderPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

func (p *CrossEncoderPipeline) Validate() error { _ = "STUB: not implemented"; return nil }

func (p *CrossEncoderPipeline) preprocessPairs(batch *backends.PipelineBatch, inputs [][2]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *CrossEncoderPipeline) forward(ctx context.Context, batch *backends.PipelineBatch) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *CrossEncoderPipeline) postprocess(batch *backends.PipelineBatch, documents []string) (*CrossEncoderOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *CrossEncoderPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

func (p *CrossEncoderPipeline) RunPipeline(ctx context.Context, query string, documents []string) (*CrossEncoderOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *CrossEncoderPipeline) runBatch(ctx context.Context, query string, documents []string, startIndex int) (*CrossEncoderOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
