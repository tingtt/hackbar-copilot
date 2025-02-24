package filesystem

import (
	"io"
	"io/fs"
	"os"
	"path"
)

type fsR interface {
	Open(name string) (fs.File, error)
}

func newFSR(baseDir string) fsR {
	return fsReadble{os.DirFS(baseDir)}
}

type fsReadble struct {
	fs fs.FS
}

func (f fsReadble) Open(name string) (fs.File, error) {
	return f.fs.Open(name)
}

type fsW interface {
	Create(name string) (io.WriteCloser, error)
}

func newFSW(baseDir string) fsW {
	return fsWritable{baseDir}
}

type fsWritable struct {
	baseDir string
}

func (f fsWritable) Create(name string) (io.WriteCloser, error) {
	return os.Create(path.Join(f.baseDir, name))
}
