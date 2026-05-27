//go:build !cgo || (!ORT && !XLA && !ALL)

package backends

type RustTokenizer struct{}

func loadRustTokenizer(_ []byte, _ *Model) error { _ = "STUB: not implemented"; return nil }

func tokenizeInputsRust(_ *PipelineBatch, _ *Tokenizer, _ []string) {
	_ = "STUB: not implemented"
	return
}

func tokenizeInputPairsRust(_ *PipelineBatch, _ *Tokenizer, _ [][2]string, _ string) {
	_ = "STUB: not implemented"
	return
}

func decodeRust(_ []uint32, _ *Tokenizer, _ bool) string { _ = "STUB: not implemented"; return "" }

func allInputTokensRust(_ *BasePipeline) error { _ = "STUB: not implemented"; return nil }
