package dense

import (
	"testing"

	"github.com/mitsuse/matrix-go/validates"
)

type constructTest struct {
	rows     int
	columns  int
	elements []float64
}

func TestNewCreatesDenseMatrix(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	_, err := New(test.rows, test.columns)(test.elements...)
	if err != nil {
		t.Error(
			"The number of \"elements\" equals to \"rows\" * \"columns\",",
			"but matrix creation failed.",
		)
		t.Errorf(
			"# elements = %v, rows = %v, columns = %v",
			test.elements,
			test.rows,
			test.columns,
		)
		t.FailNow()
	}
}

func TestNewFailsForWrongNumberOfElements(t *testing.T) {
	testSeq := []*constructTest{
		&constructTest{
			rows:     3,
			columns:  1,
			elements: []float64{0, 1, 2, 3},
		},
		&constructTest{
			rows:     1,
			columns:  3,
			elements: []float64{0},
		},
	}

	for _, test := range testSeq {
		_, err := New(test.rows, test.columns)(test.elements...)
		if err == nil {
			t.Error("The number of \"elements\" should equal to \"rows\" * \"columns\".")
			t.Errorf(
				"# elements = %v, rows = %v, columns = %v",
				test.elements,
				test.rows,
				test.columns,
			)
			t.FailNow()
		}
	}
}

func TestNewFailsForNonPositiveRows(t *testing.T) {
	test := &constructTest{
		rows:     -3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	defer func() {
		if p := recover(); p == nil || p != validates.NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should make the goroutine panic.",
			)
			t.Errorf(
				"# elements = %v, rows = %v, columns = %v",
				test.elements,
				test.rows,
				test.columns,
			)
			t.FailNow()
		}
	}()
	New(test.rows, test.columns)(test.elements...)
}

func TestNewFailsForNonPositiveColumns(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  -2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	defer func() {
		if p := recover(); p == nil || p != validates.NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should make the goroutine panic.",
			)
			t.Errorf(
				"# elements = %v, rows = %v, columns = %v",
				test.elements,
				test.rows,
				test.columns,
			)
			t.FailNow()
		}
	}()
	New(test.rows, test.columns)(test.elements...)
}
func TestNewFailsForNonPositive(t *testing.T) {
	test := &constructTest{
		rows:     -3,
		columns:  -2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	defer func() {
		if p := recover(); p == nil || p != validates.NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should make the goroutine panic.",
			)
			t.Errorf(
				"# elements = %v, rows = %v, columns = %v",
				test.elements,
				test.rows,
				test.columns,
			)
			t.FailNow()
		}
	}()
	New(test.rows, test.columns)(test.elements...)
}

func TestRowsReturnsTheNumberOfRows(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	m, _ := New(test.rows, test.columns)(test.elements...)
	if rows := m.Rows(); rows != test.rows {
		t.Fatalf("The \"rows\" should be %d, but is %d.", test.rows, rows)
	}
}

func TestColumnsReturnsTheNumberOfColumns(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	m, _ := New(test.rows, test.columns)(test.elements...)
	if columns := m.Columns(); columns != test.columns {
		t.Fatalf("The \"columns\" should be %d, but is %d.", test.columns, columns)
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
	m := Zeros(rows, columns)

	for _, test := range testSeq {
		if element := m.Get(test.row, test.column); element != 0 {
			t.Fatalf(
				"The element at (%d, %d) should be 0 before updating, but is %v.",
				test.row,
				test.column,
				test.element,
			)
		}

		m.Update(test.row, test.column, test.element)

		if element := m.Get(test.row, test.column); element != test.element {
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
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	m.Get(rows, 0)
}

func TestGetFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	m.Get(-1, 0)
}

func TestGetFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	m.Get(0, columns)
}

func TestGetFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	m.Get(0, -1)
}

func TestUpdateFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	m.Update(rows, 0, 0)
}

func TestUpdateFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	m.Update(-1, 0, 0)
}

func TestUpdateFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	m.Update(0, columns, 0)
}

func TestUpdateFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	m.Update(0, -1, 0)
}

func TestAllCreatesCursorToIterateAllElements(t *testing.T) {
	m, _ := New(2, 3)(
		0, 1, 2,
		3, 4, 5,
	)

	checkTable := [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}

	cursor := m.All()

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
	m, _ := New(2, 3)(
		0, 1, 2,
		0, 0, 3,
	)

	checkTable := [][]bool{
		[]bool{true, false, false},
		[]bool{true, true, false},
	}

	cursor := m.NonZeros()

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
	m, _ := New(3, 3)(
		1, 0, 0,
		0, 2, 0,
		0, 0, 3,
	)

	checkTable := [][]bool{
		[]bool{false, true, true},
		[]bool{true, false, true},
		[]bool{true, true, false},
	}

	cursor := m.Diagonal()

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
