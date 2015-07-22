/*
Package "matrix" provides types and operations for matrix manipulation.
*/
package matrix

import (
	"github.com/mitsuse/matrix-go/elements"
)

type Matrix interface {
	// Return the shape of matrix, which consists of the "rows" and the "columns".
	Shape() (rows, columns int)

	// Return the "rows" of matrix.
	Rows() (rows int)

	// Return the "columns" of matrix.
	Columns() (columns int)

	// Create and return an iterator for all elements.
	All() elements.Cursor

	// Create and return an iterator for non-zero elements.
	NonZeros() elements.Cursor

	// Create and return an iterator for diagonal elements.
	Diagonal() elements.Cursor

	// Get an element of matrix speficied with "row" and "column".
	Get(row, column int) (element float64)

	// Update the element of matrix speficied with "row" and "column".
	Update(row, column int, element float64) Matrix

	// Check element-wise equality of the receiver matrix and the given matrix.
	Equal(n Matrix) bool

	// Add the given matrix to the receiver matrix.
	Add(n Matrix) Matrix

	// Subtract the given matrix from the receiver matrix.
	Subtract(n Matrix) Matrix

	// Multiply the receiver matrix bythe given matrix.
	Multiply(n Matrix) Matrix

	// Multiply by scalar.
	Scalar(s float64) Matrix

	// Create the transpose matrix.
	Transpose() Matrix
}
