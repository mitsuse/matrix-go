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

func TestNewFailsForNonPositiveRowsOrColumns(t *testing.T) {
	testSeq := []*constructTest{
		&constructTest{
			rows:     -3,
			columns:  2,
			elements: []float64{0, 1, 2, 3, 4, 5},
		},
		&constructTest{
			rows:     3,
			columns:  -2,
			elements: []float64{0, 1, 2, 3, 4, 5},
		},
		&constructTest{
			rows:     -3,
			columns:  -2,
			elements: []float64{0, 1, 2, 3, 4, 5},
		},
	}

	for _, test := range testSeq {
		func() {
			defer func(test *constructTest) {
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
			}(test)
			New(test.rows, test.columns)(test.elements...)
		}()
	}
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
