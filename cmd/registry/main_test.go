package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_serveGraceful(t *testing.T) {
	t.Parallel()

	t.Run("will call recipe datasource SavePersistently() when canceled", func(t *testing.T) {
		t.Parallel()

		server := new(MockServer)
		server.On("ListenAndServe").Return(nil)
		server.On("Shutdown").Return(nil)

		datasource := new(MockRecipeDatasource)
		datasource.On("SavePersistently").Return(nil)
		depsDatasources := depsDatasources{fs: datasource}

		errChan := make(chan error, 1)
		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			errChan <- serveGraceful(ctx, server, depsDatasources)
		}()
		time.Sleep(time.Millisecond)

		server.AssertCalled(t, "ListenAndServe")
		server.AssertNotCalled(t, "Shutdown", mock.Anything)
		datasource.AssertNotCalled(t, "SavePersistently")

		cancel()
		err := <-errChan

		assert.NoError(t, err)
		server.AssertCalled(t, "Shutdown", mock.Anything)
		datasource.AssertCalled(t, "SavePersistently")
	})
}

func Test_getCLIOption(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args []string
		want option
	}{
		{
			name: "may return default options",
			args: []string{},
			want: option{
				Host:        "127.0.0.1",
				Port:        "8080",
				DataDirPath: "/var/lib/hackbar-copilot",
			},
		},
		{
			name: "may return specified options",
			args: []string{"--host", "0.0.0.0", "--port", "80", "--data", "/var/lib/test"},
			want: option{
				Host:        "0.0.0.0",
				Port:        "80",
				DataDirPath: "/var/lib/test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := getCLIOption(append([]string{"main"}, tt.args...))

			assert.Equal(t, got, tt.want)
		})
	}
}

func Test_loadDependencies(t *testing.T) {
	t.Parallel()

	t.Run("may load successfully", func(t *testing.T) {
		t.Parallel()

		deps, _, err := loadDependencies(t.TempDir())

		assert.NoError(t, err)
		assert.NotNil(t, deps.Usecase.GraphQL.Copilot)
		assert.NotNil(t, deps.Usecase.GraphQL.OrderService)
	})

	t.Run("may fail to load repository", func(t *testing.T) {
		t.Parallel()

		deps, _, err := loadDependencies("")

		assert.Error(t, err)
		assert.Nil(t, deps.Usecase.GraphQL.Copilot)
		assert.Nil(t, deps.Usecase.GraphQL.OrderService)
	})
}
