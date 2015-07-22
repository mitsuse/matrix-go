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

func TestEqualIsTrue(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	n, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if equality := m.Equal(n); !equality {
		t.Fatal("Two matrices should equal, but the result is false.")
	}
}

func TestEqualIsFalse(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	n, _ := New(4, 3)(
		0, 1, 2,
		3, 1, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if equality := m.Equal(n); equality {
		t.Fatal("Two matrices should not equal, but the result is true.")
	}
}

func TestEqualCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	n, _ := New(3, 3)(
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
				"# m.rows = %d, m.columns = %d, n.rows = %d, n.columns = %d",
				m.Rows(),
				m.Columns(),
				n.Rows(),
				n.Columns(),
			)
		}
	}()
	m.Equal(n)
}

func TestAddReturnsTheOriginal(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	n, _ := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	if r := m.Add(n); m != r {
		t.Fatal("Mutable matrix should return itself by addition.")
	}
}

func TestAddReturnsTheResultOfAddition(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	n, _ := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	r, _ := New(4, 3)(
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
	)

	m.Add(n)

	if !m.Equal(r) {
		t.Fatal("Mutable matrix should add other matrix to itself.")
	}
}

func TestAddCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	n, _ := New(3, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	)

	defer func() {
		if r := recover(); r != nil && r != validates.DIFFERENT_SIZE_PANIC {
			t.Error(
				"Addition of two matrices which have different shape",
				"should cause panic, but cause nothing.",
			)
			t.Fatalf(
				"# m.rows = %d, m.columns = %d, n.rows = %d, n.columns = %d",
				m.Rows(),
				m.Columns(),
				n.Rows(),
				n.Columns(),
			)
		}
	}()
	m.Add(n)
}

func TestSubtractReturnsTheOriginal(t *testing.T) {
	m, _ := New(4, 3)(
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
	)

	n, _ := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	if r := m.Subtract(n); m != r {
		t.Fatal("Mutable matrix should return itself by subtraction.")
	}
}

func TestSubtractReturnsTheResultOfSubtractition(t *testing.T) {
	m, _ := New(4, 3)(
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
	)

	n, _ := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	r, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	m.Subtract(n)

	if !m.Equal(r) {
		t.Fatal("Mutable matrix should subtract other matrix from itself.")
	}
}

func TestSubtractCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	n, _ := New(3, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	)

	defer func() {
		if r := recover(); r != nil && r != validates.DIFFERENT_SIZE_PANIC {
			t.Error(
				"Subtraction of two matrices which have different shape",
				"should cause panic, but cause nothing.",
			)
			t.Fatalf(
				"# m.rows = %d, m.columns = %d, n.rows = %d, n.columns = %d",
				m.Rows(),
				m.Columns(),
				n.Rows(),
				n.Columns(),
			)
		}
	}()
	m.Subtract(n)
}

func TestScalarReturnsTheOriginal(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 2, 1,
		0, 1, 2,
		3, 2, 1,
	)

	s := 3.0

	if m.Scalar(s) != m {
		t.Fatal("Mutable matrix should return itself by scalar-multiplication.")
	}
}

func TestScalarTheResultOfMultiplication(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 2, 1,
		0, 1, 2,
		3, 2, 1,
	)

	s := 3.0

	r, _ := New(4, 3)(
		0, 3, 6,
		9, 6, 3,
		0, 3, 6,
		9, 6, 3,
	)

	if !m.Scalar(s).Equal(r) {
		t.Fatal("Mutable matrix should multiply each element of itselt by scalar.")
	}
}

func TestMultiplyReturnsTheNewMatrixInstance(t *testing.T) {
	m, _ := New(2, 3)(
		2, 1, -3,
		1, -5, 2,
	)

	n, _ := New(3, 3)(
		3, 1, 0,
		2, 0, -1,
		-1, 4, 1,
	)

	if r := m.Multiply(n); m == r || n == r {
		t.Fatal("Mutable matrix should return a new instance by multiplication.")
	}
}

