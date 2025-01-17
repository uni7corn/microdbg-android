package internal

import (
	"io"
	"io/fs"
	"strings"

	"github.com/wnxd/microdbg-android/internal/virtual"
	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/filesystem"
)

type HybridFS struct {
	fs   filesystem.VirtualFS
	Base interface {
		io.Closer
		fs.FS
	}
	Sys filesystem.DirFS
}

func (h *HybridFS) Ctor(dbg debugger.Debugger) error {
	h.fs = filesystem.NewVirtualFS()
	h.fs.Link("/dev/random", virtual.RandomFS("random"))
	h.fs.Link("/dev/urandom", virtual.RandomFS("urandom"))
	// h.fs.Link("/proc/self/exe", filesystem.SoftLink("/system/bin/app_process", nil))
	dbg.AddFileHandler(h)
	return nil
}

func (h *HybridFS) Dtor(dbg debugger.Debugger) error {
	dbg.RemoveFileHandler(h)
	if h.Base != nil {
		h.Base.Close()
	}
	return nil
}

func (h *HybridFS) Open(name string) (fs.File, error) {
	return filesystem.Open(h, name)
}

func (h *HybridFS) OpenFile(name string, flag filesystem.FileFlag, perm fs.FileMode) (filesystem.File, error) {
	name = strings.TrimPrefix(name, "/")
	if h.Base != nil && flag == filesystem.O_RDONLY {
		file, err := h.Base.Open(name)
		if err == nil {
			return file, nil
		}
	}
	if h.Sys != nil {
		file, err := h.Sys.OpenFile(name, flag, perm)
		if err == nil {
			return file, nil
		}
	}
	return h.fs.OpenFile(name, flag, perm)
}

func (h *HybridFS) Stat(name string) (fs.FileInfo, error) {
	name = strings.TrimPrefix(name, "/")
	if h.Base != nil {
		info, err := fs.Stat(h.Base, name)
		if err == nil {
			return info, nil
		}
	}
	if h.Sys != nil {
		info, err := fs.Stat(h.Sys, name)
		if err == nil {
			return info, nil
		}
	}
	return fs.Stat(h.fs, name)
}

func (h *HybridFS) ReadDir(name string) ([]fs.DirEntry, error) {
	name = strings.TrimPrefix(name, "/")
	if h.Base != nil {
		arr, err := fs.ReadDir(h.Base, name)
		if err == nil {
			return arr, nil
		}
	}
	if h.Sys != nil {
		arr, err := fs.ReadDir(h.Sys, name)
		if err == nil {
			return arr, nil
		}
	}
	return fs.ReadDir(h.fs, name)
}

func (h *HybridFS) Mkdir(name string, perm fs.FileMode) (filesystem.DirFS, error) {
	name = strings.TrimPrefix(name, "/")
	if h.Sys != nil {
		dir, err := h.Sys.Mkdir(name, perm)
		if err == nil {
			return dir, nil
		}
	}
	return h.fs.Mkdir(name, perm)
}

func (h *HybridFS) Readlink(name string) (string, error) {
	return h.fs.Readlink(name)
}

func (h *HybridFS) Link(name string, handle filesystem.FS) error {
	return h.fs.Link(name, handle)
}
