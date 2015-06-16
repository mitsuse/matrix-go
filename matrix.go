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

	// Get an element of matrix speficied with "row" and "column".
	Get(row, column int) (element float64)
}

// Check whether a matrix "m" is square or not.
func IsSquare(m Matrix) bool {
	return m.Rows() == m.Columns()
}
