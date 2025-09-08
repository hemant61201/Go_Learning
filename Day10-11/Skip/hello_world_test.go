package Skip

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping test on Windows")
	}

	result := HelloWorld("Hemant")

	require.Equal(t, "Hello Hemants", result, "Result should be \"Hello Hemants\"")
}
