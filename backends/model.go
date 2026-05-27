package backends

import (
	"context"
	"io"

	"github.com/knights-analytics/hugot/options"
)

type Model struct {
	ID                    string
	ORTModel              *ORTModel
	GoMLXModel            *GoMLXModel
	Tokenizer             *Tokenizer
	Destroy               func() error
	Pipelines             map[string]Pipeline
	IDLabelMap            map[int]string
	SeparatorToken        string
	Path                  string
	OnnxFilename          string
	OnnxPath              string
	OnnxReader            io.ReadCloser
	InputsMeta            []InputOutputInfo
	OutputsMeta           []InputOutputInfo
	MaxPositionEmbeddings int
	IsGenerative          bool
}

func LoadModel(ctx context.Context, path string, onnxFilename string, options *options.Options, isGenerative bool) (*Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// creation of the session. Only one output (either token or sentence embedding).

func GetOnnxModelPath(ctx context.Context, model *Model) error {
	_ = "STUB: not implemented"
	return nil
}

func getOnnxFiles(ctx context.Context, path string) ([][]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func loadModelConfig(ctx context.Context, model *Model) error {
	_ = "STUB: not implemented"
	// load config.json if it exists, to determine max_position_embeddings
	return nil
}

// Some multimodal models store text model config under text_config, so standardise that now

// Fallback 1: tokenizer_config.json may contain sep_token (common in HF models).

// Fallback 2: tokenizer.json post_processor.special_tokens may list the separator.
// We recognise the two canonical HF separators: [SEP] (BERT family) and </s> (RoBERTa family).

func ReshapeOutput[T float32 | int64 | int32](input []T, meta InputOutputInfo, batchSize int, paddingMask [][]bool, sequenceLength int) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// If no padding mask is provided (vision models), infer middle dim.

func flatDataTo2D[T float32 | int64 | int32](input []T, batchSize int, dimension int) [][]T {
	_ = "STUB: not implemented"
	// Input string, token, dimension
	return nil
}

// it can happen in principle that the embedding dimension is -1 if it was so exported from onnx even though there
// is a fixed out dimension so we do this.

func flatDataTo3D[T float32 | int64 | int32](input []T, paddingMask [][]bool, sequenceLength int, dimension int) [][][]T {
	_ = "STUB: not implemented"
	// Input string, token, dimension
	return nil
}

// skip whole token

// valid token, create embedding

// flatDataTo3DGeneric reshapes flat data into [batchSize][N][dimension] inferring N.
func flatDataTo3DGeneric[T float32 | int64 | int32](input []T, batchSize int, dimension int) [][][]T {
	_ = "STUB: not implemented"
	return nil

	// cannot infer without last dimension; return empty
}

// fallback: best-effort

func flatDataTo4D[T float32 | int64 | int32](input []T, paddingMask [][]bool, groupSize int, dimension int) [][][][]T {
	_ = "STUB: not implemented"
	return nil
}

// B
// S

// A

// skip this entire vector

// fill with zeros or ignore
