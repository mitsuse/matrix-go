package dense

type nonZerosCursor struct {
	matrix  *matrixImpl
	element float64
	row     int
	column  int
	index   int
}

func newNonZerosCursor(matrix *matrixImpl) *nonZerosCursor {
	c := &nonZerosCursor{
		matrix:  matrix,
		element: 0,
		row:     0,
		column:  0,
		index:   0,
	}

	return c
}

func (c *nonZerosCursor) HasNext() bool {
	for c.index < len(c.matrix.elements) {
		if element := c.matrix.elements[c.index]; element != 0 {
			c.element = element

			columns := c.matrix.Columns()
			c.row = c.index / columns
			c.column = c.index % columns

			c.index++

			return true
		} else {
			c.index++
		}
	}

	return false
}

func (c *nonZerosCursor) Get() (element float64, row, column int) {
	return c.element, c.row, c.column
}
