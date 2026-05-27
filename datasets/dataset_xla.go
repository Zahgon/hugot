package datasets

import (
	"bufio"
	"io"

	"github.com/gomlx/gomlx/pkg/core/tensors"
	"github.com/gomlx/gomlx/pkg/ml/train"
	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/pipelines"
)

type Dataset interface {
	train.Dataset
	Validate() error
	SetTokenizationPipeline(pipeline backends.Pipeline) error
	SetVerbose(bool)
	Close() error
}

// SemanticSimilarityDataset is a dataset for fine-tuning a feature extraction pipeline for textual semantic similarity.
type SemanticSimilarityDataset struct {
	train.Dataset
	sourceFile       io.ReadCloser
	preprocessFunc   ExamplePreprocessFunc
	pipeline         *pipelines.FeatureExtractionPipeline
	reader           *bufio.Reader
	trainingPath     string
	trainingExamples []SemanticSimilarityExample
	batchSize        int
	batchN           int
	verbose          bool
}

// Yield returns the next batch of examples from the dataset. The examples are tokenized and converted to tensors for the training process.
func (s *SemanticSimilarityDataset) Yield() (spec any, inputs []*tensors.Tensor, labels []*tensors.Tensor, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil, nil, nil
}
