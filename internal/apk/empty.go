package apk

import "io/fs"

type emptyFS struct{}

var EmptyPackage = (&info{fs: emptyFS{}, name: "microdbg"}).init()

func (emptyFS) Close() error {
	return nil
}

func (emptyFS) Open(string) (fs.File, error) {
	return nil, fs.ErrNotExist
}
