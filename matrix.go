/*
Package "matrix" provides types and operations for matrix manipulation.
*/
package matrix

type Matrix interface {
	// Return the shape of matrix, which consists of the "rows" and the "columns".
	Shape() (rows, columns int)

	// Return the "rows" of matrix.
	Rows() (rows int)

	// Return the "columns" of matrix.
	Columns() (columns int)

	// Create and return an iterator for all elements.
	All() Elements

	// Create and return an iterator for non-zero elements.
	NonZeros() Elements

	// Create and return an iterator for diagonal elements.
	Diagonal() Elements

	// Get an element of matrix speficied with "row" and "column".
	Get(row, column int) (element float64)
}

// Check whether a matrix "m" is zero-matrix or not.
func IsZeros(m Matrix) bool {
	return !m.NonZeros().HasNext()
}

// Check whether a matrix "m" is square or not.
func IsSquare(m Matrix) bool {
	return m.Rows() == m.Columns()
}

// Check whether a matrix "m" is diagonal or not.
func IsDiagonal(m Matrix) bool {
	if !IsSquare(m) {
		return false
	}

	elements := m.All()

	for elements.HasNext() {
		element, row, column := elements.Get()
		if row != column && element != 0 {
			return false
		}
	}

	return true
}
