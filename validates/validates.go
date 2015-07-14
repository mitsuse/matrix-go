package validates

const (
	NON_POSITIVE_SIZE_PANIC = iota
	DIFFERENT_SIZE_PANIC
	NOT_MULTIPLIABLE_PANIC
	OUT_OF_RANGE_PANIC
)

func ShapeShouldBePositive(row, column int) {
	if row > 0 && column > 0 {
		return
	}

	panic(NON_POSITIVE_SIZE_PANIC)
}

type HasShape interface {
	Shape() (rows, columns int)
	Rows() int
	Columns() int
}

func ShapeShouldBeSame(m, n HasShape) {
	mRow, mColumn := m.Shape()
	nRow, nColumn := n.Shape()

	if mRow == nRow && mColumn == nColumn {
		return
	}

	panic(DIFFERENT_SIZE_PANIC)
}

func ShapeShouldBeMultipliable(m, n HasShape) {
	if m.Columns() == n.Rows() {
		return
	}

	panic(NOT_MULTIPLIABLE_PANIC)
}

func IndexShouldBeInRange(rows, columns, row, column int) {
	if (0 <= row && row < rows) && (0 <= column && column < columns) {
		return
	}

	panic(OUT_OF_RANGE_PANIC)
}
