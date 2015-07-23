package types

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

	if isZeros := IsZeros(m); !isZeros {
		t.Error("This matrix should be zeros.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsNotZerosMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		0, 1, 2,
		4, 5, 0,
		2, 3, 4,
		0, 1, 2,
	)

	if isZeros := IsZeros(m); isZeros {
		t.Error("This matrix should not be zeros.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsSquareMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		0, 1, 2, 3,
		4, 5, 0, 1,
		2, 3, 4, 5,
		0, 1, 2, 3,
	)

	if isSquare := IsSquare(m); !isSquare {
		t.Error("This matrix should be square.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsNotSquareMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		0, 1, 2,
		4, 5, 0,
		2, 3, 4,
		0, 1, 2,
	)

	if isSquare := IsSquare(m); isSquare {
		t.Error("This matrix should not be square.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsDiagonalMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		2, 0, 0, 0,
		0, 4, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 0,
	)

	if isDiagonal := IsDiagonal(m); !isDiagonal {
		t.Error("This matrix should be diagonal.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsNotDiagonalMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		0, 1, 1, 1,
		1, 0, 1, 1,
		1, 1, 0, 1,
		1, 1, 1, 0,
	)

	if isDiagonal := IsDiagonal(m); isDiagonal {
		t.Error("This matrix should not be diagonal.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsNotDiagonalNonSquareMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		2, 0, 0,
		0, 4, 0,
		0, 0, 1,
		0, 0, 0,
	)

	if isDiagonal := IsDiagonal(m); isDiagonal {
		t.Error("This matrix should not be diagonal.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsIdentityMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	if isIdentity := IsIdentity(m); !isIdentity {
		t.Error("This matrix should be identity.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsNotIdentityMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		2, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	)

	if isIdentity := IsIdentity(m); isIdentity {
		t.Error("This matrix should not be identity.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsNotIdentityNonSquareMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
		0, 0, 0,
	)

	if isIdentity := IsIdentity(m); isIdentity {
		t.Error("This matrix should not be identity.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsScalarMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		7, 0, 0, 0,
		0, 7, 0, 0,
		0, 0, 7, 0,
		0, 0, 0, 7,
	)

	if isScalar := IsScalar(m); !isScalar {
		t.Error("This matrix should be scalar.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsNotScalarMutableDense(t *testing.T) {
	m := dense.New(4, 4)(
		7, 0, 0, 0,
		0, 7, 0, 0,
		0, 0, 7, 0,
		0, 0, 0, 6,
	)

	if isScalar := IsScalar(m); isScalar {
		t.Error("This matrix should not be scalar.")
		t.Fatalf("# matrix = %+v", m)
	}
}

func TestIsNotScalarNonSquareMutableDense(t *testing.T) {
	m := dense.New(4, 3)(
		7, 0, 0,
		0, 7, 0,
		0, 0, 7,
		0, 0, 0,
	)

	if isScalar := IsScalar(m); isScalar {
		t.Error("This matrix should not be scalar.")
		t.Fatalf("# matrix = %+v", m)
	}
}
