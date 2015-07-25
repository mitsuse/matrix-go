package matrix

type Scalar float64

func (s Scalar) Multiply(m Matrix) Matrix {
	return m.Multiply(s)
}
