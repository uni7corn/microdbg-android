package extend

import gava "github.com/wnxd/microdbg-android/java"

func (ex *extend) defineSecurity() {
	Key := ex.cf.DefineClass("java.security.Key")
	Key.SetModifiers(gava.Modifier_PUBLIC | gava.Modifier_INTERFACE | gava.Modifier_ABSTRACT)
	Key.DefineMethod("getAlgorithm", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT)
	Key.DefineMethod("getEncoded", "()[B", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT)
	Key.DefineMethod("getFormat", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT)
}
