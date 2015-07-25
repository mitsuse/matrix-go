package feature

import (
	"testing"

	"github.com/mitsuse/matrix-go/mutable/dense"
)

func TestIsZerosMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
		0, 0, 0,
	)

	if IsZeros(m) {
		return
	}

	t.Fatal("This matrix should be zeros.")
}

func TestIsNotZerosMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		0, 1, 2,
		4, 5, 0,
		2, 3, 4,
		0, 1, 2,
	)

	if !IsZeros(m) {
		return
	}

	t.Fatal("This matrix should not be zeros.")
}

func TestIsSquareMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		0, 1, 2, 3,
		4, 5, 0, 1,
		2, 3, 4, 5,
		0, 1, 2, 3,
	)

	if IsSquare(m) {
		return
	}

	t.Fatal("This matrix should be square.")
}

func TestIsNotSquareMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		0, 1, 2,
		4, 5, 0,
		2, 3, 4,
		0, 1, 2,
	)

	if !IsSquare(m) {
		return
	}

	t.Fatal("This matrix should not be square.")
}

func TestIsDiagonalMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		2, 0, 0, 0,
		0, 4, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 0,
	)

	if IsDiagonal(m) {
		return
	}

	t.Fatal("This matrix should be diagonal.")
}

func TestIsNotDiagonalMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		0, 1, 1, 1,
		1, 0, 1, 1,
		1, 1, 0, 1,
		1, 1, 1, 0,
	)

	if !IsDiagonal(m) {
		return
	}

	t.Fatal("This matrix should not be diagonal.")
}

func TestIsNotDiagonalNonSquareMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		2, 0, 0,
		0, 4, 0,
		0, 0, 1,
		0, 0, 0,
	)

	if !IsDiagonal(m) {
		return
	}

	t.Fatal("This matrix should not be diagonal.")
}

func TestIsIdentityMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	if IsIdentity(m) {
		return
	}

	t.Fatal("This matrix should be identity.")
}

func TestIsNotIdentityMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		2, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	if !IsIdentity(m) {
		return
	}

	t.Fatal("This matrix should not be identity.")
}

func TestIsNotIdentityNonSquareMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
		0, 0, 0,
	)

	if !IsIdentity(m) {
		return
	}

	t.Fatal("This matrix should not be identity.")
}

func TestIsScalarMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		7, 0, 0, 0,
		0, 7, 0, 0,
		0, 0, 7, 0,
		0, 0, 0, 7,
	)

	if IsScalar(m) {
		return
	}

	t.Fatal("This matrix should be scalar.")
}

func TestIsNotScalarMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		7, 0, 0, 0,
		0, 7, 0, 0,
		0, 0, 7, 0,
		0, 0, 0, 6,
	)

	if !IsScalar(m) {
		return
	}

	t.Fatal("This matrix should not be scalar.")
}

func TestIsNotScalarNonSquareMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		7, 0, 0,
		0, 7, 0,
		0, 0, 7,
		0, 0, 0,
	)

	if !IsScalar(m) {
		return
	}

	t.Fatal("This matrix should not be scalar.")
}
