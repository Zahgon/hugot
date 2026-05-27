package safeconv

import (
	"time"
)

// IntSliceToUint32Slice converts a slice of int to uint32 with clamping to avoid overflow/underflow.
func IntSliceToUint32Slice(input []int) []uint32 { _ = "STUB: not implemented"; return nil }

// Uint32SliceToIntSlice converts a slice of uint32 to int with clamping to MaxInt when necessary.
func Uint32SliceToIntSlice(input []uint32) []int { _ = "STUB: not implemented"; return nil }

// extremely unlikely on typical platforms, but keep safe

// DurationToU64 converts a duration to an unsigned nanoseconds counter safely.
// Negative durations are mapped to 0.
func DurationToU64(d time.Duration) uint64 { _ = "STUB: not implemented"; return 0 }

// Conversion from time.Duration (int64) to uint64 is safe here because negatives are handled above.
// #nosec G115

// U64ToDuration converts an unsigned nanoseconds count to time.Duration safely.
// Values larger than MaxInt64 are clamped to time.Duration(math.MaxInt64).
func U64ToDuration(u uint64) time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }
