//go:build cgo && (XLA || ALL)

package hugot

import (
	"context"

	// import XLA backend

	"github.com/knights-analytics/hugot/options"
)

func NewXLASession(ctx context.Context, opts ...options.WithOption) (*Session, error) {
	_ = "STUB: not implemented"
	// Disabled for now until we have auto installs globally
	return nil, nil
}

func xlaDisableAutoInstall() { _ = "STUB: not implemented"; return }
