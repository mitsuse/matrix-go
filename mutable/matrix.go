package mutable

import (
	"github.com/mitsuse/matrix-go"
)

type Matrix interface {
	matrix.Matrix

	// Update the element of matrix speficied with "row" and "column".
	Update(row, column int, element float64) Matrix
}
