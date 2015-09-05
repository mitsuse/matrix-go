/*
Package "hash" provides an hash-based implementation of mutable sparse matrix.
*/
package hash

import (
	"encoding/json"
	"io"

	"github.com/mitsuse/matrix-go/internal/rewriters"
	"github.com/mitsuse/matrix-go/internal/types"
	"github.com/mitsuse/matrix-go/internal/validates"
)

type Matrix struct {
	initialized bool
	base        *types.Shape
	view        *types.Shape
	offset      *types.Index
	elements    map[int]float64
	rewriter    rewriters.Rewriter
}

type Element struct {
	Row    int
	Column int
	Value  float64
}

// Create a new matrix with given elements.
// When "rows" and "columns" is not positive,
// validates.NON_POSITIVE_SIZE_PANIC will be caused.
// In addition,
// when each element has the index out of shape,
// validates.INVALID_ELEMENTS_PANIC will be caused.
func New(rows, columns int) func(elements ...Element) *Matrix {
	validates.ShapeShouldBePositive(rows, columns)

	constructor := func(elements ...Element) *Matrix {
		shape := types.NewShape(rows, columns)
		offset := types.NewIndex(0, 0)

		m := &Matrix{
			initialized: true,
			base:        shape,
			view:        shape,
			offset:      offset,
			elements:    make(map[int]float64),
			rewriter:    rewriters.Reflect(),
		}

		for _, element := range elements {
			key := element.Row*columns + element.Column
			if _, exist := m.elements[key]; exist {
				panic(validates.INVALID_ELEMENTS_PANIC)
			}

			m.elements[key] = element.Value
		}

		return m
	}

	return constructor
}

// Create a new zero matrix.
// When "rows" and "columns" is not positive,
// validates.NON_POSITIVE_SIZE_PANIC will be caused.
func Zeros(rows, columns int) *Matrix {
	return New(rows, columns)()
}

// Convert the given matrix to *hash.Matrix.
// If the given matrix is already typed as *hash.Matrix, just returns it.
// In other cases, create a new matrix.
func Convert(m types.Matrix) *Matrix {
	// TODO: Implement.
	// d, isHash := m.(*Matrix)
	//
	// if isHash {
	// 	return d
	// }
	//
	// // TODO: Convert the other type of matrix to *Matrix.
	//
	return nil
}

// Deserialize a matrix from the given reader.
func Deserialize(reader io.Reader) (types.Matrix, error) {
	m := &Matrix{}

	if err := json.NewDecoder(reader).Decode(m); err != nil {
		return nil, err
	}

	// return m, nil
	// TODO: Implement.
	return nil, nil
}

func (m *Matrix) Serialize(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(m)
}

func (m *Matrix) MarshalJSON() ([]byte, error) {
	// TODO: Implement.
	return nil, nil
}

func (m *Matrix) UnmarshalJSON([]byte) error {
	// TODO: Implement.
	return nil
}

func (m *Matrix) Shape() (rows, columns int) {
	return m.rewriter.Rewrite(m.view.Rows(), m.view.Columns())
}

func (m *Matrix) Rows() (rows int) {
	rows, _ = m.Shape()
	return rows
}

func (m *Matrix) Columns() (columns int) {
	_, columns = m.Shape()
	return columns
}

func (m *Matrix) All() types.Cursor {
	return newAllCursor(m)
}

func (m *Matrix) NonZeros() types.Cursor {
	return newNonZerosCursor(m)
}

func (m *Matrix) Diagonal() types.Cursor {
	// TODO: Implement.
	return nil
}

func (m *Matrix) Get(row, column int) (element float64) {
	row, column = m.rewriter.Rewrite(row, column)

	validates.IndexShouldBeInRange(m.view.Rows(), m.view.Columns(), row, column)

	index := (row+m.offset.Row())*m.base.Columns() + column + m.offset.Column()

	element, exist := m.elements[index]
	if !exist {
		return 0
	}

	return element
}

func (m *Matrix) Update(row, column int, element float64) types.Matrix {
	row, column = m.rewriter.Rewrite(row, column)

	validates.IndexShouldBeInRange(m.view.Rows(), m.view.Columns(), row, column)

	index := (row+m.offset.Row())*m.base.Columns() + column + m.offset.Column()

	if element == 0 {
		delete(m.elements, index)
	} else {
		m.elements[index] = element
	}

	return m
}

func (m *Matrix) Equal(n types.Matrix) bool {
	// TODO: Implement.
	return true
}

func (m *Matrix) Add(n types.Matrix) types.Matrix {
	// TODO: Implement.
	return nil
}

func (m *Matrix) Subtract(n types.Matrix) types.Matrix {
	// TODO: Implement.
	return nil
}

func (m *Matrix) Multiply(n types.Matrix) types.Matrix {
	// TODO: Implement.
	return nil
}

func (m *Matrix) Scalar(s float64) types.Matrix {
	// TODO: Implement.
	return nil
}

func (m *Matrix) Transpose() types.Matrix {
	// TODO: Implement.
	return nil
}

func (m *Matrix) View(row, column, rows, columns int) types.Matrix {
	row, column = m.rewriter.Rewrite(row, column)
	rows, columns = m.rewriter.Rewrite(rows, columns)

	offset := types.NewIndex(m.offset.Row()+row, m.offset.Column()+column)
	view := types.NewShape(rows, columns)

	validates.ShapeShouldBePositive(rows, columns)
	validates.ViewShouldBeInBase(m.base, view, offset)

	n := &Matrix{
		initialized: true,
		base:        m.base,
		view:        view,
		offset:      offset,
		elements:    m.elements,
		rewriter:    m.rewriter,
	}

	return n
}

func (m *Matrix) Base() types.Matrix {
	n := &Matrix{
		initialized: true,
		base:        m.base,
		view:        m.base,
		offset:      types.NewIndex(0, 0),
		elements:    m.elements,
		rewriter:    m.rewriter,
	}

	return n
}

func (m *Matrix) Row(row int) types.Matrix {
	return m.View(row, 0, 1, m.view.Columns())
}

func (m *Matrix) Column(column int) types.Matrix {
	return m.View(0, column, m.view.Rows(), 1)
}

func (m *Matrix) Max() (element float64, row, column int) {
	// TODO: Implement.
	return 0, 0, 0
}

func (m *Matrix) Min() (element float64, row, column int) {
	// TODO: Implement.
	return 0, 0, 0
}
