package dense

import (
	"testing"

	"github.com/mitsuse/matrix-go/internal/validates"
)

type constructTest struct {
	rows     int
	columns  int
	elements []float64
}

type elementTest struct {
	row     int
	column  int
	element float64
}

func TestNewCreatesDenseMatrix(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("matrix-creation should not cause any panic, but causes %s.", p)
		}
	}()
	New(test.rows, test.columns)(test.elements...)
}

func TestNewFailsForTooManyElements(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  1,
		elements: []float64{0, 1, 2, 3},
	}

	defer func() {
		if p := recover(); p == validates.INVALID_ELEMENTS_PANIC {
			return
		}

		t.Fatal("The number of elements should equal to the product of rows and columns.")
	}()
	New(test.rows, test.columns)(test.elements...)
}

func TestNewFailsForTooFewElements(t *testing.T) {
	test := &constructTest{
		rows:     1,
		columns:  3,
		elements: []float64{0},
	}

	defer func() {
		if p := recover(); p == validates.INVALID_ELEMENTS_PANIC {
			return
		}

		t.Fatal("The number of elements should equal to the product of rows and columns.")
	}()
	New(test.rows, test.columns)(test.elements...)
}

func TestNewFailsForNonPositiveRows(t *testing.T) {
	test := &constructTest{
		rows:     -3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	defer func() {
		if p := recover(); p == validates.NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Non-positive rows should cause %s.",
			validates.NON_POSITIVE_SIZE_PANIC,
		)
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
		if p := recover(); p == validates.NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Non-positive columns should cause %s.",
			validates.NON_POSITIVE_SIZE_PANIC,
		)
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
		if p := recover(); p == validates.NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Non-positive rows or columns should cause %s.",
			validates.NON_POSITIVE_SIZE_PANIC,
		)
	}()
	New(test.rows, test.columns)(test.elements...)
}

func TestZerosSucceeds(t *testing.T) {
	test := &constructTest{
		rows:    3,
		columns: 2,
	}

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("matrix-creation should not cause any panic, but causes %s.", p)
		}
	}()
	Zeros(test.rows, test.columns)
}

func TestZerosCreatesZeroMatrix(t *testing.T) {
	test := &constructTest{
		rows:    3,
		columns: 2,
	}

	for _, element := range Zeros(test.rows, test.columns).(*matrixImpl).elements {
		if element == 0 {
			continue
		}

		t.Fatal("The created matrix should be zero matrix.")
	}
}

func TestZerosFailsForNonPositiveRows(t *testing.T) {
	test := &constructTest{
		rows:    -3,
		columns: 2,
	}

	defer func() {
		if p := recover(); p == validates.NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Non-positive rows should cause %s.",
			validates.NON_POSITIVE_SIZE_PANIC,
		)
	}()
	Zeros(test.rows, test.columns)
}

func TestZerosFailsForNonPositiveColumns(t *testing.T) {
	test := &constructTest{
		rows:    3,
		columns: -2,
	}

	defer func() {
		if p := recover(); p == validates.NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Non-positive columns should cause %s.",
			validates.NON_POSITIVE_SIZE_PANIC,
		)
	}()
	Zeros(test.rows, test.columns)
}

func TestZerosFailsForNonPositive(t *testing.T) {
	test := &constructTest{
		rows:    -3,
		columns: -2,
	}

	defer func() {
		if p := recover(); p == validates.NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf(
			"Non-positive rows or columns should cause %s.",
			validates.NON_POSITIVE_SIZE_PANIC,
		)
	}()
	Zeros(test.rows, test.columns)
}

func TestShapeReturnsTheNumberOfRowsAndColumns(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	rows, columns := New(test.rows, test.columns)(test.elements...).Shape()

	if rows != test.rows {
		t.Fatalf("The rows should be %d, but is %d.", test.rows, rows)
	}

	if columns != test.columns {
		t.Fatalf("The columns should be %d, but is %d.", test.columns, columns)
	}
}

