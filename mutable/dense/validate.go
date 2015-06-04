package dense

import (
	"fmt"
)

func rowsShouldBePositiveNumber(rows int) {
	shouldBePositiveNumber(rows, "rows")
}

func columnShouldBePositiveNumber(columns int) {
	shouldBePositiveNumber(columns, "columns")
}

func rowShouldBeInRows(row, rows int) {
	shouldBeNaturalNumber(row, "row")
	shouldBeSmallerThan(row, rows, "rows")
}

func columnShouldBeInColumns(column, columns int) {
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
