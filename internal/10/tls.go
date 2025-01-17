package android

const (
	ARM_TLS_SLOT_BIONIC_TLS = iota - 1
	ARM_TLS_SLOT_DTV
	ARM_TLS_SLOT_THREAD_ID
	ARM_TLS_SLOT_APP
	ARM_TLS_SLOT_OPENGL
	ARM_TLS_SLOT_OPENGL_API
	ARM_TLS_SLOT_STACK_GUARD
	ARM_TLS_SLOT_SANITIZER
	ARM_TLS_SLOT_ART_THREAD_SELF

	ARM_MIN_TLS_SLOT = ARM_TLS_SLOT_BIONIC_TLS
	ARM_MAX_TLS_SLOT = ARM_TLS_SLOT_ART_THREAD_SELF

	ARM_BIONIC_TLS_SLOTS = ARM_MAX_TLS_SLOT - ARM_MIN_TLS_SLOT + 1
)

const (
	X86_TLS_SLOT_SELF = iota
	X86_TLS_SLOT_THREAD_ID
	X86_TLS_SLOT_APP
	X86_TLS_SLOT_OPENGL
	X86_TLS_SLOT_OPENGL_API
	X86_TLS_SLOT_STACK_GUARD
	X86_TLS_SLOT_SANITIZER
	X86_TLS_SLOT_ART_THREAD_SELF
	X86_TLS_SLOT_DTV
	X86_TLS_SLOT_BIONIC_TLS

	X86_MIN_TLS_SLOT = X86_TLS_SLOT_SELF
	X86_MAX_TLS_SLOT = X86_TLS_SLOT_BIONIC_TLS

	X86_BIONIC_TLS_SLOTS = X86_MAX_TLS_SLOT - X86_MIN_TLS_SLOT + 1
)
