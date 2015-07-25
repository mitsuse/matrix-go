/*
Package "matrix" provides types which represent matrix, scalar and iterator of elements.
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
