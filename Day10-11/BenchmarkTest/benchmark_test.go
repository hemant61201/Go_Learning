package BenchmarkTest

import (
	"Simple_Programs/Day10-11/Introduction_UnitTestCase"
	"testing"
)

func BenchMarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Introduction_UnitTestCase.HelloWorld("Hemant")
	}
}

// command line

// go test -v -bench=.

// go test -v -run=NotMathUnitTest -bench=.
