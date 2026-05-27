package pipelines

// QuestionAnsweringPipeline is a go implementation of Hugging Face's question answering pipeline.
// https://github.com/huggingface/transformers/blob/main/src/transformers/pipelines/question_answering.py
//
// It supports extractive QA: given a question and a context passage the model predicts start and
// end logits over the context tokens and the top-k highest-scoring spans are returned as answers.
// By default (TopK=1) only the single best span is returned per input.
//
// The underlying ONNX model is expected to expose two outputs in order:
//
//	0 – start_logits  [batch, sequence_length]
//	1 – end_logits    [batch, sequence_length]
//
// Typical models: distilbert-base-uncased-distilled-squad, bert-large-uncased-whole-word-masking-finetuned-squad.

import (
	"context"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
)

// QuestionAnsweringInput holds a single question/context pair.
type QuestionAnsweringInput struct {
	Question string
	Context  string
}

// QuestionAnsweringOutput holds the result for a single question/context pair.
type QuestionAnsweringOutput struct {
	// Answer is the extracted answer string from Context.
	Answer string
	// Score is start_prob[best_start] * end_prob[best_end].
	Score float32
	// Start is the byte offset of the answer start inside Context.
	Start uint
	// End is the byte offset (exclusive) of the answer end inside Context.
	End uint
}

// QuestionAnsweringBatchOutput holds results for a whole batch.
// Each element of Outputs corresponds to one input and contains answers ranked by score (best first).
// With the default TopK=1 each inner slice has exactly one element.
type QuestionAnsweringBatchOutput struct {
	Outputs [][]QuestionAnsweringOutput
}

func (o *QuestionAnsweringBatchOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

// QuestionAnsweringPipeline holds the pipeline configuration.
type QuestionAnsweringPipeline struct {
	*backends.BasePipeline
	// MaxAnswerLength is the maximum number of tokens allowed in the answer span (default 15).
	MaxAnswerLength int
	// TopK is the number of ranked answer spans to return per input (default 1).
	TopK int
}

// PIPELINE OPTIONS

// WithMaxAnswerLength sets the maximum number of tokens that the answer span may cover.
func WithMaxAnswerLength(n int) backends.PipelineOption[*QuestionAnsweringPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// WithTopKAnswers sets the number of ranked answer spans to return per input.
// When k > 1 each element of QuestionAnsweringBatchOutput.Outputs is a slice of up to k answers
// sorted by score descending.
func WithTopKAnswers(k int) backends.PipelineOption[*QuestionAnsweringPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// NewQuestionAnsweringPipeline initialises a question answering pipeline.
func NewQuestionAnsweringPipeline(sessionContext context.Context, config backends.PipelineConfig[*QuestionAnsweringPipeline], s *options.Options, model *backends.Model) (*QuestionAnsweringPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// INTERFACE IMPLEMENTATIONS

func (p *QuestionAnsweringPipeline) IsGenerative() bool { _ = "STUB: not implemented"; return false }

func (p *QuestionAnsweringPipeline) GetModel() *backends.Model {
	_ = "STUB: not implemented"

	// GetMetadata returns metadata for both output tensors (start_logits, end_logits).
	return nil
}

func (p *QuestionAnsweringPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

// GetStatistics returns runtime statistics for the pipeline.
func (p *QuestionAnsweringPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

// Validate checks that the pipeline configuration is valid.
func (p *QuestionAnsweringPipeline) Validate() error { _ = "STUB: not implemented"; return nil }

// preprocess tokenises each question/context pair into a single combined sequence.
func (p *QuestionAnsweringPipeline) preprocess(batch *backends.PipelineBatch, inputs []QuestionAnsweringInput) error {
	_ = "STUB: not implemented"
	return nil
}

// forward runs the ONNX session on the batch.
func (p *QuestionAnsweringPipeline) forward(ctx context.Context, batch *backends.PipelineBatch) error {
	_ = "STUB: not implemented"
	return nil
}

// postprocess extracts the top-k answer spans for each item in the batch, ranked by score descending.
func (p *QuestionAnsweringPipeline) postprocess(batch *backends.PipelineBatch, inputs []QuestionAnsweringInput) (*QuestionAnsweringBatchOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert logits to probabilities over the sequence dimension.

// Determine the first context token (type_id == 1 or after the separator).

// Collect the top-k scoring spans and map them to character offsets.

// Run implements the Pipeline interface. Inputs are interleaved [question0, context0, question1, context1, …].
func (p *QuestionAnsweringPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

// RunPipeline is the typed entry point for the question answering pipeline.
func (p *QuestionAnsweringPipeline) RunPipeline(ctx context.Context, inputs []QuestionAnsweringInput) (*QuestionAnsweringBatchOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HELPERS

// spanCandidate holds a scored token span.
type spanCandidate struct {
	start, end int
	score      float32
}

// rankedSpans returns up to k highest-scoring valid (start, end) token spans in descending score order.
// Spans are non-overlapping: once a span is selected, any candidate that shares at least one token
// with it is discarded before the next pick.
func rankedSpans(startProbs, endProbs []float32, ctxStart, seqLen, maxAnswerLen, k int) []spanCandidate {
	_ = "STUB: not implemented"
	return nil
}

// contextStartIndex returns the index of the first token that belongs to the context
// (i.e. the answer candidate region). It prefers type_id == 1 when available; otherwise
// it falls back to finding the separator token and returning the index after it.
func contextStartIndex(input backends.TokenizedInput, sepToken string) int {
	_ = "STUB: not implemented"
	return 0
}

// Fallback: skip past the first occurrence of the separator token.

// spanToContextChars converts a [startTok, endTok] token span into byte offsets that are
// relative to the original context string (not the combined question+sep+context string).
//
// The token offsets stored in TokenizedInput are relative to the full combined string;
// we derive the context character offset from the first token whose type_id == 1 (or the
// first token after the separator, as a fallback).
func spanToContextChars(input backends.TokenizedInput, startTok, endTok int, context string) (uint, uint) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Find the character offset where the context begins in the combined string.

// Fallback: when TypeIDs are not populated (e.g. DistilBERT which has no token_type_ids input),
// derive the context's start offset from the raw combined string.
// input.Raw = question + separator + context, so the context begins at len(input.Raw) - len(context).

// non-negative by the guard above
//nolint:gosec // diff is guaranteed non-negative

// Clamp to context length.
