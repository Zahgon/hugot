package pipelines

import (
	"context"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
)

// types

type TextClassificationPipeline struct {
	*backends.BasePipeline
	AggregationFunctionName string
	ProblemType             string
	FixedPaddingLength      int
}

type TextClassificationOutput struct {
	ClassificationOutputs [][]ClassificationOutput
}

func (t *TextClassificationOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

// options

func WithSoftmax() backends.PipelineOption[*TextClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithSigmoid() backends.PipelineOption[*TextClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithSingleLabel() backends.PipelineOption[*TextClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithMultiLabel() backends.PipelineOption[*TextClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithFixedPadding(fixedPaddingLength int) backends.PipelineOption[*TextClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// NewTextClassificationPipeline initializes a new text classification pipeline.
func NewTextClassificationPipeline(sessionContext context.Context, config backends.PipelineConfig[*TextClassificationPipeline], s *options.Options, model *backends.Model) (*TextClassificationPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validate

// INTERFACE IMPLEMENTATION

func (p *TextClassificationPipeline) IsGenerative() bool { _ = "STUB: not implemented"; return false }

func (p *TextClassificationPipeline) GetModel() *backends.Model {
	_ = "STUB: not implemented"

	// GetMetadata returns metadata information about the pipeline, in particular:
	// OutputInfo: names and dimensions of the output layer used for text classification.
	return nil
}

func (p *TextClassificationPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

// GetStatistics returns the runtime statistics for the pipeline.
func (p *TextClassificationPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

// Validate checks that the pipeline is valid.
func (p *TextClassificationPipeline) Validate() error { _ = "STUB: not implemented"; return nil }

// preprocess tokenizes the input strings.
func (p *TextClassificationPipeline) preprocess(batch *backends.PipelineBatch, inputs []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TextClassificationPipeline) forward(ctx context.Context, batch *backends.PipelineBatch) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TextClassificationPipeline) postprocess(batch *backends.PipelineBatch) (*TextClassificationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Run the pipeline on a string batch.
func (p *TextClassificationPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

func (p *TextClassificationPipeline) RunPipeline(ctx context.Context, inputs []string) (*TextClassificationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
