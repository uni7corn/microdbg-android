package extend

import (
	"crypto/x509"
	"errors"
	"fmt"

	"github.com/Xuanwo/go-locale"
	android "github.com/wnxd/microdbg-android"
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

func (ex *extend) defineContent() {
	pkg := ex.art.Package()
	tag, _ := locale.Detect()

	SharedPreferences := ex.cf.DefineClass("android/content/SharedPreferences")
	SharedPreferences.DefineMethod("contains", "(Ljava/lang/String;)Z", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if sp, ok := obj.(gava.FakeObject).Value().(SharedPreference); ok {
			return sp.Contains(args[0].(java.IString).String())
		}
		return false
	})
	SharedPreferences.DefineMethod("edit", "()Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		panic(fmt.Errorf("[SharedPreferences.edit] %w", errors.ErrUnsupported))
	})
	SharedPreferences.DefineMethod("getBoolean", "(Ljava/lang/String;Z)Z", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if sp, ok := obj.(gava.FakeObject).Value().(SharedPreference); ok {
			return sp.GetBoolean(args[0].(java.IString).String(), args[1].(java.JBoolean))
		}
		return args[1]
	})
	SharedPreferences.DefineMethod("getFloat", "(Ljava/lang/String;F)F", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if sp, ok := obj.(gava.FakeObject).Value().(SharedPreference); ok {
			return sp.GetFloat(args[0].(java.IString).String(), args[1].(java.JFloat))
		}
		return args[1]
	})
	SharedPreferences.DefineMethod("getInt", "(Ljava/lang/String;I)I", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if sp, ok := obj.(gava.FakeObject).Value().(SharedPreference); ok {
			return sp.GetInt(args[0].(java.IString).String(), args[1].(java.JInt))
		}
		return args[1]
	})
	SharedPreferences.DefineMethod("getLong", "(Ljava/lang/String;J)J", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if sp, ok := obj.(gava.FakeObject).Value().(SharedPreference); ok {
			return sp.GetLong(args[0].(java.IString).String(), args[1].(java.JLong))
		}
		return args[1]
	})
	SharedPreferences.DefineMethod("getString", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if sp, ok := obj.(gava.FakeObject).Value().(SharedPreference); ok {
			return sp.GetString(args[0].(java.IString).String(), gava.ToObject[java.IString](args[1]))
		}
		return args[1]
	})

	Signature := ex.cf.DefineClass("android.content.pm.Signature")
	Signature.DefineMethod("toByteArray", "()[B", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		cert := obj.(gava.FakeObject).Value().(*x509.Certificate)
		return gava.BytesOf(cert.Raw)
	})

	Configuration := ex.cf.DefineClass("android.content.res.Configuration")
	Configuration.DefineField("locale", "Ljava/util/Locale;", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		return ex.cf.GetClass("java.util.Locale").NewObject(tag)
	})

	Resources := ex.cf.DefineClass("android.content.res.Resources")
	Resources.DefineMethod("getConfiguration", "()Landroid/content/res/Configuration;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		return Configuration.NewInstance()
	})

	ApplicationInfo := ex.cf.DefineClass("android.content.pm.ApplicationInfo", ex.cf.DefineClass("android.content.pm.PackageItemInfo"))

	PackageInfo := ex.cf.DefineClass("android.content.pm.PackageInfo")
	PackageInfo.DefineField("versionName", "Ljava/lang/String;", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		pkg := obj.(gava.FakeObject).Value().(android.Package)
		name, _ := pkg.Version()
		return gava.FakeString(name)
	})
	PackageInfo.DefineField("signatures", "[Landroid/content/pm/Signature;", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		pkg := obj.(gava.FakeObject).Value().(android.Package)
		certs := pkg.Certificate()
		if len(certs) == 0 {
			return nil
		}
		signs := make([]java.IObject, len(certs))
		for i := range certs {
			signs[i] = Signature.NewObject(certs[i])
		}
		return gava.ArrayOf(ex.cf.ArrayOf(Signature), signs)
	})

	PackageManager := ex.cf.DefineClass("android.content.pm.PackageManager")
	PackageManager.DefineMethod("getApplicationInfo", "(Ljava/lang/String;I)Landroid/content/pm/ApplicationInfo;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		packageName := args[0].(java.IString).String()
		return ApplicationInfo.NewObject(packageName)
	})
	PackageManager.DefineMethod("getApplicationLabel", "(Landroid/content/pm/ApplicationInfo;)Ljava/lang/CharSequence;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		packageName := args[0].(gava.FakeObject).Value().(string)
		if packageName == pkg.Name() {
			return gava.FakeString(pkg.Label())
		}
		return gava.FakeString("")
	})
	PackageManager.DefineMethod("getPackageInfo", "(Ljava/lang/String;I)Landroid/content/pm/PackageInfo;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		packageName := args[0].(java.IString).String()
		if packageName != pkg.Name() {
			panic(fmt.Errorf("[PackageManager.getPackageInfo] package %s not found", packageName))
		}
		return PackageInfo.NewObject(pkg)
	})

	Context := ex.cf.DefineClass("android.content.Context")
	Context.DefineMethod("getSharedPreferences", "(Ljava/lang/String;I)Landroid/content/SharedPreferences;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		val, _ := ex.sp.Load(args[0].(java.IString).String())
		return SharedPreferences.NewObject(val)
	})
	Context.DefineMethod("getFilesDir", "()Ljava/io/File;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return ex.cf.GetClass("java/io/File").NewObject(pkg.FilesDir())
	})
	Context.DefineMethod("getPackageName", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return gava.FakeString(pkg.Name())
	})
	Context.DefineMethod("getPackageManager", "()Landroid/content/pm/PackageManager;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return PackageManager.NewInstance()
	})
	Context.DefineMethod("getPackageCodePath", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return gava.FakeString(pkg.CodePath())
	})
	Context.DefineMethod("getResources", "()Landroid/content/res/Resources;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return Resources.NewInstance()
	})

	ContextWrapper := ex.cf.DefineClass("android.content.ContextWrapper", Context)
	_ = ContextWrapper
}
