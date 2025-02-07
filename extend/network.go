package extend

import (
	"net"
	"unsafe"
)

const (
	IFNAMSIZ = 16

	AF_INET  = 2
	AF_INET6 = 10
)

const (
	IFF_UP          = 1 << 0
	IFF_BROADCAST   = 1 << 1
	IFF_DEBUG       = 1 << 2
	IFF_LOOPBACK    = 1 << 3
	IFF_POINTOPOINT = 1 << 4
	IFF_NOTRAILERS  = 1 << 5
	IFF_RUNNING     = 1 << 6
	IFF_NOARP       = 1 << 7
	IFF_PROMISC     = 1 << 8
	IFF_ALLMULTI    = 1 << 9
	IFF_MASTER      = 1 << 10
	IFF_SLAVE       = 1 << 11
	IFF_MULTICAST   = 1 << 12
	IFF_PORTSEL     = 1 << 13
	IFF_AUTOMEDIA   = 1 << 14
	IFF_DYNAMIC     = 1 << 15
	IFF_LOWER_UP    = 1 << 16
	IFF_DORMANT     = 1 << 17
	IFF_ECHO        = 1 << 18
)

type Network interface {
	Interfaces() ([]NetworkInterface, error)
	InterfaceByIndex(index int) (*NetworkInterface, error)
	InterfaceByName(name string) (*NetworkInterface, error)
}

type NetworkInterface struct {
	net.Interface
	Addrs []net.Addr
}

type ifconf struct {
	ifc_len  int32
	ifc_ifcu uintptr
}

type ifreq struct {
	ifr_ifrn [IFNAMSIZ]byte
	ifr_ifru [IFNAMSIZ]byte
}

type sockaddr struct {
	sa_family sa_family_t
	sa_data   [14]byte
}

type sockaddr_in struct {
	sa_family sa_family_t
	sin_port  uint16
	sin_addr  [4]byte
}

func (req *ifreq) ifru_addr() *sockaddr {
	return (*sockaddr)(unsafe.Pointer(&req.ifr_ifru))
}

func (req *ifreq) ifr_flags() *int16 {
	return (*int16)(unsafe.Pointer(&req.ifr_ifru))
}

func (req *ifreq) ifr_ifindex() *int32 {
	return (*int32)(unsafe.Pointer(&req.ifr_ifru))
}
