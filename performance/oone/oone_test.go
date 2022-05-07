package oone

import "testing"

func BenchmarkThreeWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ThreeWords()
	}
}

func BenchmarkTenWords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TenWords()
	}
}
