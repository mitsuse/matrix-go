package hash

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/mitsuse/matrix-go/internal/rewriters"
	"github.com/mitsuse/matrix-go/internal/types"
	"github.com/mitsuse/matrix-go/internal/validates"
)

type elementTest struct {
	row     int
	column  int
	element float64
}

func TestHashMatrixSatisfiesMatrixInterface(t *testing.T) {
	var _ types.Matrix = &Matrix{}
}

func TestNewCreateHashMatrix(t *testing.T) {
	rows := 3
	columns := 2

	elements := []Element{
		{Row: 0, Column: 0, Value: 9},
		{Row: 2, Column: 1, Value: 3},
		{Row: 1, Column: 0, Value: 1},
	}

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("matrix-creation should not cause any panic, but causes %s.", p)
		}
	}()
	New(rows, columns)(elements...)
}

func TestNewFailsForOutOfShapeElement(t *testing.T) {
	rows := 3
	columns := 2

	elements := []Element{
		{Row: 0, Column: 0, Value: 9},
		{Row: 2, Column: 1, Value: 3},
		{Row: 2, Column: 2, Value: 3},
		{Row: 1, Column: 0, Value: 1},
	}

	defer func() {
		if p := recover(); p == validates.OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatal("Out-of-shaep elements should be rejected.")
	}()
	New(rows, columns)(elements...)
}

func TestNewFailsForDuplicatedIndexElements(t *testing.T) {
	rows := 3
	columns := 2

	elements := []Element{
		{Row: 0, Column: 0, Value: 9},
		{Row: 2, Column: 1, Value: 3},
		{Row: 2, Column: 1, Value: 2},
		{Row: 1, Column: 0, Value: 1},
	}

	defer func() {
		if p := recover(); p == validates.INVALID_ELEMENTS_PANIC {
			return
		}

		t.Fatal("Duplicated-index elements should be rejected.")
	}()
	New(rows, columns)(elements...)
}

