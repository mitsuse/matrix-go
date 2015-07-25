package dense

import (
	"testing"

	"github.com/mitsuse/matrix-go/validates"
)

func TestTransposeShapeReturnsTheNumberOfRowsAndColumns(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	rows, columns := New(test.rows, test.columns)(test.elements...).Transpose().Shape()

	if rows != test.columns {
		t.Fatalf("The rows should be %d, but is %d.", test.columns, rows)
	}

	if columns != test.rows {
		t.Fatalf("The columns should be %d, but is %d.", test.rows, columns)
	}
}

func TestTransposeRowsReturnsTheNumberOfRows(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	rows := New(test.rows, test.columns)(test.elements...).Transpose().Rows()
	if rows == test.columns {
		return
	}

	t.Fatalf("The rows should be %d, but is %d.", test.columns, rows)
}

func TestTransposeColumnsReturnsTheNumberOfColumns(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	columns := New(test.rows, test.columns)(test.elements...).Transpose().Columns()
	if columns == test.rows {
		return
	}

	t.Fatalf("The columns should be %d, but is %d.", test.rows, columns)
}

func TestTransposeAllCreatesCursorToIterateAllElements(t *testing.T) {
	m := New(3, 2)(
		0, 3,
		1, 4,
		2, 5,
	).Transpose()

	checkTable := [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
	}

	cursor := m.All()

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

func TestTransposeNonZerosCreatesCursorToIterateNonZeroElements(t *testing.T) {
	m := New(3, 2)(
		0, 0,
		1, 0,
		2, 3,
	).Transpose()

	checkTable := [][]bool{
		[]bool{true, false, false},
		[]bool{true, true, false},
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

func TestTransposeDiagonalCreatesCursorToIterateDiagonalElements(t *testing.T) {
	m := New(3, 3)(
		1, 0, 0,
		0, 2, 0,
		0, 0, 3,
	).Transpose()

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

func TestTransposeGetFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns).Transpose()

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The rows exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Get(columns, 0)
}

func TestTransposeGetFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns).Transpose()

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

func TestTransposeGetFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns).Transpose()

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The columns exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Get(0, rows)
}

func TestTransposeGetFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns).Transpose()

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

func TestTransposeUpdateFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns).Transpose()

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The rows exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Update(columns, 0, 0)
}

func TestTransposeUpdateFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns).Transpose()

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

func TestTransposeUpdateFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns).Transpose()

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The columns exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Update(0, rows, 0)
}

func TestTransposeUpdateFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns).Transpose()

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

func TestTransposeUpdateReplacesElement(t *testing.T) {
	test := &elementTest{row: 1, column: 4, element: 1}

	rows, columns := 8, 8
	m := Zeros(rows, columns)
	n := m.Transpose()

	if element := m.Get(test.row, test.column); element != 0 {
		t.Fatalf(
			"The element at (%d, %d) should be 0 before updating, but is %v.",
			test.row,
			test.column,
			element,
		)
	}

	if element := n.Get(test.column, test.row); element != 0 {
		t.Fatalf(
			"The transpose element at (%d, %d) should be 0 before updating, but is %v.",
			test.column,
			test.row,
			element,
		)
	}

	n.Update(test.column, test.row, test.element)

	if element := m.Get(test.row, test.column); element != test.element {
		t.Fatalf(
			"The element at (%d, %d) should be %v after updating, but is %v.",
			test.row,
			test.column,
			test.element,
			element,
		)
	}

	if element := n.Get(test.column, test.row); element != test.element {
		t.Fatalf(
			"The transpos element at (%d, %d) should be %v after updating, but is %v.",
			test.column,
			test.row,
			test.element,
			element,
		)
	}
}

func TestTransposeEqualIsTrue(t *testing.T) {
	m := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	).Transpose()

	n := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if m.Equal(n) && n.Equal(m) {
		return
	}

	t.Fatal("The equality of two matrices should be true, but the result is false.")
}