func TestRowsReturnsTheNumberOfRows(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	rows := New(test.rows, test.columns)(test.elements...).Rows()
	if rows == test.rows {
		return
	}

	t.Fatalf("The rows should be %d, but is %d.", test.rows, rows)
}

func TestColumnsReturnsTheNumberOfColumns(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	columns := New(test.rows, test.columns)(test.elements...).Columns()
	if columns == test.columns {
		return
	}

	t.Fatalf("The columns should be %d, but is %d.", test.columns, columns)
}

func TestAllCreatesCursorToIterateAllElements(t *testing.T) {
	m := New(2, 3)(
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

func TestDiagonalCreatesCursorToIterateDiagonalElements(t *testing.T) {
	m := New(3, 3)(
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
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The rows exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Get(rows, 0)
}

func TestGetFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns)

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
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The columns exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Get(0, columns)
}

func TestGetFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns)

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
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The rows exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Update(rows, 0, 0)
}

func TestUpdateFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns)

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
	m := Zeros(rows, columns)

	defer func() {
		if r := recover(); r == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf(
			"The columns exceeds the limit, but %s doesn't cause.",
			validates.OUT_OF_RANGE_PANIC,
		)
	}()
	m.Update(0, columns, 0)
}

func TestUpdateFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns)

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

	rows, columns := 8, 8
	m := Zeros(rows, columns)

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
	m := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

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

func TestEqualIsFalse(t *testing.T) {
	m := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

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

func TestEqualCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

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

func TestAddReturnsTheOriginal(t *testing.T) {
	m := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

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

func TestAddReturnsTheResultOfAddition(t *testing.T) {
	m1 := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	n1 := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	m2 := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

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

func TestAddCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

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
			"Addition of two matrices which have different shape should cause %s.",
			validates.DIFFERENT_SIZE_PANIC,
		)
	}()
	m.Add(n)
}

func TestSubtractReturnsTheOriginal(t *testing.T) {
	m := New(4, 3)(
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
	)

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

func TestSubtractReturnsTheResultOfSubtractition(t *testing.T) {
	m := New(4, 3)(
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
	)

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

func TestSubtractCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

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
			"Subtraction of two matrices which have different shape should cause %s.",
			validates.DIFFERENT_SIZE_PANIC,
		)
	}()
	m.Subtract(n)
}

func TestDotReturnsTheNewMatrixInstance(t *testing.T) {
	m := New(2, 3)(
		2, 1, -3,
		1, -5, 2,
	)

	n := New(3, 3)(
		3, 1, 0,
		2, 0, -1,
		-1, 4, 1,
	)

	if r := m.Dot(n); m != r && n != r {
		return
	}

	t.Fatal("Mutable matrix should return a new instance by multiplication.")
}

func TestDotReturnsTheResultOfMultiplication(t *testing.T) {
	m := New(2, 3)(
		2, 1, -3,
		1, -5, 2,
	)

	n := New(3, 3)(
		3, 1, 0,
		2, 0, -1,
		-1, 4, 1,
	)

	r := New(2, 3)(
		11, -10, -4,
		-9, 9, 7,
	)

	if m.Dot(n).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should multiply the receiver matrix by the given matrix.")
}

func TestMultiplyReturnsTheOriginal(t *testing.T) {
	m := New(4, 3)(
		0, 1, 2,
		3, 2, 1,
		0, 1, 2,
		3, 2, 1,
	)

	s := 3.0

	if m.Multiply(s) == m {
		return
	}

	t.Fatal("Mutable matrix should return itself by scalar-multiplication.")
}

func TestMultiplyTheResultOfMultiplication(t *testing.T) {
	m := New(4, 3)(
		0, 1, 2,
		3, 2, 1,
		0, 1, 2,
		3, 2, 1,
	)

	s := 3.0

	r := New(4, 3)(
		0, 3, 6,
		9, 6, 3,
		0, 3, 6,
		9, 6, 3,
	)

	if m.Multiply(s).Equal(r) {
		return
	}

	t.Fatal("Mutable matrix should multiply each element of itselt by scalar.")
}