func TestMultiplyReturnsTheResultOfMultiplication(t *testing.T) {
	m, _ := New(2, 3)(
		2, 1, -3,
		1, -5, 2,
	)

	n, _ := New(3, 3)(
		3, 1, 0,
		2, 0, -1,
		-1, 4, 1,
	)

	r, _ := New(2, 3)(
		11, -10, -4,
		-9, 9, 7,
	)

	if !m.Multiply(n).Equal(r) {
		t.Fatal(
			"Mutable matrix should multiply the receiver matrix by the given matrix.",
		)
	}
}

func TestTransposeTwiceEqualsToTheOriginalMatrix(t *testing.T) {
	rows, columns := 4, 3

	m := Zeros(rows, columns)
	n := m.Transpose().Transpose()

	if !m.Equal(n) {
		t.Fatal("The re-transpose matrix should equal to the original matrix.")
	}
}

func TestTransposeRowsReturnsTheNumberOfRows(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	m, _ := New(test.rows, test.columns)(test.elements...)
	n := m.Transpose()

	if rows := n.Rows(); rows != test.columns {
		t.Fatalf("The \"rows\" should be %d, but is %d.", test.columns, rows)
	}
}

func TestTransposeColumnsReturnsTheNumberOfColumns(t *testing.T) {
	test := &constructTest{
		rows:     3,
		columns:  2,
		elements: []float64{0, 1, 2, 3, 4, 5},
	}

	m, _ := New(test.rows, test.columns)(test.elements...)
	n := m.Transpose()

	if columns := n.Columns(); columns != test.rows {
		t.Fatalf("The \"columns\" should be %d, but is %d.", test.rows, columns)
	}
}

func TestTransposeUpdateReplacesElement(t *testing.T) {
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
	n := m.Transpose()

	for _, test := range testSeq {
		if element := n.Get(test.row, test.column); element != 0 {
			t.Fatalf(
				"The element at (%d, %d) should be 0 before updating, but is %v.",
				test.row,
				test.column,
				test.element,
			)
		}

		n.Update(test.row, test.column, test.element)

		if element := n.Get(test.row, test.column); element != test.element {
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

func TestTransposeGetFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns)
	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	n.Get(columns, 0)
}

func TestTransposeGetFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns)
	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	n.Get(-1, 0)
}

func TestTransposeGetFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns)
	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	n.Get(0, rows)
}

func TestTransposeGetFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns)
	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	n.Get(0, -1)
}

func TestTransposeUpdateReturnsTheReceiverMatrix(t *testing.T) {
	rows, columns := 8, 8
	m := Zeros(rows, columns)
	n := m.Transpose()

	if n.Update(0, 0, 0) != n {
		t.Fatal("Transpose matrix should return itself by updating.")
	}
}

func TestTransposeUpdateFailsByAccessingWithTooLargeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns)
	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" exceeds the limit, but no panic causes.")
		}
	}()
	n.Update(columns, 0, 0)
}

func TestTransposeUpdateFailsByAccessingWithNegativeRow(t *testing.T) {
	rows, columns := 8, 6
	m := Zeros(rows, columns)
	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"row\" is negative, but no panic causes.")
		}
	}()
	n.Update(-1, 0, 0)
}

func TestTransposeUpdateFailsByAccessingWithTooLargeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns)
	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" exceeds the limit, but no panic causes.")
		}
	}()
	n.Update(0, rows, 0)
}

func TestTransposeUpdateFailsByAccessingWithNegativeColumn(t *testing.T) {
	rows, columns := 6, 8
	m := Zeros(rows, columns)
	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.OUT_OF_RANGE_PANIC {
			t.Fatalf("The \"column\" is negative, but no panic causes.")
		}
	}()
	n.Update(0, -1, 0)
}

