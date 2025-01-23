package android

import (
	"context"
	"io"

	java "github.com/wnxd/microdbg-java"
)

type Module interface {
	io.Closer
	Name() string
	BaseAddr() uint64
	FindSymbol(name string) (Symbol, error)
	Symbols(yield func(Symbol) bool)
	CallEntry(ctx context.Context) error
	CallOnLoad(ctx context.Context) (java.JInt, error)
	FindNativeMethod(vm java.JavaVM, clazz java.IClass, name, sig string) (NativeMethod, error)
	Dump(w io.Writer) error
}

type NativeMethod = func(obj java.IObject, args ...any) any
