package filesystem

import (
	"bytes"
	"io/fs"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newFSR(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()
		assert.NotNil(t, newFSR(path.Join(t.TempDir(), "fsr")))
	})
}

func Test_fsReadble_Open(t *testing.T) {
	t.Parallel()

	type fields struct {
		fs *MockFS
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "may open exists file",
			fields: fields{
				fs: func() *MockFS {
					m := new(MockFS)
					r := bytes.NewBufferString("content")
					m.On("Open", "existFilename").Return(&MockFile{r}, nil)
					return m
				}(),
			},
			args: args{
				name: "existFilename",
			},
			wantErr: false,
		},
		{
			name: "may return error, if file not exists",
			fields: fields{
				fs: func() *MockFS {
					m := new(MockFS)
					m.On("Open", "notExistFilename").Return(&MockFile{}, fs.ErrNotExist)
					return m
				}(),
			},
			args: args{
				name: "notExistFilename",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := fsReadble{fs: tt.fields.fs}

			_, err := f.Open(tt.args.name)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_newFSW(t *testing.T) {
	t.Parallel()

	t.Run("will return non-nil struct", func(t *testing.T) {
		t.Parallel()
		assert.NotNil(t, newFSW(path.Join(t.TempDir(), "fsw")))
	})
}

func Test_fsWritable_Create(t *testing.T) {
	t.Parallel()

	t.Run("may open exists file", func(t *testing.T) {
		t.Parallel()

		baseDir := t.TempDir()
		f := fsWritable{
			baseDir: baseDir,
		}
		os.Create(path.Join(baseDir, "existsFilename"))

		got, err := f.Create("existsFilename")

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("may open exists file", func(t *testing.T) {
		t.Parallel()

		baseDir := t.TempDir()
		f := fsWritable{
			baseDir: baseDir,
		}
		os.Create(path.Join(baseDir, "existsFilename"))

		got, err := f.Create("existsFilename")

		assert.NoError(t, err)
		assert.NotNil(t, got)
	})

	t.Run("may return error, if call with uncreatable filename", func(t *testing.T) {
		t.Parallel()

		baseDir := t.TempDir()
		f := fsWritable{
			baseDir: baseDir,
		}

		got, err := f.Create("/path/to/uncreatable")

		assert.ErrorIs(t, err, os.ErrNotExist)
		assert.Nil(t, got)
	})
}
