/*
Package "transpose" provides a function to create the transpose view of a matrix.
*/
package views

import (
	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/elements"
	"github.com/mitsuse/matrix-go/mutable/dense"
	"github.com/mitsuse/matrix-go/validates"
)

type transposeView struct {
	m matrix.Matrix
}

func New(m matrix.Matrix) matrix.Matrix {
	if t, transposed := m.(*transposeView); transposed {
		return t.m
	}

	t := &transposeView{
		m: m,
	}

	return t
}

func (t *transposeView) Shape() (rows, columns int) {
	return t.Rows(), t.Columns()
}

func (t *transposeView) Rows() (rows int) {
	return t.m.Columns()
}

func (t *transposeView) Columns() (columns int) {
	return t.m.Rows()
}

func (t *transposeView) All() elements.Cursor {
	return newTransposeCursor(t.m.All())
}

func (t *transposeView) NonZeros() elements.Cursor {
	return newTransposeCursor(t.m.NonZeros())
}

func (t *transposeView) Diagonal() elements.Cursor {
	return newTransposeCursor(t.m.Diagonal())
}

func (t *transposeView) Get(row, column int) (element float64) {
	rows, columns := t.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return t.m.Get(column, row)
}

func (t *transposeView) Update(row, column int, element float64) matrix.Matrix {
	rows, columns := t.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return New(t.m.Update(column, row, element))
}

func (m *transposeView) Equal(n matrix.Matrix) bool {
	return m.m.Equal(New(n))
}

func (m *transposeView) Add(n matrix.Matrix) matrix.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	var tr matrix.Matrix = m

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		tr = tr.Update(row, column, tr.Get(row, column)+element)
	}

	return tr
}

func (m *transposeView) Subtract(n matrix.Matrix) matrix.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	var tr matrix.Matrix = m

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		tr = tr.Update(row, column, tr.Get(row, column)-element)
	}

	return tr
}

func (m *transposeView) Multiply(n matrix.Matrix) matrix.Matrix {
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

func (m *transposeView) Scalar(s float64) matrix.Matrix {
	return New(m.m.Scalar(s))
}
