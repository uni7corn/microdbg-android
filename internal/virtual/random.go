package virtual

import (
	"crypto/rand"
	"io/fs"

	"github.com/wnxd/microdbg/filesystem"
)

type RandomFS string

func (f RandomFS) Open(name string) (fs.File, error) {
	return filesystem.Open(f, name)
}

func (f RandomFS) OpenFile(name string, flag filesystem.FileFlag, perm fs.FileMode) (filesystem.File, error) {
	return f, nil
}

func (f RandomFS) Close() error {
	return nil
}

func (f RandomFS) Stat() (fs.FileInfo, error) {
	return &fileInfo{name: string(f), mode: 0666}, nil
}

func (f RandomFS) Read(b []byte) (int, error) {
	n := min(len(b), 256)
	return rand.Read(b[:n])
}
