package extend

import (
	"log"
	"strings"

	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

type SystemProps interface {
	Get(key string) (string, bool)
}

func (ex *extend) defineSystem() {
	getProperty := func(key string, defValue java.IString) java.IString {
		if ex.props == nil {
		} else if prop, ok := ex.props.Get(key); ok {
			return gava.FakeString(prop)
		}
		switch key {
		case "http.agent":
			return getHttpAgent(ex.cf)
		}
		if ex.debug {
			log.Printf("[%s] System property undefined\n", key)
		}
		return defValue
	}

	System := ex.cf.DefineClass("java.lang.System")
	System.DefineMethod("getProperty", "(Ljava/lang/String;)Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_STATIC).BindCall(func(obj java.IObject, args ...any) any {
		key := args[0].(java.IString).String()
		return getProperty(key, nil)
	})
	System.DefineMethod("getProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_STATIC).BindCall(func(obj java.IObject, args ...any) any {
		key := args[0].(java.IString).String()
		defValue := gava.ToObject[java.IString](args[1])
		return getProperty(key, defValue)
	})
}

func getHttpAgent(cf gava.ClassFactory) java.IString {
	Build := cf.GetClass("android.os.Build")
	Build_VERSION := cf.GetClass("android.os.Build$VERSION")
	var result strings.Builder
	result.WriteString("Dalvik/2.1.0 (Linux; U; Android ")
	version := gava.ToObject[java.IString](Build_VERSION.GetField("RELEASE", "Ljava/lang/String;").Get(Build_VERSION))
	if version == nil || version.Length() == 0 {
		result.WriteString("1.0")
	} else {
		result.WriteString(version.String())
	}
	if gava.FakeString("REL").Equals(Build_VERSION.GetField("CODENAME", "Ljava/lang/String;").Get(Build_VERSION)) {
		model := gava.ToObject[java.IString](Build.GetField("MODEL", "Ljava/lang/String;").Get(Build))
		if model != nil && model.Length() > 0 {
			result.WriteString("; ")
			result.WriteString(model.String())
		}
	}
	id := gava.ToObject[java.IString](Build.GetField("ID", "Ljava/lang/String;").Get(Build))
	if id != nil && id.Length() > 0 {
		result.WriteString(" Build/")
		result.WriteString(id.String())
	}
	result.WriteByte(')')
	return gava.FakeString(result.String())
}
