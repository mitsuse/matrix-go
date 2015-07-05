package transpose

import (
	"testing"

	"github.com/mitsuse/matrix-go/mutable/dense"
	"github.com/mitsuse/matrix-go/validates"
)

func TestNewTwiceReturnsTheOriginalMatrix(t *testing.T) {
	rows, columns := 4, 3

	m := dense.Zeros(rows, columns)
	tr := New(m)
	n := New(tr)

	mRow, mColumn := m.Shape()
	trRow, trColumn := tr.Shape()
	nRow, nColumn := n.Shape()

	if mRow == trRow || mColumn == trColumn {
		t.Fatalf(
			"The shape of transpose matrix should be (%d, %d), but is (%d, %d).",
			mColumn, mRow,
			trRow, trColumn,
		)
	}

	if _, isTransposed := tr.(*transposeMatrix); !isTransposed {
		t.Fatal("The type of transpose matrix should be transposed.")
	}

	if mRow != nRow || mColumn != nColumn {
		t.Fatalf(
			"The shape of re-transposed matrix should be (%d, %d), but (%d, %d).",
			mRow, mColumn,
			nRow, nColumn,
		)
	}

	if _, isTransposed := n.(*transposeMatrix); isTransposed {
		t.Fatal("The type of re-transpose matrix should not be dense.")
	}
}

type constructTest struct {
	rows     int
	columns  int
	elements []float64
}

func TestRowsReturnsTheNumberOfRows(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	m, _ := dense.New(test.rows, test.columns)(test.elements...)
	tr := New(m)

	if rows := tr.Rows(); rows != test.columns {
		t.Fatalf("The \"rows\" should be %d, but is %d.", test.columns, rows)
	}
}

func TestColumnsReturnsTheNumberOfColumns(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	m, _ := dense.New(test.rows, test.columns)(test.elements...)
	tr := New(m)

	if columns := tr.Columns(); columns != test.rows {
		t.Fatalf("The \"columns\" should be %d, but is %d.", test.rows, columns)
	}
}

type elementTest struct {
	row     int
	column  int
	element float64
}

func TestUpdateReplacesElement(t *testing.T) {
	testSeq := []*elementTest{
		&elementTest{row: 0, column: 0, element: 1},
		&elementTest{row: 1, column: 0, element: 2},
		&elementTest{row: 0, column: 1, element: 3},
		&elementTest{row: 3, column: 6, element: 4},
		&elementTest{row: 7, column: 5, element: 5},
		&elementTest{row: 5, column: 7, element: 6},
		&elementTest{row: 7, column: 7, element: 7},
	}

	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	for _, test := range testSeq {
		if element := tr.Get(test.row, test.column); element != 0 {
			t.Fatalf(
				"The element at (%d, %d) should be 0 before updating, but is %v.",
				test.row,
				test.column,
				test.element,
			)
		}

		tr.Update(test.row, test.column, test.element)

		if element := tr.Get(test.row, test.column); element != test.element {
			t.Fatalf(
				"The element at (%d, %d) should be %v after updating, but is %v.",
				test.row,
				test.column,
				test.element,
				element,
			)
		}
	}
}

func TestGetFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	tr.Get(columns, 0)
}

func TestGetFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	tr.Get(-1, 0)
}

func TestGetFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	tr.Get(0, rows)
}

func TestGetFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	tr.Get(0, -1)
}

func TestUpdateFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	tr.Update(columns, 0, 0)
}

func TestUpdateFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	tr.Update(-1, 0, 0)
}

func TestUpdateFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	tr.Update(0, rows, 0)
}

func TestUpdateFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := dense.Zeros(rows, columns)
	tr := New(m)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	tr.Update(0, -1, 0)
}
