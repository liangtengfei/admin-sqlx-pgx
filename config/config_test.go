package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	config, err := NewConfig("../resources")
	require.NoError(t, err)
	require.NotEmpty(t, config)

	require.Equal(t, "postgresql", config.Server.DBDriver)
	require.Equal(t, "debug", config.Logger.Level)
	require.Equal(t, 3306, config.DB.Port)
}
