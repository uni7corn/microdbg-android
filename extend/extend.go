package extend

import (
	"sync"

	android "github.com/wnxd/microdbg-android"
	gava "github.com/wnxd/microdbg-android/java"
)

type Extend interface {
	EnableDebug()
	SharedPreference(name string, handler SharedPreference)
	RegisterIntent(name string, handler Intent)
}

type extend struct {
	debug  bool
	art    android.Runtime
	cf     gava.ClassFactory
	pref   sync.Map
	intent sync.Map
}

func Define(art android.Runtime, cf gava.ClassFactory) Extend {
	ex := &extend{art: art, cf: cf}
	ex.defineLocale()
	ex.defineFile()
	ex.defineSecurity()
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

func (ex *extend) SharedPreference(name string, handler SharedPreference) {
	ex.pref.Store(name, handler)
}

func (ex *extend) RegisterIntent(name string, handler Intent) {
	ex.intent.Store(name, &intent{
		action:  name,
		handler: handler,
	})
}
