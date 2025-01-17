package extend

import (
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
	"golang.org/x/text/language"
)

func (ex *extend) defineLocale() {
	Locale := ex.cf.DefineClass("java.util.Locale")
	Locale.DefineMethod("getLanguage", "()Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		tag := obj.(gava.FakeObject).Value().(language.Tag)
		lang, _ := tag.Base()
		return gava.FakeString(lang.String())
	})
}
