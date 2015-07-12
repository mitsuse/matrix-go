package validates

import (
	"github.com/mitsuse/matrix-go"
)

const (
	NON_POSITIVE_SIZE_PANIC = iota
	DIFFERENT_SIZE_PANIC
	OUT_OF_RANGE_PANIC
)

func ShapeShouldBePositive(row, column int) {
	if row > 0 && column > 0 {
		return
	}

	panic(NON_POSITIVE_SIZE_PANIC)
}

func ShapeShouldBeSame(m, n matrix.Matrix) {
	mRow, mColumn := m.Shape()
	nRow, nColumn := n.Shape()

	if mRow == nRow && mColumn == nColumn {
		return
	}

	panic(DIFFERENT_SIZE_PANIC)
}

func IndexShouldBeInRange(rows, columns, row, column int) {
	if (0 <= row && row < rows) && (0 <= column && column < columns) {
		return
	}

	panic(OUT_OF_RANGE_PANIC)
}
