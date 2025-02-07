package extend

import (
	"github.com/Xuanwo/go-locale"
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
	"golang.org/x/text/language"
)

func (ex *extend) defineUtil() {
	Locale := ex.cf.DefineClass("java.util.Locale")
	tag, _ := locale.Detect()
	defaultLocale := Locale.NewObject(tag)
	Locale.DefineMethod("getDefault", "()Ljava/util/Locale;", gava.Modifier_PUBLIC|gava.Modifier_STATIC).BindCall(func(obj java.IObject, args ...any) any {
		return defaultLocale
	})
	Locale.DefineMethod("getLanguage", "()Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		tag := obj.(gava.FakeObject).Value().(language.Tag)
		lang, _ := tag.Base()
		return gava.FakeString(lang.String())
	})
}
