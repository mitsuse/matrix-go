package dense

type diagonalCursor struct {
	matrix  *matrixImpl
	element float64
	row     int
	column  int
	index   int
}

func newDiagonalCursor(matrix *matrixImpl) *diagonalCursor {
	c := &diagonalCursor{
		matrix:  matrix,
		element: 0,
		row:     0,
		column:  0,
		index:   0,
	}

	return c
}

func (c *diagonalCursor) HasNext() bool {
	for c.index < len(c.matrix.elements) {
		columns := c.matrix.Columns()

		row := c.index / columns
		column := c.index % columns

		if row == column {
			c.element = c.matrix.elements[c.index]
			c.row = row
			c.column = column

			c.index++

			return true
		} else {
			c.index++
		}
	}

	return false
}

func (c *diagonalCursor) Get() (element float64, row, column int) {
	row, column = c.matrix.rewriter.Rewrite(c.row, c.column)
	return c.element, row, column
}
