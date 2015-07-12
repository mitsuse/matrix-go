/*
Package "transpose" provides a function to create the transpose view of a matrix.
*/
package transpose

import (
	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/elements"
	"github.com/mitsuse/matrix-go/validates"
)

type transposeMatrix struct {
	m matrix.Matrix
}

func New(m matrix.Matrix) matrix.Matrix {
	if t, transposed := m.(*transposeMatrix); transposed {
		return t.m
	}

	t := &transposeMatrix{
		m: m,
	}

	return t
}

func (t *transposeMatrix) Shape() (rows, columns int) {
	return t.Rows(), t.Columns()
}

func (t *transposeMatrix) Rows() (rows int) {
	return t.m.Columns()
}

func (t *transposeMatrix) Columns() (columns int) {
	return t.m.Rows()
}

func (t *transposeMatrix) All() elements.Cursor {
	return newTransposeCursor(t.m.All())
}

func (t *transposeMatrix) NonZeros() elements.Cursor {
	return newTransposeCursor(t.m.NonZeros())
}

func (t *transposeMatrix) Diagonal() elements.Cursor {
	return newTransposeCursor(t.m.Diagonal())
}

func (t *transposeMatrix) Get(row, column int) (element float64) {
	rows, columns := t.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return t.m.Get(column, row)
}

func (t *transposeMatrix) Update(row, column int, element float64) matrix.Matrix {
	rows, columns := t.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return t.m.Update(column, row, element)
}

func (m *transposeMatrix) Equal(n matrix.Matrix) bool {
	// TODO: Implement.
	return false
}

func (m *transposeMatrix) Add(n matrix.Matrix) matrix.Matrix {
	// TODO: Implement.
	return m
}

func (m *transposeMatrix) Subtract(n matrix.Matrix) matrix.Matrix {
	// TODO: Implement.
	return m
}
