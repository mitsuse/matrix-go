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

func TestNonZerosCreatesCursorToIterateNonZeroElements(t *testing.T) {
	m := New(2, 3)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 2, Value: 3},
	).View(1, 1, 1, 2)

	checkTable := [][]bool{
		[]bool{true, false},
	}

	cursor := m.NonZeros()

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

func TestEqualIsTrue(t *testing.T) {
	m := New(5, 4)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 1, Value: 4},
		Element{Row: 1, Column: 2, Value: 5},
		Element{Row: 2, Column: 0, Value: 6},
		Element{Row: 2, Column: 1, Value: 7},
		Element{Row: 2, Column: 2, Value: 8},
		Element{Row: 3, Column: 0, Value: 9},
		Element{Row: 3, Column: 2, Value: 1},
	).View(0, 0, 4, 3)

	n := New(7, 5)(
		Element{Row: 1, Column: 3, Value: 1},
		Element{Row: 1, Column: 4, Value: 2},
		Element{Row: 2, Column: 2, Value: 3},
		Element{Row: 2, Column: 3, Value: 4},
		Element{Row: 2, Column: 4, Value: 5},
		Element{Row: 3, Column: 2, Value: 6},
		Element{Row: 3, Column: 3, Value: 7},
		Element{Row: 3, Column: 4, Value: 8},
		Element{Row: 4, Column: 2, Value: 9},
		Element{Row: 4, Column: 4, Value: 1},
	).View(1, 2, 4, 3)

	if m.Equal(n) && n.Equal(m) {
		return
	}

	t.Fatal("The equality of two matrices should be true, but the result is false.")
}

func TestEqualIsFalse(t *testing.T) {
	m := New(5, 4)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 1, Value: 4},
		Element{Row: 1, Column: 2, Value: 5},
		Element{Row: 2, Column: 0, Value: 6},
		Element{Row: 2, Column: 1, Value: 7},
		Element{Row: 2, Column: 2, Value: 8},
		Element{Row: 3, Column: 0, Value: 9},
		Element{Row: 3, Column: 2, Value: 1},
	).View(0, 0, 4, 3)

	n := New(7, 5)(
		Element{Row: 1, Column: 3, Value: 1},
		Element{Row: 1, Column: 4, Value: 2},
		Element{Row: 2, Column: 2, Value: 3},
		Element{Row: 2, Column: 3, Value: 4},
		Element{Row: 3, Column: 2, Value: 6},
		Element{Row: 3, Column: 3, Value: 7},
		Element{Row: 3, Column: 4, Value: 8},
		Element{Row: 4, Column: 2, Value: 9},
		Element{Row: 4, Column: 4, Value: 1},
	).View(1, 2, 4, 3)

	if !m.Equal(n) && !n.Equal(m) {
		return
	}

	t.Fatal("The equality of two matrices should be false, but the result is true.")
}

func TestEqualCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(5, 4)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 1, Value: 4},
		Element{Row: 1, Column: 2, Value: 5},
		Element{Row: 2, Column: 0, Value: 6},
		Element{Row: 2, Column: 1, Value: 7},
		Element{Row: 2, Column: 2, Value: 8},
		Element{Row: 3, Column: 0, Value: 9},
		Element{Row: 3, Column: 2, Value: 1},
	).View(0, 0, 4, 4)

	n := New(7, 5)(
		Element{Row: 1, Column: 3, Value: 1},
		Element{Row: 1, Column: 4, Value: 2},
		Element{Row: 2, Column: 2, Value: 3},
		Element{Row: 2, Column: 3, Value: 4},
		Element{Row: 2, Column: 4, Value: 5},
		Element{Row: 3, Column: 2, Value: 6},
		Element{Row: 3, Column: 3, Value: 7},
		Element{Row: 3, Column: 4, Value: 8},
		Element{Row: 4, Column: 2, Value: 9},
		Element{Row: 4, Column: 4, Value: 1},
	).View(1, 2, 4, 3)

	defer func() {
		if r := recover(); r == validates.DIFFERENT_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Checking equality of matrices which have different shape should cause %s.",
			validates.DIFFERENT_SIZE_PANIC,
		)
	}()
	m.Equal(n)
}

