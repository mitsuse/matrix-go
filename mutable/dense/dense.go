/*
Package "dense" provides an implementation of "Matrix" which stores elements in a slide.
*/
package dense

import (
	"errors"
	"fmt"

	"github.com/mitsuse/matrix-go"
)

type matrixImpl struct {
	rows     int
	columns  int
	elements []float64
}

func New(rows, columns int) func(elements ...float64) (matrix.Matrix, error) {
	rowsShouldBePositiveNumber(rows)
	columnShouldBePositiveNumber(rows)

	constructor := func(elements ...float64) (matrix.Matrix, error) {
		size := rows * columns

		if len(elements) != size {
			template := "The number of %q should equal to %q * %q."
			message := fmt.Sprintf(template, "elements", "rows", "columns")

			return nil, errors.New(message)
		}

		m := &matrixImpl{
			rows:     rows,
			columns:  columns,
			elements: make([]float64, size),
		}
		copy(m.elements, elements)

		return m, nil
	}

	return constructor
}

func (m *matrixImpl) Shape() (rows, columns int) {
	return m.Rows(), m.Columns()
}

func (m *matrixImpl) Rows() (rows int) {
	return m.rows
}

func (m *matrixImpl) Columns() (columns int) {
	return m.columns
}

func (m *matrixImpl) Row(row int) matrix.Row {
	// TODO: Implement.
	return nil
}

func (m *matrixImpl) Column(column int) matrix.Column {
	// TODO: Implement.
	return nil
}