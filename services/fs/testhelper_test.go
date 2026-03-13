package fs

import (
	"github.com/rgglez/go-storage/v5/types"
)

// newStorager is a test helper that wraps NewStorager for use in package-level tests.
func newStorager(pairs ...types.Pair) (types.Storager, error) {
	return NewStorager(pairs...)
}
