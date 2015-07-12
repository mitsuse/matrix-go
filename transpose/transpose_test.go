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

func TestAllCreatesCursorToIterateAllElements(t *testing.T) {
	m, _ := dense.New(2, 3)(
		0, 1, 2,
		3, 4, 5,
	)

	tr := New(m)

	checkTable := [][]bool{
		[]bool{false, false},
		[]bool{false, false},
		[]bool{false, false},
	}

	cursor := tr.All()

	for cursor.HasNext() {
		element, row, column := cursor.Get()

		if e := m.Get(column, row); element != e {
			t.Fatalf(
				"The element at (%d, %d) should be %v, but the cursor say it is %v.",
				row,
				column,
				e,
				element,
			)
		}

		if checked := checkTable[row][column]; checked {
			t.Error("Cursor should visit each element only once, but visits some twice.")
			t.Fatalf(
				"# element = %v, row = %d, column = %d",
				element,
				row,
				column,
			)
		}
		checkTable[row][column] = true
	}

	for row, checkRow := range checkTable {
		for column, checked := range checkRow {
			if !checked {
				t.Error(
					"Cursor should visit each element only once, but never visits some.",
				)
				t.Fatalf(
					"# row = %d, column = %d",
					row,
					column,
				)
			}
		}
	}
}

func TestNonZerosCreatesCursorToIterateNonZeroElements(t *testing.T) {
	m, _ := dense.New(2, 3)(
		0, 1, 2,
		0, 0, 3,
	)

	tr := New(m)

	checkTable := [][]bool{
		[]bool{true, true},
		[]bool{false, true},
		[]bool{false, false},
	}

	cursor := tr.NonZeros()

	for cursor.HasNext() {
		element, row, column := cursor.Get()

		if e := m.Get(column, row); element != e {
			t.Fatalf(
				"The element at (%d, %d) should be %v, but the cursor say it is %v.",
				row,
				column,
				e,
				element,
			)
		}

		if checked := checkTable[row][column]; checked {
			t.Error(
				"Cursor should visit each non-zero element only once,",
				"but visits some twice or zero-element.",
			)
			t.Fatalf(
				"# element = %v, row = %d, column = %d",
				element,
				row,
				column,
			)
		}
		checkTable[row][column] = true
	}

	for row, checkRow := range checkTable {
		for column, checked := range checkRow {
			if !checked {
				t.Error(
					"Cursor should visit each non-zero element only once,",
					"but never visits some.",
				)
				t.Fatalf(
					"# row = %d, column = %d",
					row,
					column,
				)
			}
		}
	}
}

func TestDiagonalCreatesCursorToIterateDiagonalElements(t *testing.T) {
	m, _ := dense.New(3, 3)(
		1, 0, 0,
		0, 2, 0,
		0, 0, 3,
	)

	tr := New(m)

	checkTable := [][]bool{
		[]bool{false, true, true},
		[]bool{true, false, true},
		[]bool{true, true, false},
	}

	cursor := tr.Diagonal()

	for cursor.HasNext() {
		element, row, column := cursor.Get()

		if e := m.Get(row, column); element != e {
			t.Fatalf(
				"The element at (%d, %d) should be %v, but the cursor say it is %v.",
				row,
				column,
				e,
				element,
			)
		}

		if checked := checkTable[row][column]; checked {
			t.Error(
				"Cursor should visit each diagonal element only once,",
				"but visits some twice or non-diagonal element.",
			)
			t.Fatalf(
				"# element = %v, row = %d, column = %d",
				element,
				row,
				column,
			)
		}
		checkTable[row][column] = true
	}

	for row, checkRow := range checkTable {
		for column, checked := range checkRow {
			if !checked {
				t.Error(
					"Cursor should visit each diagonal element only once,",
					"but never visits some.",
				)
				t.Fatalf(
					"# row = %d, column = %d",
					row,
					column,
				)
			}
		}
	}
}

func TestEqualIsTrue(t *testing.T) {
	m, _ := dense.New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	tr := New(m)

	n, _ := dense.New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if equality := tr.Equal(n); !equality {
		t.Fatal("Two matrices should equal, but the result is false.")
	}
}

func TestEqualIsFalse(t *testing.T) {
	m, _ := dense.New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	tr := New(m)

	n, _ := dense.New(4, 3)(
		0, 1, 2,
		3, 1, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if equality := tr.Equal(n); equality {
		t.Fatal("Two matrices should not equal, but the result is true.")
	}
}

func TestEqualCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m, _ := dense.New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	tr := New(m)

	n, _ := dense.New(3, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	)

	defer func() {
		if r := recover(); r != nil && r != validates.DIFFERENT_SIZE_PANIC {
			t.Error(
				"Checking equality of matrices which have different shape",
				"should cause panic, but cause nothing.",
			)
			t.Fatalf(
				"# tr.rows = %d, tr.columns = %d, tr.rows = %d, tr.columns = %d",
				tr.Rows(),
				tr.Columns(),
				n.Rows(),
				n.Columns(),
			)
		}
	}()
	m.Equal(n)
}
