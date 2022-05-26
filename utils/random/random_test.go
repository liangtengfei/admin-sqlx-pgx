package random

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomString(t *testing.T) {
	str := RandomString(32)
	require.NotEmpty(t, str)
	require.Equal(t, 32, len(str))
}
