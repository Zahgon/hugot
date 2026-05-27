package pipelines

import (
	"context"
	"image"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
	"github.com/knights-analytics/hugot/util/imageutil"
)

// FeatureExtractionPipeline A feature extraction pipeline is a go version of
// https://github.com/huggingface/transformers/blob/main/src/transformers/pipelines/feature_extraction.py
// It supports both text and image inputs for embedding extraction.
type FeatureExtractionPipeline struct {
	*backends.BasePipeline
	OutputName         string
	imageFormat        string
	Output             backends.InputOutputInfo
	preprocessSteps    []imageutil.PreprocessStep
	normalizationSteps []imageutil.NormalizationStep
	OutputIndex        int // Record the index of the output selected, defaults to first (0)
	Normalization      bool
	// Image mode fields (for vision encoders like CLIP visual)
	imageMode bool // true if this is a vision model
}

type FeatureExtractionOutput struct {
	Embeddings [][]float32
}

func (t *FeatureExtractionOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

func (p *FeatureExtractionPipeline) addPreprocessSteps(steps ...imageutil.PreprocessStep) {
	_ = "STUB: not implemented"
	return
}

func (p *FeatureExtractionPipeline) addNormalizationSteps(steps ...imageutil.NormalizationStep) {
	_ = "STUB: not implemented"
	return
}

func (p *FeatureExtractionPipeline) setImageFormat(format string) {
	_ = "STUB: not implemented"
	return

	// PIPELINE OPTIONS
}

// WithNormalization applies normalization to the mean pooled output of the feature pipeline.
func WithNormalization() backends.PipelineOption[*FeatureExtractionPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithOutputName if there are multiple outputs from the underlying model, which output should
// be returned. If not passed, the first output from the feature pipeline is returned.
func WithOutputName(outputName string) backends.PipelineOption[*FeatureExtractionPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithImageMode enables image feature extraction mode for vision encoders (e.g., CLIP visual encoder).
// When enabled, the pipeline accepts images instead of text and skips tokenization.
func WithImageMode() backends.PipelineOption[*FeatureExtractionPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// NewFeatureExtractionPipeline init a feature extraction pipeline.
func NewFeatureExtractionPipeline(sessionContext context.Context, config backends.PipelineConfig[*FeatureExtractionPipeline], s *options.Options, model *backends.Model) (*FeatureExtractionPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set default image format if in image mode

// filter outputs

// we take the first output otherwise, like transformers does

// validate pipeline

// INTERFACE IMPLEMENTATIONS.

func (p *FeatureExtractionPipeline) IsGenerative() bool { _ = "STUB: not implemented"; return false }

func (p *FeatureExtractionPipeline) GetModel() *backends.Model {
	_ = "STUB: not implemented"

	// GetMetadata returns metadata information about the pipeline, in particular:
	// OutputInfo: names and dimensions of the output layer.
	return nil
}

func (p *FeatureExtractionPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

// GetStatistics returns the runtime statistics for the pipeline.
func (p *FeatureExtractionPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

// Validate checks that the pipeline is valid.
func (p *FeatureExtractionPipeline) Validate() error { _ = "STUB: not implemented"; return nil }

// Tokenizer is only required for text mode

// Allow up to 4D inputs for all pipelines:
// - Text models: [batch, sequence] or [batch, sequence, hidden]
// - Image models: [batch, channels, height, width]
// - Multimodal models may have various dimensional structures

// preprocess tokenizes the input strings.
func (p *FeatureExtractionPipeline) preprocess(batch *backends.PipelineBatch, inputs []string) error {
	_ = "STUB: not implemented"
	return nil
}

// forward performs the forward inference of the feature extraction pipeline.
func (p *FeatureExtractionPipeline) forward(ctx context.Context, batch *backends.PipelineBatch) error {
	_ = "STUB: not implemented"
	return nil
}

// postprocess parses the first output from the network similar to the transformers' implementation.
func (p *FeatureExtractionPipeline) postprocess(batch *backends.PipelineBatch) (*FeatureExtractionOutput, error) {
	_ = "STUB: not implemented"
	// TODO: this works if token embeddings are returned or sentence embeddings are returned.
	// in the former case embeddings are mean pooled. In the latter they are just returned.
	// to make this more general for other pipelines and to allow return of raw token embeddings,
	// we need an ndarray type that can be the return type of this pipeline. Need to think
	// about how to do this in a lightweight manner.
	return nil, nil
}

// Use the index of the output we want to return

// Normalize embeddings (if asked), like in https://huggingface.co/sentence-transformers/all-mpnet-base-v2

func meanPooling(tokens [][]float32, input backends.TokenizedInput, maxSequence int, dimensions int) []float32 {
	_ = "STUB: not implemented"
	return nil
}

// if there is no attention mask, take all tokens

// Run the pipeline on a batch of strings.
func (p *FeatureExtractionPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

// RunPipeline is like Run, but returns the concrete feature extraction output type rather than the interface.
func (p *FeatureExtractionPipeline) RunPipeline(ctx context.Context, inputs []string) (*FeatureExtractionOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IMAGE MODE METHODS

// PreprocessImages converts images to input tensors for vision models.
func (p *FeatureExtractionPipeline) PreprocessImages(batch *backends.PipelineBatch, inputs []image.Image) error {
	_ = "STUB: not implemented"
	return nil
}

// RunWithImages runs the pipeline on a batch of images (for vision models).
// Use this method when ImageMode is enabled.
func (p *FeatureExtractionPipeline) RunWithImages(ctx context.Context, images []image.Image) (*FeatureExtractionOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunWithImagePaths loads images from file paths and runs the pipeline.
// Convenience method that combines image loading with RunWithImages.
func (p *FeatureExtractionPipeline) RunWithImagePaths(ctx context.Context, paths []string) (*FeatureExtractionOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
