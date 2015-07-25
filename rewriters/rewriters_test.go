package rewriters

import (
	"testing"
)

type testPair struct {
	X int
	Y int
}

func TestReflect(t *testing.T) {
	test := testPair{
		X: 1,
		Y: 2,
	}

	if x, y := Reflect().Rewrite(test.X, test.Y); x == test.X && y == test.Y {
		return
	}

	t.Fatal("The result pair should equal to the input pair.")
}

func TestReverse(t *testing.T) {
	test := testPair{
		X: 1,
		Y: 2,
	}

	if x, y := Reverse().Rewrite(test.X, test.Y); x == test.Y && y == test.X {
		return
	}

	t.Fatal("The result pair should be reversed.")
}

func TestReflectTranspose(t *testing.T) {
	if Reflect().Transpose() == Reverse() {
		return
	}

	t.Fatal("The transpose of Reflect should be Reverse.")
}

func TestReverseTranspose(t *testing.T) {
	if Reverse().Transpose() == Reflect() {
		return
	}

	t.Fatal("The transpose of Reverse should be Reflect.")
}
