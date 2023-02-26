package Upper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCapitalize(t *testing.T) {
	expected := "HELLO"
	result := Capitalize("hello")
	require.Equal(t, expected, result, fmt.Sprintf("result must be %v", expected))
}
