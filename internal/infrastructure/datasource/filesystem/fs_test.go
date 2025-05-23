package filesystem

import (
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()

		fs, err := NewRepository(t.TempDir())

		assert.NoError(t, err)
		assert.NotNil(t, fs)
	})

	t.Run("will return non-nil implements", func(t *testing.T) {
		t.Parallel()

		fs, err := NewRepository(t.TempDir())

		assert.NoError(t, err)
		assert.NotNil(t, fs.BarCounterGateway())
		assert.NotNil(t, fs.CashierGateway())
		assert.NotNil(t, fs.CopilotGateway())
		assert.NotNil(t, fs.OrderGateway())
	})

	t.Run("may return error, if fail to create data files", func(t *testing.T) {
		t.Parallel()

		fs, err := NewRepository(path.Join(t.TempDir(), "path/to/not/writable"))

		assert.Error(t, err)
		assert.Nil(t, fs)
	})
}
