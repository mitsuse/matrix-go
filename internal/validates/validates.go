package validates

import "github.com/mitsuse/matrix-go/internal/types"

const (
	NON_POSITIVE_SIZE_PANIC Panic = iota
	DIFFERENT_SIZE_PANIC
	NOT_MULTIPLIABLE_PANIC
	OUT_OF_RANGE_PANIC
	INVALID_ELEMENTS_PANIC
	INVALID_VIEW_PANIC
)

//go:generate stringer -type=Panic
type Panic int

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

func ViewShouldBeInBase(base, view types.Shape, offset types.Index) {
	rows := offset.Row() + view.Rows()
	columns := offset.Column() + view.Columns()

	if 0 < rows && rows <= base.Rows() && 0 < columns && columns <= base.Columns() {
		return
	}

	panic(INVALID_VIEW_PANIC)
}
