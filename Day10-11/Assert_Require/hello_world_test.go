package Assert_Require

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {

	result := HelloWorld("Hemant")

	assert.Equal(t, "Hello Hemants!", result, "Result should be \"Hello Hemants!\"")

	fmt.Println("Test Case Stoped...")
}

func TestHelloWorldRequire(t *testing.T) {

	result := HelloWorld("Hemant")

	require.Equal(t, "Hello Hemants!", result, "Result should be \"Hello Hemants!\"")

	fmt.Println("Test Case Stoped...")
}
