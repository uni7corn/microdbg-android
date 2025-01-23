package arm64

import linux "github.com/wnxd/microdbg-linux"

const (
	NR_io_setup = iota
	NR_io_destroy
	NR_io_submit
	NR_io_cancel
	NR_io_getevents
	NR_setxattr
	NR_lsetxattr
	NR_fsetxattr
	NR_getxattr
	NR_lgetxattr
	NR_fgetxattr
	NR_listxattr
	NR_llistxattr
	NR_flistxattr
	NR_removexattr
	NR_lremovexattr
	NR_fremovexattr
	NR_getcwd
	NR_lookup_dcookie
	NR_eventfd2
	NR_epoll_create1
	NR_epoll_ctl
	NR_epoll_pwait
	NR_dup
	NR_dup3
	NR_fcntl
	NR_inotify_init1
	NR_inotify_add_watch
	NR_inotify_rm_watch
	NR_ioctl
	NR_ioprio_set
	NR_ioprio_get
	NR_flock
	NR_mknodat
	NR_mkdirat
	NR_unlinkat
	NR_symlinkat
	NR_linkat
	NR_renameat
	NR_umount2
	NR_mount
	NR_pivot_root
	NR_nfsservctl
	NR_statfs
	NR_fstatfs
	NR_truncate
	NR_ftruncate
	NR_fallocate
	NR_faccessat
	NR_chdir
	NR_fchdir
	NR_chroot
	NR_fchmod
	NR_fchmodat
	NR_fchownat
	NR_fchown
	NR_openat
	NR_close
	NR_vhangup
	NR_pipe2
	NR_quotactl
	NR_getdents64
	NR_lseek
	NR_read
	NR_write
	NR_readv
	NR_writev
	NR_pread64
	NR_pwrite64
	NR_preadv
	NR_pwritev
	NR_sendfile
	NR_pselect6
	NR_ppoll
	NR_signalfd4
	NR_vmsplice
	NR_splice
	NR_tee
	NR_readlinkat
	NR_newfstatat
	NR_fstat
	NR_sync
	NR_fsync
	NR_fdatasync
	NR_sync_file_range
	NR_timerfd_create
	NR_timerfd_settime
	NR_timerfd_gettime
	NR_utimensat
	NR_acct
	NR_capget
	NR_capset
	NR_personality
	NR_exit
	NR_exit_group
	NR_waitid
	NR_set_tid_address
	NR_unshare
	NR_futex
	NR_set_robust_list
	NR_get_robust_list
	NR_nanosleep
	NR_getitimer
	NR_setitimer
	NR_kexec_load
	NR_init_module
	NR_delete_module
	NR_timer_create
	NR_timer_gettime
	NR_timer_getoverrun
	NR_timer_settime
	NR_timer_delete
	NR_clock_settime
	NR_clock_gettime
	NR_clock_getres
	NR_clock_nanosleep
	NR_syslog
	NR_ptrace
	NR_sched_setparam
	NR_sched_setscheduler
	NR_sched_getscheduler
	NR_sched_getparam
	NR_sched_setaffinity
	NR_sched_getaffinity
	NR_sched_yield
	NR_sched_get_priority_max
	NR_sched_get_priority_min
	NR_sched_rr_get_interval
	NR_restart_syscall
	NR_kill
	NR_tkill
	NR_tgkill
	NR_sigaltstack
	NR_rt_sigsuspend
	NR_rt_sigaction
	NR_rt_sigprocmask
	NR_rt_sigpending
	NR_rt_sigtimedwait
	NR_rt_sigqueueinfo
	NR_rt_sigreturn
	NR_setpriority
	NR_getpriority
	NR_reboot
	NR_setregid
	NR_setgid
	NR_setreuid
	NR_setuid
	NR_setresuid
	NR_getresuid
	NR_setresgid
	NR_getresgid
	NR_setfsuid
	NR_setfsgid
	NR_times
	NR_setpgid
	NR_getpgid
	NR_getsid
	NR_setsid
	NR_getgroups
	NR_setgroups
	NR_uname
	NR_sethostname
	NR_setdomainname
	NR_getrlimit
	NR_setrlimit
	NR_getrusage
	NR_umask
	NR_prctl
	NR_getcpu
	NR_gettimeofday
	NR_settimeofday
	NR_adjtimex
	NR_getpid
	NR_getppid
	NR_getuid
	NR_geteuid
	NR_getgid
	NR_getegid
	NR_gettid
	NR_sysinfo
	NR_mq_open
	NR_mq_unlink
	NR_mq_timedsend
	NR_mq_timedreceive
	NR_mq_notify
	NR_mq_getsetattr
	NR_msgget
	NR_msgctl
	NR_msgrcv
	NR_msgsnd
	NR_semget
	NR_semctl
	NR_semtimedop
	NR_semop
	NR_shmget
	NR_shmctl
	NR_shmat
	NR_shmdt
	NR_socket
	NR_socketpair
	NR_bind
	NR_listen
	NR_accept
	NR_connect
	NR_getsockname
	NR_getpeername
	NR_sendto
	NR_recvfrom
	NR_setsockopt
	NR_getsockopt
	NR_shutdown
	NR_sendmsg
	NR_recvmsg
	NR_readahead
	NR_brk
	NR_munmap
	NR_mremap
	NR_add_key
	NR_request_key
	NR_keyctl
	NR_clone
	NR_execve
	NR_mmap
	NR_fadvise64
	NR_swapon
	NR_swapoff
	NR_mprotect
	NR_msync
	NR_mlock
	NR_munlock
	NR_mlockall
	NR_munlockall
	NR_mincore
	NR_madvise
	NR_remap_file_pages
	NR_mbind
	NR_get_mempolicy
	NR_set_mempolicy
	NR_migrate_pages
	NR_move_pages
	NR_rt_tgsigqueueinfo
	NR_perf_event_open
	NR_accept4
	NR_recvmmsg
	NR_arch_specific_syscall
)
const (
	NR_wait4 = 260 + iota
	NR_prlimit64
	NR_fanotify_init
	NR_fanotify_mark
	NR_name_to_handle_at
	NR_open_by_handle_at
	NR_clock_adjtime
	NR_syncfs
	NR_setns
	NR_sendmmsg
	NR_process_vm_readv
	NR_process_vm_writev
	NR_kcmp
	NR_finit_module
	NR_sched_setattr
	NR_sched_getattr
	NR_renameat2
	NR_seccomp
	NR_getrandom
	NR_memfd_create
	NR_bpf
	NR_execveat
	NR_userfaultfd
	NR_membarrier
	NR_mlock2
	NR_copy_file_range
	NR_preadv2
	NR_pwritev2
	NR_pkey_mprotect
	NR_pkey_alloc
	NR_pkey_free
	NR_statx
	NR_io_pgetevents
	NR_rseq
	NR_kexec_file_load

	NR_syscalls
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
		NR_msgget:                 linux.NR_msgget,
		NR_msgctl:                 linux.NR_msgctl,
		NR_msgrcv:                 linux.NR_msgrcv,
		NR_msgsnd:                 linux.NR_msgsnd,
		NR_semget:                 linux.NR_semget,
		NR_semctl:                 linux.NR_semctl,
		NR_semtimedop:             linux.NR_semtimedop,
		NR_semop:                  linux.NR_semop,
		NR_shmget:                 linux.NR_shmget,
		NR_shmctl:                 linux.NR_shmctl,
		NR_shmat:                  linux.NR_shmat,
		NR_shmdt:                  linux.NR_shmdt,
		NR_socket:                 linux.NR_socket,
		NR_socketpair:             linux.NR_socketpair,
		NR_bind:                   linux.NR_bind,
		NR_listen:                 linux.NR_listen,
		NR_accept:                 linux.NR_accept,
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
		NR_arch_specific_syscall:  linux.NR_arch_specific_syscall,
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
		NR_kexec_file_load:        linux.NR_kexec_file_load,
	}
)
