/*
Package "matrix" is an experimental library for matrix manipulation implemented in Golang.
*/
package matrix

import (
	"github.com/mitsuse/matrix-go/internal/types"
)

type Matrix interface {
	types.Matrix
}

type Cursor interface {
	types.Cursor
}
