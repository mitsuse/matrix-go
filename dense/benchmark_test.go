package dense

import (
	"testing"
)

func BenchmarkGet(b *testing.B) {
	m := Zeros(8, 8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Get(0, 0)
	}
}

func BenchmarkUpdate(b *testing.B) {
	m := Zeros(8, 8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Update(0, 0, 1)
	}
}

func BenchmarkAddition(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		m := New(4, 3)(
			0, -1, 2, 3,
			4.1, 5, -6, 7.4,
			-8, 9.2, 0, -1.1,
		)

		n := New(4, 3)(
			0, 5.3, 1.7, -0.1,
			2.7, -8, 0, 1.1,
			3.4, -0.123, 99, 2.3141,
		)

		b.StartTimer()

		m.Add(n)
	}
}

func BenchmarkSubtraction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		m := New(4, 3)(
			0, -1, 2, 3,
			4.1, 5, -6, 7.4,
			-8, 9.2, 0, -1.1,
		)

		n := New(4, 3)(
			0, 5.3, 1.7, -0.1,
			2.7, -8, 0, 1.1,
			3.4, -0.123, 99, 2.3141,
		)

		b.StartTimer()

		m.Subtract(n)
	}
}

func BenchmarkDot(b *testing.B) {
	m := New(4, 3)(
		0, -1, 2, 3,
		4.1, 5, -6, 7.4,
		-8, 9.2, 0, -1.1,
	)

	n := New(3, 4)(
		0, 2.7, -3.4,
		5.3, -8, -0.123,
		1.7, 0, 99,
		-0.1, 1.1, 2.3141,
	)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Dot(n)
	}
}

func BenchmarkMultiply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		m := New(4, 3)(
			0, -1, 2, 3,
			4.1, 5, -6, 7.4,
			-8, 9.2, 0, -1.1,
		)

		s := 0.1

		b.StartTimer()

		m.Multiply(s)
	}
}
