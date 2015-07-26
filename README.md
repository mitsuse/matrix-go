# Matrix

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
![Version](https://img.shields.io/github/tag/mitsuse/matrix-go-green.svg?style=flat-square)
[![Wercker](http://img.shields.io/wercker/ci/55672222ee357fac39001a2a.svg?style=flat-square)][wercker]
[![Coverage](https://img.shields.io/codecov/c/github/mitsuse/matrix-go/develop.svg?style=flat-square)][coverage]

[license]: LICENSE.txt
[godoc]: http://godoc.org/github.com/mitsuse/matrix-go
[wercker]: https://app.wercker.com/project/bykey/093a5cff0964f0f4ba5fcf9117e940e4
[coverage]: https://codecov.io/github/mitsuse/matrix-go

An experimental library for matrix manipulation implemented in Golang.


## Motivations

1. Portability - Implement in pure Golang to achieve cgo-free.
1. Efficiency - Pursue performance as possible without highly optimized back-ends like blas.
1. Simplicity - Provide clean API. It is also not trivial.


## Installation

For installation, execute the following command:

```
$ go get github.com/mitsuse/matrix-go
```


## Example

```go
package main

import (
	"fmt"

	. "github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
)

func main() {
	// Create 3x3 zero matrix.
	m := dense.Zeros(3, 3)

	// true
	fmt.Println("IsZeros(m) =", IsZeros(m))

	// Update the element at (1, 0) to 1.
	m.Update(1, 0, 1)

	// false
	fmt.Println("IsZeros(m) =", IsZeros(m))

	// Create (3x3) matrix with given elements.
	n := dense.New(3, 3)(
		1, 0, 0,
		0, 1, 0,
		0, 0, 1,
	)

	// Add n to m.
	m.Add(n)

	r1 := dense.New(3, 3)(
		1, 0, 0,
		1, 1, 0,
		0, 0, 1,
	)

	// true
	fmt.Println("m.Equal(r1) =", m.Equal(r1))

	// Caluculate scalar-right multiplication.
	m.Multiply(2)

	r2 := dense.New(3, 3)(
		2, 0, 0,
		2, 2, 0,
		0, 0, 2,
	)

	// true
	fmt.Println("m.Equal(r2) =", m.Equal(r2))

	// Caluculate scalar-left multiplication.
	Scalar(0.5).Multiply(m)

	// true
	fmt.Println("m.Equal(r1) =", m.Equal(r1))

	// Operations are chainable.
	m.Multiply(-1).Add(n).Subtract(Scalar(0.5).Multiply(n))

	r3 := dense.New(3, 3)(
		-0.5, 0, 0,
		-1, -0.5, 0,
		0, 0, -0.5,
	)

	// true
	fmt.Println("m.Equal(r3) =", m.Equal(r3))
}
```
