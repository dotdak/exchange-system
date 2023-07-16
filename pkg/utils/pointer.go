package utils

import (
	"golang.org/x/exp/constraints"
)

func NewPointer[T constraints.Signed | constraints.Unsigned | constraints.Float](v T) *T {
	return &v
}
