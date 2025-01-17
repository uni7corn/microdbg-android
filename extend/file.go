package extend

import (
	"path/filepath"

	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

func (ex *extend) defineFile() {
	File := ex.cf.DefineClass("java/io/File")
	File.DefineMethod("toString", "()Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(gava.FakeObject)
		name := fake.Value().(string)
		return gava.FakeString(name)
	})
	File.DefineMethod("getName", "()Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(gava.FakeObject)
		name := fake.Value().(string)
		return gava.FakeString(filepath.Base(name))
	})
	File.DefineMethod("getPath", "()Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(gava.FakeObject)
		name := fake.Value().(string)
		return gava.FakeString(name)
	})
	File.DefineMethod("getAbsolutePath", "()Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		fake := obj.(gava.FakeObject)
		name := fake.Value().(string)
		return gava.FakeString(name)
	})
}
