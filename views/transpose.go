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
	if v, transposed := m.(*transposeView); transposed {
		return v.m
	}

	v := &transposeView{
		m: m,
	}

	return v
}

func (v *transposeView) Shape() (rows, columns int) {
	return v.Rows(), v.Columns()
}

func (v *transposeView) Rows() (rows int) {
	return v.m.Columns()
}

func (v *transposeView) Columns() (columns int) {
	return v.m.Rows()
}

func (v *transposeView) All() elements.Cursor {
	return newTransposeCursor(v.m.All())
}

func (v *transposeView) NonZeros() elements.Cursor {
	return newTransposeCursor(v.m.NonZeros())
}

func (v *transposeView) Diagonal() elements.Cursor {
	return newTransposeCursor(v.m.Diagonal())
}

func (v *transposeView) Get(row, column int) (element float64) {
	rows, columns := v.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return v.m.Get(column, row)
}

func (v *transposeView) Update(row, column int, element float64) matrix.Matrix {
	rows, columns := v.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return New(v.m.Update(column, row, element))
}

func (v *transposeView) Equal(n matrix.Matrix) bool {
	return v.m.Equal(New(n))
}

func (v *transposeView) Add(m matrix.Matrix) matrix.Matrix {
	validates.ShapeShouldBeSame(v, m)

	cursor := m.NonZeros()

	var tr matrix.Matrix = v

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		tr = tr.Update(row, column, tr.Get(row, column)+element)
	}

	return tr
}

func (v *transposeView) Subtract(m matrix.Matrix) matrix.Matrix {
	validates.ShapeShouldBeSame(v, m)

	cursor := m.NonZeros()

	var tr matrix.Matrix = v

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		tr = tr.Update(row, column, tr.Get(row, column)-element)
	}

	return tr
}

func (v *transposeView) Multiply(m matrix.Matrix) matrix.Matrix {
	// TODO: Avoid to use "dense.Zeros" or implement transpose view for each matrix.
	validates.ShapeShouldBeMultipliable(v, m)

	rows := v.Rows()
	columns := m.Columns()

	r := dense.Zeros(rows, columns)

	cursor := m.NonZeros()

	for cursor.HasNext() {
		element, j, k := cursor.Get()

		for i := 0; i < rows; i++ {
			r.Update(i, k, r.Get(i, k)+v.Get(i, j)*element)
		}
	}

	return r
}

func (v *transposeView) Scalar(s float64) matrix.Matrix {
	return New(v.m.Scalar(s))
}
