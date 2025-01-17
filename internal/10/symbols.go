package android

import (
	"github.com/wnxd/microdbg/debugger"
)

type symbols struct {
	symbols map[string]debugger.ControlHandler
}

func (s *symbols) ctor(dbg debugger.Debugger, handler map[string]any) error {
	s.symbols = make(map[string]debugger.ControlHandler)
	var err error
	for name, fn := range handler {
		s.symbols[name], err = dbg.AddControl(s.handleSymbol, fn)
		if err != nil {
			s.dtor()
			return err
		}
	}
	return nil
}

func (s *symbols) dtor() error {
	for _, sym := range s.symbols {
		sym.Close()
	}
	s.symbols = nil
	return nil
}

func (s *symbols) find(name string) (uint64, error) {
	if sym, ok := s.symbols[name]; ok {
		return sym.Addr(), nil
	}
	return 0, debugger.ErrSymbolNotFound
}

func (s *symbols) handleSymbol(ctx debugger.Context, data any) {
	if fn, ok := data.(func(debugger.Context) any); ok {
		r := fn(ctx)
		ctx.RetWrite(r)
	} else if fn, ok := data.(func(debugger.Context)); ok {
		fn(ctx)
	}
	ctx.Return()
}

func (dbg *dbg) loader_shared_globals(ctx debugger.Context) any {
	return uintptr(dbg.linker.globals)
}

func (dbg *dbg) loader_android_get_application_target_sdk_version(ctx debugger.Context) any {
	const __ANDROID_API__ = 10000

	return int32(__ANDROID_API__)
}

func (dbg *dbg) loader_dlopen(ctx debugger.Context) any {
	var filename string
	ctx.ArgExtract(debugger.Calling_Default, &filename)
	module, err := dbg.FindModule(filename)
	if err != nil {
		return uintptr(0)
	}
	return uintptr(module.BaseAddr())
}

func (dbg *dbg) loader_dlsym(ctx debugger.Context) any {
	var handle uintptr
	var symbol string
	ctx.ArgExtract(debugger.Calling_Default, &handle, &symbol)
	module := dbg.GetModule(uint64(handle))
	if module == nil {
		return uintptr(0)
	}
	addr, err := module.FindSymbol(symbol)
	if err != nil {
		return uintptr(0)
	}
	return handle + uintptr(addr)
}
