package dense

import (
	"fmt"
	"testing"
)

type constructTest struct {
	rows     int
	columns  int
	elements []float64
}

func createErroneousConstructTestSeq() []*constructTest {
	testSeq := []*constructTest{
		&constructTest{rows: 2, columns: 2, elements: []float64{0, 1, 2}},
		&constructTest{rows: 1, columns: 2, elements: []float64{0, 1, 2}},
		&constructTest{rows: 2, columns: 1, elements: []float64{0}},
		&constructTest{rows: 3, columns: 1, elements: []float64{0, 1, 2, 3}},
		&constructTest{rows: 1, columns: 3, elements: []float64{0}},
		&constructTest{rows: 3, columns: 2, elements: []float64{0, 1, 2}},
		&constructTest{rows: 2, columns: 3, elements: []float64{0, 1, 2, 3}},
	}

	return testSeq
}

func createShapeTestSeq() []*constructTest {
	testSeq := []*constructTest{
		&constructTest{rows: 2, columns: 2, elements: []float64{0, 1, 2, 3}},
		&constructTest{rows: 1, columns: 2, elements: []float64{0, 1}},
		&constructTest{rows: 2, columns: 1, elements: []float64{0, 1}},
		&constructTest{rows: 3, columns: 1, elements: []float64{0, 1, 2}},
		&constructTest{rows: 1, columns: 3, elements: []float64{0, 1, 2}},
		&constructTest{rows: 3, columns: 2, elements: []float64{0, 1, 2, 3, 4, 5}},
		&constructTest{rows: 2, columns: 3, elements: []float64{0, 1, 2, 3, 4, 5}},
	}

	return testSeq
}

func TestNewFailedWithTheWrongNumberOfElements(t *testing.T) {
	testSeq := createErroneousConstructTestSeq()

	for _, test := range testSeq {
		_, err := New(test.rows, test.columns)(test.elements...)
		if err == nil {
			template := "The number of %q doesn't equal to %q * %q, but an error caused."
			t.Fatalf(template, "elements", "rows", "columns")
		}
	}
}

func TestRowsSucceedsAlways(t *testing.T) {
	testSeq := createShapeTestSeq()

	for _, test := range testSeq {
		m, err := New(test.rows, test.columns)(test.elements...)
		if err != nil {
			template := "The number of %q equals to %q * %q, but an error caused."
			t.Fatalf(template, "elements", "rows", "columns")
		}

		if rows := m.Rows(); rows != test.rows {
			template := "The \"rows\" should be %d, but is %d."
			t.Fatalf(template, test.rows, rows)
		}
	}
}

func TestColumnsSucceedsAlways(t *testing.T) {
	testSeq := createShapeTestSeq()

	for _, test := range testSeq {
		m, err := New(test.rows, test.columns)(test.elements...)
		if err != nil {
			template := "The number of %q equals to %q * %q, but an error caused."
			t.Fatalf(template, "elements", "rows", "columns")
		}

		if columns := m.Columns(); columns != test.columns {
			template := "The \"columns\" should be %d, but is %d."
			t.Fatalf(template, test.columns, columns)
		}
	}
}

type elementTest struct {
	row     int
	column  int
	element float64
}

func TestUpdateAndGetSucceed(t *testing.T) {
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
			template := "\"m.Get(%d, %d)\" should return 0, but returns %v."
			t.Fatalf(template, test.row, test.column, element)
		}

		m.Update(test.row, test.column, test.element)

		if element := m.Get(test.row, test.column); element != test.element {
			template := "\"m.Get(%d, %d)\" should return %v, but returns %v."
			t.Fatalf(template, test.row, test.column, test.element, element)
		}
	}
}

func TestGetFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		message := fmt.Sprintf("%d should be smaller than %q %d.", rows, "rows", rows)
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	m.Get(rows, 0)
}

func TestGetFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		message := fmt.Sprintf("%q should be a natural number.", "row")
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	m.Get(-1, 0)
}

func TestGetFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		message := fmt.Sprintf("%d should be smaller than %q %d.", columns, "columns", columns)
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	m.Get(0, columns)
}

func TestGetFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		message := fmt.Sprintf("%q should be a natural number.", "column")
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	m.Get(0, -1)
}

func TestUpdateFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		message := fmt.Sprintf("%d should be smaller than %q %d.", rows, "rows", rows)
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	m.Update(rows, 0, 0)
}

func TestUpdateFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		message := fmt.Sprintf("%q should be a natural number.", "row")
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	m.Update(-1, 0, 0)
}

func TestUpdateFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		message := fmt.Sprintf("%d should be smaller than %q %d.", columns, "columns", columns)
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	m.Update(0, columns, 0)
}

func TestUpdateFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)

	defer func() {
		message := fmt.Sprintf("%q should be a natural number.", "column")
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	m.Update(0, -1, 0)
}
