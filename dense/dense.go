/*
Package "dense" provides an implementation of mutable dense matrix.
*/
package dense

import (
	"io"
	"math"

	"github.com/mitsuse/matrix-go/internal/rewriters"
	"github.com/mitsuse/matrix-go/internal/types"
	"github.com/mitsuse/matrix-go/internal/validates"
	"github.com/mitsuse/serial-go"
)

const (
	id      string = "github.com/mitsuse/matrix-go/dense"
	version byte   = 1
)

type denseMatrix struct {
	base     types.Shape
	view     types.Shape
	offset   types.Index
	elements []float64
	rewriter rewriters.Rewriter
}

// Create a new matrix with given elements.
// When "rows" and "columns" is not positive,
// validates.NON_POSITIVE_SIZE_PANIC will be caused.
// In addition,
// when the product of "row"s and "column" doesn't equal to the size of "elements",
// validates.INVALID_ELEMENTS_PANIC will be caused.
func New(rows, columns int) func(elements ...float64) types.Matrix {
	validates.ShapeShouldBePositive(rows, columns)

	constructor := func(elements ...float64) types.Matrix {
		size := rows * columns

		if len(elements) != size {
			panic(validates.INVALID_ELEMENTS_PANIC)
		}

		shape := types.NewShape(rows, columns)
		offset := types.NewIndex(0, 0)

		m := &denseMatrix{
			base:     shape,
			view:     shape,
			offset:   offset,
			elements: make([]float64, size),
			rewriter: rewriters.Reflect(),
		}
		copy(m.elements, elements)

		return m
	}

	return constructor
}

// Create a new zero matrix.
// When "rows" and "columns" is not positive,
// validates.NON_POSITIVE_SIZE_PANIC will be caused.
func Zeros(rows, columns int) types.Matrix {
	return New(rows, columns)(make([]float64, rows*columns)...)
}

// Deserialize a matrix from the given reader.
// This accepts data generated with (*denseMatrix).Serialize.
func Deserialize(reader io.Reader) (types.Matrix, error) {
	r := serial.NewReader(id, version, reader)

	r.ReadId()
	r.ReadVersion()
	r.ReadArch()

	var baseRows int64
	r.Read(&baseRows)

	var baseColumns int64
	r.Read(&baseColumns)

	var viewRows int64
	r.Read(&viewRows)

	var viewColumns int64
	r.Read(&viewColumns)

	var offsetRow int64
	r.Read(&offsetRow)

	var offsetColumn int64
	r.Read(&offsetColumn)

	var size int64
	r.Read(&size)

	elements := make([]float64, int(size))
	for index := 0; index < len(elements); index++ {
		var element float64
		r.Read(&element)
		elements[index] = element
	}

	if err := r.Error(); err != nil {
		return nil, err
	}

	rewriter, err := rewriters.Deserialize(reader)
	if err != nil {
		return nil, err
	}

	m := &denseMatrix{
		base:     types.NewShape(int(baseRows), int(baseColumns)),
		view:     types.NewShape(int(viewRows), int(viewColumns)),
		offset:   types.NewIndex(int(offsetRow), int(offsetColumn)),
		elements: elements,
		rewriter: rewriter,
	}

	if len(elements) != m.base.Rows()*m.base.Columns() {
		panic(validates.INVALID_ELEMENTS_PANIC)
	}

	return m, nil
}

func (m *denseMatrix) Serialize(writer io.Writer) error {
	w := serial.NewWriter(id, version, writer)

	w.WriteId()
	w.WriteVersion()
	w.WriteArch()

	w.Write(int64(m.base.Rows()))
	w.Write(int64(m.base.Columns()))
	w.Write(int64(m.view.Rows()))
	w.Write(int64(m.view.Columns()))
	w.Write(int64(m.offset.Row()))
	w.Write(int64(m.offset.Column()))

	w.Write(int64(len(m.elements)))
	for _, element := range m.elements {
		w.Write(int64(element))
	}

	if err := w.Error(); err != nil {
		return err
	}

	if err := m.rewriter.Serialize(writer); err != nil {
		return err
	}

	return nil
}

func (m *denseMatrix) Shape() (rows, columns int) {
	return m.rewriter.Rewrite(m.view.Rows(), m.view.Columns())
}