func TestNewFailsForNonPositiveRows(t *testing.T) {
	rows := -3
	columns := 2

	elements := []Element{
		{Row: 0, Column: 0, Value: 9},
		{Row: 2, Column: 1, Value: 3},
		{Row: 1, Column: 0, Value: 1},
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
	New(rows, columns)(elements...)
}

func TestNewFailsForNonPositiveColumns(t *testing.T) {
	rows := 3
	columns := -2

	elements := []Element{
		{Row: 0, Column: 0, Value: 9},
		{Row: 2, Column: 1, Value: 3},
		{Row: 1, Column: 0, Value: 1},
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
	New(rows, columns)(elements...)
}

func TestNewFailsForNonPositive(t *testing.T) {
	rows := -3
	columns := -2

	elements := []Element{
		{Row: 0, Column: 0, Value: 9},
		{Row: 2, Column: 1, Value: 3},
		{Row: 1, Column: 0, Value: 1},
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
	New(rows, columns)(elements...)
}

func TestZerosCreatesZeroMatrix(t *testing.T) {
	rows := 3
	columns := 2

	if len(Zeros(rows, columns).elements) == 0 {
		return
	}
	t.Fatal("The created matrix should be zero matrix.")
}

func TestSerialize(t *testing.T) {
	m := New(3, 3)(
		Element{Row: 0, Column: 0, Value: 1.0},
		Element{Row: 0, Column: 1, Value: 0.1},
		Element{Row: 0, Column: 2, Value: 0.9},
		Element{Row: 1, Column: 0, Value: 0.1},
		Element{Row: 1, Column: 1, Value: 2.5},
		Element{Row: 1, Column: 2, Value: 0.2},
		Element{Row: 2, Column: 0, Value: 0.2},
		Element{Row: 2, Column: 1, Value: 0.1},
		Element{Row: 2, Column: 2, Value: 3.1},
	).View(1, 1, 2, 1).Transpose()

	writer := bytes.NewBuffer([]byte{})

	if err := m.Serialize(writer); err != nil {
		t.Fatalf("An expected error occured on serialization: %s", err)
	}

	reader := bytes.NewReader(writer.Bytes())

	n, err := Deserialize(reader)

	if err != nil {
		t.Fatalf("An expected error occured on deserialization: %s", err)
	}

	if !m.Base().Equal(n.Base()) || !m.Equal(n) {
		t.Fatal("Deserialization failed for a serialized matrix.")
	}
}

func TestUnmarshalJSONFailsWithAlreadyInitializedMatrix(t *testing.T) {
	m := New(3, 3)(
		Element{Row: 0, Column: 0, Value: 1.0},
		Element{Row: 0, Column: 1, Value: 0.1},
		Element{Row: 0, Column: 2, Value: 0.9},
		Element{Row: 1, Column: 0, Value: 0.1},
		Element{Row: 1, Column: 1, Value: 2.5},
		Element{Row: 1, Column: 2, Value: 0.2},
		Element{Row: 2, Column: 0, Value: 0.2},
		Element{Row: 2, Column: 1, Value: 0.1},
		Element{Row: 2, Column: 2, Value: 3.1},
	).View(1, 1, 2, 1).Transpose()

	n := New(3, 3)(
		Element{Row: 0, Column: 0, Value: 1.0},
		Element{Row: 0, Column: 1, Value: 0.1},
		Element{Row: 0, Column: 2, Value: 0.9},
		Element{Row: 1, Column: 0, Value: 0.1},
		Element{Row: 1, Column: 1, Value: 2.5},
		Element{Row: 1, Column: 2, Value: 0.2},
		Element{Row: 2, Column: 0, Value: 0.2},
		Element{Row: 2, Column: 1, Value: 0.1},
		Element{Row: 2, Column: 2, Value: 3.1},
	).View(1, 1, 2, 1).Transpose()

	b, _ := json.Marshal(m)

	if err := json.Unmarshal(b, n); err == nil || err.Error() != AlreadyInitializedError {
		t.Fatalf("Unmarshal can be applied to uninitialized matrix.")
	}
}

func TestUnmarshalJSONFailsWithIncompatibleVersion(t *testing.T) {
	m := &matrixJson{
		Version: 99999,
		Base:    types.NewShape(3, 3),
		View:    types.NewShape(2, 1),
		Offset:  types.NewIndex(1, 1),
		Elements: []Element{
			{Row: 0, Column: 0, Value: 1.0},
			{Row: 0, Column: 1, Value: 0.1},
			{Row: 0, Column: 2, Value: 0.9},
			{Row: 1, Column: 0, Value: 0.1},
			{Row: 1, Column: 1, Value: 2.5},
			{Row: 1, Column: 2, Value: 0.2},
			{Row: 2, Column: 0, Value: 0.2},
			{Row: 2, Column: 1, Value: 0.1},
			{Row: 2, Column: 2, Value: 3.1},
		},
		Rewriter: rewriters.Reverse().Type(),
	}

	n := &Matrix{}

	b, _ := json.Marshal(m)

	if err := json.Unmarshal(b, n); err == nil || err.Error() != IncompatibleVersionError {
		t.Fatalf("Unmarshal can be applied to compatible-version matrix.")
	}
}

func TestUnmarshalJSONFailsWithUnknownRewriter(t *testing.T) {
	m := &matrixJson{
		Version: version,
		Base:    types.NewShape(3, 3),
		View:    types.NewShape(2, 1),
		Offset:  types.NewIndex(1, 1),
		Elements: []Element{
			{Row: 0, Column: 0, Value: 1.0},
			{Row: 0, Column: 1, Value: 0.1},
			{Row: 0, Column: 2, Value: 0.9},
			{Row: 1, Column: 0, Value: 0.1},
			{Row: 1, Column: 1, Value: 2.5},
			{Row: 1, Column: 2, Value: 0.2},
			{Row: 2, Column: 0, Value: 0.2},
			{Row: 2, Column: 1, Value: 0.1},
			{Row: 2, Column: 2, Value: 3.1},
		},
		Rewriter: 255,
	}

	n := &Matrix{}

	b, _ := json.Marshal(m)

	// TODO: Use an exported error constant.
	if err := json.Unmarshal(b, n); err == nil || err.Error() != "UNKNOWN_REWRITER_ERROR" {
		t.Fatalf("Unmarshal can be applied to matrix which has a valid rewriter.")
	}
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
		Element{Row: 3, Column: 2, Value: 6},
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
