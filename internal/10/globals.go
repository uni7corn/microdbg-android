package android

import "github.com/wnxd/microdbg/emulator"

type globals32 struct {
	fd_table                 fdTable[emulator.Uintptr32]
	initial_linker_arg_count int32
	auxv                     emulator.Uintptr32
	abort_msg_lock           pthread_mutex32
	abort_msg                emulator.Uintptr32
	static_tls_layout        staticTlsLayout[uint32]
	tls_modules              tlsModules32
	tls_allocator            bionicAllocator[emulator.Uintptr32]
	init_progname            emulator.Uintptr32
	init_environ             emulator.Uintptr32
}

type globals64 struct {
	fd_table                 fdTable[emulator.Uintptr64]
	initial_linker_arg_count int32
	auxv                     emulator.Uintptr64
	abort_msg_lock           pthread_mutex64
	abort_msg                emulator.Uintptr64
	static_tls_layout        staticTlsLayout[uint64]
	tls_modules              tlsModules64
	tls_allocator            bionicAllocator[emulator.Uintptr64]
	init_progname            emulator.Uintptr64
	init_environ             emulator.Uintptr64
}
