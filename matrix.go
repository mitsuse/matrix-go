/*
Package "matrix" is an experimental library for matrix manipulation implemented in Golang.
*/
package matrix

import (
	"github.com/mitsuse/matrix-go/internal/types"
)

/*
"Matrix" is the interface for implementations of various matrix types.
For more details, refer "types.Matrix".
*/
type Matrix interface {
	types.Matrix
}

/*
"Cursor" is the interface for iterator for elements of matrix.
Some implementations of "Cursor" iterate all elements,
and others iterates elements satisfying conditions.
For more details, refer "types.Cursor".
*/
type Cursor interface {
	types.Cursor
}
