//go:build !cgo || (!XLA && !ALL)

package hugot

import (
	"context"

	"github.com/knights-analytics/hugot/options"
)

func NewXLASession(_ context.Context, _ ...options.WithOption) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func xlaDisableAutoInstall() { _ = "STUB: not implemented"; return }
