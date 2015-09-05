package hash

import (
	"github.com/mitsuse/matrix-go/internal/types"
)

type nonZerosCursor struct {
	matrix  *Matrix
	element float64
	current *types.Index
	next    *types.Index
}

func newNonZerosCursor(matrix *Matrix) *nonZerosCursor {
	c := &nonZerosCursor{
		matrix:  matrix,
		element: 0,
		current: types.NewIndex(0, 0),
		next:    types.NewIndex(0, 0),
	}

	return c
}

func (c *nonZerosCursor) HasNext() bool {
	for {
		c.current = c.next

		if c.current.Row() >= c.matrix.view.Rows() || c.current.Column() >= c.matrix.view.Columns() {
			return false
		}

		index := c.matrix.base.Columns()*(c.matrix.offset.Row()+c.current.Row()) + c.matrix.offset.Column() + c.current.Column()
		c.next = types.NewIndex(c.current.Row()+1, c.current.Column())

		if element, exist := c.matrix.elements[index]; exist {
			c.element = element
			return true
		}

		if c.next.Row() >= c.matrix.view.Rows() {
			c.next = types.NewIndex(0, c.current.Column()+1)
		}
	}

	return false
}

func (c *nonZerosCursor) Get() (element float64, row, column int) {
	row, column = c.matrix.rewriter.Rewrite(c.current.Row(), c.current.Column())
	return c.element, row, column
}
