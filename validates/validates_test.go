package validates

import (
	"testing"
)

type shapeTest struct {
	Rows    int
	Columns int
}

func TestShapeShouldBePositiveCausesNothing(t *testing.T) {
	test := &shapeTest{Rows: 1, Columns: 1}

	defer func() {
		if p := recover(); p != nil {
			t.Error("Positive rows or columns should be valid, but causes panic.")
			t.Fatalf(
				"# rows = %d, columns = %d",
				test.Rows,
				test.Columns,
			)
		}
	}()
	ShapeShouldBePositive(test.Rows, test.Columns)
}

func TestShapeShouldBePositiveCausePanicWithZeroRwos(t *testing.T) {
	test := shapeTest{Rows: 0, Columns: 1}

	defer func() {
		if p := recover(); p != NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should be invalid, but causes nothing.",
			)
			t.Fatalf(
				"# rows = %d, columns = %d",
				test.Rows,
				test.Columns,
			)
		}
	}()
	ShapeShouldBePositive(test.Rows, test.Columns)
}

func TestShapeShouldBePositiveCausePanicWithZeroColumns(t *testing.T) {
	test := shapeTest{Rows: 1, Columns: 0}

	defer func() {
		if p := recover(); p != NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should be invalid, but causes nothing.",
			)
			t.Fatalf(
				"# rows = %d, columns = %d",
				test.Rows,
				test.Columns,
			)
		}
	}()
	ShapeShouldBePositive(test.Rows, test.Columns)
}

func TestShapeShouldBePositiveCausePanicWithZeros(t *testing.T) {
	test := shapeTest{Rows: 0, Columns: 0}

	defer func() {
		if p := recover(); p != NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should be invalid, but causes nothing.",
			)
			t.Fatalf(
				"# rows = %d, columns = %d",
				test.Rows,
				test.Columns,
			)
		}
	}()
	ShapeShouldBePositive(test.Rows, test.Columns)
}

func TestShapeShouldBePositiveCausePanicWithNegativeRows(t *testing.T) {
	test := shapeTest{Rows: -1, Columns: 0}

	defer func() {
		if p := recover(); p != NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should be invalid, but causes nothing.",
			)
			t.Fatalf(
				"# rows = %d, columns = %d",
				test.Rows,
				test.Columns,
			)
		}
	}()
	ShapeShouldBePositive(test.Rows, test.Columns)
}

func TestShapeShouldBePositiveCausePanicWithNegativeColumnns(t *testing.T) {
	test := shapeTest{Rows: 0, Columns: -1}

	defer func() {
		if p := recover(); p != NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should be invalid, but causes nothing.",
			)
			t.Fatalf(
				"# rows = %d, columns = %d",
				test.Rows,
				test.Columns,
			)
		}
	}()
	ShapeShouldBePositive(test.Rows, test.Columns)
}

func TestShapeShouldBePositiveCausePanicWithNegative(t *testing.T) {
	test := shapeTest{Rows: -1, Columns: -1}

	defer func() {
		if p := recover(); p != NON_POSITIVE_SIZE_PANIC {
			t.Error(
				"Non-positive rows or columns should be invalid, but causes nothing.",
			)
			t.Fatalf(
				"# rows = %d, columns = %d",
				test.Rows,
				test.Columns,
			)
		}
	}()
	ShapeShouldBePositive(test.Rows, test.Columns)
}

type indexTest struct {
	Row    int
	Column int
}

type rangeTest struct {
	Shape shapeTest
	Index indexTest
}

func TestIndexShouldBeInRangeCausesNothing(t *testing.T) {
	testSeq := []*rangeTest{
		&rangeTest{
			Shape: shapeTest{Rows: 4, Columns: 3},
			Index: indexTest{Row: 0, Column: 0},
		},

		&rangeTest{
			Shape: shapeTest{Rows: 4, Columns: 3},
			Index: indexTest{Row: 0, Column: 1},
		},

		&rangeTest{
			Shape: shapeTest{Rows: 4, Columns: 3},
			Index: indexTest{Row: 1, Column: 0},
		},

		&rangeTest{
			Shape: shapeTest{Rows: 4, Columns: 3},
			Index: indexTest{Row: 3, Column: 0},
		},

		&rangeTest{
			Shape: shapeTest{Rows: 4, Columns: 3},
			Index: indexTest{Row: 0, Column: 2},
		},

		&rangeTest{
			Shape: shapeTest{Rows: 4, Columns: 3},
			Index: indexTest{Row: 3, Column: 2},
		},
	}

	for _, test := range testSeq {
		defer func(test *rangeTest) {
			if p := recover(); p != nil {
				t.Error("Inside-of-range index should be valid, but causes panic.")
				t.Fatalf(
					"# rows = %d, columns = %d, row = %d, column = %d",
					test.Shape.Rows,
					test.Shape.Columns,
					test.Index.Row,
					test.Index.Column,
				)
			}
		}(test)

		IndexShouldBeInRange(
			test.Shape.Rows,
			test.Shape.Columns,
			test.Index.Row,
			test.Index.Column,
		)
	}
}

