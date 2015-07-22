/*
Package "rewriters" provideds functions to rewrite the indexes or shape of matrix.
*/
package rewriters

type Rewrite func(int, int) (int, int)

func Reflect(x, y int) (int, int) {
	return x, y
}

func Reverse(x, y int) (int, int) {
	return y, x
}
