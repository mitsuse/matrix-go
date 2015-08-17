/*
Package "rewriters" provides functions to rewrite the indexes or shape of matrix.
*/
package rewriters

import (
	"errors"
	"io"

	"github.com/mitsuse/serial-go"
)

const (
	id      string = "github.com/mitsuse/matrix-go/internal/rewriters"
	version byte   = 0
)

const (
	typeReflect byte = iota
	typeReverse
)

const (
	UNKNOWN_REWRITER_ERROR = "UNKNOWN_REWRITER_ERROR"
)

var (
	reflect *reflectImpl
	reverse *reverseImpl
)

func init() {
	reflect = &reflectImpl{}
	reverse = &reverseImpl{}
}

func Get(typeId byte) (Rewriter, error) {
	switch typeId {
	case typeReflect:
		return Reflect(), nil
	case typeReverse:
		return Reverse(), nil
	}

	return nil, errors.New(UNKNOWN_REWRITER_ERROR)
}

func Reflect() Rewriter {
	return reflect
}

func Reverse() Rewriter {
	return reverse
}

// Deserialize a rewriter from the given reader.
// This accepts data generated with (Rewriter).Serialize.
func Deserialize(reader io.Reader) (Rewriter, error) {
	var rewriterType byte

	r := serial.NewReader(id, version, reader)

	r.ReadId()
	r.ReadVersion()
	r.Read(&rewriterType)

	if err := r.Error(); err != nil {
		return nil, err
	}

	switch rewriterType {
	case typeReflect:
		return Reflect(), nil
	case typeReverse:
		return Reverse(), nil
	}

	return nil, errors.New(UNKNOWN_REWRITER_ERROR)
}

type Rewriter interface {
	Type() byte
	Serialize(writer io.Writer) error
	Rewrite(int, int) (int, int)
	Transpose() Rewriter
}

type reflectImpl struct {
}

func (r *reflectImpl) Type() byte {
	return typeReflect
}

func (r *reflectImpl) Serialize(writer io.Writer) error {
	w := serial.NewWriter(id, version, writer)

	w.WriteId()
	w.WriteVersion()

	w.Write(typeReflect)

	return w.Error()
}

func (r *reflectImpl) Rewrite(x, y int) (int, int) {
	return x, y
}

func (r *reflectImpl) Transpose() Rewriter {
	return Reverse()
}

type reverseImpl struct {
}

func (r *reverseImpl) Type() byte {
	return typeReverse
}

func (r *reverseImpl) Serialize(writer io.Writer) error {
	w := serial.NewWriter(id, version, writer)

	w.WriteId()
	w.WriteVersion()

	w.Write(typeReverse)

	return w.Error()
}

func (r *reverseImpl) Rewrite(x, y int) (int, int) {
	return y, x
}

func (r *reverseImpl) Transpose() Rewriter {
	return Reflect()
}
