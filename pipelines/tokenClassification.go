package pipelines

import (
	"context"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
)

// TokenClassificationPipeline is a go version of huggingface tokenClassificationPipeline.
// https://github.com/huggingface/transformers/blob/main/src/transformers/pipelines/token_classification.py
type TokenClassificationPipeline struct {
	*backends.BasePipeline
	IDLabelMap          map[int]string
	AggregationStrategy string
	IgnoreLabels        []string
	SplitWords          bool
}
type Entity struct {
	Entity    string
	Word      string
	Scores    []float32
	TokenIDs  []uint32
	Index     int
	Start     uint
	End       uint
	Score     float32
	IsSubword bool
}
type TokenClassificationOutput struct {
	Entities [][]Entity
}

func (t *TokenClassificationOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

// options

// WithSimpleAggregation sets the aggregation strategy for the token labels to simple
// It reproduces simple aggregation from the huggingface implementation.
func WithSimpleAggregation() backends.PipelineOption[*TokenClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithAverageAggregation sets the aggregation strategy for the token labels to average
// It reproduces simple aggregation from the huggingface implementation.
func WithAverageAggregation() backends.PipelineOption[*TokenClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithMaxAggregation sets the aggregation strategy for the token labels to Max
// It reproduces max aggregation from the huggingface implementation.
func WithMaxAggregation() backends.PipelineOption[*TokenClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithFirstAggregation sets the aggregation strategy for the token labels to first
// It reproduces first aggregation from the huggingface implementation.
func WithFirstAggregation() backends.PipelineOption[*TokenClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithoutAggregation returns the token labels.
func WithoutAggregation() backends.PipelineOption[*TokenClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithIgnoreLabels(ignoreLabels []string) backends.PipelineOption[*TokenClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithSplitWords enables word-level alignment like Hugging Face's is_split_into_words.
func WithSplitWords() backends.PipelineOption[*TokenClassificationPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// NewTokenClassificationPipeline Initializes a feature extraction pipeline.
func NewTokenClassificationPipeline(sessionContext context.Context, config backends.PipelineConfig[*TokenClassificationPipeline], s *options.Options, model *backends.Model) (*TokenClassificationPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Id label map

// default strategies if not set

// Additional options needed for postprocessing

// INTERFACE IMPLEMENTATION

func (p *TokenClassificationPipeline) IsGenerative() bool { _ = "STUB: not implemented"; return false }

func (p *TokenClassificationPipeline) GetModel() *backends.Model {
	_ = "STUB: not implemented"

	// GetMetadata returns metadata information about the pipeline, in particular:
	// OutputInfo: names and dimensions of the output layer used for token classification.
	return nil
}

func (p *TokenClassificationPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

// GetStatistics returns the runtime statistics for the pipeline.
func (p *TokenClassificationPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

// Validate checks that the pipeline is valid.
func (p *TokenClassificationPipeline) Validate() error { _ = "STUB: not implemented"; return nil }

// preprocess tokenizes the input strings.
func (p *TokenClassificationPipeline) preprocess(batch *backends.PipelineBatch, inputs []string) error {
	_ = "STUB: not implemented"
	return nil
}

// preprocessWords tokenizes pre-split words and maps tokens to word IDs via offsets.
func (p *TokenClassificationPipeline) preprocessWords(batch *backends.PipelineBatch, inputs [][]string) error {
	_ = "STUB: not implemented"

	// Join words with single spaces to simulate pretokenized behavior
	return nil
}

// local helper to convert non-negative int to uint safely

// compute boundaries in joined string

// clamp to non-negative and convert safely to uint
// ensure non-negative before converting to uint

// add one space after every word except last

// set raw to joined string for offsets consistency

// forward performs the forward inference of the pipeline.
func (p *TokenClassificationPipeline) forward(ctx context.Context, batch *backends.PipelineBatch) error {
	_ = "STUB: not implemented"
	return nil
}

// postprocess function for a token classification pipeline.
func (p *TokenClassificationPipeline) postprocess(batch *backends.PipelineBatch) (*TokenClassificationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// now convert the logits to the predictions of actual entities

// Filter anything that is in ignore_labels

// gatherPreEntities from batch of logits to list of pre-aggregated outputs.
func (p *TokenClassificationPipeline) gatherPreEntities(input backends.TokenizedInput, output [][]float32) []Entity {
	_ = "STUB: not implemented"
	return nil
}

// filter out special tokens (skip them)

// TODO: the python code uses id_to_token to get the token here which is a method on the rust tokenizer, check if it's better

// TODO: the determination of subword can probably be better done by exporting the words field from the tokenizer directly

// In split-words mode, grouping will use offsets between tokens rather than IsSubword.
// TODO: check for unknown token here, it's in the config and can be loaded and compared with the token
// in that case set the subword as in the python code

func (p *TokenClassificationPipeline) aggregateWord(entities []Entity) (Entity, error) {
	_ = "STUB: not implemented"
	return *new(Entity), nil
}

func (p *TokenClassificationPipeline) aggregateWords(entities []Entity) ([]Entity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Default behavior: group by IsSubword boundaries

// In split-words mode, we group by contiguous tokens of the same word boundary since we pretokenized.
// Since preEntities don’t carry word IDs directly, simulate grouping by contiguous offsets:
// break group if there is a gap between previous End and current Start (space) or heuristic non-subword.
// TODO: eventually we should export word IDs from the tokenizer to avoid this heuristic but the rust tokenizer bindings don't expose this yet
// and we also use other tokenizers in go backend.

// if there is a gap in offsets consider it a new word

func (p *TokenClassificationPipeline) aggregate(input backends.TokenizedInput, preEntities []Entity) ([]Entity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *TokenClassificationPipeline) getTag(entityName string) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

// defaulting to "I" if string is not in B- I- format

func (p *TokenClassificationPipeline) groupSubEntities(entities []Entity) (Entity, error) {
	_ = "STUB: not implemented"
	return *new(Entity), nil
}

// note: here we directly appeal to the tokenizer decoder with the tokenIds
// in the python code they pass the words to a token_to_string_method

// groupEntities group together adjacent tokens with the same entity predicted.
func (p *TokenClassificationPipeline) groupEntities(entities []Entity) ([]Entity, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create the grouped entity

// last entity remaining

// Run the pipeline on a string batch.
func (p *TokenClassificationPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

// RunPipeline is like Run but returns the concrete type rather than the interface.
func (p *TokenClassificationPipeline) RunPipeline(ctx context.Context, inputs []string) (*TokenClassificationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RunWords runs the pipeline for pre-split word inputs.
// Each input is a slice of words representing a pretokenized sentence.
// This is particularly useful when the user wants to control tokenization because of special tokens,
// hashtags, or other domain-specific tokenization needs.
func (p *TokenClassificationPipeline) RunWords(ctx context.Context, inputs [][]string) (*TokenClassificationOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
