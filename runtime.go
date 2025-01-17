package android

import (
	"context"
	"io"
	"io/fs"

	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/emulator"
	"github.com/wnxd/microdbg/filesystem"
)

type Runtime interface {
	io.Closer
	Debugger() debugger.Debugger
	Emulator() emulator.Emulator
	LoadModule(ctx context.Context, file fs.File) (Module, error)
	FindModule(name string) (Module, error)
	LinkFS(name string, fs filesystem.FS) error
	JavaVM() java.JavaVM
	Package() Package
	ClassFactory() gava.ClassFactory
	RegisterNatives(clazz java.IClass, methods []java.JNINativeMethod) java.JInt
	UnregisterNatives(clazz java.IClass) java.JInt
}
