package types

import (
	"io"
)

type Matrix interface {
	// Serialize the receiver matrix by using the given writer.
	Serialize(wrtier io.Writer) error

	// Return the shape of matrix, which consists of the "rows" and the "columns".
	Shape() (rows, columns int)

	// Return the number of "rows".
	Rows() (rows int)

	// Return the number of "columns".
	Columns() (columns int)

	// Create and return an iterator for all elements.
	All() Cursor

	// Create and return an iterator for non-zero elements.
	NonZeros() Cursor

	// Create and return an iterator for diagonal elements.
	Diagonal() Cursor

	// Get an element of matrix specified with "row" and "column".
	// When "row" or "column" is lower than the number of rows or columns,
	// validates.OUT_OF_RANGE_PANIC will be caused.
	Get(row, column int) (element float64)

	// Update the element of matrix specified with "row" and "column".
	// When "row" or "column" is lower than the number of rows or columns,
	// validates.OUT_OF_RANGE_PANIC will be caused.
	Update(row, column int, element float64) Matrix

	// Check element-wise equality of the receiver matrix and the given matrix.
	// When the shape of the receiver and the argument is different,
	// validates.DIFFERENT_SIZE_PANIC will be caused.
	Equal(n Matrix) bool

	// Add the given matrix to the receiver matrix.
	// When the shape of the receiver and the argument is different,
	// validates.DIFFERENT_SIZE_PANIC will be caused.
	Add(n Matrix) Matrix

	// Subtract the given matrix from the receiver matrix.
	// When the shape of the receiver and the argument is different,
	// validates.DIFFERENT_SIZE_PANIC will be caused.
	Subtract(n Matrix) Matrix

	// Multiply the receiver matrix by the given matrix.
	// When the number of columns of the receiver doesn't equal to
	// the number of rows of the argument,
	// validates.NOT_MULTIPLIABLE_PANIC will be caused.
	Multiply(n Matrix) Matrix

	// Multiply by scalar value.
	Scalar(s float64) Matrix

	// Create the transpose matrix.
	Transpose() Matrix

	// Create a arbitrary view.
	View(row, column, rows, columns int) Matrix

	// Create a row view.
	Row(row int) Matrix

	// Create a column view.
	Column(column int) Matrix

	// Find and return the first one of maximum elements.
	Max() (element float64, row, column int)

	// Find and return the first one of minimum elements.
	Min() (element float64, row, column int)
}
