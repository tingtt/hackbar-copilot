package main

import (
	"context"
	"hackbar-copilot/internal/infrastructure/datasource/filesystem"

	"github.com/stretchr/testify/mock"
)

type MockRecipeDatasource struct {
	mock.Mock
	filesystem.Filesystem
}

func (m *MockRecipeDatasource) SavePersistently() error {
	args := m.Called()
	return args.Error(0)
}

var _ server = new(MockServer)

type MockServer struct {
	mock.Mock
}

func (m *MockServer) ListenAndServe() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockServer) Shutdown(ctx context.Context) error {
	args := m.Called()
	return args.Error(0)
}
