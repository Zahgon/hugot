package backends

import (
	"github.com/gomlx/go-huggingface/tokenizers/api"
)

type GoTokenizer struct {
	Tokenizer     api.Tokenizer
	TypeIDs       bool
	AttentionMask bool
}

func loadGoTokenizer(tokenizerBytes []byte, model *Model) error {
	_ = "STUB: not implemented"
	return nil
}

func getGoTokenizerOptions(model *Model) (api.EncodeOptions, bool, bool, error) {
	_ = "STUB: not implemented"
	return *new(api.EncodeOptions), false, false, nil
}

// Skip inputs that are handled at the model level

func tokenizeInputsGo(batch *PipelineBatch, tk *Tokenizer, inputs []string) {
	_ = "STUB: not implemented"
	return
}

// defaults to 0

// we need the offsets here for postprocessing later

func tokenizeInputPairsGo(batch *PipelineBatch, tk *Tokenizer, inputs [][2]string, sepToken string) {
	_ = "STUB: not implemented"
	return
}

func decodeGo(tokens []uint32, tokenizer *Tokenizer) string { _ = "STUB: not implemented"; return "" }

func getGoTokens(ids []int, tokenizer *Tokenizer) []string { _ = "STUB: not implemented"; return nil }

func convertGoOffsets(spans []api.TokenSpan) [][2]uint { _ = "STUB: not implemented"; return nil }

func allInputTokensGo(pipeline *BasePipeline) error { _ = "STUB: not implemented"; return nil }