func TestAddReturnsTheOriginal(t *testing.T) {
	m := New(4, 3)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 1, Value: 4},
		Element{Row: 1, Column: 2, Value: 5},
		Element{Row: 2, Column: 0, Value: 6},
		Element{Row: 2, Column: 1, Value: 7},
		Element{Row: 2, Column: 2, Value: 8},
		Element{Row: 3, Column: 0, Value: 9},
		Element{Row: 3, Column: 2, Value: 1},
	).View(1, 1, 2, 2)

	n := New(3, 4)(
		Element{Row: 1, Column: 0, Value: 5},
		Element{Row: 1, Column: 1, Value: 4},
		Element{Row: 2, Column: 0, Value: 2},
		Element{Row: 2, Column: 1, Value: 1},
	).View(1, 0, 2, 2)

	if r := m.Add(n); m == r {
		return
	}

	t.Fatal("Mutable matrix should return itself by addition.")
}

func TestAddReturnsTheResultOfAddition(t *testing.T) {
	m1 := New(4, 3)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 1, Value: 4},
		Element{Row: 1, Column: 2, Value: 5},
		Element{Row: 2, Column: 0, Value: 6},
		Element{Row: 2, Column: 1, Value: 7},
		Element{Row: 2, Column: 2, Value: 8},
		Element{Row: 3, Column: 0, Value: 9},
		Element{Row: 3, Column: 2, Value: 1},
	).View(1, 1, 2, 2)

	n1 := New(3, 3)(
		Element{Row: 0, Column: 1, Value: 5},
		Element{Row: 0, Column: 2, Value: 4},
		Element{Row: 1, Column: 1, Value: 2},
		Element{Row: 1, Column: 2, Value: 1},
	).View(0, 1, 2, 2)

	m2 := New(4, 3)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 1, Value: 4},
		Element{Row: 1, Column: 2, Value: 5},
		Element{Row: 2, Column: 0, Value: 6},
		Element{Row: 2, Column: 1, Value: 7},
		Element{Row: 2, Column: 2, Value: 8},
		Element{Row: 3, Column: 0, Value: 9},
		Element{Row: 3, Column: 2, Value: 1},
	).View(1, 1, 2, 2)

	n2 := New(3, 3)(
		Element{Row: 0, Column: 1, Value: 5},
		Element{Row: 0, Column: 2, Value: 4},
		Element{Row: 1, Column: 1, Value: 2},
		Element{Row: 1, Column: 2, Value: 1},
	).View(0, 1, 2, 2)

	r := New(3, 2)(
		Element{Row: 0, Column: 0, Value: 9},
		Element{Row: 0, Column: 1, Value: 9},
		Element{Row: 1, Column: 0, Value: 9},
		Element{Row: 1, Column: 1, Value: 9},
	).View(0, 0, 2, 2)

	if m1.Add(n1).Equal(r) && n2.Add(m2).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should add other matrix to itself.")
}

func TestAddCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(4, 5)(
		Element{Row: 0, Column: 2, Value: 1},
		Element{Row: 0, Column: 3, Value: 2},
		Element{Row: 1, Column: 2, Value: 3},
		Element{Row: 1, Column: 3, Value: 4},
		Element{Row: 1, Column: 4, Value: 5},
		Element{Row: 2, Column: 2, Value: 6},
		Element{Row: 2, Column: 3, Value: 7},
		Element{Row: 2, Column: 4, Value: 8},
		Element{Row: 3, Column: 2, Value: 9},
		Element{Row: 3, Column: 4, Value: 1},
	).View(0, 1, 4, 3)

	n := New(4, 4)(
		Element{Row: 1, Column: 2, Value: 1},
		Element{Row: 1, Column: 3, Value: 2},
		Element{Row: 2, Column: 2, Value: 3},
		Element{Row: 2, Column: 3, Value: 4},
		Element{Row: 3, Column: 4, Value: 5},
		Element{Row: 3, Column: 2, Value: 6},
		Element{Row: 4, Column: 3, Value: 7},
		Element{Row: 5, Column: 4, Value: 8},
	).View(1, 1, 3, 3)

	defer func() {
		if r := recover(); r == validates.DIFFERENT_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Addition of two matrices which have different shape should cause %s.",
			validates.DIFFERENT_SIZE_PANIC,
		)
	}()
	m.Add(n)
}

