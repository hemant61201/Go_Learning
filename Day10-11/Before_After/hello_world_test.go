package Before_After

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("Test Case Started...")

	m.Run()

	//after
	fmt.Println("Test Case Stoped...")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Hemant")

	require.Equal(t, "Hello Hemants", result, "Result should be \"Hello Hemants\"")
}