func TestTransposeAllCreatesCursorToIterateAllElements(t *testing.T) {
	m, _ := New(2, 3)(
		0, 1, 2,
		3, 4, 5,
	)

	n := m.Transpose()

	checkTable := [][]bool{
		[]bool{false, false},
		[]bool{false, false},
		[]bool{false, false},
	}

	cursor := n.All()

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

func TestTransposeNonZerosCreatesCursorToIterateNonZeroElements(t *testing.T) {
	m, _ := New(2, 3)(
		0, 1, 2,
		0, 0, 3,
	)

	n := m.Transpose()

	checkTable := [][]bool{
		[]bool{true, true},
		[]bool{false, true},
		[]bool{false, false},
	}

	cursor := n.NonZeros()

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

func TestTransposeDiagonalCreatesCursorToIterateDiagonalElements(t *testing.T) {
	m, _ := New(3, 3)(
		1, 0, 0,
		0, 2, 0,
		0, 0, 3,
	)

	n := m.Transpose()

	checkTable := [][]bool{
		[]bool{false, true, true},
		[]bool{true, false, true},
		[]bool{true, true, false},
	}

	cursor := n.Diagonal()

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

func TestTransposeEqualIsTrue(t *testing.T) {
	m, _ := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	n := m.Transpose()

	r, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if equality := n.Equal(r); !equality {
		t.Fatal("Two matrices should equal, but the result is false.")
	}
}

func TestTransposeEqualIsFalse(t *testing.T) {
	m, _ := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	n := m.Transpose()

	r, _ := New(4, 3)(
		0, 1, 2,
		3, 1, 5,
		6, 7, 8,
		9, 0, 1,
	)

	if equality := n.Equal(r); equality {
		t.Fatal("Two matrices should not equal, but the result is true.")
	}
}

func TestTransposeEqualCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m, _ := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	n := m.Transpose()

	defer func() {
		if r := recover(); r != nil && r != validates.DIFFERENT_SIZE_PANIC {
			t.Error(
				"Checking equality of matrices which have different shape",
				"should cause panic, but cause nothing.",
			)
			t.Fatalf(
				"# m.rows = %d, m.columns = %d, n.rows = %d, n.columns = %d",
				m.Rows(),
				m.Columns(),
				n.Rows(),
				n.Columns(),
			)
		}
	}()
	m.Equal(n)
}

func TestTransposeAddReturnsTheOriginal(t *testing.T) {
	m, _ := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	n := m.Transpose()

	o, _ := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	if n.Add(o) != n {
		t.Fatal("Transpose matrix should return itself by addition.")
	}
}

func TestTransposeAddReturnsTheResultOfAddition(t *testing.T) {
	m, _ := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	n := m.Transpose()

	o, _ := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	r, _ := New(4, 3)(
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
		9, 9, 9,
	)

	b, _ := New(3, 4)(
		9, 9, 9, 9,
		9, 9, 9, 9,
		9, 9, 9, 9,
	)

	if !n.Add(o).Equal(r) || !m.Equal(b) {
		t.Fatal("Transpose matrix should add other matrix to its base matrix.")
	}
}

func TestTransposeAddCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m, _ := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	n := m.Transpose()

	o, _ := New(3, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	)

	defer func() {
		if r := recover(); r != nil && r != validates.DIFFERENT_SIZE_PANIC {
			t.Error(
				"Addition of two matrices which have different shape",
				"should cause panic, but cause nothing.",
			)
			t.Fatalf(
				"# n.rows = %d, n.columns = %d, o.rows = %d, o.columns = %d",
				n.Rows(),
				n.Columns(),
				o.Rows(),
				o.Columns(),
			)
		}
	}()
	n.Add(o)
}

func TestTransposeSubtractReturnsTheOriginal(t *testing.T) {
	m, _ := New(3, 4)(
		9, 9, 9, 9,
		9, 9, 9, 9,
		9, 9, 9, 9,
	)

	n := m.Transpose()

	o, _ := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	if n.Subtract(o) != n {
		t.Fatal("Transpose matrix should return itself by subtraction.")
	}
}

func TestTransposeSubtractReturnsTheResultOfSubtractition(t *testing.T) {
	m, _ := New(3, 4)(
		9, 9, 9, 9,
		9, 9, 9, 9,
		9, 9, 9, 9,
	)

	n := m.Transpose()

	o, _ := New(4, 3)(
		9, 8, 7,
		6, 5, 4,
		3, 2, 1,
		0, 9, 8,
	)

	r, _ := New(4, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		9, 0, 1,
	)

	b, _ := New(3, 4)(
		0, 3, 6, 9,
		1, 4, 7, 0,
		2, 5, 8, 1,
	)

	if !n.Subtract(o).Equal(r) || !b.Equal(m) {
		t.Fatal("Transpose matrix should subtract other matrix from itself.")
	}
}

func TestTransposeSubtractCausesPanicForDifferentShapeMatrices(t *testing.T) {
	m, _ := New(3, 4)(
		9, 9, 9, 9,
		9, 9, 9, 9,
		9, 9, 9, 9,
	)

	n := m.Transpose()

	o, _ := New(3, 3)(
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
	)

	defer func() {
		if r := recover(); r != nil && r != validates.DIFFERENT_SIZE_PANIC {
			t.Error(
				"Subtraction of two matrices which have different shape",
				"should cause panic, but cause nothing.",
			)
			t.Fatalf(
				"# tr.rows = %d, tr.columns = %d, n.rows = %d, n.columns = %d",
				n.Rows(),
				n.Columns(),
				o.Rows(),
				o.Columns(),
			)
		}
	}()
	n.Subtract(o)
}

func TestTransposeScalarReturnsTheOriginal(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 2, 1,
		0, 1, 2,
		3, 2, 1,
	)

	n := m.Transpose()

	s := 3.0

	if n.Scalar(s) != n {
		t.Fatal("Transpose matrix should return itself by scalar-multiplication.")
	}
}

