package validates

import (
	"testing"
)

type shapeTest struct {
	rows    int
	columns int
}

func (t *shapeTest) Shape() (rows, columns int) {
	return t.Rows(), t.Columns()
}

func (t *shapeTest) Rows() int {
	return t.rows
}

func (t *shapeTest) Columns() int {
	return t.columns
}

type indexTest struct {
	Row    int
	Column int
}

type rangeTest struct {
	Shape shapeTest
	Index indexTest
}

func TestShapeShouldBePositiveCausesNothing(t *testing.T) {
	test := &shapeTest{rows: 1, columns: 1}

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("Positive rows or columns should be valid, but causes %s.", p)
		}
	}()
	ShapeShouldBePositive(test.Rows(), test.Columns())
}

func TestShapeShouldBePositiveCausePanicWithZeroRwos(t *testing.T) {
	test := shapeTest{rows: 0, columns: 1}

	defer func() {
		if p := recover(); p == NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf("Non-positive rows or columns should cause %s.", NON_POSITIVE_SIZE_PANIC)
	}()
	ShapeShouldBePositive(test.Rows(), test.Columns())
}

func TestShapeShouldBePositiveCausePanicWithZeroColumns(t *testing.T) {
	test := shapeTest{rows: 1, columns: 0}

	defer func() {
		if p := recover(); p == NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf("Non-positive rows or columns should cause %s.", NON_POSITIVE_SIZE_PANIC)
	}()
	ShapeShouldBePositive(test.Rows(), test.Columns())
}

func TestShapeShouldBePositiveCausePanicWithZeros(t *testing.T) {
	test := shapeTest{rows: 0, columns: 0}

	defer func() {
		if p := recover(); p == NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf("Non-positive rows or columns should cause %s.", NON_POSITIVE_SIZE_PANIC)
	}()
	ShapeShouldBePositive(test.Rows(), test.Columns())
}

func TestShapeShouldBePositiveCausePanicWithNegativeRows(t *testing.T) {
	test := shapeTest{rows: -1, columns: 0}

	defer func() {
		if p := recover(); p == NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf("Non-positive rows or columns should cause %s.", NON_POSITIVE_SIZE_PANIC)
	}()
	ShapeShouldBePositive(test.Rows(), test.Columns())
}

func TestShapeShouldBePositiveCausePanicWithNegativeColumnns(t *testing.T) {
	test := shapeTest{rows: 0, columns: -1}

	defer func() {
		if p := recover(); p == NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf("Non-positive rows or columns should cause %s.", NON_POSITIVE_SIZE_PANIC)
	}()
	ShapeShouldBePositive(test.Rows(), test.Columns())
}

func TestShapeShouldBePositiveCausePanicWithNegative(t *testing.T) {
	test := shapeTest{rows: -1, columns: -1}

	defer func() {
		if p := recover(); p == NON_POSITIVE_SIZE_PANIC {
			return
		}

		t.Fatalf("Non-positive rows or columns should cause %s.", NON_POSITIVE_SIZE_PANIC)
	}()
	ShapeShouldBePositive(test.Rows(), test.Columns())
}

func TestShapeShouldBeSameCauseNathing(t *testing.T) {
	m := &shapeTest{rows: 4, columns: 3}
	n := &shapeTest{rows: 4, columns: 3}

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("Two matrices have same size, but cause %s.", p)
		}
	}()
	ShapeShouldBeSame(m, n)
}

func TestShapeShouldBeSameCausePanicWithDifferenceShape(t *testing.T) {
	m := &shapeTest{rows: 4, columns: 3}
	n := &shapeTest{rows: 6, columns: 2}

	defer func() {
		if p := recover(); p == DIFFERENT_SIZE_PANIC {
			return
		}

		t.Fatalf("Two matrices with different size should cause %s.", DIFFERENT_SIZE_PANIC)
	}()
	ShapeShouldBeSame(m, n)
}

func TestShapeShouldBeSameCausePanicWithDifferenceRows(t *testing.T) {
	m := &shapeTest{rows: 4, columns: 3}
	n := &shapeTest{rows: 6, columns: 3}

	defer func() {
		if p := recover(); p == DIFFERENT_SIZE_PANIC {
			return
		}

		t.Fatalf("Two matrices with different size should cause %s.", DIFFERENT_SIZE_PANIC)
	}()
	ShapeShouldBeSame(m, n)
}

func TestShapeShouldBeSameCausePanicWithDifferenceColumns(t *testing.T) {
	m := &shapeTest{rows: 4, columns: 3}
	n := &shapeTest{rows: 4, columns: 2}

	defer func() {
		if p := recover(); p == DIFFERENT_SIZE_PANIC {
			return
		}

		t.Fatalf("Two matrices with different size should cause %s.", DIFFERENT_SIZE_PANIC)
	}()
	ShapeShouldBeSame(m, n)
}

