package validates

import (
	"fmt"
)

const (
	NON_POSITIVE_SIZE_PANIC = iota
)

func RowsShouldBePositiveNumber(rows int) {
	if isPositiveNumber(rows) {
		return
	}

	panic(NON_POSITIVE_SIZE_PANIC)
}

func ColumnShouldBePositiveNumber(columns int) {
	if isPositiveNumber(columns) {
		return
	}

	panic(NON_POSITIVE_SIZE_PANIC)
}

func RowShouldBeInRows(row, rows int) {
	shouldBeNaturalNumber(row, "row")
	shouldBeSmallerThan(row, rows, "rows")
}

func ColumnShouldBeInColumns(column, columns int) {
	shouldBeNaturalNumber(column, "column")
	shouldBeSmallerThan(column, columns, "columns")
}

func isPositiveNumber(x int) bool {
	return x > 0
}

func shouldBeNaturalNumber(x int, name string) {
	if x >= 0 {
		return
	}

	message := fmt.Sprintf("%q should be a natural number.", name)
	panic(message)
}

func shouldBeSmallerThan(x, limit int, name string) {
	if x < limit {
		return
	}

	message := fmt.Sprintf("%d should be smaller than %q %d.", x, name, limit)
	panic(message)
}
