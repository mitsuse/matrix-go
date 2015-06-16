/*
Package "dense" provides an implementation of "Matrix" which stores elements in a slide.
*/
package dense

import (
	"errors"
	"fmt"

	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/mutable"
)

type matrixImpl struct {
	rows     int
	columns  int
	elements []float64
}

func New(rows, columns int) func(elements ...float64) (mutable.Matrix, error) {
	rowsShouldBePositiveNumber(rows)
	columnShouldBePositiveNumber(rows)

	constructor := func(elements ...float64) (mutable.Matrix, error) {
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

func Zeros(rows, columns int) mutable.Matrix {
	m, _ := New(rows, columns)(make([]float64, rows*columns)...)
	return m
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

func (m *matrixImpl) Get(row, column int) (element float64) {
	rows, columns := m.Shape()

	rowShouldBeInRows(row, rows)
	columnShouldBeInColumns(column, columns)

	return m.elements[row*columns+column]
}

func (m *matrixImpl) Update(row, column int, element float64) mutable.Matrix {
	rows, columns := m.Shape()

	rowShouldBeInRows(row, rows)
	columnShouldBeInColumns(column, columns)

	m.elements[row*columns+column] = element

	return m
}
