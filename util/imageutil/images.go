package imageutil

import (
	"context"
	"image"

	_ "image/gif"  // adds gif support
	_ "image/jpeg" // adds jpeg support
	_ "image/png"  // adds png support

	_ "golang.org/x/image/webp" // adds webp support
)

func LoadImagesFromPaths(ctx context.Context, paths []string) ([]image.Image, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type PreprocessStep interface {
	Apply(img image.Image) (image.Image, error)
}

type ResizePreprocessor struct {
	targetSize int
}

func ResizeStep(targetSize int) *ResizePreprocessor { _ = "STUB: not implemented"; return nil }

func (s *ResizePreprocessor) Apply(img image.Image) (image.Image, error) {
	_ = "STUB: not implemented"
	return *new(image.Image), nil
}

func CenterCropStep(targetWidth, targetHeight int) *CenterCropPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

type CenterCropPreprocessor struct {
	targetWidth  int
	targetHeight int
}

func (s *CenterCropPreprocessor) Apply(img image.Image) (image.Image, error) {
	_ = "STUB: not implemented"
	return *new(image.Image), nil
}

type NormalizationStep interface {
	Apply(r, g, b float32) (float32, float32, float32)
}

type PixelNormalizationPreprocessor struct {
	mean [3]float32
	std  [3]float32
}

func (s *PixelNormalizationPreprocessor) Apply(r, g, b float32) (float32, float32, float32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func PixelNormalizationStep(mean, std [3]float32) *PixelNormalizationPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

func ImagenetPixelNormalizationStep() *PixelNormalizationPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

// CLIPPixelNormalizationStep returns CLIP's normalization values.
// Use after RescaleStep() to normalize to 0-1 range first.
func CLIPPixelNormalizationStep() *PixelNormalizationPreprocessor {
	_ = "STUB: not implemented"
	return nil
}

type RescalePreprocessor struct{}

func (s *RescalePreprocessor) Apply(r, g, b float32) (float32, float32, float32) {
	_ = "STUB: not implemented"
	return 0, 0, 0
}

func RescaleStep() *RescalePreprocessor { _ = "STUB: not implemented"; return nil }

// resizeImage resizes an image to the given width and height using nearest neighbor (simple, replace with better if needed).
func resizeImage(img image.Image, newW, newH int) image.Image {
	_ = "STUB: not implemented"
	return *new(image.Image)
}
