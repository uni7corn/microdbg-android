package virtual

import (
	"errors"
	"io/fs"

	"github.com/wnxd/microdbg/filesystem"
)

type vfs struct {
	fs fs.FS
}

func FS(fs fs.FS) filesystem.FS {
	return vfs{fs: fs}
}

func (f vfs) Sub(dir string) (fs.FS, error) {
	if sub, ok := f.fs.(fs.SubFS); ok {
		return sub.Sub(dir)
	}
	return nil, errors.ErrUnsupported
}

func (f vfs) Open(name string) (fs.File, error) {
	return filesystem.Open(f, name)
}

func (f vfs) OpenFile(name string, flag filesystem.FileFlag, perm fs.FileMode) (filesystem.File, error) {
	if flag != filesystem.O_RDONLY {
		return nil, fs.ErrPermission
	}
	return f.fs.Open(name)
}

func (f vfs) ReadDir(name string) ([]fs.DirEntry, error) {
	if dir, ok := f.fs.(fs.ReadDirFS); ok {
		return dir.ReadDir(name)
	}
	return nil, errors.ErrUnsupported
}

func (f vfs) Mkdir(name string, perm fs.FileMode) (filesystem.DirFS, error) {
	if dir, ok := f.fs.(filesystem.DirFS); ok {
		return dir.Mkdir(name, perm)
	}
	return nil, errors.ErrUnsupported
}

func (f vfs) Readlink(name string) (string, error) {
	if dir, ok := f.fs.(filesystem.ReadlinkFS); ok {
		return dir.Readlink(name)
	}
	return "", errors.ErrUnsupported
}
