/*
Package "types" provides functions to descriminate the type of given matrix.
*/
package types

import (
	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/elements"
)

// Check whether "m" is zero matrix or not.
func IsZeros(m matrix.Matrix) bool {
	return !m.NonZeros().HasNext()
}

// Check whether "m" is square matrix or not.
func IsSquare(m matrix.Matrix) bool {
	return m.Rows() == m.Columns()
}

// Check whether "m" is special diagonal matrix satisfying arbitary condition.
func IsSpecialDiagonal(m matrix.Matrix, match elements.Match) bool {
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
func IsDiagonal(m matrix.Matrix) bool {
	match := func(element float64, row, column int) bool {
		return row == column || element == 0
	}

	return IsSpecialDiagonal(m, match)
}

// Check whether "m" is identity matrix or not.
func IsIdentity(m matrix.Matrix) bool {
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
func IsScalar(m matrix.Matrix) bool {
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
