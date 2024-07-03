package main

import "testing"

func BenchmarkMathBigBasicAddition(b *testing.B) {
	for i := 0; i < 1000; i++ {
		MathBigBasicAddition("4.2938293734535", "0.00034345354345353534535")
}
}

func BenchmarkDecimalAdd(b *testing.B) {
	for i := 0; i < 100; i++ {
		DecimalAdd("4.2938293734535", "0.00034345354345353534535") 
	}
}

func BenchmarkMathBigMul(b *testing.B) {
	for i := 0; i < 100; i++ {
		MathBigBasicMultiplication("4.29382937", "0.00034") 
	}
}

func BenchmarkDecimalMul(b *testing.B) {
	for i := 0; i < 100; i++ {
		DecimalMul("4.29382937", "0.00034") 
	}
}

func BenchmarkCockroachAdd(b *testing.B) {
	for i := 0; i < 100; i++ {
		CockroachAdd("4.2938293734535", "0.00034345354345353534535") 
	}
}

func BenchmarkCockroachMul(b *testing.B) {
	for i := 0; i < 100; i++ {
		CockroachMul("4.29382937", "0.00034")
	}
}