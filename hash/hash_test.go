package hash

import (
	"testing"

	"github.com/mitsuse/matrix-go/internal/validates"
)

type elementTest struct {
	row     int
	column  int
	element float64
}

func TestAllCreatesCursorToIterateAllElements(t *testing.T) {
	m := New(2, 3)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 2, Value: 5},
	).View(1, 1, 1, 2)

	checkTable := [][]bool{
		[]bool{false, false},
	}

	cursor := m.All()
	_ = checkTable

	for cursor.HasNext() {
		element, row, column := cursor.Get()

		if e := m.Get(row, column); element != e {
			t.Fatalf(
				"The element at (%d, %d) should be %v, but the cursor returns %v.",
				row,
				column,
				e,
				element,
			)
		}

		if checked := checkTable[row][column]; checked {
			t.Fatalf("Cursor should visit (%d, %d) more than necessary.", row, column)
		}
		checkTable[row][column] = true
	}

	for row, checkRow := range checkTable {
		for column, checked := range checkRow {
			if checked {
				continue
			}

			t.Fatalf("Cursor didn't visit (%d, %d).", row, column)
		}
	}
}

func TestGetFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 6
	viewRows, viewColumns := 4, 3
	offsetRow, offsetColumn := 2, 3

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The rows exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Get(viewRows, 0)
}

func TestGetFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 6
	viewRows, viewColumns := 4, 3
	offsetRow, offsetColumn := 2, 3

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The rows is negative, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Get(-1, 0)
}

func TestGetFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 6, 8
	viewRows, viewColumns := 3, 4
	offsetRow, offsetColumn := 3, 2

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The columns exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Get(0, viewColumns)
}

func TestGetFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 6, 8
	viewRows, viewColumns := 3, 4
	offsetRow, offsetColumn := 3, 2

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The columns is negative, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Get(0, -1)
}

func TestUpdateFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 6
	viewRows, viewColumns := 4, 3
	offsetRow, offsetColumn := 2, 3

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The rows exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Update(viewRows, 0, 0)
}

func TestUpdateFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 6
	viewRows, viewColumns := 4, 3
	offsetRow, offsetColumn := 2, 3

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The rows is negative, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Update(-1, 0, 0)
}

func TestUpdateFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 6, 8
	viewRows, viewColumns := 3, 4
	offsetRow, offsetColumn := 3, 2

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The columns exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Update(0, viewColumns, 0)
}

func TestUpdateFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 6, 8
	viewRows, viewColumns := 3, 4
	offsetRow, offsetColumn := 3, 2

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The columns is negative, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Update(0, -1, 0)
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

	rows, columns := 12, 11
	viewRows, viewColumns := 8, 8
	offsetRow, offsetColumn := 1, 2

	m := Zeros(rows, columns).View(offsetRow, offsetColumn, viewRows, viewColumns)

	for _, test := range testSeq {
		if element := m.Get(test.row, test.column); element != 0 {
			t.Fatalf(
				"The element at (%d, %d) should be 0 before updating, but is %v.",
				test.row,
				test.column,
				element,
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
