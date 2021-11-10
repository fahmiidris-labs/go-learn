package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Fahmi",
			request: "Fahmi",
		},
		{
			name:    "Idris",
			request: "Idris",
		},
		{
			name:    "FahmiIdrisA",
			request: "Fahmi Idris A",
		},
		{
			name:    "Budi",
			request: "Budi Nugraha",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Fahmi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Fahmi")
		}
	})
	b.Run("Idris", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Idris")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Fahmi")
	}
}

func BenchmarkHelloWorldIdris(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Idris")
	}
}

func TestHelloWorldtable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Fahmi",
			request:  "Fahmi",
			expected: "Hello Fahmi",
		},
		{
			name:     "Idris",
			request:  "Idris",
			expected: "Hello Idris",
		},
		{
			name:     "Budi",
			request:  "Budi",
			expected: "Hello Budi",
		},
		{
			name:     "Joko",
			request:  "Joko",
			expected: "Hello Joko",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Fahmi", func(t *testing.T) {
		result := HelloWorld("Fahmi")
		assert.Equal(t, "Hello Fahmi", result, "Result must be 'Hello Fahmi'")
	})
	t.Run("Idris", func(t *testing.T) {
		result := HelloWorld("Idris")
		assert.Equal(t, "Hello Idris", result, "Result must be 'Hello Idris'")
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Mulai")
	m.Run()
	fmt.Println("Selesai")
}

func TestSkipTest(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Can't run on Linux")
	}
	result := HelloWorld("Skip")
	assert.Equal(t, "Hello Skip", result, "Result must be 'Hello Skip'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be 'Hello Assert'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be 'Hello Require'")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")
	if result != "Hello World" {
		t.Error("Result must be 'Hello World'")
	}
}

func TestHelloWorldFahmi(t *testing.T) {
	result := HelloWorld("Fahmi")
	if result != "Hello Fahmi" {
		t.Fatal("Result must be 'Hello Fahmi'")
	}
}
