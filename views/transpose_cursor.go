package views

import (
	"github.com/mitsuse/matrix-go/elements"
)

type transposeCursor struct {
	cursor elements.Cursor
}

func newTransposeCursor(cursor elements.Cursor) *transposeCursor {
	c := &transposeCursor{
		cursor: cursor,
	}

	return c
}

func (c *transposeCursor) HasNext() bool {
	return c.cursor.HasNext()
}

func (c *transposeCursor) Get() (element float64, row, column int) {
	element, column, row = c.cursor.Get()
	return element, row, column
}
