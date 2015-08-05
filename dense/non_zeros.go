package dense

import (
	"github.com/mitsuse/matrix-go/internal/types"
)

type nonZerosCursor struct {
	cursor types.Cursor
}

func newNonZerosCursor(matrix *denseMatrix) *nonZerosCursor {
	c := &nonZerosCursor{
		cursor: matrix.All(),
	}

	return c
}

func (c *nonZerosCursor) HasNext() bool {
	for c.cursor.HasNext() {
		if element, _, _ := c.cursor.Get(); element != 0 {
			return true
		}
	}

	return false
}

func (c *nonZerosCursor) Get() (element float64, row, column int) {
	return c.cursor.Get()
}
