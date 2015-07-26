package matrix

/*
"Scalar" is used to calculate the product with scalar-left multiplication.
This type is underlying float64.
*/
type Scalar float64

func (s Scalar) Multiply(m Matrix) Matrix {
	return m.Scalar(float64(s))
}
