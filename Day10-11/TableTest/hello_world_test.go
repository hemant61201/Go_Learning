package TableTest

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHelloWorld(t *testing.T) {

	tests := []struct {
		name     string
		requst   string
		expected string
	}{
		{
			name:     "Hemant",
			requst:   "Hemant",
			expected: "Hello Hemant!",
		},
		{
			name:     "Aakash",
			requst:   "Aakash",
			expected: "Hello Aakash!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.requst)
			require.Equal(t, test.expected, result, "Result should be \"Hello Aakash\"")
		})
	}
}
