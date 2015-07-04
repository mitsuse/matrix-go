package matrix

import (
	"github.com/mitsuse/matrix-go/elements"
	"github.com/mitsuse/matrix-go/validates"
)

type transposeMatrix struct {
	m Matrix
}

func Transpose(m Matrix) Matrix {
	t, transposed := m.(*transposeMatrix)
	if transposed {
		return t.m
	}

	t = &transposeMatrix{
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

func (t *transposeMatrix) All() elements.Curor {
	// TODO: Implement this.
	return nil
}

func (t *transposeMatrix) NonZeros() elements.Curor {
	// TODO: Implement this.
	return nil
}

func (t *transposeMatrix) Diagonal() elements.Curor {
	// TODO: Implement this.
	return nil
}

func (t *transposeMatrix) Get(row, column int) (element float64) {
	rows, columns := t.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return t.m.Get(column, row)
}

func (t *transposeMatrix) Update(row, column int, element float64) Matrix {
	rows, columns := t.Shape()

	validates.IndexShouldBeInRange(rows, columns, row, column)

	return t.m.Update(column, row, element)
}

func (t *transposeMatrix) Transpose() Matrix {
	return Transpose(t)
}
