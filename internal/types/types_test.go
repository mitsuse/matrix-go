package types

import (
	"testing"
)

func TestShapeRetainsRowsAndColumns(t *testing.T) {
	rows, columns := 2, 1

	s := NewShape(rows, columns)

	if s.Rows() == rows && s.Columns() == columns {
		return
	}

	t.Fatal("Shape should retains the rows and columns equivalent to the given ones.")
}

func TestIndexRetainsRowAndColumn(t *testing.T) {
	row, column := 1, 0

	i := NewIndex(row, column)

	if i.Row() == row && i.Column() == column {
		return
	}

	t.Fatal("Index should retains the row and column equivalent to the given ones.")
}
