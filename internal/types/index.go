package types

type Index struct {
	row    int
	column int
}

func NewIndex(row, column int) Index {
	i := Index{
		row:    row,
		column: column,
	}

	return i
}

func (i Index) Row() int {
	return i.row
}

func (i Index) Column() int {
	return i.column
}