func (m *denseMatrix) Rows() (rows int) {
	rows, _ = m.Shape()
	return rows
}

func (m *denseMatrix) Columns() (columns int) {
	_, columns = m.Shape()
	return columns
}

func (m *denseMatrix) All() types.Cursor {
	return newAllCursor(m)
}

func (m *denseMatrix) NonZeros() types.Cursor {
	return newNonZerosCursor(m)
}

func (m *denseMatrix) Diagonal() types.Cursor {
	return newDiagonalCursor(m)
}

func (m *denseMatrix) Get(row, column int) (element float64) {
	row, column = m.rewriter.Rewrite(row, column)

	validates.IndexShouldBeInRange(m.view.Rows(), m.view.Columns(), row, column)

	index := (row+m.offset.Row())*m.base.Columns() + column + m.offset.Column()

	return m.elements[index]
}

func (m *denseMatrix) Update(row, column int, element float64) types.Matrix {
	row, column = m.rewriter.Rewrite(row, column)

	validates.IndexShouldBeInRange(m.view.Rows(), m.view.Columns(), row, column)

	index := (row+m.offset.Row())*m.base.Columns() + column + m.offset.Column()
	m.elements[index] = element

	return m
}

func (m *denseMatrix) Equal(n types.Matrix) bool {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.All()

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		if m.Get(row, column) != element {
			return false
		}
	}

	return true
}

func (m *denseMatrix) Add(n types.Matrix) types.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		m.Update(row, column, m.Get(row, column)+element)
	}

	return m
}

func (m *denseMatrix) Subtract(n types.Matrix) types.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		m.Update(row, column, m.Get(row, column)-element)
	}

	return m
}

func (m *denseMatrix) Multiply(n types.Matrix) types.Matrix {
	validates.ShapeShouldBeMultipliable(m, n)

	rows := m.Rows()
	columns := n.Columns()

	r := Zeros(rows, columns)

	cursor := n.NonZeros()

	for cursor.HasNext() {
		element, j, k := cursor.Get()

		for i := 0; i < rows; i++ {
			r.Update(i, k, r.Get(i, k)+m.Get(i, j)*element)
		}
	}

	return r
}

func (m *denseMatrix) Scalar(s float64) types.Matrix {
	for index, element := range m.elements {
		m.elements[index] = element * s
	}

	return m
}

func (m *denseMatrix) Transpose() types.Matrix {
	n := &denseMatrix{
		base:     m.base,
		view:     m.view,
		offset:   m.offset,
		elements: m.elements,
		rewriter: m.rewriter.Transpose(),
	}

	return n
}

func (m *denseMatrix) View(row, column, rows, columns int) types.Matrix {
	offset := types.NewIndex(m.offset.Row()+row, m.offset.Column()+column)
	view := types.NewShape(rows, columns)

	validates.ViewShouldBeInBase(m.base, view, offset)

	n := &denseMatrix{
		base:     m.base,
		view:     view,
		offset:   offset,
		elements: m.elements,
		rewriter: m.rewriter,
	}

	return n
}

func (m *denseMatrix) Base() types.Matrix {
	n := &denseMatrix{
		base:     m.base,
		view:     m.base,
		offset:   m.offset,
		elements: m.elements,
		rewriter: m.rewriter,
	}

	return n
}

func (m *denseMatrix) Row(row int) types.Matrix {
	return m.View(row, 0, 1, m.view.Columns())
}

func (m *denseMatrix) Column(column int) types.Matrix {
	return m.View(0, column, m.view.Rows(), 1)
}

func (m *denseMatrix) Max() (element float64, row, column int) {
	max := math.Inf(-1)
	index := 0

	for i, c := range m.elements {
		if max >= c {
			continue
		}

		max = c
		index = i
	}

	row, column = m.convertToRowColumn(index)

	return max, row, column
}

func (m *denseMatrix) Min() (element float64, row, column int) {
	min := math.Inf(0)
	index := 0

	for i, c := range m.elements {
		if min <= c {
			continue
		}

		min = c
		index = i
	}

	row, column = m.convertToRowColumn(index)

	return min, row, column
}

func (m *denseMatrix) convertToRowColumn(index int) (row, column int) {
	columns := m.Columns()

	row = index / columns
	column = index - (row * columns)

	return row, column
}
