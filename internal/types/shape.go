package types

type Shape struct {
	rows    int
	columns int
}

func NewShape(rows, columns int) Shape {
	s := Shape{
		rows:    rows,
		columns: columns,
	}

	return s
}

func (s Shape) Rows() int {
	return s.rows
}

func (s Shape) Columns() int {
	return s.columns
}
