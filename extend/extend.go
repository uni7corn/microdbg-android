package extend

import (
	"sync"

	android "github.com/wnxd/microdbg-android"
	gava "github.com/wnxd/microdbg-android/java"
	"github.com/wnxd/microdbg/debugger"
	"github.com/wnxd/microdbg/socket"
)

type Extend interface {
	EnableDebug()
	Network(handler Network)
	SystemProperties(handler SystemProps)
	SharedPreference(name string, handler SharedPreference)
	RegisterIntent(name string, handler Intent)
}

type extend struct {
	debugger.DefaultFileHandler
	debug  bool
	art    android.Runtime
	cf     gava.ClassFactory
	net    Network
	props  SystemProps
	pref   sync.Map
	intent sync.Map
}

func Define(art android.Runtime, cf gava.ClassFactory) Extend {
	ex := &extend{art: art, cf: cf}
	ex.defineSystem()
	ex.defineIO()
	ex.defineSecurity()
	ex.defineUtil()
	ex.defineCrypto()
	ex.defineException()
	ex.defineBuild()
	ex.defineContent()
	ex.defineApp()
	return ex
}

func (ex *extend) EnableDebug() {
	ex.debug = true
}

func (ex *extend) Network(handler Network) {
	ex.net = handler
}

func (ex *extend) SystemProperties(handler SystemProps) {
	ex.props = handler
}

func (ex *extend) SharedPreference(name string, handler SharedPreference) {
	ex.pref.Store(name, handler)
}

func (ex *extend) RegisterIntent(name string, handler Intent) {
	ex.intent.Store(name, handler)
}

func (ex *extend) NewSocket(network socket.Network) (socket.Socket, error) {
	return virtualSocket{
		Socket: socket.New(network),
		ex:     ex,
	}, nil
}
