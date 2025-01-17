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
	Dump(w io.Writer) error
}
