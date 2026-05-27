//go:build !cgo || (!ORT && !ALL)

package hugot

import (
	"context"

	"github.com/knights-analytics/hugot/options"
)

func NewORTSession(_ context.Context, _ ...options.WithOption) (*Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
