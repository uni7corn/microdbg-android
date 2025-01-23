package x86_64

import linux "github.com/wnxd/microdbg-linux"

const (
	NR_read                   = 0
	NR_write                  = 1
	NR_open                   = 2
	NR_close                  = 3
	NR_stat                   = 4
	NR_fstat                  = 5
	NR_lstat                  = 6
	NR_poll                   = 7
	NR_lseek                  = 8
	NR_mmap                   = 9
	NR_mprotect               = 10
	NR_munmap                 = 11
	NR_brk                    = 12
	NR_rt_sigaction           = 13
	NR_rt_sigprocmask         = 14
	NR_rt_sigreturn           = 15
	NR_ioctl                  = 16
	NR_pread64                = 17
	NR_pwrite64               = 18
	NR_readv                  = 19
	NR_writev                 = 20
	NR_access                 = 21
	NR_pipe                   = 22
	NR_select                 = 23
	NR_sched_yield            = 24
	NR_mremap                 = 25
	NR_msync                  = 26
	NR_mincore                = 27
	NR_madvise                = 28
	NR_shmget                 = 29
	NR_shmat                  = 30
	NR_shmctl                 = 31
	NR_dup                    = 32
	NR_dup2                   = 33
	NR_pause                  = 34
	NR_nanosleep              = 35
	NR_getitimer              = 36
	NR_alarm                  = 37
	NR_setitimer              = 38
	NR_getpid                 = 39
	NR_sendfile               = 40
	NR_socket                 = 41
	NR_connect                = 42
	NR_accept                 = 43
	NR_sendto                 = 44
	NR_recvfrom               = 45
	NR_sendmsg                = 46
	NR_recvmsg                = 47
	NR_shutdown               = 48
	NR_bind                   = 49
	NR_listen                 = 50
	NR_getsockname            = 51
	NR_getpeername            = 52
	NR_socketpair             = 53
	NR_setsockopt             = 54
	NR_getsockopt             = 55
	NR_clone                  = 56
	NR_fork                   = 57
	NR_vfork                  = 58
	NR_execve                 = 59
	NR_exit                   = 60
	NR_wait4                  = 61
	NR_kill                   = 62
	NR_uname                  = 63
	NR_semget                 = 64
	NR_semop                  = 65
	NR_semctl                 = 66
	NR_shmdt                  = 67
	NR_msgget                 = 68
	NR_msgsnd                 = 69
	NR_msgrcv                 = 70
	NR_msgctl                 = 71
	NR_fcntl                  = 72
	NR_flock                  = 73
	NR_fsync                  = 74
	NR_fdatasync              = 75
	NR_truncate               = 76
	NR_ftruncate              = 77
	NR_getdents               = 78
	NR_getcwd                 = 79
	NR_chdir                  = 80
	NR_fchdir                 = 81
	NR_rename                 = 82
	NR_mkdir                  = 83
	NR_rmdir                  = 84
	NR_creat                  = 85
	NR_link                   = 86
	NR_unlink                 = 87
	NR_symlink                = 88
	NR_readlink               = 89
	NR_chmod                  = 90
	NR_fchmod                 = 91
	NR_chown                  = 92
	NR_fchown                 = 93
	NR_lchown                 = 94
	NR_umask                  = 95
	NR_gettimeofday           = 96
	NR_getrlimit              = 97
	NR_getrusage              = 98
	NR_sysinfo                = 99
	NR_times                  = 100
	NR_ptrace                 = 101
	NR_getuid                 = 102
	NR_syslog                 = 103
	NR_getgid                 = 104
	NR_setuid                 = 105
	NR_setgid                 = 106
	NR_geteuid                = 107
	NR_getegid                = 108
	NR_setpgid                = 109
	NR_getppid                = 110
	NR_getpgrp                = 111
	NR_setsid                 = 112
	NR_setreuid               = 113
	NR_setregid               = 114
	NR_getgroups              = 115
	NR_setgroups              = 116
	NR_setresuid              = 117
	NR_getresuid              = 118
	NR_setresgid              = 119
	NR_getresgid              = 120
	NR_getpgid                = 121
	NR_setfsuid               = 122
	NR_setfsgid               = 123
	NR_getsid                 = 124
	NR_capget                 = 125
	NR_capset                 = 126
	NR_rt_sigpending          = 127
	NR_rt_sigtimedwait        = 128
	NR_rt_sigqueueinfo        = 129
	NR_rt_sigsuspend          = 130
	NR_sigaltstack            = 131
	NR_utime                  = 132
	NR_mknod                  = 133
	NR_uselib                 = 134
	NR_personality            = 135
	NR_ustat                  = 136
	NR_statfs                 = 137
	NR_fstatfs                = 138
	NR_sysfs                  = 139
	NR_getpriority            = 140
	NR_setpriority            = 141
	NR_sched_setparam         = 142
	NR_sched_getparam         = 143
	NR_sched_setscheduler     = 144
	NR_sched_getscheduler     = 145
	NR_sched_get_priority_max = 146
	NR_sched_get_priority_min = 147
	NR_sched_rr_get_interval  = 148
	NR_mlock                  = 149
	NR_munlock                = 150
	NR_mlockall               = 151
	NR_munlockall             = 152
	NR_vhangup                = 153
	NR_modify_ldt             = 154
	NR_pivot_root             = 155
	NR__sysctl                = 156
	NR_prctl                  = 157
	NR_arch_prctl             = 158
	NR_adjtimex               = 159
	NR_setrlimit              = 160
	NR_chroot                 = 161
	NR_sync                   = 162
	NR_acct                   = 163
	NR_settimeofday           = 164
	NR_mount                  = 165
	NR_umount2                = 166
	NR_swapon                 = 167
	NR_swapoff                = 168
	NR_reboot                 = 169
	NR_sethostname            = 170
	NR_setdomainname          = 171
	NR_iopl                   = 172
	NR_ioperm                 = 173
	NR_create_module          = 174
	NR_init_module            = 175
	NR_delete_module          = 176
	NR_get_kernel_syms        = 177
	NR_query_module           = 178
	NR_quotactl               = 179
	NR_nfsservctl             = 180
	NR_getpmsg                = 181
	NR_putpmsg                = 182
	NR_afs_syscall            = 183
	NR_tuxcall                = 184
	NR_security               = 185
	NR_gettid                 = 186
	NR_readahead              = 187
	NR_setxattr               = 188
	NR_lsetxattr              = 189
	NR_fsetxattr              = 190
	NR_getxattr               = 191
	NR_lgetxattr              = 192
	NR_fgetxattr              = 193
	NR_listxattr              = 194
	NR_llistxattr             = 195
	NR_flistxattr             = 196
	NR_removexattr            = 197
	NR_lremovexattr           = 198
	NR_fremovexattr           = 199
	NR_tkill                  = 200
	NR_time                   = 201
	NR_futex                  = 202
	NR_sched_setaffinity      = 203
	NR_sched_getaffinity      = 204
	NR_set_thread_area        = 205
	NR_io_setup               = 206
	NR_io_destroy             = 207
	NR_io_getevents           = 208
	NR_io_submit              = 209
	NR_io_cancel              = 210
	NR_get_thread_area        = 211
	NR_lookup_dcookie         = 212
	NR_epoll_create           = 213
	NR_epoll_ctl_old          = 214
	NR_epoll_wait_old         = 215
	NR_remap_file_pages       = 216
	NR_getdents64             = 217
	NR_set_tid_address        = 218
	NR_restart_syscall        = 219
	NR_semtimedop             = 220
	NR_fadvise64              = 221
	NR_timer_create           = 222
	NR_timer_settime          = 223
	NR_timer_gettime          = 224
	NR_timer_getoverrun       = 225
	NR_timer_delete           = 226
	NR_clock_settime          = 227
	NR_clock_gettime          = 228
	NR_clock_getres           = 229
	NR_clock_nanosleep        = 230
	NR_exit_group             = 231
	NR_epoll_wait             = 232
	NR_epoll_ctl              = 233
	NR_tgkill                 = 234
	NR_utimes                 = 235
	NR_vserver                = 236
	NR_mbind                  = 237
	NR_set_mempolicy          = 238
	NR_get_mempolicy          = 239
	NR_mq_open                = 240
	NR_mq_unlink              = 241
	NR_mq_timedsend           = 242
	NR_mq_timedreceive        = 243
	NR_mq_notify              = 244
	NR_mq_getsetattr          = 245
	NR_kexec_load             = 246
	NR_waitid                 = 247
	NR_add_key                = 248
	NR_request_key            = 249
	NR_keyctl                 = 250
	NR_ioprio_set             = 251
	NR_ioprio_get             = 252
	NR_inotify_init           = 253
	NR_inotify_add_watch      = 254
	NR_inotify_rm_watch       = 255
	NR_migrate_pages          = 256
	NR_openat                 = 257
	NR_mkdirat                = 258
	NR_mknodat                = 259
	NR_fchownat               = 260
	NR_futimesat              = 261
	NR_newfstatat             = 262
	NR_unlinkat               = 263
	NR_renameat               = 264
	NR_linkat                 = 265
	NR_symlinkat              = 266
	NR_readlinkat             = 267
	NR_fchmodat               = 268
	NR_faccessat              = 269
	NR_pselect6               = 270
	NR_ppoll                  = 271
	NR_unshare                = 272
	NR_set_robust_list        = 273
	NR_get_robust_list        = 274
	NR_splice                 = 275
	NR_tee                    = 276
	NR_sync_file_range        = 277
	NR_vmsplice               = 278
	NR_move_pages             = 279
	NR_utimensat              = 280
	NR_epoll_pwait            = 281
	NR_signalfd               = 282
	NR_timerfd_create         = 283
	NR_eventfd                = 284
	NR_fallocate              = 285
	NR_timerfd_settime        = 286
	NR_timerfd_gettime        = 287
	NR_accept4                = 288
	NR_signalfd4              = 289
	NR_eventfd2               = 290
	NR_epoll_create1          = 291
	NR_dup3                   = 292
	NR_pipe2                  = 293
	NR_inotify_init1          = 294
	NR_preadv                 = 295
	NR_pwritev                = 296
	NR_rt_tgsigqueueinfo      = 297
	NR_perf_event_open        = 298
	NR_recvmmsg               = 299
	NR_fanotify_init          = 300
	NR_fanotify_mark          = 301
	NR_prlimit64              = 302
	NR_name_to_handle_at      = 303
	NR_open_by_handle_at      = 304
	NR_clock_adjtime          = 305
	NR_syncfs                 = 306
	NR_sendmmsg               = 307
	NR_setns                  = 308
	NR_getcpu                 = 309
	NR_process_vm_readv       = 310
	NR_process_vm_writev      = 311
	NR_kcmp                   = 312
	NR_finit_module           = 313
	NR_sched_setattr          = 314
	NR_sched_getattr          = 315
	NR_renameat2              = 316
	NR_seccomp                = 317
	NR_getrandom              = 318
	NR_memfd_create           = 319
	NR_kexec_file_load        = 320
	NR_bpf                    = 321
	NR_execveat               = 322
	NR_userfaultfd            = 323
	NR_membarrier             = 324
	NR_mlock2                 = 325
	NR_copy_file_range        = 326
	NR_preadv2                = 327
	NR_pwritev2               = 328
	NR_pkey_mprotect          = 329
	NR_pkey_alloc             = 330
	NR_pkey_free              = 331
	NR_statx                  = 332
	NR_io_pgetevents          = 333
	NR_rseq                   = 334
)

