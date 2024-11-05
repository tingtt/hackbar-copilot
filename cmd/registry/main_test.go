package main

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_serveGraceful(t *testing.T) {
	t.Parallel()

	type errorBehaviors struct {
		ListenAndServe   error
		Shutdown         error
		SavePersistently error
	}
	tests := []struct {
		name           string
		errorBehaviors errorBehaviors
	}{
		{
			name: "may success graceful shutdown",
			errorBehaviors: errorBehaviors{
				ListenAndServe:   nil,
				Shutdown:         nil,
				SavePersistently: nil,
			},
		},
		{
			name: "may success graceful shutdown",
			errorBehaviors: errorBehaviors{
				ListenAndServe:   nil,
				Shutdown:         http.ErrServerClosed,
				SavePersistently: nil,
			},
		},
		{
			name: "may fail to shutdown server",
			errorBehaviors: errorBehaviors{
				ListenAndServe:   nil,
				Shutdown:         errors.New("wanted error"),
				SavePersistently: nil,
			},
		},
		{
			name: "may fail to save data",
			errorBehaviors: errorBehaviors{
				ListenAndServe:   nil,
				Shutdown:         nil,
				SavePersistently: errors.New("wanted error"),
			},
		},
		{
			name: "may fail to shutdown server and save data",
			errorBehaviors: errorBehaviors{
				ListenAndServe:   nil,
				Shutdown:         errors.New("wanted error 1"),
				SavePersistently: errors.New("wanted error 2"),
			},
		},
		{
			name: "may fail to serve",
			errorBehaviors: errorBehaviors{
				ListenAndServe:   errors.New("wanted error"),
				Shutdown:         nil,
				SavePersistently: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(context.Background())
			server := new(MockServer)
			server.On("ListenAndServe").Return(tt.errorBehaviors.ListenAndServe)
			server.On("Shutdown").Return(tt.errorBehaviors.Shutdown)
			recipeDatasource := new(MockRecipeDatasource)
			recipeDatasource.On("SavePersistently").Return(tt.errorBehaviors.SavePersistently)
			depsDatasources := depsDatasources{
				Recipes: recipeDatasource,
			}

			errChan := make(chan error, 1)
			go func() {
				errChan <- serveGraceful(ctx, server, depsDatasources)
			}()
			cancel()
			time.Sleep(time.Millisecond)
			err := <-errChan

			if tt.errorBehaviors.ListenAndServe == nil && tt.errorBehaviors.Shutdown == nil && tt.errorBehaviors.SavePersistently == nil {
				assert.NoError(t, err)
			} else {
				if tt.errorBehaviors.ListenAndServe != nil {
					assert.ErrorIs(t, err, tt.errorBehaviors.ListenAndServe)
				}
				if tt.errorBehaviors.Shutdown != nil {
					assert.ErrorIs(t, err, tt.errorBehaviors.Shutdown)
				}
				if tt.errorBehaviors.SavePersistently != nil {
					assert.ErrorIs(t, err, tt.errorBehaviors.SavePersistently)
				}
			}
			server.AssertNumberOfCalls(t, "ListenAndServe", 1)
			server.AssertNumberOfCalls(t, "Shutdown", 1)
			recipeDatasource.AssertNumberOfCalls(t, "SavePersistently", 1)
		})
	}
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

		deps, err := loadDependencies(t.TempDir())

		assert.NoError(t, err)
		assert.NotNil(t, deps.Datasources.Recipes)
		assert.NotNil(t, deps.Usecase.GraphQL.Recipes)
		// assert.NotNil(t, deps.Usecase.GraphQL.Orders)
	})

	t.Run("may fail to load repository", func(t *testing.T) {
		t.Parallel()

		deps, err := loadDependencies("")

		assert.Error(t, err)
		assert.Nil(t, deps.Datasources.Recipes)
		assert.Nil(t, deps.Usecase.GraphQL.Recipes)
		assert.Nil(t, deps.Usecase.GraphQL.Orders)
	})
}
