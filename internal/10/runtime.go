package android

import (
	"archive/zip"
	"context"
	"io/fs"

	android "github.com/wnxd/microdbg-android"
	"github.com/wnxd/microdbg-android/internal"
	"github.com/wnxd/microdbg-android/internal/apk"
	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/emulator"
	"github.com/wnxd/microdbg/filesystem"
)

type art struct {
	dbg    *dbg
	hybrid internal.HybridFS
	internal.Environ
}

func NewRuntime(emu emulator.Emulator, options ...android.Option) (android.Runtime, error) {
	dbg, err := newDbg(emu)
	if err != nil {
		return nil, err
	}
	releases := []func() error{dbg.Close}
	defer func() {
		for i := len(releases) - 1; i >= 0; i-- {
			releases[i]()
		}
	}()
	r := &art{dbg: dbg}
	r.Environ.APK = apk.EmptyPackage
	err = r.hybrid.Ctor(dbg)
	if err != nil {
		return nil, err
	}
	releases = append(releases, func() error { return r.hybrid.Dtor(dbg) })
	err = android.SetOption(r, options...)
	if err != nil {
		return nil, err
	}
	err = r.dbg.linker.init(r)
	if err != nil {
		return nil, err
	}
	err = r.Environ.Init(dbg)
	if err != nil {
		return nil, err
	}
	releases = nil
	return r, nil
}

func (r *art) Close() error {
	r.Environ.Close()
	r.hybrid.Dtor(r.dbg)
	return r.dbg.Close()
}

func (r *art) Debugger() debugger.Debugger {
	return r.dbg
}

func (r *art) Emulator() emulator.Emulator {
	return r.dbg.Emulator()
}

func (r *art) LoadModule(ctx context.Context, file fs.File) (android.Module, error) {
	module, err := r.dbg.loadModule(ctx, file)
	if err != nil {
		return nil, err
	}
	return internal.ModuleOf(module, r), nil
}

func (r *art) FindModule(name string) (android.Module, error) {
	module, err := r.dbg.FindModule(name)
	if err != nil {
		return nil, err
	}
	return internal.ModuleOf(module, r), nil
}

func (r *art) LinkFS(name string, handler filesystem.FS) error {
	return r.hybrid.Link(name, handler)
}

func (r *art) setApkPath(name string) error {
	info, err := apk.Load(name)
	if err != nil {
		return err
	}
	r.Environ.APK = info
	info.Link(r)
	return nil
}

func (r *art) setRuntimeDir(name string) error {
	fs, err := zip.OpenReader(name)
	if err != nil {
		return err
	}
	if r.hybrid.Base != nil {
		r.hybrid.Base.Close()
	}
	r.hybrid.Base = fs
	return nil
}

func (r *art) setRootDir(name string) error {
	r.hybrid.Sys = filesystem.SysDirFS(name)
	return nil
}

func (r *art) setJNIEnv(env android.JNIEnv) error {
	r.Environ.JNI = env
	return nil
}
