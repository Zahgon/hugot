package pipelines

import (
	"context"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
)

type ZeroShotClassificationPipeline struct {
	*backends.BasePipeline
	HypothesisTemplate string
	Sequences          []string
	Labels             []string
	EntailmentID       int
	Multilabel         bool
}
type ZeroShotClassificationOutput struct {
	Sequence     string
	SortedValues []struct {
		Key   string
		Value float64
	}
}
type ZeroShotOutput struct {
	ClassificationOutputs []ZeroShotClassificationOutput
}

// options

// WithMultilabel can be used to set whether the pipeline is multilabel.
func WithMultilabel(multilabel bool) backends.PipelineOption[*ZeroShotClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithLabels can be used to set the labels to classify the examples.
func WithLabels(labels []string) backends.PipelineOption[*ZeroShotClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithHypothesisTemplate can be used to set the hypothesis template for classification.
func WithHypothesisTemplate(hypothesisTemplate string) backends.PipelineOption[*ZeroShotClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// GetOutput converts raw output to readable output.
func (t *ZeroShotOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

// create all pairs between input sequences and labels.
func createSequencePairs(sequences interface{}, labels []string, hypothesisTemplate string) ([][][]string, []string, error) {
	_ = "STUB: not implemented"
	// Check if labels or sequences are empty
	return nil, nil, nil
}

// Check if hypothesisTemplate can be formatted with labels

// Convert sequences to []string if it's a single string

// Create sequence_pairs

// NewZeroShotClassificationPipeline create new Zero Shot Classification Pipeline.
func NewZeroShotClassificationPipeline(sessionContext context.Context, config backends.PipelineConfig[*ZeroShotClassificationPipeline], s *options.Options, model *backends.Model) (*ZeroShotClassificationPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default value

// Find entailment ID

func (p *ZeroShotClassificationPipeline) preprocessPairs(batch *backends.PipelineBatch, inputs [][2]string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *ZeroShotClassificationPipeline) forward(ctx context.Context, batch *backends.PipelineBatch) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *ZeroShotClassificationPipeline) postprocess(outputTensors [][][]float32, labels []string, sequences []string) *ZeroShotOutput {
	_ = "STUB: not implemented"
	return nil
}

// Define ss as a slice of anonymous structs

// Sort the slice by the value field

// Define ss as a slice of anonymous structs

// Sort the slice by the value field

func (p *ZeroShotClassificationPipeline) RunPipeline(ctx context.Context, inputs []string) (*ZeroShotOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PIPELINE INTERFACE IMPLEMENTATION

func (p *ZeroShotClassificationPipeline) IsGenerative() bool {
	_ = "STUB: not implemented"
	return false
}

func (p *ZeroShotClassificationPipeline) GetModel() *backends.Model {
	_ = "STUB: not implemented"
	return nil
}

func (p *ZeroShotClassificationPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

// GetStatistics returns the runtime statistics for the pipeline.
func (p *ZeroShotClassificationPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

func (p *ZeroShotClassificationPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

func (p *ZeroShotClassificationPipeline) Validate() error { _ = "STUB: not implemented"; return nil }
