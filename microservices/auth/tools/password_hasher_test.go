package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPasswordHasher(t *testing.T) {
	password := "my$ecretP@ssw0rd"

	t.Run("hash and compare success", func(t *testing.T) {
		hash, err := Hash(password)
		require.NoError(t, err)
		assert.NotEmpty(t, hash)
		assert.NotEqual(t, password, hash)

		err = Compare(hash, password)
		assert.NoError(t, err)
	})

	t.Run("compare fail", func(t *testing.T) {
		hash, err := Hash(password)
		require.NoError(t, err)

		err = Compare(hash, "wrongpassword")
		assert.Error(t, err)
	})
}
