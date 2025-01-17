package android

import (
	"unsafe"

	"github.com/wnxd/microdbg-android/internal"
	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/emulator"
	emu_arm "github.com/wnxd/microdbg/emulator/arm"
	emu_arm64 "github.com/wnxd/microdbg/emulator/arm64"
	emu_x86 "github.com/wnxd/microdbg/emulator/x86"
)

var environ = []string{
	"ANDROID_DATA=/data",
	"ANDROID_ROOT=/system",
	"PATH=/sbin:/vendor/bin:/system/sbin:/system/bin:/system/xbin",
	"NO_ADDR_COMPAT_LAYOUT_FIXUP=1",
}

type linker struct {
	mem       emulator.MemRegion
	addr      uint64
	progname  uint64
	env       []uint64
	auxv      map[uint64]uint64
	bionicTls uint64
	errno     emulator.Pointer
	globals   uint64
}

func (l *linker) ctor(dbg debugger.Debugger) (err error) {
	l.mem, err = dbg.MapAlloc(0x2000, emulator.MEM_PROT_READ|emulator.MEM_PROT_WRITE)
	return
}

func (l *linker) dtor(dbg debugger.Debugger) error {
	if l.bionicTls != 0 {
		dbg.MemFree(l.bionicTls)
	}
	return dbg.MapFree(l.mem.Addr, l.mem.Size)
}

func (l *linker) init(art internal.Runtime) error {
	l.addr = l.mem.Addr
	err := l.initEnvironment(art)
	if err != nil {
		return err
	}
	err = l.initTLS(art)
	if err != nil {
		return err
	}
	return l.initSharedGlobals(art)
}

func (l *linker) pushData(emu emulator.Emulator, ptr unsafe.Pointer, size uint64) (uint64, error) {
	addr := l.addr
	err := emu.MemWritePtr(addr, size, ptr)
	if err != nil {
		return 0, err
	}
	l.addr += size
	return addr, nil
}

func (l *linker) pushString(emu emulator.Emulator, str string) (uint64, error) {
	addr, err := l.pushData(emu, unsafe.Pointer(unsafe.StringData(str)), uint64(len(str)))
	if err != nil {
		return 0, err
	}
	l.addr += 1
	return addr, nil
}

func (l *linker) initEnvironment(art internal.Runtime) error {
	emu := art.Emulator()
	processName := art.Package().Name()
	addr, err := l.pushString(emu, processName)
	if err != nil {
		return err
	}
	l.progname = addr
	for _, v := range environ {
		addr, err = l.pushString(emu, v)
		if err != nil {
			return err
		}
		l.env = append(l.env, addr)
	}
	l.auxv = make(map[uint64]uint64)
	l.auxv[AT_PAGESZ] = emu.PageSize()
	var random [16]byte
	addr, err = l.pushData(emu, unsafe.Pointer(unsafe.SliceData(random[:])), uint64(len(random)))
	if err != nil {
		return err
	}
	l.auxv[AT_RANDOM] = addr
	return nil
}

