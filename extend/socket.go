package extend

import (
	"iter"
	"net"
	"strings"
	"unsafe"

	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/socket"
)

type virtualSocket struct {
	socket.Socket
	ex *extend
}

func (vs virtualSocket) Control(op int, arg any) error {
	const (
		SIOCADDRT    = 0x890B
		SIOCDELRT    = 0x890C
		SIOCRTMSG    = 0x890D
		SIOCGIFNAME  = 0x8910
		SIOCSIFLINK  = 0x8911
		SIOCGIFCONF  = 0x8912
		SIOCGIFFLAGS = 0x8913
	)

	if vs.ex.net != nil {
		switch op {
		case SIOCGIFNAME:
			return vs.gifname(arg)
		case SIOCGIFCONF:
			return vs.gifconf(arg)
		case SIOCGIFFLAGS:
			return vs.gifflags(arg)
		}
	}
	return vs.Socket.Control(op, arg)
}

func (vs virtualSocket) gifname(arg any) error {
	addr, ok := arg.(emuptr)
	if !ok {
		return debugger.ErrArgumentInvalid
	}
	dbg := vs.ex.art.Debugger()
	var req ifreq
	err := dbg.MemExtract(addr, &req)
	if err != nil {
		return err
	}
	ifi, err := vs.ex.net.InterfaceByIndex(int(*req.ifr_ifindex()))
	if err != nil {
		return err
	}
	n := copy(req.ifr_ifrn[:], []byte(ifi.Name))
	if n != IFNAMSIZ {
		req.ifr_ifrn[n] = 0
	}
	dbg.MemWrite(addr, req)
	return nil
}

func (vs virtualSocket) gifconf(arg any) error {
	addr, ok := arg.(emuptr)
	if !ok {
		return debugger.ErrArgumentInvalid
	}
	dbg := vs.ex.art.Debugger()
	var conf ifconf
	err := dbg.MemExtract(addr, &conf)
	if err != nil {
		return err
	}
	ifs, err := vs.ex.net.Interfaces()
	if err != nil {
		return err
	}
	size := int32(IFNAMSIZ + IFNAMSIZ)
	if dbg.PointerSize() == 8 {
		size += 8
	}
	count := conf.ifc_len / size
	ptr := emuptr(conf.ifc_ifcu)
	var i int32
	for ifi, ip := range rangeInterfaces(ifs) {
		if i == count {
			break
		} else if ip = ip.To4(); ip == nil { //skip IPv6
			continue
		}
		var req ifreq
		copy(req.ifr_ifrn[:], []byte(ifi.Name))
		in := (*sockaddr_in)(unsafe.Pointer(req.ifru_addr()))
		in.sa_family = AF_INET
		copy(in.sin_addr[:], ip)
		dbg.MemWrite(ptr, req)
		ptr += emuptr(size)
		i++
	}
	conf.ifc_len = i * size
	dbg.MemWrite(addr, conf)
	return nil
}

func (vs virtualSocket) gifflags(arg any) error {
	addr, ok := arg.(emuptr)
	if !ok {
		return debugger.ErrArgumentInvalid
	}
	dbg := vs.ex.art.Debugger()
	var req ifreq
	err := dbg.MemExtract(addr, &req)
	if err != nil {
		return err
	}
	name := strings.TrimRight(string(req.ifr_ifrn[:]), "\x00")
	ifi, err := vs.ex.net.InterfaceByName(name)
	if err != nil {
		return err
	}
	flags := req.ifr_flags()
	if ifi.Flags&net.FlagUp != 0 {
		*flags |= IFF_UP
	}
	if ifi.Flags&net.FlagBroadcast != 0 {
		*flags |= IFF_BROADCAST
	}
	if ifi.Flags&net.FlagLoopback != 0 {
		*flags |= IFF_LOOPBACK
	}
	if ifi.Flags&net.FlagPointToPoint != 0 {
		*flags |= IFF_POINTOPOINT
	}
	if ifi.Flags&net.FlagMulticast != 0 {
		*flags |= IFF_MULTICAST
	}
	if ifi.Flags&net.FlagRunning != 0 {
		*flags |= IFF_RUNNING
	}
	dbg.MemWrite(addr, req)
	return nil
}

func rangeInterfaces(list []NetworkInterface) iter.Seq2[*NetworkInterface, net.IP] {
	return func(yield func(*NetworkInterface, net.IP) bool) {
		for i := range list {
			for _, addr := range list[i].Addrs {
				var ip net.IP
				if in, ok := addr.(*net.IPNet); ok {
					ip = in.IP
				} else if ia, ok := addr.(*net.IPAddr); ok {
					ip = ia.IP
				}
				if !yield(&list[i], ip) {
					return
				}
			}
		}
	}
}
