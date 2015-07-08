# Matrix

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
[![Wercker](http://img.shields.io/wercker/ci/55672222ee357fac39001a2a.svg?style=flat-square)][wercker]
[![Coverage](https://img.shields.io/codecov/c/github/mitsuse/matrix-go/develop.svg?style=flat-square)][coverage]

[license]: http://opensource.org/licenses/MIT
[godoc]: http://godoc.org/github.com/mitsuse/matrix-go
[wercker]: https://app.wercker.com/project/bykey/093a5cff0964f0f4ba5fcf9117e940e4
[coverage]: https://codecov.io/github/mitsuse/matrix-go

An experimental library for matrix manipulation implemented in Golang.


## Motivations

1. Portability - Implement in pure Golang to achieve cgo-free.
1. Efficiency - Pursue performance as possible without highly optimized back-ends like blas.
1. Simplicity - Provide clean API. It is also not trivial.


## License

The MIT License (MIT)

Copyright (c) 2015 Tomoya Kose.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
