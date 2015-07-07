package dense

type allCursor struct {
	matrix  *matrixImpl
	element float64
	row     int
	column  int
	index   int
}

func newAllCursor(matrix *matrixImpl) *allCursor {
	c := &allCursor{
		matrix:  matrix,
		element: 0,
		row:     0,
		column:  0,
		index:   0,
	}

	return c
}

func (c *allCursor) HasNext() bool {
	if c.index >= len(c.matrix.elements) {
		return false
	}

	columns := c.matrix.Columns()

	c.element = c.matrix.elements[c.index]
	c.row = c.index / columns
	c.column = c.index % columns

	c.index++

	return true
}

func (c *allCursor) Get() (element float64, row, column int) {
	return c.element, c.row, c.column
}
