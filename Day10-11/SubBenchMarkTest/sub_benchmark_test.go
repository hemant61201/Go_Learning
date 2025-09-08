package SubBenchMarkTest

import (
	"Simple_Programs/Day10-11/Introduction_UnitTestCase"
	"testing"
)

func BenchMarkHelloWorld(b *testing.B) {

	b.Run("Hemant", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Introduction_UnitTestCase.HelloWorld("Hemant")
		}
	})

	b.Run("Aakash", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Introduction_UnitTestCase.HelloWorld("Aakash")
		}
	})
}
