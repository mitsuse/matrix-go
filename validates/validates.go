package validates

const (
	NON_POSITIVE_SIZE_PANIC = iota
	OUT_OF_RANGE_PANIC
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
	if isNaturalNumber(row) && isSmallerThan(row, rows) {
		return
	}

	panic(OUT_OF_RANGE_PANIC)
}

func ColumnShouldBeInColumns(column, columns int) {
	if isNaturalNumber(column) && isSmallerThan(column, columns) {
		return
	}

	panic(OUT_OF_RANGE_PANIC)
}

func isPositiveNumber(x int) bool {
	return x > 0
}

func isNaturalNumber(x int) bool {
	return x >= 0
}

func isSmallerThan(x, limit int) bool {
	return x < limit
}