func TestTransposeEqualIsFalse(t *testing.T) {
	m := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	).Transpose()

	n := New(4, 3)(
		0, 1, 2,
		3, 1, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if !m.Equal(n) && !n.Equal(m) {
		return
	}

	t.Fatal("The equality of two matrices should be false, but the result is true.")
}

func TestTransposeEqualCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	).Transpose()

	n := New(3, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	)

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

func TestTransposeAddReturnsTheOriginal(t *testing.T) {
	m := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	).Transpose()

	n := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	if r := m.Add(n); m == r {
		return
	}

	t.Fatal("Mutable matrix should return itself by addition.")
}

func TestTransposeAddReturnsTheResultOfAddition(t *testing.T) {
	m1 := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	).Transpose()

	n1 := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	m2 := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	).Transpose()

	n2 := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	r := New(4, 3)(
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
	)

	if m1.Add(n1).Equal(r) && n2.Add(m2).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should add other matrix to itself.")
}

func TestTransposeAddCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	).Transpose()

	n := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

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

func TestTransposeSubtractReturnsTheOriginal(t *testing.T) {
	m := New(3, 4)(
		9, 9, 9, 9,
		9, 9, 9, 9,
		9, 9, 9, 9,
	).Transpose()

	n := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	if r := m.Subtract(n); m == r {
		return
	}

	t.Fatal("Mutable matrix should return itself by subtraction.")
}

func TestTransposeSubtractReturnsTheResultOfSubtractition(t *testing.T) {
	m := New(3, 4)(
		9, 9, 9, 9,
		9, 9, 9, 9,
		9, 9, 9, 9,
	).Transpose()

	n := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	r := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if m.Subtract(n).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should subtract other matrix from itself.")
}

func TestTransposeSubtractCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(4, 3)(
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
	).Transpose()

	n := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

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

func TestTransposeMultiplyReturnsTheNewMatrixInstance(t *testing.T) {
	m := New(3, 2)(
		2, 1,
		1, -5,
		-3, 2,
	).Transpose()

	n := New(3, 3)(
		3, 1, 0,
		2, 0, -1,
		-1, 4, 1,
	)

	if r := m.Multiply(n); m != r && n != r {
		return
	}

	t.Fatal("Mutable matrix should return a new instance by multiplication.")
}

func TestTransposeMultiplyReturnsTheResultOfMultiplication(t *testing.T) {
	m := New(3, 2)(
		2, 1,
		1, -5,
		-3, 2,
	).Transpose()

	n := New(3, 3)(
		3, 1, 0,
		2, 0, -1,
		-1, 4, 1,
	)

	r := New(2, 3)(
		11, -10, -4,
		-9, 9, 7,
	)

	if m.Multiply(n).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should multiply the receiver matrix by the given matrix.")
}

func TestTransposeScalarReturnsTheOriginal(t *testing.T) {
	m := New(3, 4)(
		0, 3, 0, 3,
		1, 2, 1, 2,
		2, 1, 2, 1,
	).Transpose()

	s := 3.0

	if m.Scalar(s) == m {
		return
	}

	t.Fatal("Mutable matrix should return itself by scalar-multiplication.")
}

func TestTransposeScalarTheResultOfMultiplication(t *testing.T) {
	m := New(3, 4)(
		0, 3, 0, 3,
		1, 2, 1, 2,
		2, 1, 2, 1,
	).Transpose()

	s := 3.0

	r := New(4, 3)(
		0, 3, 6,
		9, 6, 3,
		0, 3, 6,
		9, 6, 3,
	)

	if m.Scalar(s).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should multiply each element of itselt by scalar.")
}

func TestTransposeTwiceEqualsToTheOriginalMatrix(t *testing.T) {
	rows, columns := 4, 3

	m := Zeros(rows, columns)
	n := m.Transpose().Transpose()

	if !m.Equal(n) {
		t.Fatal("The re-transpose matrix should equal to the original matrix.")
	}
}
