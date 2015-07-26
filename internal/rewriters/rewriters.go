/*
Package "rewriters" provides functions to rewrite the indexes or shape of matrix.
*/
package rewriters

var (
	reflect *reflectImpl
	reverse *reverseImpl
)

func init() {
	reflect = &reflectImpl{}
	reverse = &reverseImpl{}
}

func Reflect() Rewriter {
	return reflect
}

func Reverse() Rewriter {
	return reverse
}

type Rewriter interface {
	Rewrite(int, int) (int, int)
	Transpose() Rewriter
}

type reflectImpl struct {
}

func (r *reflectImpl) Rewrite(x, y int) (int, int) {
	return x, y
}

func (r *reflectImpl) Transpose() Rewriter {
	return Reverse()
}

type reverseImpl struct {
}

func (r *reverseImpl) Rewrite(x, y int) (int, int) {
	return y, x
}

func (r *reverseImpl) Transpose() Rewriter {
	return Reflect()
}
