package datasets

import (
	"context"

	"github.com/knights-analytics/hugot/backends"
)

func (s *SemanticSimilarityDataset) SetVerbose(v bool) { _ = "STUB: not implemented"; return }

func (s *SemanticSimilarityDataset) SetTokenizationPipeline(pipeline backends.Pipeline) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *SemanticSimilarityDataset) Validate() error { _ = "STUB: not implemented"; return nil }

// SemanticSimilarityExample is a single example for the semantic similarity dataset.
type SemanticSimilarityExample struct {
	Data      map[string]any // to store any additional data for the example. Not used by the dataset.
	Sentence1 string         `json:"sentence1"`
	Sentence2 string         `json:"sentence2"`
	Score     float32        `json:"score"`
}
type ExamplePreprocessFunc func([]SemanticSimilarityExample) ([]SemanticSimilarityExample, error)

// NewSemanticSimilarityDataset creates a new SemanticSimilarityDataset.
// The trainingPath must be a .jsonl file where each line has the following format:
// {"sentence1":"A plane is taking off.","sentence2":"An air plane is taking off.","score":1.0}
// The score is a float value between 0 and 1.
// preprocessFunc here must be a function that takes a slice of SemanticSimilarityExample and returns a slice of SemanticSimilarityExample.
// This function can be used to apply any custom preprocessing to the example batch before they are passed to the model.
func NewSemanticSimilarityDataset(ctx context.Context, trainingPath string, batchSize int, preprocessFunc ExamplePreprocessFunc) (*SemanticSimilarityDataset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewInMemorySemanticSimilarityDataset creates a new SemanticSimilarityDataset in memory from a slice of examples.
// preprocessFunc here must be a function that takes a slice of SemanticSimilarityExample and returns a slice of SemanticSimilarityExample.
// This function can be used to apply any custom preprocessing to the example batch before they are passed to the model.
func NewInMemorySemanticSimilarityDataset(examples []SemanticSimilarityExample, batchSize int, preprocessFunc ExamplePreprocessFunc) (*SemanticSimilarityDataset, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Reset resets the dataset to the beginning of the training data (after the epoch is done).
func (s *SemanticSimilarityDataset) Reset() { _ = "STUB: not implemented"; return }

// note: these panics will be catched later with the TryExcept

// restart the reader

// YieldRaw returns the next raw batch of examples from the dataset. Note that if a preprocessing function has been
// provided at creation time, the examples will be preprocessed before being returned.
func (s *SemanticSimilarityDataset) YieldRaw() ([]SemanticSimilarityExample, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// in memory dataset

// return error for reset

func (s *SemanticSimilarityDataset) Close() error { _ = "STUB: not implemented"; return nil }
