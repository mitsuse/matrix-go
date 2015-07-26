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


## Features

### Implementations

Currently, the following types are implemented:

- mutable dense matrix


### Creation

Use `dense.New` to create a new dense matrix with given elements.

```go
// Create a 2 x 3 matrix.
m := dense.New(2, 3)(
    0, 1, 2,
    3, 4, 5,
)
```

To create zero matrix, call `dense.Zeros` instead.

```go
// Create a 2 x 3 zero matrix.
m := dense.Zeros(2, 3)
```


### Operations

#### Addition & Subtraction

Add a matrix to other with `(Matrix).Add`:

```go
m := dense.New(2, 3)(
    0, 1, 2,
    3, 4, 5,
)

n := dense.New(2, 3)(
    5, 4, 3,
    2, 1, 0,
)

r := dense.New(2, 3)(
    5, 5, 5,
    5, 5, 5,
)

// true
m.Add(n).Equal(r)
```

Similarly, `(Matrix).Subtract` is used for subtraction on two matrix.


#### Scalar Multiplication

`(Matrix).Multiply` is available for Scalar multiplication (scalar-left multiplication).

```go
m := dense.New(2, 2)(
    0, 1,
    2, 3,
)

r := dense.New(2, 2)(
    0, 2,
    4, 6,
)

// true
m.Multiply(-1).Equal(r)
```

For scalar-right multiplication, use `(Scalar).Multiply`.

```go
m := dense.New(2, 2)(
    0, 1,
    2, 3,
)

r := dense.New(2, 2)(
    0, 2,
    4, 6,
)

// true
Scalar(-1).Multiply(m).Equal(r)
```


### Cursor

`Matrix` has several methods to iterate elements.
They return a value typed as `Cursor` which is a refernce the element to visit.

```go
m := dense.New(2, 3)(
    0, 1, 2,
    3, 4, 5,
)

c := m.All()

for c.HasNext() {
    element, row, column := c.Get()

    fmt.Printf(
        "element = %d, row = $d, column = %d",
        element,
        row,
        column,
    )
}
```

Currently, three methods are implemented as follow:

- `(Matrix).All`
- `(Matrix).NonZeros`
- `(Matrix).Diagonals


### More Details

Please read the [documentation][godoc].
