/*
Package "transpose" provides a function to create the transpose view of a matrix.
*/
package transpose

import (
	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/elements"
	"github.com/mitsuse/matrix-go/mutable/dense"
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

	return New(t.m.Update(column, row, element))
}

func (m *transposeMatrix) Equal(n matrix.Matrix) bool {
	return m.m.Equal(New(n))
}

func (m *transposeMatrix) Add(n matrix.Matrix) matrix.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	var tr matrix.Matrix = m

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		tr = tr.Update(row, column, tr.Get(row, column)+element)
	}

	return tr
}

func (m *transposeMatrix) Subtract(n matrix.Matrix) matrix.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	var tr matrix.Matrix = m

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		tr = tr.Update(row, column, tr.Get(row, column)-element)
	}

	return tr
}

func (m *transposeMatrix) Multiply(n matrix.Matrix) matrix.Matrix {
	// TODO: Avoid to use "dense.Zeros" or implement transpose view for each matrix.
	validates.ShapeShouldBeMultipliable(m, n)

	rows := m.Rows()
	columns := n.Columns()

	r := dense.Zeros(rows, columns)

	cursor := n.NonZeros()

	for cursor.HasNext() {
		element, j, k := cursor.Get()

		for i := 0; i < rows; i++ {
			r.Update(i, k, r.Get(i, k)+m.Get(i, j)*element)
		}
	}

	return r
}

func (m *transposeMatrix) Scalar(s float64) matrix.Matrix {
	return New(m.m.Scalar(s))
}
