package backends

import (
	"context"
	"io"

	"github.com/gomlx/gomlx/backends"
	"github.com/gomlx/gomlx/pkg/core/graph"
	gomlxcontext "github.com/gomlx/gomlx/pkg/ml/context"
	"github.com/gomlx/onnx-gomlx/onnx"
	"github.com/knights-analytics/hugot/options"
)

// Default shape buckets for batch size and sequence length.
// These provide a good balance between padding overhead and JIT cache size.
var (
	defaultBatchBuckets    []int
	defaultSequenceBuckets []int
)

type GoMLXModel struct {
	Backend         backends.Backend
	OnnxModel       onnx.Model
	Ctx             *gomlxcontext.Context // ctx with the model's weights.
	Exec            *gomlxcontext.Exec    // exec is used to execute the model with a context.
	Call            func(ctx *gomlxcontext.Context, inputs []*graph.Node) []*graph.Node
	Destroy         func()
	BatchBuckets    []int // BatchBuckets defines bucket sizes for batch dimension padding.
	SequenceBuckets []int // SequenceBuckets defines bucket sizes for sequence length padding.
	MaxCache        int   // MaxCache sets the maximum number of unique input shapes to cache.
}

func createGoMLXModelBackend(model *Model, options *options.Options) error {
	_ = "STUB: not implemented"
	return nil
}

// Mark it to reuse variables: it will be an error to create a new variable – for safety.

// Read variables from ONNX model.

// Create model executor.

func getCacheAndBucketSizes(options *options.Options, backend string) (int, []int, []int) {
	_ = "STUB: not implemented"
	return 0,

		// Use configured buckets or fall back to defaults.
		nil, nil
}

// If using simpleGo, and user hasnt specified custom buckets, set max cache to unlimitted and disable bucketing

func loadInputOutputMetaGoMLX(model onnx.Model) ([]InputOutputInfo, []InputOutputInfo) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createInputTensorsGoMLX(batch *PipelineBatch, model *Model, padBatchDimension bool, padSequenceDimension bool) error {
	_ = "STUB: not implemented"
	return nil
}

// 1) prepare result containers - now we use all inputs, not filtering

// 2) build each tensor

// 3) assign and prepare cleanup

func runGoMLXSessionOnBatch(ctx context.Context, batch *PipelineBatch, p *BasePipeline) error {
	_ = "STUB: not implemented"
	return nil
}

// C code does not support context, so cancelling a context and/or session will usually trigger a segfault(panic).
// recover this here, so context can be cancelled gracefully and return an error.

// Transfer output tensors from device (TPU/GPU) to local (CPU) memory immediately.
//
// Go's GC doesn't see the large device-side allocations, so it doesn't feel pressure
// to reclaim tensor wrappers. By explicitly transferring to local memory and
// invalidating device copies, we release TPU/GPU memory as soon as the computation
// completes rather than waiting for eventual GC.

// Copy data from device to local memory
// Free device memory immediately

// When the batch dimension was padded to a bucket (e.g. batch=2 → bucket=8),
// the flat output contains paddedBatch × paddedSeq elements. Passing paddedBatch
// rows to ReshapeOutput with batch.Size=2 causes flatDataTo2D to compute the
// wrong row width (totalElems/2 instead of paddedSeq), mapping row-1 onto
// padding-batch rows whose NaN logits (from all-zero attention masks) propagate
// through softmax and make every score comparison with bestScore(-1) false.
// Strip excess batch rows here so ReshapeOutput only sees batch.Size rows.

// shapeBucket quantizes input dimensions to coarse buckets to reduce JIT cache pressure.
//
// XLA compiles a separate program for each unique input shape.
//
// By default we use "two-bit" bucketing, which provides a good balance between
// padding overhead and JIT cache size (approx. 1.41x increase between buckets).
//
//	Two-bit buckets: 1, 2, 3, 4, 6, 8, 12, 16, 24, 32, 48, 64, 96, 128, ...
//
// Trade-off: Slightly more padding waste for small inputs (e.g., batch=5 pads to 6),
// but dramatically fewer compiled programs in memory.
func shapeBucket(n int, buckets []int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (goMLXModel *GoMLXModel) Save(w io.Writer) error { _ = "STUB: not implemented"; return nil }

func createImageTensorsGoXLA(batch *PipelineBatch, model *Model, preprocessed [][][][]float32) error {
	_ = "STUB: not implemented"
	return nil
}

// Optionally add pixel_mask as ones if required by model.

// Infer mask dims

// Default to 3D [n,H,W]; if meta has 4 dims, use [n,1,H,W]

func createTabularTensorsGoMLX(batch *PipelineBatch, model *Model, features [][]float32) error {
	_ = "STUB: not implemented"
	return nil
}

// Assume first input is the tabular data.

// dynamic feature dim: infer from first sample

// Validate feature lengths

// No padding mask for tabular
