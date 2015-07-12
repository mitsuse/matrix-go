/*
Package "dense" provides an implementation of "Matrix" which stores elements in a slide.
*/
package dense

import (
	"errors"
	"fmt"

	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/elements"
	"github.com/mitsuse/matrix-go/validates"
)

type matrixImpl struct {
	rows     int
	columns  int
	elements []float64
}

func New(rows, columns int) func(elements ...float64) (matrix.Matrix, error) {
	validates.ShapeShouldBePositive(rows, columns)

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

func Zeros(rows, columns int) matrix.Matrix {
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

func (m *matrixImpl) All() elements.Cursor {
	return newAllCursor(m)
}

func (m *matrixImpl) NonZeros() elements.Cursor {
	return newNonZerosCursor(m)
}

func (m *matrixImpl) Diagonal() elements.Cursor {
	return newDiagonalCursor(m)
}

func (m *matrixImpl) Get(row, column int) (element float64) {
	rows, columns := m.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return m.elements[row*columns+column]
}

func (m *matrixImpl) Update(row, column int, element float64) matrix.Matrix {
	rows, columns := m.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	m.elements[row*columns+column] = element

	return m
}

func (m *matrixImpl) Equal(n matrix.Matrix) bool {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.All()

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		if m.Get(row, column) != element {
			return false
		}
	}

	return true
}

func (m *matrixImpl) Add(n matrix.Matrix) matrix.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		m.Update(row, column, m.Get(row, column)+element)
	}

	return m
}

func (m *matrixImpl) Sub(n matrix.Matrix) matrix.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		m.Update(row, column, m.Get(row, column)-element)
	}

	return m
}
