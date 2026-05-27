package hugot

import (
	"context"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/datasets"
)

type earlyStopping struct {
	patience  int     // number of epochs to wait for improvement before stopping
	tolerance float32 // tolerance for loss comparison
}
type TrainingStatistics struct {
	EpochTrainLosses []float32 `json:"epochTrainLosses"` // stores the training loss for each epoch
	EpochEvalLosses  []float32 `json:"epochEvalLosses"`  // stores the evaluation loss for each epoch
}
type TrainingSession struct {
	pipeline         backends.Pipeline
	earlyStopping    *earlyStopping
	backend          string
	statistics       TrainingStatistics
	freezeLayers     []int // freeze the layers of the transformer model, 0 is the first layer etc. Set [-1] to freeze all layers apart from the last one
	config           TrainingConfig
	maxEpochs        int
	cuda             bool
	freezeEmbeddings bool // freeze the embedding layers of the transfomer model
}

// GetPipeline returns the pipeline used in the training session.
func (s *TrainingSession) GetPipeline() backends.Pipeline {
	_ = "STUB: not implemented"
	return *new(backends.Pipeline)
}

func (s *TrainingSession) Destroy() error { _ = "STUB: not implemented"; return nil }

type TrainingOption func(eo *TrainingSession) error

func WithEpochs(epochs int) TrainingOption { _ = "STUB: not implemented"; return *new(TrainingOption) }

func WithFreezeEmbeddings() TrainingOption { _ = "STUB: not implemented"; return *new(TrainingOption) }

func WithFreezeLayers(layers []int) TrainingOption {
	_ = "STUB: not implemented"
	return *new(TrainingOption)
}

func WithCuda() TrainingOption { _ = "STUB: not implemented"; return *new(TrainingOption) }

func WithEarlyStopping() TrainingOption { _ = "STUB: not implemented"; return *new(TrainingOption) }

// default patience and tolerance

func WithEarlyStoppingParams(patience int, tolerance float32) TrainingOption {
	_ = "STUB: not implemented"
	return *new(TrainingOption)
}

type TrainingConfig struct {
	TrainDataset         datasets.Dataset
	TrainEvalDataset     datasets.Dataset // used to evaluate on training
	EvalDataset          datasets.Dataset // optional, used for early stopping and eval metrics
	GOMLXTrainingOptions *GOMLXTrainingOptions
	ModelPath            string
	OnnxFilename         string
	Options              []TrainingOption
	Verbose              bool
}

func newTrainingSession[T backends.Pipeline](sessionContext context.Context, backend string, config TrainingConfig) (*TrainingSession, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// default to 100 epochs if not set

// hook the datasets up with the pipeline for tokenization

func (s *TrainingSession) Train() error { _ = "STUB: not implemented"; return nil }

// Save serializes the trained model as an onnx model.
// If a tokenizer is present, the tokenizer files are copied from the untrained model directory to the trained model.
// Path is the full path to the directory where the model will be saved.
func (s *TrainingSession) Save(ctx context.Context, path string) error {
	_ = "STUB: not implemented"
	return nil
}

// copy tokenizer files from original model

func copyTokenizer(ctx context.Context, from, to string) error {
	_ = "STUB: not implemented"
	return nil
}
