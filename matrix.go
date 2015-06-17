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

// Check whether "m" is zero matrix or not.
func IsZeros(m Matrix) bool {
	return !m.NonZeros().HasNext()
}

// Check whether "m" is square matrix or not.
func IsSquare(m Matrix) bool {
	return m.Rows() == m.Columns()
}

// Check whether "m" is special diagonal matrix satisfying arbitary condition.
func IsSpecialDiagonal(m Matrix, match ElementMatcher) bool {
	if !IsSquare(m) {
		return false
	}

	elements := m.All()

	for elements.HasNext() {
		element, row, column := elements.Get()
		if !match(element, row, column) {
			return false
		}
	}

	return true
}

// Check whether "m" is diagonal matrix or not.
func IsDiagonal(m Matrix) bool {
	match := func(element float64, row, column int) bool {
		return row != column && element != 0
	}

	return IsSpecialDiagonal(m, match)
}

// Check whether "m" is identity matrix or not.
func IsIdentity(m Matrix) bool {
	match := func(element float64, row, column int) bool {
		if row == column {
			return element == 1
		} else {
			return element == 0
		}
	}

	return IsSpecialDiagonal(m, match)
}

// Check whether "m" is scalar matrix or not.
func IsScalar(m Matrix) bool {
	if rows, columns := m.Shape(); rows == 0 && columns == 0 {
		return true
	}

	scalar := m.Get(0, 0)

	match := func(element float64, row, column int) bool {
		if row == column {
			return element == scalar
		} else {
			return element == 0
		}
	}

	return IsSpecialDiagonal(m, match)
}