func TestIndexShouldBeInRangeCausesPanicWithNegativeRow(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{Rows: 4, Columns: 3},
		Index: indexTest{Row: -1, Column: 0},
	}

	defer func() {
		if p := recover(); p != OUT_OF_RANGE_PANIC {
			t.Error("Outside-of-range index should be invalid, but causes nothing.")
			t.Fatalf(
				"# rows = %d, columns = %d, row = %d, column = %d",
				test.Shape.Rows,
				test.Shape.Columns,
				test.Index.Row,
				test.Index.Column,
			)
		}
	}()

	IndexShouldBeInRange(
		test.Shape.Rows,
		test.Shape.Columns,
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithNegativeColumn(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{Rows: 4, Columns: 3},
		Index: indexTest{Row: 0, Column: -1},
	}

	defer func() {
		if p := recover(); p != OUT_OF_RANGE_PANIC {
			t.Error("Outside-of-range index should be invalid, but causes nothing.")
			t.Fatalf(
				"# rows = %d, columns = %d, row = %d, column = %d",
				test.Shape.Rows,
				test.Shape.Columns,
				test.Index.Row,
				test.Index.Column,
			)
		}
	}()

	IndexShouldBeInRange(
		test.Shape.Rows,
		test.Shape.Columns,
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithNegative(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{Rows: 4, Columns: 3},
		Index: indexTest{Row: -1, Column: -1},
	}

	defer func() {
		if p := recover(); p != OUT_OF_RANGE_PANIC {
			t.Error("Outside-of-range index should be invalid, but causes nothing.")
			t.Fatalf(
				"# rows = %d, columns = %d, row = %d, column = %d",
				test.Shape.Rows,
				test.Shape.Columns,
				test.Index.Row,
				test.Index.Column,
			)
		}
	}()

	IndexShouldBeInRange(
		test.Shape.Rows,
		test.Shape.Columns,
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithLargeRow(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{Rows: 4, Columns: 3},
		Index: indexTest{Row: 4, Column: 0},
	}

	defer func() {
		if p := recover(); p != OUT_OF_RANGE_PANIC {
			t.Error("Outside-of-range index should be invalid, but causes nothing.")
			t.Fatalf(
				"# rows = %d, columns = %d, row = %d, column = %d",
				test.Shape.Rows,
				test.Shape.Columns,
				test.Index.Row,
				test.Index.Column,
			)
		}
	}()

	IndexShouldBeInRange(
		test.Shape.Rows,
		test.Shape.Columns,
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithLargeColumn(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{Rows: 4, Columns: 3},
		Index: indexTest{Row: 0, Column: 3},
	}

	defer func() {
		if p := recover(); p != OUT_OF_RANGE_PANIC {
			t.Error("Outside-of-range index should be invalid, but causes nothing.")
			t.Fatalf(
				"# rows = %d, columns = %d, row = %d, column = %d",
				test.Shape.Rows,
				test.Shape.Columns,
				test.Index.Row,
				test.Index.Column,
			)
		}
	}()

	IndexShouldBeInRange(
		test.Shape.Rows,
		test.Shape.Columns,
		test.Index.Row,
		test.Index.Column,
	)
}

func TestIndexShouldBeInRangeCausesPanicWithLarge(t *testing.T) {
	test := &rangeTest{
		Shape: shapeTest{Rows: 4, Columns: 3},
		Index: indexTest{Row: 4, Column: 3},
	}

	defer func() {
		if p := recover(); p != OUT_OF_RANGE_PANIC {
			t.Error("Outside-of-range index should be invalid, but causes nothing.")
			t.Fatalf(
				"# rows = %d, columns = %d, row = %d, column = %d",
				test.Shape.Rows,
				test.Shape.Columns,
				test.Index.Row,
				test.Index.Column,
			)
		}
	}()

	IndexShouldBeInRange(
		test.Shape.Rows,
		test.Shape.Columns,
		test.Index.Row,
		test.Index.Column,
	)
}
