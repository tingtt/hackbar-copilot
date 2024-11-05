package filesystem

import (
	"bytes"
	"io"
	"io/fs"

	"github.com/stretchr/testify/mock"
)

var _ fs.FS = new(MockFS)

type MockFS struct {
	mock.Mock
}

func (m *MockFS) Open(name string) (fs.File, error) {
	args := m.Called(name)
	return args.Get(0).(fs.File), args.Error(1)
}

var _ fsR = new(MockFSR)

type MockFSR struct {
	mock.Mock
}

func (m *MockFSR) Open(name string) (io.ReadCloser, error) {
	args := m.Called(name)
	return args.Get(0).(io.ReadCloser), args.Error(1)
}

var _ fs.File = new(MockFile)

type MockFile struct {
	*bytes.Buffer
}

func (m *MockFile) Close() error {
	return nil
}

func (m *MockFile) Stat() (fs.FileInfo, error) {
	return nil, nil
}
