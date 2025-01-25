package android

import (
	"context"

	"github.com/wnxd/microdbg/debugger"
)

type Symbol interface {
	Name() string
	Address() uint64
	Call(ctx context.Context, calling debugger.Calling, ret any, args ...any) error
	MainCall(ctx context.Context, calling debugger.Calling, ret any, args ...any) error
}