func (l *linker) initTLS(art internal.Runtime) error {
	dbg := art.Debugger()
	emu := art.Emulator()
	arch := emu.Arch()
	switch arch {
	case emulator.ARCH_ARM, emulator.ARCH_ARM64, emulator.ARCH_X86, emulator.ARCH_X86_64:
	default:
		return emulator.ErrArchUnsupported
	}
	var err error
	var thread, stackGuard, pointerSize uint64
	switch arch {
	case emulator.ARCH_ARM, emulator.ARCH_X86:
		l.bionicTls, err = dbg.MemAlloc(0x2b44)
		if err != nil {
			return err
		}
		l.addr = align(l.addr, 8)
		t := pthread[emulator.Uintptr32]{tid: 1}
		thread, err = l.pushData(emu, unsafe.Pointer(&t), uint64(unsafe.Sizeof(t)))
		if err != nil {
			return err
		}
		l.errno = dbg.ToPointer(thread).Add(0x260)
		pointerSize = 4
	case emulator.ARCH_ARM64, emulator.ARCH_X86_64:
		l.bionicTls, err = dbg.MemAlloc(0x2fa0)
		if err != nil {
			return err
		}
		l.addr = align(l.addr, 16)
		t := pthread[emulator.Uintptr64]{tid: 1}
		thread, err = l.pushData(emu, unsafe.Pointer(&t), uint64(unsafe.Sizeof(t)))
		if err != nil {
			return err
		}
		l.errno = dbg.ToPointer(thread).Add(0x2B8)
		pointerSize = 8
	}
	stackGuard, err = l.pushData(emu, unsafe.Pointer(&stackGuard), pointerSize)
	if err != nil {
		return err
	}
	bionicTls := l.bionicTls
	var tls uint64
	switch arch {
	case emulator.ARCH_ARM, emulator.ARCH_ARM64:
		tls = l.mem.Addr + l.mem.Size - (ARM_MAX_TLS_SLOT * pointerSize)
		emu.MemWritePtr(tls+(unsigned(ARM_TLS_SLOT_BIONIC_TLS)*pointerSize), pointerSize, unsafe.Pointer(&bionicTls))
		emu.MemWritePtr(tls+(ARM_TLS_SLOT_THREAD_ID*pointerSize), pointerSize, unsafe.Pointer(&thread))
		emu.MemWritePtr(tls+(ARM_TLS_SLOT_STACK_GUARD*pointerSize), pointerSize, unsafe.Pointer(&stackGuard))
	case emulator.ARCH_X86, emulator.ARCH_X86_64:
		tls = l.mem.Addr + l.mem.Size - (X86_MAX_TLS_SLOT * pointerSize)
		emu.MemWritePtr(tls+(X86_TLS_SLOT_BIONIC_TLS*pointerSize), pointerSize, unsafe.Pointer(&bionicTls))
		emu.MemWritePtr(tls+(X86_TLS_SLOT_THREAD_ID*pointerSize), pointerSize, unsafe.Pointer(&thread))
		emu.MemWritePtr(tls+(X86_TLS_SLOT_STACK_GUARD*pointerSize), pointerSize, unsafe.Pointer(&stackGuard))
	}
	switch arch {
	case emulator.ARCH_ARM:
		emu.RegWrite(emu_arm.ARM_REG_C13_C0_3, tls)
	case emulator.ARCH_ARM64:
		emu.RegWrite(emu_arm64.ARM64_REG_TPIDR_EL0, tls)
	case emulator.ARCH_X86:
		emu.RegWrite(emu_x86.X86_REG_GS, tls)
	case emulator.ARCH_X86_64:
		emu.RegWrite(emu_x86.X86_REG_FS, tls)
	}
	return nil
}

func (l *linker) initSharedGlobals(art internal.Runtime) error {
	emu := art.Emulator()
	switch emu.Arch() {
	case emulator.ARCH_ARM, emulator.ARCH_X86:
		var globals globals32
		globals.auxv = emulator.Uintptr32(l.addr)
		for typ, val := range l.auxv {
			l.pushData(emu, unsafe.Pointer(&typ), 4)
			l.pushData(emu, unsafe.Pointer(&val), 4)
		}
		{
			null := uint32(AT_NULL)
			l.pushData(emu, unsafe.Pointer(&null), 4)
			l.pushData(emu, unsafe.Pointer(&null), 4)
		}
		globals.init_progname = emulator.Uintptr32(l.progname)
		globals.init_environ = emulator.Uintptr32(l.addr)
		for _, v := range l.env {
			l.pushData(emu, unsafe.Pointer(&v), 4)
		}
		{
			var null uint32
			l.pushData(emu, unsafe.Pointer(&null), 4)
		}
		l.addr = align(l.addr, 8)
		l.globals, _ = l.pushData(emu, unsafe.Pointer(&globals), uint64(unsafe.Sizeof(globals)))
	case emulator.ARCH_ARM64, emulator.ARCH_X86_64:
		var globals globals64
		globals.auxv = l.addr
		for typ, val := range l.auxv {
			l.pushData(emu, unsafe.Pointer(&typ), 8)
			l.pushData(emu, unsafe.Pointer(&val), 8)
		}
		{
			null := uint64(AT_NULL)
			l.pushData(emu, unsafe.Pointer(&null), 8)
			l.pushData(emu, unsafe.Pointer(&null), 8)
		}
		globals.init_progname = l.progname
		globals.init_environ = l.addr
		for _, v := range l.env {
			l.pushData(emu, unsafe.Pointer(&v), 8)
		}
		{
			var null uint64
			l.pushData(emu, unsafe.Pointer(&null), 8)
		}
		l.addr = align(l.addr, 16)
		l.globals, _ = l.pushData(emu, unsafe.Pointer(&globals), uint64(unsafe.Sizeof(globals)))
	}
	return nil
}

func unsigned(v int64) uint64 {
	return uint64(v)
}

func align(a, b uint64) uint64 {
	a += b - 1
	mask := -b
	a &= mask
	return a
}
