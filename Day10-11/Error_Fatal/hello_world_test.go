package Error_Fatal

import (
	"fmt"
	"testing"
)

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Hemant")

	if result != "Hello Hemants!" {
		t.Error("Result is wrong The Output should be Hemants")
		fmt.Println("Result: ", result)
	}
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("Hemant")

	if result != "Hello Hemants!" {
		t.Fatal("Result is wrong The Output should be Hemants")
		fmt.Println("Result: ", result)
	}
}
