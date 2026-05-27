package pipelines

import (
	"context"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
)

// TabularPipeline supports classic ML models (e.g., decision trees, random forests)
// exported to ONNX that take numeric feature vectors and output either class logits
// or regression values.
type TabularPipeline struct {
	*backends.BasePipeline
	AggregationFunctionName string         // for classification: SOFTMAX or SIGMOID
	ProblemType             string         // "classification" or "regression"
	IDLabelMap              map[int]string // optional mapping from class IDs to labels
}

type TabularClassificationOutput struct {
	PredictedClass string
	Probabilities  []ClassificationOutput
}

// TabularOutput returns per-input results.
// - For classification: []TabularClassificationOutput
// - For regression: float32
// for each input.
type TabularOutput struct {
	ClassificationResults []TabularClassificationOutput
	RegressionResults     []float32
}

func (o *TabularOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

// Options

func WithRegression() backends.PipelineOption[*TabularPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithClassification() backends.PipelineOption[*TabularPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithTabularSoftmax() backends.PipelineOption[*TabularPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithTabularSigmoid() backends.PipelineOption[*TabularPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithIDLabelMap(labels map[int]string) backends.PipelineOption[*TabularPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// NewTabularPipeline initializes the pipeline.
func NewTabularPipeline(sessionContext context.Context, config backends.PipelineConfig[*TabularPipeline], s *options.Options, model *backends.Model) (*TabularPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// build default IDLabelMap

// Interface implementation

func (p *TabularPipeline) IsGenerative() bool { _ = "STUB: not implemented"; return false }

func (p *TabularPipeline) GetModel() *backends.Model { _ = "STUB: not implemented"; return nil }

func (p *TabularPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

func (p *TabularPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

func (p *TabularPipeline) Validate() error { _ = "STUB: not implemented"; return nil }

// on the outputs we are now strict:
// if it's a classification model we expect either:
// - one output with shape (batch)
// - or two outputs (class labels and probabilities) with shape (batch, 1) and (batch, num_classes)
// if it's a regression model we expect one output with shape (batch, 1)

// preprocess parses inputs strings into [][]float32 and builds tensors.
func (p *TabularPipeline) preprocess(batch *backends.PipelineBatch, inputs [][]float32) error {
	_ = "STUB: not implemented"

	// Build tensors
	return nil
}

// measured but not recorded; tokenizer not used

func (p *TabularPipeline) forward(ctx context.Context, batch *backends.PipelineBatch) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *TabularPipeline) postprocess(batch *backends.PipelineBatch) (*TabularOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// two outputs: class labels and probabilities

// one output: logits

// find predicted class

// Run executes the pipeline over inputs.
func (p *TabularPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

func (p *TabularPipeline) RunPipeline(ctx context.Context, inputs [][]float32) (*TabularOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// parseFeatures accepts each input only as a JSON array ("[1,2,3]").
func parseFeatures(inputs []string) ([][]float32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
