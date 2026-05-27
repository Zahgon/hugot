//go:build cgo && (ORT || ALL)

package backends

import (
	"context"
	"sync"

	ort "github.com/yalue/onnxruntime_go"

	"github.com/knights-analytics/hugot/options"
	"github.com/knights-analytics/ortgenai"
)

type ORTModel struct {
	Session           *ort.DynamicAdvancedSession
	GenerativeSession *ortgenai.Session
	GenerativeEngine  *ortgenai.Engine
	SessionOptions    *ort.SessionOptions
	Options           *options.OrtOptions
	Destroy           func() error
}

var generativeBackendMutex = sync.Mutex{}

func mapORTOptions(options *options.Options) ([]string, map[string]map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// let default EPs be used

// CUDA

// CoreML

// DirectML

// Map device id to string map expected by advanced session

// OpenVINO

// TensorRT

// TensorRT

// Extra EPs

func createORTGenerativeSession(ctx context.Context, model *Model, options *options.Options) error {
	_ = "STUB: not implemented"
	return nil
}

func initialiseORTGenAI(ctx context.Context, options *options.Options) error {
	_ = "STUB: not implemented"
	return nil
}

func runGenerativeORTSessionOnBatch(ctx context.Context, batch *PipelineBatch, p *BasePipeline, maxLength int, stopSequences []string, temperature *float64, topP *float64, seed *int, tools []string, guidance *Guidance) (chan SequenceDelta, chan error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Check if we have multimodal tensors to use instead of text tokenization

// Map optional guidance config to ortgenai type.

// Precompute stop-sequence data once (avoid scanning growing buffers).

// Keep only a rolling tail per sequence (enough to detect stop sequences).
// This avoids O(n^2) behavior from (a) growing strings and (b) repeatedly searching the whole prefix.

// Already complete; ignore further tokens for this sequence.

// EOS terminates sequence; no token content to forward.

// Forward token immediately.

// Detect stop sequences using a bounded rolling window.
// Window size is maxStopLen + len(current token) to catch:
//  - stops spanning the boundary between previous tail and this token
//  - stops fully contained inside a single (possibly long) token

func createORTModelBackend(model *Model, options *options.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// TODO: currently models with external data can only load from regular filesystems, and require dir change

func loadInputOutputMetaORTReader(onnxBytes []byte) ([]InputOutputInfo, []InputOutputInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func loadInputOutputMetaORTFile(onnxPath string) ([]InputOutputInfo, []InputOutputInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func createInputTensorsORT(batch *PipelineBatch, model *Model) error {
	_ = "STUB: not implemented"
	return nil
}

// 1) prepare result containers - now we use all inputs, not filtering

// 2) build each tensor

// Handle regular input tensors

// 1-indexed positions

// create the ONNX Runtime tensor for regular inputs

// 3) assign and prepare cleanup

func runORTSessionOnBatch(ctx context.Context, batch *PipelineBatch, p *BasePipeline) error {
	_ = "STUB: not implemented"
	return nil
}

// C code does not support context, so cancelling a context and/or session will usually trigger a segfault(panic).
// recover this here, so context can be cancelled gracefully and return an error.

// store resulting tensors

func convertORTInputOutputs(inputOutputs []ort.InputOutputInfo) []InputOutputInfo {
	_ = "STUB: not implemented"
	return nil
}

// createTabularTensorsORT flattens [][]float32 features into a [batch, feature_dim] tensor.
// Currently supports models with a single input of 2D shape (batch, features).
func createTabularTensorsORT(batch *PipelineBatch, model *Model, features [][]float32) error {
	_ = "STUB: not implemented"
	return nil
}

// Assume first input is the tabular data.

// dynamic feature dim: infer from first sample

// Validate feature lengths

// No padding mask for tabular

func createImageTensorsORT(batch *PipelineBatch, model *Model, preprocessed [][][][]float32) error {
	_ = "STUB: not implemented"
	return nil
}

// Prepare inputs slice according to model input metadata order.

// Helper to infer mask dims

// Try to find known H and W; fallback to image h,w

// Build pixel_mask tensor of ones using int64 dtype, shape [n, H, W] or [n,1,H,W] depending on meta.

// Default to 3D [n,H,W]

// Some models expect [n,1,H,W]

// If creating 4D fails, try 3D fallback

// Only destroy once; avoid double-destroy if multiple inputs map to same tensor

// If only one input, just that tensor

func CreateMessagesORT(batch *PipelineBatch, inputs any, systemPrompt string) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if any messages contain images