func TestShapeShouldBeMultipliableCausesNothing(t *testing.T) {
	m := &shapeTest{rows: 4, columns: 2}
	n := &shapeTest{rows: 2, columns: 3}

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf("Two matrices are mulpliable, but cause %s.", p)
		}
	}()
	ShapeShouldBeMultipliable(m, n)
}

func TestShapeShouldBeMultipliableCausesPanic(t *testing.T) {
	m := &shapeTest{rows: 4, columns: 2}
	n := &shapeTest{rows: 5, columns: 3}

	defer func() {
		if p := recover(); p == NOT_MULTIPLIABLE_PANIC {
			return
		}

		t.Fatalf("Two matrices should not be mulpliable.")
	}()
	ShapeShouldBeMultipliable(m, n)
}

func TestIndexShouldBeInRangeCausesNothing(t *testing.T) {
	testSeq := []*rangeTest{
		&rangeTest{
			Shape: shapeTest{rows: 4, columns: 3},
			Index: indexTest{Row: 0, Column: 0},
		},

		&rangeTest{
			Shape: shapeTest{rows: 4, columns: 3},
			Index: indexTest{Row: 0, Column: 1},
		},

		&rangeTest{
			Shape: shapeTest{rows: 4, columns: 3},
			Index: indexTest{Row: 1, Column: 0},
		},

		&rangeTest{
			Shape: shapeTest{rows: 4, columns: 3},
			Index: indexTest{Row: 3, Column: 0},
		},

		&rangeTest{
			Shape: shapeTest{rows: 4, columns: 3},
			Index: indexTest{Row: 0, Column: 2},
		},

		&rangeTest{
			Shape: shapeTest{rows: 4, columns: 3},
			Index: indexTest{Row: 3, Column: 2},
		},
	}

	for _, test := range testSeq {
		defer func(test *rangeTest) {
			if p := recover(); p != nil {
				t.Fatalf(
					"An index %v is contained in shape %v, but cause %s.",
					test.Index,
					test.Shape,
					p,
				)
			}
		}(test)

		IndexShouldBeInRange(
			test.Shape.Rows(),
			test.Shape.Columns(),
			test.Index.Row,
			test.Index.Column,
		)
	}
}

func TestIndexShouldBeInRangeCausesPanicWithNegativeRow(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{rows: 4, columns: 3},
		Index: indexTest{Row: -1, Column: 0},
	}

	defer func() {
		if p := recover(); p == OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf("Outside-of-range index should cause %s.", OUT_OF_RANGE_PANIC)
	}()

	IndexShouldBeInRange(
		test.Shape.Rows(),
		test.Shape.Columns(),
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithNegativeColumn(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{rows: 4, columns: 3},
		Index: indexTest{Row: 0, Column: -1},
	}

	defer func() {
		if p := recover(); p == OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf("Outside-of-range index should cause %s.", OUT_OF_RANGE_PANIC)
	}()

	IndexShouldBeInRange(
		test.Shape.Rows(),
		test.Shape.Columns(),
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithNegative(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{rows: 4, columns: 3},
		Index: indexTest{Row: -1, Column: -1},
	}

	defer func() {
		if p := recover(); p == OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf("Outside-of-range index should cause %s.", OUT_OF_RANGE_PANIC)
	}()

	IndexShouldBeInRange(
		test.Shape.Rows(),
		test.Shape.Columns(),
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithLargeRow(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{rows: 4, columns: 3},
		Index: indexTest{Row: 4, Column: 0},
	}

	defer func() {
		if p := recover(); p == OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf("Outside-of-range index should cause %s.", OUT_OF_RANGE_PANIC)
	}()

	IndexShouldBeInRange(
		test.Shape.Rows(),
		test.Shape.Columns(),
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithLargeColumn(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{rows: 4, columns: 3},
		Index: indexTest{Row: 0, Column: 3},
	}

	defer func() {
		if p := recover(); p == OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf("Outside-of-range index should cause %s.", OUT_OF_RANGE_PANIC)
	}()

	IndexShouldBeInRange(
		test.Shape.Rows(),
		test.Shape.Columns(),
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithLarge(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{rows: 4, columns: 3},
		Index: indexTest{Row: 4, Column: 3},
	}

	defer func() {
		if p := recover(); p == OUT_OF_RANGE_PANIC {
			return
		}

		t.Fatalf("Outside-of-range index should cause %s.", OUT_OF_RANGE_PANIC)
	}()

	IndexShouldBeInRange(
		test.Shape.Rows(),
		test.Shape.Columns(),
		test.Index.Row,
		test.Index.Column,
	)
}
