package Introduction_UnitTestCase

import "testing"

func TestHelloWorld(t *testing.T) {

	result := HelloWorld("Hemant")

	if result != "Hello Hemant!" {
		panic("Result Wrong!")
	}
}

// Command line

// go test
// go test -v : detailed output
// go test -v -run=TestHelloWorld
// go test -v ./ ... : To run all test cases
