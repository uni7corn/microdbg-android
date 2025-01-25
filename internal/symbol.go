package internal

import (
	"context"

	android "github.com/wnxd/microdbg-android"
	"github.com/wnxd/microdbg/debugger"
)

type Symbol struct {
	dbg  debugger.Debugger
	name string
	addr uint64
}

func NewSymbol(dbg debugger.Debugger, name string, addr uint64) android.Symbol {
	return &Symbol{dbg: dbg, name: name, addr: addr}
}

func (sym Symbol) Name() string {
	return sym.name
}

func (sym Symbol) Address() uint64 {
	return sym.addr
}

func (sym Symbol) Call(ctx context.Context, calling debugger.Calling, ret any, args ...any) error {
	task, err := sym.dbg.GetMainTask(ctx)
	if err != nil {
		task, err = sym.dbg.CreateTask(ctx)
		if err != nil {
			return err
		}
	}
	defer task.Close()
	return sym.call(task, calling, ret, args...)
}

func (sym Symbol) call(task debugger.Task, calling debugger.Calling, ret any, args ...any) error {
	taskCtx := task.Context()
	taskCtx.ArgWrite(calling, args...)
	err := sym.dbg.CallTaskOf(task, sym.addr)
	if err != nil {
		return err
	}
	err = task.SyncRun()
	if err != nil {
		return err
	}
	return taskCtx.RetExtract(ret)
}
