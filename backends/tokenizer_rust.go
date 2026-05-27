//go:build cgo && (ORT || XLA || ALL)

package backends

import (
	"github.com/daulet/tokenizers"
)

type RustTokenizer struct {
	Tokenizer *tokenizers.Tokenizer
	Options   []tokenizers.EncodeOption
}

func loadRustTokenizer(tokenizerBytes []byte, model *Model) error {
	_ = "STUB: not implemented"
	return nil
}

// tokenizer init

func getRustTokenizerOptions(model *Model) ([]tokenizers.EncodeOption, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip inputs that are handled at the model level

func tokenizeInputsRust(batch *PipelineBatch, tk *Tokenizer, inputs []string) {
	_ = "STUB: not implemented"
	return
}

// we need the offsets here for postprocessing later

func tokenizeInputPairsRust(batch *PipelineBatch, tk *Tokenizer, inputs [][2]string, sepToken string) {
	_ = "STUB: not implemented"
	return
}

// Adjust type IDs. Since we manually concatenated, we should try to find where the second part starts.
// However, the most robust way if we don't know the separator is to tokenize separately,
// but daulet/tokenizers EncodeWithOptions is better if it supports pairs.
// Given I cannot find WithEncodePair, I will use the patch logic if TypeIDs are all zero.

// we need the offsets here for postprocessing later

func decodeRust(tokens []uint32, tokenizer *Tokenizer, skipSpecialTokens bool) string {
	_ = "STUB: not implemented"
	return ""
}

func convertRustOffsets(input []tokenizers.Offset) [][2]uint { _ = "STUB: not implemented"; return nil }

func allInputTokensRust(pipeline *BasePipeline) error { _ = "STUB: not implemented"; return nil }
