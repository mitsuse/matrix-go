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

	c.element = c.matrix.elements[c.index]
	c.row = c.index / c.matrix.columns
	c.column = c.index % c.matrix.columns

	c.index++

	return true
}

func (c *allCursor) Get() (element float64, row, column int) {
	row, column = c.matrix.rewriter.Rewrite(c.row, c.column)
	return c.element, row, column
}
