package extend

import (
	"sync"

	android "github.com/wnxd/microdbg-android"
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

type Extend interface {
	SharedPreference(name string, handler SharedPreference)
}

type SharedPreference interface {
	Contains(key string) java.JBoolean
	GetBoolean(key string, defValue java.JBoolean) java.JBoolean
	GetFloat(key string, defValue java.JFloat) java.JFloat
	GetInt(key string, defValue java.JInt) java.JInt
	GetLong(key string, defValue java.JLong) java.JLong
	GetString(key string, defValue java.IString) java.IString
}

type extend struct {
	art android.Runtime
	cf  gava.ClassFactory
	sp  sync.Map
}

func Define(art android.Runtime, cf gava.ClassFactory) Extend {
	ex := &extend{art: art, cf: cf}
	ex.defineLocale()
	ex.defineFile()
	ex.defineBuild()
	ex.defineContext()
	return ex
}

func (ex *extend) SharedPreference(name string, handler SharedPreference) {
	ex.sp.Store(name, handler)
}
