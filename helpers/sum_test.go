package helpers

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test MAIN
func TestMain(m *testing.M) {
	fmt.Println("Before Main Test")

	m.Run()

	fmt.Println("After Main Test")
}

func TestSum(t *testing.T) {
	result := Sum(2, 2)

	if result != 4 {
		t.Error("Result is not 4")
	}
}

func TestSum6(t *testing.T) {
	result := Sum(2, 4)

	assert.Equal(t, 6, result, "Result Must 6")

	fmt.Println("Done")
}

func TestSum8(t *testing.T) {
	result := Sum(3, 5)

	require.Equal(t, 8, result, "Result Must 8")
}

// SUB TEST
func TestSubTest(t *testing.T) {
	t.Run("Sum 8", func(t *testing.T) {
		result := Sum(4, 4)

		require.Equal(t, 8, result, "Result Must 8")
	})

	t.Run("Sum 10", func(t *testing.T) {
		result := Sum(4, 6)

		require.Equal(t, 10, result, "Result Must 10")
	})
}

// Table Test
func TestTableSum(t *testing.T) {
	tests := []struct {
		name     string
		param1   int
		param2   int
		expected int
	}{
		{name: "sum 10",
			param1:   3,
			param2:   7,
			expected: 10},
		{name: "sum 14",
			param1:   5,
			param2:   9,
			expected: 14},
		{name: "sum 25",
			param1:   16,
			param2:   9,
			expected: 25},
	}

	for _, test := range tests {
		t.Run(test.name, func(*testing.T) {
			result := Sum(test.param1, test.param2)
			require.Equal(t, test.expected, result, "Result Must "+strconv.Itoa(test.expected))
		})
	}
}

// BENCHMARK
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(4, 5)
	}
}

func BenchmarkSum100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(45, 55)
	}
}

// Benchmark SUB
func BenchmarkSub(b *testing.B) {
	b.Run("Benchmark 10", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Sum(6, 4)
		}
	})

	b.Run("Benchmark 50", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Sum(20, 30)
		}
	})
}

// Benchmark TABLE
func BenchmarkSumTable(b *testing.B) {
	benchmarks := []struct {
		Name   string
		Param1 int
		Param2 int
	}{
		{
			Name:   "SUM 10",
			Param1: 4,
			Param2: 6,
		},
		{
			Name:   "SUM 25",
			Param1: 10,
			Param2: 15,
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Sum(benchmark.Param1, benchmark.Param2)
			}
		})
	}
}
