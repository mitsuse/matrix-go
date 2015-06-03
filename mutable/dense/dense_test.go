package dense

import (
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
