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

### Matrix Types

Currently, the following types are implemented:

- mutable dense matrix


### More Details

Please read the [documentation][godoc].
