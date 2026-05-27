package pipelines

import (
	"context"
	"image"

	"github.com/knights-analytics/hugot/backends"
	"github.com/knights-analytics/hugot/options"
	"github.com/knights-analytics/hugot/util/imageutil"
)

// ObjectDetectionPipeline implements a Hugging Face-like object detection pipeline.
// It supports models that output bounding boxes and class scores.
type ObjectDetectionPipeline struct {
	*backends.BasePipeline
	IDLabelMap      map[int]string
	imageFormat     string
	BoxesOutput     string
	ScoresOutput    string
	preprocessSteps []imageutil.PreprocessStep
	normalizeSteps  []imageutil.NormalizationStep
	TopK            int
	ScoreThreshold  float32
	IouThreshold    float32
}

func (p *ObjectDetectionPipeline) addPreprocessSteps(steps ...imageutil.PreprocessStep) {
	_ = "STUB: not implemented"
	return
}

func (p *ObjectDetectionPipeline) addNormalizationSteps(steps ...imageutil.NormalizationStep) {
	_ = "STUB: not implemented"
	return
}

func (p *ObjectDetectionPipeline) setImageFormat(format string) { _ = "STUB: not implemented"; return }

func (p *ObjectDetectionPipeline) GetStatistics() backends.PipelineStatistics {
	_ = "STUB: not implemented"
	return *new(backends.PipelineStatistics)
}

type Detection struct {
	Label string
	Class int
	Box   [4]float32 // [xmin, ymin, xmax, ymax] in pixels
	Score float32
}
type ObjectDetectionOutput struct {
	Detections [][]Detection
}

func (o *ObjectDetectionOutput) GetOutput() []any { _ = "STUB: not implemented"; return nil }

// Options.

func WithBoxesOutput(name string) backends.PipelineOption[*ObjectDetectionPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithScoresOutput(name string) backends.PipelineOption[*ObjectDetectionPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithDetectionScoreThreshold(th float32) backends.PipelineOption[*ObjectDetectionPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithDetectionIouThreshold(th float32) backends.PipelineOption[*ObjectDetectionPipeline] {
	_ = "STUB: not implemented"
	return nil
}

func WithDetectionTopK(k int) backends.PipelineOption[*ObjectDetectionPipeline] {
	_ = "STUB: not implemented"
	return nil
}

// NewObjectDetectionPipeline initializes an object detection pipeline.
func NewObjectDetectionPipeline(sessionContext context.Context, config backends.PipelineConfig[*ObjectDetectionPipeline], s *options.Options, model *backends.Model) (*ObjectDetectionPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Set sensible default normalization for vision models (Rescale + ImageNet mean/std)

// Interface implementations.

func (p *ObjectDetectionPipeline) IsGenerative() bool { _ = "STUB: not implemented"; return false }

func (p *ObjectDetectionPipeline) GetModel() *backends.Model { _ = "STUB: not implemented"; return nil }

func (p *ObjectDetectionPipeline) GetMetadata() backends.PipelineMetadata {
	_ = "STUB: not implemented"
	return *new(backends.PipelineMetadata)
}

func (p *ObjectDetectionPipeline) GetStats() []string { _ = "STUB: not implemented"; return nil }

func (p *ObjectDetectionPipeline) Validate() error { _ = "STUB: not implemented"; return nil }

// preprocess images into tensors.
func (p *ObjectDetectionPipeline) preprocess(batch *backends.PipelineBatch, inputs []image.Image) error {
	_ = "STUB: not implemented"
	return nil
}

// forward inference.
func (p *ObjectDetectionPipeline) forward(ctx context.Context, batch *backends.PipelineBatch) error {
	_ = "STUB: not implemented"
	return nil
}

// postprocess parses boxes/scores, applies thresholds and NMS.
func (p *ObjectDetectionPipeline) postprocess(batch *backends.PipelineBatch) (*ObjectDetectionOutput, error) {
	_ = "STUB: not implemented"
	// Locate outputs by name in OutputValues
	// We assume order of Model.OutputsMeta corresponds to OutputValues
	return nil, nil
}

// Expected shapes: boxes [batch][num][4], scores [batch][num][num_classes]

func decodeDetections(boxes [][]float32, scores [][]float32, labels map[int]string, scoreTh float32, topK int) []Detection {
	_ = "STUB: not implemented"
	return nil
}

// Activate logits with softmax per box and pick best class (skip N/A class id 0).

// skip no-object class

// convertBoxToCorners assumes input box is [cx, cy, w, h] normalized to [0,1].
func convertBoxToCorners(b []float32) []float32 { _ = "STUB: not implemented"; return nil }

func iou(a, b [4]float32) float32 { _ = "STUB: not implemented"; return 0 }

func nonMaxSuppress(dets []Detection, iouTh float32) []Detection {
	_ = "STUB: not implemented"
	return nil
}

// Run with file paths.
func (p *ObjectDetectionPipeline) Run(ctx context.Context, inputs []string) (backends.PipelineBatchOutput, error) {
	_ = "STUB: not implemented"
	return *new(backends.PipelineBatchOutput), nil
}

func (p *ObjectDetectionPipeline) RunPipeline(ctx context.Context, inputs []string) (*ObjectDetectionOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *ObjectDetectionPipeline) RunWithImages(ctx context.Context, inputs []image.Image) (*ObjectDetectionOutput, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
