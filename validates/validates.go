package validates

import (
	"fmt"
)

func RowsShouldBePositiveNumber(rows int) {
	shouldBePositiveNumber(rows, "rows")
}

func ColumnShouldBePositiveNumber(columns int) {
	shouldBePositiveNumber(columns, "columns")
}

func RowShouldBeInRows(row, rows int) {
	shouldBeNaturalNumber(row, "row")
	shouldBeSmallerThan(row, rows, "rows")
}

func ColumnShouldBeInColumns(column, columns int) {
	shouldBeNaturalNumber(column, "column")
	shouldBeSmallerThan(column, columns, "columns")
}

func shouldBePositiveNumber(x int, name string) {
	if x > 0 {
		return
	}

	message := fmt.Sprintf("%q should be a positive number.", name)
	panic(message)
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
