package dense

type allCursor struct {
	matrix *matrixImpl
	index  int
}

func newAllCursor(matrix *matrixImpl) *allCursor {
	c := &allCursor{
		matrix: matrix,
		index:  0,
	}

	return c
}

func (c *allCursor) HasNext() bool {
	return c.index < len(c.matrix.elements)
}

func (c *allCursor) Get() (element float64, row, column int) {
	columns := c.matrix.Columns()

	element = c.matrix.elements[c.index]
	row = c.index / columns
	column = c.index % columns

	c.index++

	return element, row, column
}
