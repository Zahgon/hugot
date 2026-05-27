//go:build cgo && (ORT || ALL)

package hugot

import (
	"context"

	"github.com/knights-analytics/hugot/options"
)

func NewORTSession(ctx context.Context, opts ...options.WithOption) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// set session options and initialise

func (s *Session) initialiseORT(ctx context.Context) (bool, error) {
	_ = "STUB: not implemented"
	return false,

		// Set pre-initialisation options
		nil
}

// Onnx runtime does not provide a default for Mac, so do that here instead.

// Start OnnxRuntime

// Create session options for use in all pipelines
