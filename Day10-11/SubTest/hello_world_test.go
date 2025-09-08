package SubTest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {

	t.Run("Hemant", func(t *testing.T) {
		result := HelloWorld("Hemant")
		require.Equal(t, "Hello Hemant!", result, "Result should be \"Hello Hemant\"")
	})

	t.Run("Aakash", func(t *testing.T) {
		result := HelloWorld("Aakash")
		require.Equal(t, "Hello Aakash!", result, "Result should be \"Hello Aakash\"")
	})
}