var (
	NRMap = map[uint64]linux.NR{
		NR_io_setup:               linux.NR_io_setup,
		NR_io_destroy:             linux.NR_io_destroy,
		NR_io_submit:              linux.NR_io_submit,
		NR_io_cancel:              linux.NR_io_cancel,
		NR_io_getevents:           linux.NR_io_getevents,
		NR_setxattr:               linux.NR_setxattr,
		NR_lsetxattr:              linux.NR_lsetxattr,
		NR_fsetxattr:              linux.NR_fsetxattr,
		NR_getxattr:               linux.NR_getxattr,
		NR_lgetxattr:              linux.NR_lgetxattr,
		NR_fgetxattr:              linux.NR_fgetxattr,
		NR_listxattr:              linux.NR_listxattr,
		NR_llistxattr:             linux.NR_llistxattr,
		NR_flistxattr:             linux.NR_flistxattr,
		NR_removexattr:            linux.NR_removexattr,
		NR_lremovexattr:           linux.NR_lremovexattr,
		NR_fremovexattr:           linux.NR_fremovexattr,
		NR_getcwd:                 linux.NR_getcwd,
		NR_lookup_dcookie:         linux.NR_lookup_dcookie,
		NR_eventfd2:               linux.NR_eventfd2,
		NR_epoll_create1:          linux.NR_epoll_create1,
		NR_epoll_ctl:              linux.NR_epoll_ctl,
		NR_epoll_pwait:            linux.NR_epoll_pwait,
		NR_dup:                    linux.NR_dup,
		NR_dup3:                   linux.NR_dup3,
		NR_fcntl:                  linux.NR_fcntl,
		NR_inotify_init1:          linux.NR_inotify_init1,
		NR_inotify_add_watch:      linux.NR_inotify_add_watch,
		NR_inotify_rm_watch:       linux.NR_inotify_rm_watch,
		NR_ioctl:                  linux.NR_ioctl,
		NR_ioprio_set:             linux.NR_ioprio_set,
		NR_ioprio_get:             linux.NR_ioprio_get,
		NR_flock:                  linux.NR_flock,
		NR_mknodat:                linux.NR_mknodat,
		NR_mkdirat:                linux.NR_mkdirat,
		NR_unlinkat:               linux.NR_unlinkat,
		NR_symlinkat:              linux.NR_symlinkat,
		NR_linkat:                 linux.NR_linkat,
		NR_renameat:               linux.NR_renameat,
		NR_umount2:                linux.NR_umount2,
		NR_mount:                  linux.NR_mount,
		NR_pivot_root:             linux.NR_pivot_root,
		NR_nfsservctl:             linux.NR_nfsservctl,
		NR_statfs:                 linux.NR_statfs,
		NR_fstatfs:                linux.NR_fstatfs,
		NR_truncate:               linux.NR_truncate,
		NR_ftruncate:              linux.NR_ftruncate,
		NR_fallocate:              linux.NR_fallocate,
		NR_faccessat:              linux.NR_faccessat,
		NR_chdir:                  linux.NR_chdir,
		NR_fchdir:                 linux.NR_fchdir,
		NR_chroot:                 linux.NR_chroot,
		NR_fchmod:                 linux.NR_fchmod,
		NR_fchmodat:               linux.NR_fchmodat,
		NR_fchownat:               linux.NR_fchownat,
		NR_fchown:                 linux.NR_fchown,
		NR_open:                   linux.NR_open,
		NR_openat:                 linux.NR_openat,
		NR_close:                  linux.NR_close,
		NR_vhangup:                linux.NR_vhangup,
		NR_pipe2:                  linux.NR_pipe2,
		NR_quotactl:               linux.NR_quotactl,
		NR_getdents64:             linux.NR_getdents64,
		NR_lseek:                  linux.NR_lseek,
		NR_read:                   linux.NR_read,
		NR_write:                  linux.NR_write,
		NR_readv:                  linux.NR_readv,
		NR_writev:                 linux.NR_writev,
		NR_pread64:                linux.NR_pread64,
		NR_pwrite64:               linux.NR_pwrite64,
		NR_preadv:                 linux.NR_preadv,
		NR_pwritev:                linux.NR_pwritev,
		NR_sendfile:               linux.NR_sendfile,
		NR_pselect6:               linux.NR_pselect6,
		NR_ppoll:                  linux.NR_ppoll,
		NR_signalfd4:              linux.NR_signalfd4,
		NR_vmsplice:               linux.NR_vmsplice,
		NR_splice:                 linux.NR_splice,
		NR_tee:                    linux.NR_tee,
		NR_readlinkat:             linux.NR_readlinkat,
		NR_newfstatat:             linux.NR_fstatat64,
		NR_fstat:                  linux.NR_fstat64,
		NR_sync:                   linux.NR_sync,
		NR_fsync:                  linux.NR_fsync,
		NR_fdatasync:              linux.NR_fdatasync,
		NR_sync_file_range:        linux.NR_sync_file_range,
		NR_timerfd_create:         linux.NR_timerfd_create,
		NR_timerfd_settime:        linux.NR_timerfd_settime,
		NR_timerfd_gettime:        linux.NR_timerfd_gettime,
		NR_utimensat:              linux.NR_utimensat,
		NR_acct:                   linux.NR_acct,
		NR_capget:                 linux.NR_capget,
		NR_capset:                 linux.NR_capset,
		NR_personality:            linux.NR_personality,
		NR_exit:                   linux.NR_exit,
		NR_exit_group:             linux.NR_exit_group,
		NR_waitid:                 linux.NR_waitid,
		NR_set_tid_address:        linux.NR_set_tid_address,
		NR_unshare:                linux.NR_unshare,
		NR_futex:                  linux.NR_futex,
		NR_set_robust_list:        linux.NR_set_robust_list,
		NR_get_robust_list:        linux.NR_get_robust_list,
		NR_nanosleep:              linux.NR_nanosleep,
		NR_getitimer:              linux.NR_getitimer,
		NR_setitimer:              linux.NR_setitimer,
		NR_kexec_load:             linux.NR_kexec_load,
		NR_init_module:            linux.NR_init_module,
		NR_delete_module:          linux.NR_delete_module,
		NR_timer_create:           linux.NR_timer_create,
		NR_timer_gettime:          linux.NR_timer_gettime,
		NR_timer_getoverrun:       linux.NR_timer_getoverrun,
		NR_timer_settime:          linux.NR_timer_settime,
		NR_timer_delete:           linux.NR_timer_delete,
		NR_clock_settime:          linux.NR_clock_settime,
		NR_clock_gettime:          linux.NR_clock_gettime,
		NR_clock_getres:           linux.NR_clock_getres,
		NR_clock_nanosleep:        linux.NR_clock_nanosleep,
		NR_syslog:                 linux.NR_syslog,
		NR_ptrace:                 linux.NR_ptrace,
		NR_sched_setparam:         linux.NR_sched_setparam,
		NR_sched_setscheduler:     linux.NR_sched_setscheduler,
		NR_sched_getscheduler:     linux.NR_sched_getscheduler,
		NR_sched_getparam:         linux.NR_sched_getparam,
		NR_sched_setaffinity:      linux.NR_sched_setaffinity,
		NR_sched_getaffinity:      linux.NR_sched_getaffinity,
		NR_sched_yield:            linux.NR_sched_yield,
		NR_sched_get_priority_max: linux.NR_sched_get_priority_max,
		NR_sched_get_priority_min: linux.NR_sched_get_priority_min,
		NR_sched_rr_get_interval:  linux.NR_sched_rr_get_interval,
		NR_restart_syscall:        linux.NR_restart_syscall,
		NR_kill:                   linux.NR_kill,
		NR_tkill:                  linux.NR_tkill,
		NR_tgkill:                 linux.NR_tgkill,
		NR_sigaltstack:            linux.NR_sigaltstack,
		NR_rt_sigsuspend:          linux.NR_rt_sigsuspend,
		NR_rt_sigaction:           linux.NR_rt_sigaction,
		NR_rt_sigprocmask:         linux.NR_rt_sigprocmask,
		NR_rt_sigpending:          linux.NR_rt_sigpending,
		NR_rt_sigtimedwait:        linux.NR_rt_sigtimedwait,
		NR_rt_sigqueueinfo:        linux.NR_rt_sigqueueinfo,
		NR_rt_sigreturn:           linux.NR_rt_sigreturn,
		NR_setpriority:            linux.NR_setpriority,
		NR_getpriority:            linux.NR_getpriority,
		NR_reboot:                 linux.NR_reboot,
		NR_setregid:               linux.NR_setregid,
		NR_setgid:                 linux.NR_setgid,
		NR_setreuid:               linux.NR_setreuid,
		NR_setuid:                 linux.NR_setuid,
		NR_setresuid:              linux.NR_setresuid,
		NR_getresuid:              linux.NR_getresuid,
		NR_setresgid:              linux.NR_setresgid,
		NR_getresgid:              linux.NR_getresgid,
		NR_setfsuid:               linux.NR_setfsuid,
		NR_setfsgid:               linux.NR_setfsgid,
		NR_times:                  linux.NR_times,
		NR_setpgid:                linux.NR_setpgid,
		NR_getpgid:                linux.NR_getpgid,
		NR_getsid:                 linux.NR_getsid,
		NR_setsid:                 linux.NR_setsid,
		NR_getgroups:              linux.NR_getgroups,
		NR_setgroups:              linux.NR_setgroups,
		NR_uname:                  linux.NR_uname,
		NR_sethostname:            linux.NR_sethostname,
		NR_setdomainname:          linux.NR_setdomainname,
		NR_getrlimit:              linux.NR_getrlimit,
		NR_setrlimit:              linux.NR_setrlimit,
		NR_getrusage:              linux.NR_getrusage,
		NR_umask:                  linux.NR_umask,
		NR_prctl:                  linux.NR_prctl,
		NR_getcpu:                 linux.NR_getcpu,
		NR_gettimeofday:           linux.NR_gettimeofday,
		NR_settimeofday:           linux.NR_settimeofday,
		NR_adjtimex:               linux.NR_adjtimex,
		NR_getpid:                 linux.NR_getpid,
		NR_getppid:                linux.NR_getppid,
		NR_getuid:                 linux.NR_getuid,
		NR_geteuid:                linux.NR_geteuid,
		NR_getgid:                 linux.NR_getgid,
		NR_getegid:                linux.NR_getegid,
		NR_gettid:                 linux.NR_gettid,
		NR_sysinfo:                linux.NR_sysinfo,
		NR_mq_open:                linux.NR_mq_open,
		NR_mq_unlink:              linux.NR_mq_unlink,
		NR_mq_timedsend:           linux.NR_mq_timedsend,
		NR_mq_timedreceive:        linux.NR_mq_timedreceive,
		NR_mq_notify:              linux.NR_mq_notify,
		NR_mq_getsetattr:          linux.NR_mq_getsetattr,
		NR_socket:                 linux.NR_socket,
		NR_socketpair:             linux.NR_socketpair,
		NR_bind:                   linux.NR_bind,
		NR_listen:                 linux.NR_listen,
		NR_connect:                linux.NR_connect,
		NR_getsockname:            linux.NR_getsockname,
		NR_getpeername:            linux.NR_getpeername,
		NR_sendto:                 linux.NR_sendto,
		NR_recvfrom:               linux.NR_recvfrom,
		NR_setsockopt:             linux.NR_setsockopt,
		NR_getsockopt:             linux.NR_getsockopt,
		NR_shutdown:               linux.NR_shutdown,
		NR_sendmsg:                linux.NR_sendmsg,
		NR_recvmsg:                linux.NR_recvmsg,
		NR_readahead:              linux.NR_readahead,
		NR_brk:                    linux.NR_brk,
		NR_munmap:                 linux.NR_munmap,
		NR_mremap:                 linux.NR_mremap,
		NR_add_key:                linux.NR_add_key,
		NR_request_key:            linux.NR_request_key,
		NR_keyctl:                 linux.NR_keyctl,
		NR_clone:                  linux.NR_clone,
		NR_execve:                 linux.NR_execve,
		NR_mmap:                   linux.NR_mmap,
		NR_fadvise64:              linux.NR_fadvise64,
		NR_swapon:                 linux.NR_swapon,
		NR_swapoff:                linux.NR_swapoff,
		NR_mprotect:               linux.NR_mprotect,
		NR_msync:                  linux.NR_msync,
		NR_mlock:                  linux.NR_mlock,
		NR_munlock:                linux.NR_munlock,
		NR_mlockall:               linux.NR_mlockall,
		NR_munlockall:             linux.NR_munlockall,
		NR_mincore:                linux.NR_mincore,
		NR_madvise:                linux.NR_madvise,
		NR_remap_file_pages:       linux.NR_remap_file_pages,
		NR_mbind:                  linux.NR_mbind,
		NR_get_mempolicy:          linux.NR_get_mempolicy,
		NR_set_mempolicy:          linux.NR_set_mempolicy,
		NR_migrate_pages:          linux.NR_migrate_pages,
		NR_move_pages:             linux.NR_move_pages,
		NR_rt_tgsigqueueinfo:      linux.NR_rt_tgsigqueueinfo,
		NR_perf_event_open:        linux.NR_perf_event_open,
		NR_accept4:                linux.NR_accept4,
		NR_recvmmsg:               linux.NR_recvmmsg,
		NR_wait4:                  linux.NR_wait4,
		NR_prlimit64:              linux.NR_prlimit64,
		NR_fanotify_init:          linux.NR_fanotify_init,
		NR_fanotify_mark:          linux.NR_fanotify_mark,
		NR_name_to_handle_at:      linux.NR_name_to_handle_at,
		NR_open_by_handle_at:      linux.NR_open_by_handle_at,
		NR_clock_adjtime:          linux.NR_clock_adjtime,
		NR_syncfs:                 linux.NR_syncfs,
		NR_setns:                  linux.NR_setns,
		NR_sendmmsg:               linux.NR_sendmmsg,
		NR_process_vm_readv:       linux.NR_process_vm_readv,
		NR_process_vm_writev:      linux.NR_process_vm_writev,
		NR_kcmp:                   linux.NR_kcmp,
		NR_finit_module:           linux.NR_finit_module,
		NR_sched_setattr:          linux.NR_sched_setattr,
		NR_sched_getattr:          linux.NR_sched_getattr,
		NR_renameat2:              linux.NR_renameat2,
		NR_seccomp:                linux.NR_seccomp,
		NR_getrandom:              linux.NR_getrandom,
		NR_memfd_create:           linux.NR_memfd_create,
		NR_bpf:                    linux.NR_bpf,
		NR_execveat:               linux.NR_execveat,
		NR_userfaultfd:            linux.NR_userfaultfd,
		NR_membarrier:             linux.NR_membarrier,
		NR_mlock2:                 linux.NR_mlock2,
		NR_copy_file_range:        linux.NR_copy_file_range,
		NR_preadv2:                linux.NR_preadv2,
		NR_pwritev2:               linux.NR_pwritev2,
		NR_pkey_mprotect:          linux.NR_pkey_mprotect,
		NR_pkey_alloc:             linux.NR_pkey_alloc,
		NR_pkey_free:              linux.NR_pkey_free,
		NR_statx:                  linux.NR_statx,
		NR_io_pgetevents:          linux.NR_io_pgetevents,
		NR_rseq:                   linux.NR_rseq,
	}
)
