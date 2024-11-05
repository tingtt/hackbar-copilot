package toml

import (
	"io"

	"github.com/stretchr/testify/mock"
)

var _ io.Reader = new(MockIOReader)

type MockIOReader struct {
	mock.Mock
}

// Read implements io.Reader.
func (m *MockIOReader) Read(p []byte) (n int, err error) {
	args := m.Called(p)
	return args.Int(0), args.Error(1)
}
