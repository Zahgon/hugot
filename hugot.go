package hugot

import (
	"context"
	"sync"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
	"github.com/knights-analytics/hugot/pipelines"
)

// Session allows for the creation of new pipelines and holds the pipeline already created.
type Session struct {
	featureExtractionPipelines      pipelineMap[*pipelines.FeatureExtractionPipeline]
	tokenClassificationPipelines    pipelineMap[*pipelines.TokenClassificationPipeline]
	textClassificationPipelines     pipelineMap[*pipelines.TextClassificationPipeline]
	zeroShotClassificationPipelines pipelineMap[*pipelines.ZeroShotClassificationPipeline]
	crossEncoderPipelines           pipelineMap[*pipelines.CrossEncoderPipeline]
	imageClassificationPipelines    pipelineMap[*pipelines.ImageClassificationPipeline]
	objectDetectionPipelines        pipelineMap[*pipelines.ObjectDetectionPipeline]
	textGenerationPipelines         pipelineMap[*pipelines.TextGenerationPipeline]
	tabularPipelines                pipelineMap[*pipelines.TabularPipeline]
	questionAnsweringPipelines      pipelineMap[*pipelines.QuestionAnsweringPipeline]
	models                          map[string]*backends.Model
	modelLocks                      map[string]*sync.Mutex
	modelLocksMu                    sync.Mutex
	pipelineLocks                   map[string]*sync.Mutex
	pipelineLocksMu                 sync.Mutex
	options                         *options.Options
	environmentDestroy              func() error
	sessionContext                  context.Context
	cancelSessionContext            context.CancelFunc
}

func (s *Session) GetModels() map[string]*backends.Model { _ = "STUB: not implemented"; return nil }

func (s *Session) getModelLock(modelID string) *sync.Mutex { _ = "STUB: not implemented"; return nil }

func (s *Session) removeModelLock(modelID string) { _ = "STUB: not implemented"; return }

func (s *Session) getPipelineLock(name string) *sync.Mutex { _ = "STUB: not implemented"; return nil }

func newSession(ctx context.Context, backend string, opts ...options.WithOption) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Collect options into a struct, so they can be applied in the correct order later

type pipelineMap[T backends.Pipeline] map[string]T

func (m pipelineMap[T]) GetStatistics() map[string]backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return nil
}

// FeatureExtractionConfig is the configuration for a feature extraction pipeline.
type FeatureExtractionConfig = backends.PipelineConfig[*pipelines.FeatureExtractionPipeline]

// FeatureExtractionOption is an option for a feature extraction pipeline.
type FeatureExtractionOption = backends.PipelineOption[*pipelines.FeatureExtractionPipeline]

// TextClassificationConfig is the configuration for a text classification pipeline.
type TextClassificationConfig = backends.PipelineConfig[*pipelines.TextClassificationPipeline]

// TextClassificationOption is an option for a text classification pipeline.
type TextClassificationOption = backends.PipelineOption[*pipelines.TextClassificationPipeline]

// ZeroShotClassificationConfig is the configuration for a zero shot classification pipeline.
type ZeroShotClassificationConfig = backends.PipelineConfig[*pipelines.ZeroShotClassificationPipeline]

// ZeroShotClassificationOption is an option for a zero shot classification pipeline.
type ZeroShotClassificationOption = backends.PipelineOption[*pipelines.ZeroShotClassificationPipeline]

// TokenClassificationConfig is the configuration for a token classification pipeline.
type TokenClassificationConfig = backends.PipelineConfig[*pipelines.TokenClassificationPipeline]

// TokenClassificationOption is an option for a token classification pipeline.
type TokenClassificationOption = backends.PipelineOption[*pipelines.TokenClassificationPipeline]

// CrossEncoderConfig is the configuration for a cross encoder pipeline.
type CrossEncoderConfig = backends.PipelineConfig[*pipelines.CrossEncoderPipeline]

// CrossEncoderOption is an option for a cross encoder pipeline.
type CrossEncoderOption = backends.PipelineOption[*pipelines.CrossEncoderPipeline]

// ImageClassificationConfig is the configuration for an image classification pipeline.
type ImageClassificationConfig = backends.PipelineConfig[*pipelines.ImageClassificationPipeline]

// ImageClassificationOption is an option for an image classification pipeline.
type ImageClassificationOption = backends.PipelineOption[*pipelines.ImageClassificationPipeline]

// ObjectDetectionConfig is the configuration for an object detection pipeline.
type ObjectDetectionConfig = backends.PipelineConfig[*pipelines.ObjectDetectionPipeline]

// ObjectDetectionOption is an option for an object detection pipeline.
type ObjectDetectionOption = backends.PipelineOption[*pipelines.ObjectDetectionPipeline]

// TextGenerationConfig is the configuration for a text generation pipeline.
type TextGenerationConfig = backends.PipelineConfig[*pipelines.TextGenerationPipeline]

// TextGenerationOption is an option for a text generation pipeline.
type TextGenerationOption = backends.PipelineOption[*pipelines.TextGenerationPipeline]

// TabularConfig is the configuration for a tabular pipeline.
type TabularConfig = backends.PipelineConfig[*pipelines.TabularPipeline]

// TabularOption is an option for a tabular pipeline.
type TabularOption = backends.PipelineOption[*pipelines.TabularPipeline]

// QuestionAnsweringConfig is the configuration for a question answering pipeline.
type QuestionAnsweringConfig = backends.PipelineConfig[*pipelines.QuestionAnsweringPipeline]

// QuestionAnsweringOption is an option for a question answering pipeline.
type QuestionAnsweringOption = backends.PipelineOption[*pipelines.QuestionAnsweringPipeline]

// NewPipeline can be used to create a new pipeline of type T. The initialised pipeline will be returned and it
// will also be stored in the session object so that all created pipelines can be destroyed with session.Destroy()
// at once.
func NewPipeline[T backends.Pipeline](s *Session, pipelineConfig backends.PipelineConfig[T]) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

// Load model if it has not been loaded already

func initializePipeline[T backends.Pipeline](sessionContext context.Context, p T, pipelineConfig backends.PipelineConfig[T], options *options.Options, model *backends.Model) (T, string, error) {
	_ = "STUB: not implemented"
	return *new(T), "", nil
}

// GetPipeline can be used to retrieve a pipeline of type T with the given name from the session.
func GetPipeline[T backends.Pipeline](s *Session, name string) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}

func GetPipelines[T backends.Pipeline](s *Session) (map[string]T, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ClosePipeline[T backends.Pipeline](s *Session, name string) error {
	_ = "STUB: not implemented"
	return nil
}

type pipelineNotFoundError struct {
	pipelineName string
}

func (e *pipelineNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

// GetStatistics returns runtime statistics for all initialized pipelines for profiling purposes. We currently record for each pipeline:
// the total runtime of the tokenization step
// the number of batch calls to the tokenization step
// the average time per tokenization batch call
// the total runtime of the inference (i.e. onnxruntime) step
// the number of batch calls to the onnxruntime inference
// the average time per onnxruntime inference batch call.
func (s *Session) GetStatistics() map[string]backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return nil
}

// PrintStatistics prints runtime statistics for all initialized pipelines to stdout.
func (s *Session) PrintStatistics() { _ = "STUB: not implemented"; return }

// Destroy deletes the hugot session and onnxruntime environment and all initialized pipelines, freeing memory.
// A hugot session should be destroyed when not neeeded any more, preferably with a defer() call.
func (s *Session) Destroy() error { _ = "STUB: not implemented"; return nil }
