package dense

import (
	"github.com/mitsuse/matrix-go/internal/types"
)

type diagonalCursor struct {
	matrix  *denseMatrix
	element float64
	current types.Index
	next    types.Index
}

func newDiagonalCursor(matrix *denseMatrix) *diagonalCursor {
	c := &diagonalCursor{
		matrix:  matrix,
		element: 0,
		current: types.NewIndex(0, 0),
		next:    types.NewIndex(0, 0),
	}

	return c
}

func (c *diagonalCursor) HasNext() bool {
	c.current = c.next

	if c.current.Row() >= c.matrix.view.Rows() || c.current.Column() >= c.matrix.view.Columns() {
		return false
	}

	index := c.matrix.base.Columns()*(c.matrix.offset.Row()+c.current.Row()) + c.matrix.offset.Column() + c.current.Column()
	c.element = c.matrix.elements[index]

	c.next = types.NewIndex(c.current.Row()+1, c.current.Column()+1)

	return true
}

func (c *diagonalCursor) Get() (element float64, row, column int) {
	row, column = c.matrix.rewriter.Rewrite(c.current.Row(), c.current.Column())
	return c.element, row, column
}
