# Matrix

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
[![Version](https://img.shields.io/github/tag/mitsuse/matrix-go.svg?style=flat-square)][release]
[![Wercker](http://img.shields.io/wercker/ci/55672222ee357fac39001a2a.svg?style=flat-square)][wercker]
[![Coverage](https://img.shields.io/codecov/c/github/mitsuse/matrix-go/develop.svg?style=flat-square)][coverage]

[license]: LICENSE.txt
[godoc]: http://godoc.org/github.com/mitsuse/matrix-go
[release]: https://github.com/mitsuse/matrix-go/releases
[wercker]: https://app.wercker.com/project/bykey/093a5cff0964f0f4ba5fcf9117e940e4
[coverage]: https://codecov.io/github/mitsuse/matrix-go

An experimental library for matrix manipulation implemented in [Golang][golang].

[golang]: http://golang.org/


## Motivations

1. Portability - Implement in pure Golang to achieve cgo-free.
1. Efficiency - Pursue performance as possible without highly optimized back-ends like blas.
1. Simplicity - Provide clean API.


## Installation

For installation, execute the following command:

```
$ go get github.com/mitsuse/matrix-go
```


## Features

### Matrix Types

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

When the receiver is mutable,
`(Matrix).Add` and `(Matrix).Subtract` return the receiver itself,
the elements of which is rewritten.


#### Matrix Multiplication

The product of two matrices can be calculated by `(Matrix).Multiply`.

```go
m := dense.New(3, 2)(
    0, 1,
    2, 3,
    4, 5,
)

n := dense.New(2, 1)(
    0,
    -1,
)

r := dense.New(3, 1)(
    -1,
    -3,
    -5,
)

m.Multiply(n).Equal(r)
```

Matrix multiplication always create a new matrix.
The type of the result matrix is same as the type of the receiver.


#### Scalar Multiplication

`(Matrix).Scalar` is available for Scalar multiplication (scalar-left multiplication).

```go
m := dense.New(2, 2)(
    0, 1,
    2, 3,
)

r := dense.New(2, 2)(
    0, -1,
    -2, -3,
)

// true
m.Scalar(-1).Equal(r)
```

For scalar-right multiplication, use `(Scalar).Multiply`.

```go
m := dense.New(2, 2)(
    0, 1,
    2, 3,
)

r := dense.New(2, 2)(
    0, -1,
    -2, -3,
)

// true
Scalar(-1).Multiply(m).Equal(r)
```

When the matrix used for scalar multiplication is mutable,
`(Matrix).Scalar` and `(Scalar).Multiply` rewrite elements of the matrix.


### Cursor

`Matrix` has several methods to iterate elements.
They return a value typed as `Cursor` which is a reference to the element to visit.

```go
m := dense.New(2, 3)(
    0, 1, 2,
    3, 4, 5,
)

// Create a cursor to iterate all elements of matrix m.
c := m.All()

// Check whether the element to visit exists or not.
for c.HasNext() {
    element, row, column := c.Get()

    fmt.Printf(
        "element = %d, row = %d, column = %d\n",
        element,
        row,
        column,
    )
}
```

Currently, three methods are implemented which return a cursor:

- `(Matrix).All`
- `(Matrix).NonZeros`
- `(Matrix).Diagonal`

For details, please read the documentation of
[`types.Matrix`](http://godoc.org/github.com/mitsuse/matrix-go/internal/types/#Matrix).


## More Details

Please read the [documentation][godoc].


## License

Please read [LICENSE.txt](LICENSE.txt).
