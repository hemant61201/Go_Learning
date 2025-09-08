package Fail_FailNow

import (
	"fmt"
	"testing"
)

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("Hemant")

	if result != "Hello Hemants!" {
		t.Fail()
		fmt.Println("Execution is running...")
	}
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("Hemant")

	if result != "Hello Hemants!" {
		t.FailNow()
		fmt.Println("Execution is stopped above...")
	}
}
