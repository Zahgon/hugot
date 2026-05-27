package backends

import (
	"context"

	"github.com/knights-analytics/hugot/options"
)

type Tokenizer struct {
	RustTokenizer    *RustTokenizer
	GoTokenizer      *GoTokenizer
	TokenizerTimings *timings
	Destroy          func() error
	Runtime          string
	MaxAllowedTokens int
}

func LoadTokenizer(ctx context.Context, model *Model, s *options.Options) error {
	_ = "STUB: not implemented"
	return nil
}

func TokenizeInputs(batch *PipelineBatch, tk *Tokenizer, inputs []string) {
	_ = "STUB: not implemented"
	return
}

func TokenizeInputPairs(batch *PipelineBatch, tk *Tokenizer, inputs [][2]string, sepToken string) {
	_ = "STUB: not implemented"
	return
}

func patchBertSequenceTokenTypeIDs(batch *PipelineBatch, sepToken string) {
	_ = "STUB: not implemented"
	// Fix token_type_ids for BERT-style models when we manually concatenated the pair as a single sequence.
	// Pattern expected: [CLS] query [SEP] doc [SEP]
	// HF sets token_type_ids=0 up to and including first [SEP], then 1 for remainder (including final [SEP]).
	return
}

// Only adjust if type ids exist and are all zero

// Find first [SEP] token index (skip position 0 which should be [CLS])

// nothing to split

func AllInputTokens(pipeline *BasePipeline) error { _ = "STUB: not implemented"; return nil }

func Decode(tokens []uint32, tokenizer *Tokenizer) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