func TestSubtractReturnsTheOriginal(t *testing.T) {
	m := New(4, 3)(
		Element{Row: 0, Column: 0, Value: 9},
		Element{Row: 0, Column: 1, Value: 9},
		Element{Row: 0, Column: 2, Value: 9},
		Element{Row: 1, Column: 0, Value: 9},
		Element{Row: 1, Column: 1, Value: 9},
		Element{Row: 1, Column: 2, Value: 9},
		Element{Row: 2, Column: 0, Value: 9},
		Element{Row: 2, Column: 1, Value: 9},
		Element{Row: 2, Column: 2, Value: 9},
		Element{Row: 3, Column: 0, Value: 9},
		Element{Row: 3, Column: 1, Value: 9},
		Element{Row: 3, Column: 2, Value: 9},
	).View(1, 1, 2, 2)

	n := New(2, 3)(
		Element{Row: 0, Column: 0, Value: 5},
		Element{Row: 0, Column: 1, Value: 4},
		Element{Row: 1, Column: 0, Value: 2},
		Element{Row: 1, Column: 1, Value: 1},
	).View(0, 0, 2, 2)

	if r := m.Subtract(n); m == r {
		return
	}

	t.Fatal("Mutable matrix should return itself by subtraction.")
}

func TestSubtractReturnsTheResultOfSubtractition(t *testing.T) {
	m := New(4, 3)(
		Element{Row: 0, Column: 0, Value: 9},
		Element{Row: 0, Column: 1, Value: 9},
		Element{Row: 0, Column: 2, Value: 9},
		Element{Row: 1, Column: 0, Value: 9},
		Element{Row: 1, Column: 1, Value: 9},
		Element{Row: 1, Column: 2, Value: 9},
		Element{Row: 2, Column: 0, Value: 9},
		Element{Row: 2, Column: 1, Value: 9},
		Element{Row: 2, Column: 2, Value: 9},
		Element{Row: 3, Column: 0, Value: 9},
		Element{Row: 3, Column: 1, Value: 9},
		Element{Row: 3, Column: 2, Value: 9},
	).View(1, 1, 2, 2)

	n := New(3, 3)(
		Element{Row: 0, Column: 1, Value: 5},
		Element{Row: 0, Column: 2, Value: 4},
		Element{Row: 1, Column: 1, Value: 2},
		Element{Row: 1, Column: 2, Value: 1},
	).View(0, 1, 2, 2)

	r := New(3, 3)(
		Element{Row: 1, Column: 0, Value: 4},
		Element{Row: 1, Column: 1, Value: 5},
		Element{Row: 2, Column: 0, Value: 7},
		Element{Row: 2, Column: 1, Value: 8},
	).View(1, 0, 2, 2)

	if m.Subtract(n).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should subtract other matrix from itself.")
}

func TestSubtractCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(4, 4)(
		Element{Row: 0, Column: 2, Value: 1},
		Element{Row: 0, Column: 3, Value: 2},
		Element{Row: 1, Column: 1, Value: 3},
		Element{Row: 1, Column: 2, Value: 4},
		Element{Row: 1, Column: 3, Value: 5},
		Element{Row: 2, Column: 1, Value: 6},
		Element{Row: 2, Column: 2, Value: 7},
		Element{Row: 2, Column: 3, Value: 8},
		Element{Row: 3, Column: 1, Value: 9},
		Element{Row: 3, Column: 3, Value: 1},
	).View(0, 1, 4, 3)

	n := New(4, 4)(
		Element{Row: 1, Column: 1, Value: 1},
		Element{Row: 1, Column: 2, Value: 2},
		Element{Row: 2, Column: 0, Value: 3},
		Element{Row: 2, Column: 1, Value: 4},
		Element{Row: 2, Column: 2, Value: 5},
		Element{Row: 3, Column: 0, Value: 6},
		Element{Row: 3, Column: 1, Value: 7},
		Element{Row: 3, Column: 2, Value: 8},
	).View(1, 0, 3, 3)

	defer func() {
		if r := recover(); r == validates.DIFFERENT_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Subtraction of two matrices which have different shape should cause %s.",
			validates.DIFFERENT_SIZE_PANIC,
		)
	}()
	m.Subtract(n)
}

