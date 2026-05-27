package pipelines

import (
	"context"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
)

type TextGenerationPipeline struct {
	*backends.BasePipeline
	SystemPrompt  string
	MaxLength     int
	Streaming     bool
	Temperature   *float64
	TopP          *float64
	Seed          *int
	StopSequences []string
	Tools         []string
	Guidance      *backends.Guidance
}

type TextGenerationOutput struct {
	TokenStream chan backends.SequenceDelta
	ErrorStream chan error
	Responses   []string
}

func (t *TextGenerationOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

// WithSystemPrompt allows the user to define a system prompt that will be prepended to every input.
func WithSystemPrompt(systemPrompt string) backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithMaxLength allows the user to define the maximum generated tokens.
func WithMaxLength(maxLength int) backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithStreaming allows the user to receive generated tokens as a stream instead of waiting for the entire response.
func WithStreaming() backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithTemperature(temperature float64) backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithTopP(topP float64) backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithSeed(seed int) backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithStopSequences allows the user to define stop sequences that will end the generation when encountered.
// If the model produces any of the provided strings in the output, generation for that sequence will stop and the stop string will be excluded.
func WithStopSequences(stopSequences []string) backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithTools sets the list of Hermes-style tool definition JSON strings to include in the chat template.
func WithTools(tools []string) backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithGuidance enables constrained (guided) generation using a lark grammar, JSON schema, or regex.
func WithGuidance(guidance *backends.Guidance) backends.PipelineOption[*TextGenerationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// NewTextGenerationPipeline initializes a new text generation pipeline.
func NewTextGenerationPipeline(sessionContext context.Context, config backends.PipelineConfig[*TextGenerationPipeline], s *options.Options, model *backends.Model) (*TextGenerationPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default value if not set as per Python

// INTERFACE IMPLEMENTATION

func (p *TextGenerationPipeline) IsGenerative() bool { _ = "STUB: not implemented"; return false }

func (p *TextGenerationPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

func (p *TextGenerationPipeline) GetModel() *backends.Model {
	_ = "STUB: not implemented"

	// GetStatistics returns the runtime statistics for the pipeline.
	return nil
}

func (p *TextGenerationPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

func (p *TextGenerationPipeline) Validate() error { _ = "STUB: not implemented"; return nil }

func (p *TextGenerationPipeline) preprocess(batch *backends.PipelineBatch, inputs any) error {
	_ = "STUB: not implemented"
	return nil
}

// forward initiates the generation loop with explicit tools and guidance, allowing per-call overrides.
func (p *TextGenerationPipeline) forward(ctx context.Context, batch *backends.PipelineBatch, tools []string, guidance *backends.Guidance) (chan backends.SequenceDelta, chan error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func (p *TextGenerationPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

func (p *TextGenerationPipeline) RunPipeline(ctx context.Context, inputs []string) (*TextGenerationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect responses and errors

// runMessages processes a batch of message inputs.
// tools and guidance override the pipeline-level defaults for this call only.
// If the model produces any of the provided strings in the output, generation for that sequence will stop and the stop string will be excluded.
// If multimodal, the images should be added to the messages.
func (p *TextGenerationPipeline) runMessages(ctx context.Context, inputs [][]backends.Message, tools []string, guidance *backends.Guidance) (*TextGenerationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect responses and errors

func (p *TextGenerationPipeline) RunMessages(ctx context.Context, inputs [][]backends.Message) (*TextGenerationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *TextGenerationPipeline) RunMessagesWithOverrides(ctx context.Context, inputs [][]backends.Message, tools []string, guidance *backends.Guidance) (*TextGenerationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func collectResponses(tokenStream chan backends.SequenceDelta, errorStream chan error, batchSize int) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
