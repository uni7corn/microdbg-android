package android

import (
	"unsafe"

	"github.com/wnxd/microdbg/emulator"
)

type Pointer interface {
	~uint32 | ~uint64
}

type fdTable[P Pointer] struct {
	version     uint32
	error_level int32
	entries     [128]uint64
	overflow    P
}

type staticTlsLayout[P Pointer] struct {
	offset_            P
	alignment_         P
	overflowed_        bool
	offset_bionic_tcb_ P
	offset_bionic_tls_ P
}

type bionicSmallObjectAllocator[P Pointer] struct {
	type_            uint32
	block_size_      P
	blocks_per_page_ P
	free_pages_cnt_  P
	page_list_       P
}

type bionicAllocator[P Pointer] struct {
	allocators_     P
	allocators_buf_ [7]bionicSmallObjectAllocator[P]
}

type mntent[P Pointer] struct {
	mnt_fsname P
	mnt_dir    P
	mnt_type   P
	mnt_opts   P
	mnt_freq   int32
	mnt_passno int32
}

type group[P Pointer] struct {
	gr_name   P
	gr_passwd P
	gr_gid    uint32
	gr_mem    P
}

type group_state[P Pointer] struct {
	group_             group[P]
	group_members_     [2]P
	group_name_buffer_ [32]byte
	getgrent_idx       P
}

type tlsModules32 struct {
	generation         uint32
	generation_libc_so uint32
	rwlock             pthread_rwlock32
	module_count       uint32
	module_table       emulator.Uintptr32
}

type tlsModules64 struct {
	generation         uint64
	generation_libc_so uint64
	rwlock             pthread_rwlock64
	module_count       uint64
	module_table       emulator.Uintptr64
}

type passwd32 struct {
	pw_name   emulator.Uintptr32
	pw_passwd emulator.Uintptr32
	pw_uid    uint32
	pw_gid    uint32
	pw_dir    emulator.Uintptr32
	pw_shell  emulator.Uintptr32
}

type passwd64 struct {
	pw_name   emulator.Uintptr64
	pw_passwd emulator.Uintptr64
	pw_uid    uint32
	pw_gid    uint32
	pw_gecos  emulator.Uintptr64
	pw_dir    emulator.Uintptr64
	pw_shell  emulator.Uintptr64
}

type passwd_state32 struct {
	passwd_      passwd32
	name_buffer_ [32]byte
	dir_buffer_  [32]byte
	sh_buffer_   [32]byte
	getpwent_idx int32
}

type passwd_state64 struct {
	passwd_      passwd64
	name_buffer_ [32]byte
	dir_buffer_  [32]byte
	sh_buffer_   [32]byte
	getpwent_idx int64
}

type bionicTls32 struct {
	key_data       [130]pthread_key_data[uint32]
	locale         emulator.Uintptr32
	basename_buf   [4096]byte
	dirname_buf    [4096]byte
	mntent_buf     mntent[emulator.Uintptr32]
	mntent_strings [1024]byte
	ptsname_buf    [32]byte
	ttyname_buf    [64]byte
	strerror_buf   [255]byte
	strsignal_buf  [255]byte
	group          group_state[emulator.Uintptr32]
	passwd         passwd_state32
}

type bionicTls64 struct {
	key_data       [130]pthread_key_data[uint64]
	locale         emulator.Uintptr64
	basename_buf   [4096]byte
	dirname_buf    [4096]byte
	mntent_buf     mntent[emulator.Uintptr64]
	mntent_strings [1024]byte
	ptsname_buf    [32]byte
	ttyname_buf    [64]byte
	strerror_buf   [255]byte
	strsignal_buf  [255]byte
	group          group_state[emulator.Uintptr64]
	passwd         passwd_state64
}

func (layout *staticTlsLayout[P]) reserve_bionic_tls() {
	if unsafe.Sizeof(P(0)) == 4 {
		layout.offset_bionic_tls_ = layout.reserve(P(unsafe.Sizeof(bionicTls32{})), P(unsafe.Alignof(bionicTls32{})))
	} else {
		layout.offset_bionic_tls_ = layout.reserve(P(unsafe.Sizeof(bionicTls64{})), P(unsafe.Alignof(bionicTls64{})))
	}
}

func (layout *staticTlsLayout[P]) reserve(size, alignment P) P {
	offset := align(layout.offset_, alignment)
	if offset < layout.offset_ {
		layout.overflowed_ = true
	}
	layout.offset_ = offset + size
	if layout.offset_ < offset {
		layout.overflowed_ = true
	}
	layout.alignment_ = max(layout.alignment_, alignment)
	return offset
}
