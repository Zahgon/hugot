package backends

import (
	"context"
	"time"

	"github.com/knights-analytics/hugot/options"
)

// BasePipeline can be embedded by a pipeline.
type BasePipeline struct {
	Model           *Model
	PipelineTimings *timings
	PipelineName    string
	Runtime         string
	SessionContext  context.Context
}

type InputOutputInfo struct {
	// The name of the input or output
	Name string
	// The input or output's dimensions, if it's a tensor. This should be
	// ignored for non-tensor types.
	Dimensions Shape
}
type Shape []int64

func (s Shape) String() string { _ = "STUB: not implemented"; return "" }

func (s Shape) ValuesInt() []int { _ = "STUB: not implemented"; return nil }

// NewShape Returns a Shape, with the given dimensions.
func NewShape(dimensions ...int64) Shape { _ = "STUB: not implemented"; return *new(Shape) }

type OutputInfo struct {
	Name       string
	Dimensions []int64
}
type PipelineMetadata struct {
	OutputsInfo []OutputInfo
}
type PipelineBatchOutput interface {
	GetOutput() []any
}

// Pipeline is the interface that any pipeline must implement.
type Pipeline interface {
	GetStatistics() PipelineStatistics                          // Get the pipeline running statistics
	Validate() error                                            // Validate the pipeline for correctness
	GetMetadata() PipelineMetadata                              // Return metadata information for the pipeline
	GetModel() *Model                                           // Return the model used by the pipeline
	IsGenerative() bool                                         // Return whether the pipeline is generative
	Run(context.Context, []string) (PipelineBatchOutput, error) // Run the pipeline on an input
}

type PipelineStatistics struct {
	TokenizerTotalTime             time.Duration
	TokenizerExecutionCount        uint64
	TokenizerAvgQueryTime          time.Duration
	OnnxTotalTime                  time.Duration
	OnnxExecutionCount             uint64
	OnnxAvgQueryTime               time.Duration
	TotalQueries                   uint64
	TotalDocuments                 uint64
	AverageLatency                 time.Duration
	AverageBatchSize               float64
	FilteredResults                uint64
	AvgPrefillSeconds              float64
	TokensPerSecond                float64
	CumulativePrefillSum           float64
	CumulativePrefillCount         int
	CumulativeTokens               int
	CumulativeTokenDurationSeconds float64
}

func (p *PipelineStatistics) ComputeTokenizerStatistics(timings *timings) {
	_ = "STUB: not implemented"
	return
}

func (p *PipelineStatistics) ComputeOnnxStatistics(timings *timings) {
	_ = "STUB: not implemented"
	return
}

func (p *PipelineStatistics) Print() { _ = "STUB: not implemented"; return }

// PipelineOption is an option for a pipeline type.
type PipelineOption[T Pipeline] func(eo T) error

// PipelineConfig is a configuration for a pipeline type that can be used
// to create that pipeline.
type PipelineConfig[T Pipeline] struct {
	ModelPath    string
	Name         string
	OnnxFilename string
	Options      []PipelineOption[T]
}
type timings struct {
	NumCalls uint64
	TotalNS  uint64
}

// TokenizedInput holds the result of running tokenizer on an input.
type TokenizedInput struct {
	Raw               string
	Tokens            []string
	TokenIDs          []uint32
	TypeIDs           []uint32
	AttentionMask     []uint32
	SpecialTokensMask []uint32
	Offsets           [][2]uint
	MaxAttentionIndex int
}

// PipelineBatch represents a batch of inputs that runs through the pipeline.
type PipelineBatch struct {
	InputValues       any
	DestroyInputs     func() error
	Input             []TokenizedInput
	PaddingMask       [][]bool
	OutputValues      []any
	Size              int
	MaxSequenceLength int
	MaxNewTokens      int
	// PaddedBatchSize is the bucketed batch size used when XLA pads the batch dimension.
	// Zero means no batch padding was applied (ORT / GO backend).
	PaddedBatchSize int
	// Multimodal support
	Images            any // Will hold *ortgenai.Images for generative models
	DestroyMultimodal func() error
}

func (b *PipelineBatch) Destroy() error { _ = "STUB: not implemented"; return nil }

// NewBatch initializes a new batch for inference.
func NewBatch(size int) *PipelineBatch { _ = "STUB: not implemented"; return nil }

func GetNames(info []InputOutputInfo) []string { _ = "STUB: not implemented"; return nil }

func RunSessionOnBatch(ctx context.Context, batch *PipelineBatch, p *BasePipeline) error {
	_ = "STUB: not implemented"
	return nil
}

func RunGenerativeSessionOnBatch(ctx context.Context, batch *PipelineBatch, p *BasePipeline, maxLength int, stopSequences []string, temperature *float64, topP *float64, seed *int, tools []string, guidance *Guidance) (chan SequenceDelta, chan error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func CreateMessages(batch *PipelineBatch, p *BasePipeline, inputs any, systemPrompt string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateInputTensorsTraining creates input tensors for training. Same as CreateInputTensors but
// we never pad the batch size as we expect regular batch sizes from the dataset.
func CreateInputTensorsTraining(batch *PipelineBatch, model *Model, runtime string) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateInputTensors(batch *PipelineBatch, model *Model, runtime string) error {
	_ = "STUB: not implemented"
	return nil
}

// only pad the batch dimension if we have a limited cache

// CreateTabularTensors builds input tensors for classic ML/tabular models.
func CreateTabularTensors(batch *PipelineBatch, model *Model, features [][]float32, runtime string) error {
	_ = "STUB: not implemented"
	return nil
}

func NewBasePipeline[T Pipeline](sessionContext context.Context, config PipelineConfig[T], s *options.Options, model *Model) (*BasePipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateModelBackend(ctx context.Context, model *Model, s *options.Options) error {
	_ = "STUB: not implemented"
	return nil
}