func TestMultiplyReturnsTheNewMatrixInstance(t *testing.T) {
	m := New(4, 4)(
		Element{Row: 0, Column: 1, Value: 2},
		Element{Row: 0, Column: 2, Value: 1},
		Element{Row: 0, Column: 3, Value: -3},
		Element{Row: 1, Column: 1, Value: 1},
		Element{Row: 1, Column: 2, Value: -5},
		Element{Row: 1, Column: 3, Value: 2},
	).View(0, 1, 2, 3)

	n := New(4, 5)(
		Element{Row: 0, Column: 2, Value: 3},
		Element{Row: 0, Column: 3, Value: 1},
		Element{Row: 1, Column: 2, Value: 2},
		Element{Row: 1, Column: 4, Value: -1},
		Element{Row: 2, Column: 2, Value: -1},
		Element{Row: 2, Column: 3, Value: 4},
		Element{Row: 2, Column: 4, Value: 1},
	).View(0, 2, 3, 3)

	if r := m.Multiply(n); m != r && n != r {
		return
	}

	t.Fatal("Mutable matrix should return a new instance by multiplication.")
}

func TestMultiplyReturnsTheResultOfMultiplication(t *testing.T) {
	m := New(4, 4)(
		Element{Row: 0, Column: 1, Value: 2},
		Element{Row: 0, Column: 2, Value: 1},
		Element{Row: 0, Column: 3, Value: -3},
		Element{Row: 1, Column: 1, Value: 1},
		Element{Row: 1, Column: 2, Value: -5},
		Element{Row: 1, Column: 3, Value: 2},
	).View(0, 1, 2, 3)

	n := New(4, 5)(
		Element{Row: 0, Column: 2, Value: 3},
		Element{Row: 0, Column: 3, Value: 1},
		Element{Row: 1, Column: 2, Value: 2},
		Element{Row: 1, Column: 4, Value: -1},
		Element{Row: 2, Column: 2, Value: -1},
		Element{Row: 2, Column: 3, Value: 4},
		Element{Row: 2, Column: 4, Value: 1},
	).View(0, 2, 3, 3)

	r := New(2, 3)(
		Element{Row: 0, Column: 0, Value: 11},
		Element{Row: 0, Column: 1, Value: -10},
		Element{Row: 0, Column: 2, Value: -4},
		Element{Row: 1, Column: 0, Value: -9},
		Element{Row: 1, Column: 1, Value: 9},
		Element{Row: 1, Column: 2, Value: 7},
	)

	if m.Multiply(n).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should multiply the receiver matrix by the given matrix.")
}

func TestScalarReturnsTheOriginal(t *testing.T) {
	m := New(5, 4)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 1, Value: 2},
		Element{Row: 1, Column: 2, Value: 1},
		Element{Row: 2, Column: 1, Value: 1},
		Element{Row: 2, Column: 2, Value: 2},
		Element{Row: 3, Column: 0, Value: 3},
		Element{Row: 3, Column: 1, Value: 2},
		Element{Row: 3, Column: 2, Value: 1},
	).View(0, 0, 4, 3)

	s := 3.0

	if m.Scalar(s) == m {
		return
	}

	t.Fatal("Mutable matrix should return itself by scalar-multiplication.")
}

func TestScalarTheResultOfMultiplication(t *testing.T) {
	m := New(5, 4)(
		Element{Row: 0, Column: 1, Value: 1},
		Element{Row: 0, Column: 2, Value: 2},
		Element{Row: 1, Column: 0, Value: 3},
		Element{Row: 1, Column: 1, Value: 2},
		Element{Row: 1, Column: 2, Value: 1},
		Element{Row: 2, Column: 1, Value: 1},
		Element{Row: 2, Column: 2, Value: 2},
		Element{Row: 3, Column: 0, Value: 3},
		Element{Row: 3, Column: 1, Value: 2},
		Element{Row: 3, Column: 2, Value: 1},
	).View(0, 0, 4, 3)

	s := 3.0

	r := New(5, 4)(
		Element{Row: 1, Column: 2, Value: 3},
		Element{Row: 1, Column: 3, Value: 6},
		Element{Row: 2, Column: 1, Value: 9},
		Element{Row: 2, Column: 2, Value: 6},
		Element{Row: 2, Column: 3, Value: 3},
		Element{Row: 3, Column: 2, Value: 3},
		Element{Row: 3, Column: 3, Value: 6},
		Element{Row: 4, Column: 1, Value: 9},
		Element{Row: 4, Column: 2, Value: 6},
		Element{Row: 4, Column: 3, Value: 3},
	).View(1, 1, 4, 3)

	if m.Scalar(s).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should multiply each element of itselt by scalar.")
}
