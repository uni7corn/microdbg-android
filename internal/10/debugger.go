package android

import (
	"context"
	"errors"
	"io/fs"
	"sync"
	"unsafe"

	"github.com/wnxd/microdbg-android/internal"
	"github.com/wnxd/microdbg-android/internal/arm"
	"github.com/wnxd/microdbg-android/internal/arm64"
	"github.com/wnxd/microdbg-android/internal/x86"
	"github.com/wnxd/microdbg-android/internal/x86_64"
	linux "github.com/wnxd/microdbg-linux"
	kernel "github.com/wnxd/microdbg-linux/kernel"
	"github.com/wnxd/microdbg-loader/elf"
	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/debugger/extend"
	"github.com/wnxd/microdbg/emulator"
	"github.com/wnxd/microdbg/filesystem"
)

type dbg struct {
	extend.ExtendDebugger
	*kernel.Kernel
	linker
	symbols
	nrMap      map[uint64]linux.NR
	errModuels sync.Map
}

func newDbg(emu emulator.Emulator) (*dbg, error) {
	dbg, err := extend.New[*dbg](emu)
	if err != nil {
		return nil, err
	}
	releases := []func() error{dbg.Close}
	defer func() {
		for i := len(releases) - 1; i >= 0; i-- {
			releases[i]()
		}
	}()
	dbg.Kernel, err = kernel.NewKernel(dbg)
	if err != nil {
		return nil, err
	}
	releases = append(releases, dbg.Kernel.Close)
	err = dbg.linker.ctor(dbg)
	if err != nil {
		return nil, err
	}
	releases = append(releases, func() error { return dbg.linker.dtor(dbg) })
	err = dbg.symbols.ctor(dbg, map[string]any{
		"__loader_shared_globals":                             dbg.loader_shared_globals,
		"__loader_android_get_application_target_sdk_version": dbg.loader_android_get_application_target_sdk_version,
		"__loader_dlopen":                                     dbg.loader_dlopen,
		"__loader_dlsym":                                      dbg.loader_dlsym,
	})
	if err != nil {
		return nil, err
	}
	releases = nil
	switch emu.Arch() {
	case emulator.ARCH_ARM:
		dbg.nrMap = arm.NRMap
	case emulator.ARCH_ARM64:
		dbg.nrMap = arm64.NRMap
	case emulator.ARCH_X86:
		dbg.nrMap = x86.NRMap
	case emulator.ARCH_X86_64:
		dbg.nrMap = x86_64.NRMap
	}
	return dbg, nil
}

func (dbg *dbg) Close() error {
	dbg.symbols.dtor()
	dbg.linker.dtor(dbg)
	dbg.Kernel.Close()
	return dbg.ExtendDebugger.Close()
}

func (dbg *dbg) FindModule(name string) (debugger.Module, error) {
	module, err := dbg.ExtendDebugger.FindModule(name)
	if err == nil {
		return module, nil
	}
	if err, ok := dbg.errModuels.Load(name); ok {
		return nil, err.(error)
	}
	var file filesystem.File
	switch dbg.Arch() {
	case emulator.ARCH_ARM, emulator.ARCH_X86:
		file, err = dbg.OpenFile("/system/lib/"+name, filesystem.O_RDONLY, 0)
	case emulator.ARCH_ARM64, emulator.ARCH_X86_64:
		file, err = dbg.OpenFile("/system/lib64/"+name, filesystem.O_RDONLY, 0)
	}
	if err != nil {
		dbg.errModuels.Store(name, debugger.ErrModuleNotFound)
		return nil, debugger.ErrModuleNotFound
	}
	module, err = dbg.loadModule(context.TODO(), file.(fs.File))
	file.Close()
	if err != nil {
		err = errors.Join(debugger.ErrModuleNotFound, err)
		dbg.errModuels.Store(name, err)
		return nil, err
	}
	return module, nil
}

func (dbg *dbg) FindSymbol(name string) (debugger.Module, uint64, error) {
	addr, err := dbg.symbols.find(name)
	if err == nil {
		return debugger.InternalModule, addr, nil
	}
	return dbg.ExtendDebugger.FindSymbol(name)
}

func (dbg *dbg) NR(no uint64) linux.NR {
	return dbg.nrMap[no]
}

func (dbg *dbg) Errno() linux.Errno {
	if dbg.linker.errno.IsNil() {
		return 0
	}
	var err int32
	dbg.linker.errno.MemReadPtr(4, unsafe.Pointer(&err))
	return linux.Errno(err)
}

func (dbg *dbg) SetErrno(err linux.Errno) {
	if dbg.linker.errno.IsNil() {
		return
	}
	dbg.linker.errno.MemWritePtr(4, unsafe.Pointer(&err))
}

func (dbg *dbg) loadModule(ctx context.Context, file fs.File) (elf.Module, error) {
	module, err := elf.ImportFile(dbg, file)
	if err != nil {
		return nil, err
	}
	dbg.Load(module)
	internal.AndroidReloc(dbg, module)
	err = module.Init(ctx)
	if err != nil {
		dbg.Unload(module)
		module.Close()
		return nil, err
	}
	return module, nil
}
