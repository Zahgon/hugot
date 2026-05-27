package backends

import (
	"image"

	"github.com/knights-analytics/hugot/util/imageutil"
)

// DetectImageTensorFormat inspects the first image-like input and infers NHWC or NCHW.
func DetectImageTensorFormat(model *Model) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Prefer a typical image input name

// Unexpected for image tensors; default to NCHW

// If we see channel=3 in second dim -> NCHW; if 3 in last dim -> NHWC.

// Dynamic or unknown — default to NCHW

func CreateImageTensors(batch *PipelineBatch, model *Model, preprocessed [][][][]float32, runtime string) error {
	_ = "STUB: not implemented"
	return nil
}

// PreprocessImages preprocesses images into a 4D tensor slice according to format and steps.
func PreprocessImages(format string, images []image.Image, preprocess []imageutil.PreprocessStep, normalize []imageutil.NormalizationStep) ([][][][]float32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
