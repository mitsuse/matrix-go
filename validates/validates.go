package validates

const (
	NON_POSITIVE_SIZE_PANIC = iota
	OUT_OF_RANGE_PANIC
)

func ShapeShouldBePositive(row, column int) {
	if row > 0 && column > 0 {
		return
	}

	panic(NON_POSITIVE_SIZE_PANIC)
}

func IndexShouldBeInRange(rows, columns, row, column int) {
	if (0 <= row && row < rows) && (0 <= column && column < columns) {
		return
	}

	panic(OUT_OF_RANGE_PANIC)
}
