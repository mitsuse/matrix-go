/*
Package "dense" provides an implementation of mutable dense matrix.
*/
package dense

import (
	"encoding/json"
	"errors"
	"io"
	"math"

	"github.com/mitsuse/matrix-go/internal/rewriters"
	"github.com/mitsuse/matrix-go/internal/types"
	"github.com/mitsuse/matrix-go/internal/validates"
)

type Matrix struct {
	initialized bool
	base        *types.Shape
	view        *types.Shape
	offset      *types.Index
	elements    []float64
	rewriter    rewriters.Rewriter
}

// Create a new matrix with given elements.
// When "rows" and "columns" is not positive,
// validates.NON_POSITIVE_SIZE_PANIC will be caused.
// In addition,
// when the product of "row"s and "column" doesn't equal to the size of "elements",
// validates.INVALID_ELEMENTS_PANIC will be caused.
func New(rows, columns int) func(elements ...float64) *Matrix {
	validates.ShapeShouldBePositive(rows, columns)

	constructor := func(elements ...float64) *Matrix {
		size := rows * columns

		if len(elements) != size {
			panic(validates.INVALID_ELEMENTS_PANIC)
		}

		shape := types.NewShape(rows, columns)
		offset := types.NewIndex(0, 0)

		m := &Matrix{
			initialized: true,
			base:        shape,
			view:        shape,
			offset:      offset,
			elements:    make([]float64, size),
			rewriter:    rewriters.Reflect(),
		}
		copy(m.elements, elements)

		return m
	}

	return constructor
}

// Create a new zero matrix.
// When "rows" and "columns" is not positive,
// validates.NON_POSITIVE_SIZE_PANIC will be caused.
func Zeros(rows, columns int) *Matrix {
	return New(rows, columns)(make([]float64, rows*columns)...)
}

// Convert the given matrix to *dense.Matrix.
// If the given matrix is already typed as *dense.Matrix, just returns it.
// In other cases, create a new matrix.
func Convert(m types.Matrix) *Matrix {
	d, isDense := m.(*Matrix)

	if isDense {
		return d
	}

	// TODO: Convert the other type of matrix to *Matrix.

	return nil
}

// Deserialize a matrix from the given reader.
func Deserialize(reader io.Reader) (types.Matrix, error) {
	m := &Matrix{}

	if err := json.NewDecoder(reader).Decode(m); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Matrix) Serialize(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(m)
}

func (m *Matrix) MarshalJSON() ([]byte, error) {
	jsonObject := matrixJson{
		Version:  version,
		Base:     m.base,
		View:     m.view,
		Offset:   m.offset,
		Elements: m.elements,
		Rewriter: m.rewriter.Type(),
	}

	return json.Marshal(&jsonObject)
}

func (m *Matrix) UnmarshalJSON(b []byte) error {
	if m.initialized {
		return errors.New(AlreadyInitializedError)
	}

	jsonObject := &matrixJson{}

	if err := json.Unmarshal(b, jsonObject); err != nil {
		return err
	}

	if jsonObject.Version < minVersion || maxVersion < jsonObject.Version {
		return errors.New(IncompatibleVersionError)
	}

	m.base = jsonObject.Base
	m.view = jsonObject.View
	m.offset = jsonObject.Offset
	m.elements = jsonObject.Elements

	rewriter, err := rewriters.Get(jsonObject.Rewriter)
	if err != nil {
		return err
	}
	m.rewriter = rewriter

	// TODO: Return error value instead of causing panic.
	validates.ShapeShouldBePositive(m.base.Rows(), m.base.Columns())

	validates.IndexShouldBeInRange(
		m.base.Rows(),
		m.base.Columns(),
		m.offset.Row(),
		m.offset.Column(),
	)

	validates.ViewShouldBeInBase(m.base, m.view, m.offset)

	m.initialized = true

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
	return newDiagonalCursor(m)
}

func (m *Matrix) Get(row, column int) (element float64) {
	row, column = m.rewriter.Rewrite(row, column)

	validates.IndexShouldBeInRange(m.view.Rows(), m.view.Columns(), row, column)

	index := (row+m.offset.Row())*m.base.Columns() + column + m.offset.Column()

	return m.elements[index]
}

func (m *Matrix) Update(row, column int, element float64) types.Matrix {
	row, column = m.rewriter.Rewrite(row, column)

	validates.IndexShouldBeInRange(m.view.Rows(), m.view.Columns(), row, column)

	index := (row+m.offset.Row())*m.base.Columns() + column + m.offset.Column()
	m.elements[index] = element

	return m
}

func (m *Matrix) Equal(n types.Matrix) bool {
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

func (m *Matrix) Add(n types.Matrix) types.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		m.Update(row, column, m.Get(row, column)+element)
	}

	return m
}

func (m *Matrix) Subtract(n types.Matrix) types.Matrix {
	validates.ShapeShouldBeSame(m, n)

	cursor := n.NonZeros()

	for cursor.HasNext() {
		element, row, column := cursor.Get()
		m.Update(row, column, m.Get(row, column)-element)
	}

	return m
}

func (m *Matrix) Multiply(n types.Matrix) types.Matrix {
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

func (m *Matrix) Scalar(s float64) types.Matrix {
	for index, element := range m.elements {
		m.elements[index] = element * s
	}

	return m
}

func (m *Matrix) Transpose() types.Matrix {
	n := &Matrix{
		initialized: true,
		base:        m.base,
		view:        m.view,
		offset:      m.offset,
		elements:    m.elements,
		rewriter:    m.rewriter.Transpose(),
	}

	return n
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
	max := math.Inf(-1)
	index := types.NewIndex(0, 0)

	cursor := m.All()

	for cursor.HasNext() {
		element, row, column := cursor.Get()

		if max >= element {
			continue
		}

		max = element
		index = types.NewIndex(row, column)
	}

	return max, index.Row(), index.Column()
}

func (m *Matrix) Min() (element float64, row, column int) {
	max := math.Inf(1)
	index := types.NewIndex(0, 0)

	cursor := m.All()

	for cursor.HasNext() {
		element, row, column := cursor.Get()

		if max <= element {
			continue
		}

		max = element
		index = types.NewIndex(row, column)
	}

	return max, index.Row(), index.Column()
}
