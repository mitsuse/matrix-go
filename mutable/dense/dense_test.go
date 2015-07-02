package dense

import (
	"fmt"
	"testing"

	"github.com/mitsuse/matrix-go/validates"
)

type constructTest struct {
	rows     int
	columns  int
	elements []float64
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

func TestNewFailsNonPositiveRowsOrColumns(t *testing.T) {
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

func TestTransposeTwiceShouldReturnsTheOriginalMatrix(t *testing.T) {
	rows, columns := 4, 3

	m := Zeros(rows, columns)
	tr := m.Transpose()
	n := tr.Transpose()

	mRow, mColumn := m.Shape()
	trRow, trColumn := tr.Shape()
	nRow, nColumn := n.Shape()

	if mRow == trRow || mColumn == trColumn {
		template := "The shape of transpose matrix should be (%d, %d), but (%d, %d)."
		t.Fatalf(template, mColumn, mRow, trRow, trColumn)
	}

	if _, isDense := tr.(*matrixImpl); isDense {
		t.Fatalf("The type of transpose matrix should not be dense.")
	}

	if mRow != nRow || mColumn != nColumn {
		template := "The shape of re-transposed matrix should be (%d, %d), but (%d, %d)."
		t.Fatalf(template, mRow, mColumn, nRow, nColumn)
	}

	if _, isDense := n.(*matrixImpl); !isDense {
		t.Fatalf("The type of re-transpose matrix should not be dense.")
	}
}

func TestTransposeRowsSucceedsAlways(t *testing.T) {
	testSeq := createShapeTestSeq()

	for _, test := range testSeq {
		m, err := New(test.rows, test.columns)(test.elements...)
		if err != nil {
			template := "The number of %q equals to %q * %q, but an error caused."
			t.Fatalf(template, "elements", "rows", "columns")
		}

		tr := m.Transpose()

		if rows := tr.Rows(); rows != test.columns {
			template := "The \"rows\" should be %d, but is %d."
			t.Fatalf(template, test.columns, rows)
		}
	}
}

func TestTransposeColumnsSucceedsAlways(t *testing.T) {
	testSeq := createShapeTestSeq()

	for _, test := range testSeq {
		m, err := New(test.rows, test.columns)(test.elements...)
		if err != nil {
			template := "The number of %q equals to %q * %q, but an error caused."
			t.Fatalf(template, "elements", "rows", "columns")
		}

		tr := m.Transpose()

		if columns := tr.Columns(); columns != test.rows {
			template := "The \"columns\" should be %d, but is %d."
			t.Fatalf(template, test.rows, columns)
		}
	}
}

func TestTransposeUpdateAndGetSucceed(t *testing.T) {
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
	tr := m.Transpose()

	for _, test := range testSeq {
		if element := m.Get(test.row, test.column); element != 0 {
			template := "\"m.Get(%d, %d)\" should return 0, but returns %v."
			t.Fatalf(template, test.row, test.column, element)
		}

		if element := tr.Get(test.column, test.row); element != 0 {
			template := "\"tr.Get(%d, %d)\" should return 0, but returns %v."
			t.Fatalf(template, test.column, test.row, element)
		}

		tr.Update(test.column, test.row, test.element)

		if element := m.Get(test.row, test.column); element != test.element {
			template := "\"m.Get(%d, %d)\" should return %v, but returns %v."
			t.Fatalf(template, test.row, test.column, test.element, element)
		}

		if element := tr.Get(test.column, test.row); element != test.element {
			template := "\"tr.Get(%d, %d)\" should return %v, but returns %v."
			t.Fatalf(template, test.column, test.row, test.element, element)
		}
	}
}

func TestTransposeGetFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	tr := m.Transpose()

	defer func() {
		message := fmt.Sprintf("%d should be smaller than %q %d.", columns, "rows", columns)
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	tr.Get(columns, 0)
}

func TestTransposeGetFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	tr := m.Transpose()

	defer func() {
		message := fmt.Sprintf("%q should be a natural number.", "row")
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	tr.Get(-1, 0)
}

func TestTransposeGetFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	tr := m.Transpose()

	defer func() {
		message := fmt.Sprintf("%d should be smaller than %q %d.", rows, "columns", rows)
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	tr.Get(0, rows)
}

func TestTransposeGetFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	tr := m.Transpose()

	defer func() {
		message := fmt.Sprintf("%q should be a natural number.", "column")
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	tr.Get(0, -1)
}

func TestTransposeUpdateFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	tr := m.Transpose()

	defer func() {
		message := fmt.Sprintf("%d should be smaller than %q %d.", columns, "rows", columns)
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	tr.Update(columns, 0, 0)
}

func TestTransposeUpdateFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	tr := m.Transpose()

	defer func() {
		message := fmt.Sprintf("%q should be a natural number.", "row")
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	tr.Update(-1, 0, 0)
}

func TestTransposeUpdateFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	tr := m.Transpose()

	defer func() {
		message := fmt.Sprintf("%d should be smaller than %q %d.", rows, "columns", rows)
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	tr.Update(0, rows, 0)
}

func TestTransposeUpdateFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	tr := m.Transpose()

	defer func() {
		message := fmt.Sprintf("%q should be a natural number.", "column")
		r := recover()

		if r != nil && r != message {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	tr.Update(0, -1, 0)
}
