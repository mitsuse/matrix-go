package matrix

import (
	"testing"

	"github.com/mitsuse/matrix-go/dense"
)

func TestScaleMultiplyIsCommutative(t *testing.T) {
	m := dense.New(2, 2)(
		0, 1,
		2, 3,
	)

	n := dense.New(2, 2)(
		0, 1,
		2, 3,
	)

	s := 2.0

	if Scalar(s).Scale(m).Equal(n.Scale(s)) {
		return
	}

	t.Fatal("Scalar multiplication should be commutative.")
}