func TestTransposeScalarTheResultOfMultiplication(t *testing.T) {
	m, _ := New(4, 3)(
		0, 1, 2,
		3, 2, 1,
		0, 1, 2,
		3, 2, 1,
	)

	n := m.Transpose()

	s := 3.0

	r, _ := New(3, 4)(
		0, 9, 0, 9,
		3, 6, 3, 6,
		6, 3, 6, 3,
	)

	b, _ := New(4, 3)(
		0, 3, 6,
		9, 6, 3,
		0, 3, 6,
		9, 6, 3,
	)

	if !n.Scalar(s).Equal(r) || !m.Equal(b) {
		t.Fatal("Transpose matrix should multiply each element of itselt by scalar.")
	}
}

func TestTransposeMultiplyReturnsTheNewMatrixInstance(t *testing.T) {
	m, _ := New(3, 2)(
		2, 1,
		1, -5,
		-3, 2,
	)

	n := m.Transpose()

	o, _ := New(3, 3)(
		3, 1, 0,
		2, 0, -1,
		-1, 4, 1,
	)

	if r := n.Multiply(o); n == r || m == r || o == r {
		t.Fatal("Transpose matrix should return a new instance by multiplication.")
	}
}

func TestTransposeMultiplyReturnsTheResultOfMultiplication(t *testing.T) {
	m, _ := New(3, 2)(
		2, 1,
		1, -5,
		-3, 2,
	)

	n := m.Transpose()

	o, _ := New(3, 3)(
		3, 1, 0,
		2, 0, -1,
		-1, 4, 1,
	)

	r, _ := New(2, 3)(
		11, -10, -4,
		-9, 9, 7,
	)

	if !n.Multiply(o).Equal(r) {
		t.Fatal(
			"Transpose matrix should multiply the receiver matrix by the given matrix.",
		)
	}
}
