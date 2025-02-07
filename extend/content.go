package extend

import (
	"crypto/x509"
	"log"
	"time"

	"github.com/Xuanwo/go-locale"
	android "github.com/wnxd/microdbg-android"
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

type intentInfo struct {
	action  string
	handler Intent
}

func (ex *extend) defineContent() {
	pkg := ex.art.Package()
	tag, _ := locale.Detect()
	installTime := java.JLong(time.Now().UnixMilli())

	SharedPreferencesEditor := ex.cf.DefineClass("android.content.SharedPreferences$Editor")
	SharedPreferencesEditor.DefineMethod("apply", "()V", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return nil
	})
	SharedPreferencesEditor.DefineMethod("clear", "()Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if spe, ok := obj.(gava.FakeObject).Value().(SharedPreferenceEditor); ok {
			spe.Clear()
		}
		return obj
	})
	SharedPreferencesEditor.DefineMethod("commit", "()Z", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return true
	})
	SharedPreferencesEditor.DefineMethod("putBoolean", "(Ljava/lang/String;Z)Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if spe, ok := obj.(gava.FakeObject).Value().(SharedPreferenceEditor); ok {
			spe.SetBoolean(args[0].(java.IString).String(), args[1].(java.JBoolean))
		}
		return obj
	})
	SharedPreferencesEditor.DefineMethod("putFloat", "(Ljava/lang/String;F)Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if spe, ok := obj.(gava.FakeObject).Value().(SharedPreferenceEditor); ok {
			spe.SetFloat(args[0].(java.IString).String(), args[1].(java.JFloat))
		}
		return obj
	})
	SharedPreferencesEditor.DefineMethod("putInt", "(Ljava/lang/String;I)Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if spe, ok := obj.(gava.FakeObject).Value().(SharedPreferenceEditor); ok {
			spe.SetInt(args[0].(java.IString).String(), args[1].(java.JInt))
		}
		return obj
	})
	SharedPreferencesEditor.DefineMethod("putLong", "(Ljava/lang/String;J)Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if spe, ok := obj.(gava.FakeObject).Value().(SharedPreferenceEditor); ok {
			spe.SetLong(args[0].(java.IString).String(), args[1].(java.JLong))
		}
		return obj
	})
	SharedPreferencesEditor.DefineMethod("putString", "(Ljava/lang/String;Ljava/lang/String;)Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if spe, ok := obj.(gava.FakeObject).Value().(SharedPreferenceEditor); ok {
			spe.SetString(args[0].(java.IString).String(), gava.ToObject[java.IString](args[1]))
		}
		return obj
	})
	SharedPreferencesEditor.DefineMethod("remove", "(Ljava/lang/String;)Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if spe, ok := obj.(gava.FakeObject).Value().(SharedPreferenceEditor); ok {
			spe.Remove(args[0].(java.IString).String())
		}
		return obj
	})

	SharedPreferences := ex.cf.DefineClass("android.content.SharedPreferences")
	SharedPreferences.DefineMethod("contains", "(Ljava/lang/String;)Z", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if sp, ok := obj.(gava.FakeObject).Value().(SharedPreference); ok {
			return sp.Contains(args[0].(java.IString).String())
		}
		return false
	})
	SharedPreferences.DefineMethod("edit", "()Landroid/content/SharedPreferences$Editor;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		if spe, ok := obj.(gava.FakeObject).Value().(SharedPreferenceEditor); ok {
			return SharedPreferencesEditor.NewObject(spe)
		}
		return SharedPreferencesEditor.NewObject(DefaultPreference{})
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

	IntentFilter := ex.cf.DefineClass("android.content.IntentFilter")
	IntentFilter.DefineMethod(gava.ConstructorMethodName, "(Ljava/lang/String;)V", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		action := args[0].(java.IString).String()
		return IntentFilter.NewObject(action)
	})
	IntentFilter.DefineMethod("addAction", "(Ljava/lang/String;)V", gava.Modifier_PUBLIC|gava.Modifier_FINAL).BindCall(func(obj java.IObject, args ...any) any {
		return nil
	})
	IntentFilter.DefineMethod("addCategory", "(Ljava/lang/String;)V", gava.Modifier_PUBLIC|gava.Modifier_FINAL).BindCall(func(obj java.IObject, args ...any) any {
		return nil
	})

	ContentIntent := ex.cf.DefineClass("android.content.Intent")
	ContentIntent.DefineMethod("getAction", "()Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		return gava.FakeString(intent.action)
	})
	ContentIntent.DefineMethod("getData", "()Landroid/net/Uri;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		data := intent.handler.GetData()
		if data == nil {
			return nil
		}
		return ex.cf.GetClass("android.net.Uri").NewObject(*data)
	})
	ContentIntent.DefineMethod("getBooleanExtra", "(Ljava/lang/String;Z)Z", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		defValue := args[1].(java.JBoolean)
		return intent.handler.GetBooleanExtra(name, defValue)
	})
	ContentIntent.DefineMethod("getByteExtra", "(Ljava/lang/String;B)B", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		defValue := args[1].(java.JByte)
		return intent.handler.GetByteExtra(name, defValue)
	})
	ContentIntent.DefineMethod("getCharExtra", "(Ljava/lang/String;C)C", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		defValue := args[1].(java.JChar)
		return intent.handler.GetCharExtra(name, defValue)
	})
	ContentIntent.DefineMethod("getShortExtra", "(Ljava/lang/String;S)S", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		defValue := args[1].(java.JShort)
		return intent.handler.GetShortExtra(name, defValue)
	})
	ContentIntent.DefineMethod("getIntExtra", "(Ljava/lang/String;I)I", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		defValue := args[1].(java.JInt)
		return intent.handler.GetIntExtra(name, defValue)
	})
	ContentIntent.DefineMethod("getLongExtra", "(Ljava/lang/String;J)J", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		defValue := args[1].(java.JLong)
		return intent.handler.GetLongExtra(name, defValue)
	})
	ContentIntent.DefineMethod("getFloatExtra", "(Ljava/lang/String;F)F", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		defValue := args[1].(java.JFloat)
		return intent.handler.GetFloatExtra(name, defValue)
	})
	ContentIntent.DefineMethod("getDoubleExtra", "(Ljava/lang/String;D)D", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		defValue := args[1].(java.JDouble)
		return intent.handler.GetDoubleExtra(name, defValue)
	})
	ContentIntent.DefineMethod("getStringExtra", "(Ljava/lang/String;)Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		intent := obj.(gava.FakeObject).Value().(*intentInfo)
		name := args[0].(java.IString).String()
		return intent.handler.GetStringExtra(name)
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
	Configuration.DefineField("screenLayout", "I", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		const (
			SCREENLAYOUT_SIZE_NORMAL   = 0x00000002
			SCREENLAYOUT_LONG_YES      = 0x00000020
			SCREENLAYOUT_LAYOUTDIR_LTR = 0x00000040
			SCREENLAYOUT_ROUND_NO      = 0x00000100
		)

		return java.JInt(SCREENLAYOUT_SIZE_NORMAL | SCREENLAYOUT_LONG_YES | SCREENLAYOUT_LAYOUTDIR_LTR | SCREENLAYOUT_ROUND_NO)
	})

	Resources := ex.cf.DefineClass("android.content.res.Resources")
	Resources.DefineMethod("getConfiguration", "()Landroid/content/res/Configuration;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		return Configuration.NewInstance()
	})

	ApplicationInfo := ex.cf.DefineClass("android.content.pm.ApplicationInfo", ex.cf.DefineClass("android.content.pm.PackageItemInfo"))
	ApplicationInfo.DefineField("minSdkVersion", "I", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		pkg := obj.(gava.FakeObject).Value().(android.Package)
		min, _ := pkg.UsesSdk()
		return java.JInt(min)
	})
	ApplicationInfo.DefineField("targetSdkVersion", "I", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		pkg := obj.(gava.FakeObject).Value().(android.Package)
		_, target := pkg.UsesSdk()
		return java.JInt(target)
	})

	PackageInfo := ex.cf.DefineClass("android.content.pm.PackageInfo")
	PackageInfo.DefineField("versionName", "Ljava/lang/String;", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		pkg := obj.(gava.FakeObject).Value().(android.Package)
		name, _ := pkg.Version()
		return gava.FakeString(name)
	})
	PackageInfo.DefineField("versionCode", "I", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		pkg := obj.(gava.FakeObject).Value().(android.Package)
		_, code := pkg.Version()
		return java.JInt(code)
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
	PackageInfo.DefineField("firstInstallTime", "J", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		return installTime
	})
	PackageInfo.DefineField("lastUpdateTime", "J", gava.Modifier_PUBLIC).BindGet(func(obj java.IObject) any {
		return installTime
	})

	NameNotFoundException := ex.cf.DefineClass("android.content.pm.PackageManager$NameNotFoundException", ex.cf.GetClass("android.util.AndroidException"))

	PackageManager := ex.cf.DefineClass("android.content.pm.PackageManager")
	PackageManager.DefineMethod("getApplicationInfo", "(Ljava/lang/String;I)Landroid/content/pm/ApplicationInfo;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		packageName := args[0].(java.IString).String()
		if packageName != pkg.Name() {
			panic(NameNotFoundException.NewThrowable(packageName))
		}
		return ApplicationInfo.NewObject(pkg)
	})
	PackageManager.DefineMethod("getApplicationLabel", "(Landroid/content/pm/ApplicationInfo;)Ljava/lang/CharSequence;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		pkg := args[0].(gava.FakeObject).Value().(android.Package)
		return gava.FakeString(pkg.Label())
	})
	PackageManager.DefineMethod("getPackageInfo", "(Ljava/lang/String;I)Landroid/content/pm/PackageInfo;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		packageName := args[0].(java.IString).String()
		if packageName != pkg.Name() {
			panic(NameNotFoundException.NewThrowable(packageName))
		}
		return PackageInfo.NewObject(pkg)
	})

	Context := ex.cf.DefineClass("android.content.Context")
	Context.DefineMethod("getSharedPreferences", "(Ljava/lang/String;I)Landroid/content/SharedPreferences;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		name := args[0].(java.IString).String()
		val, ok := ex.pref.Load(name)
		if !ok && ex.debug {
			log.Printf("[%s] Preferences undefined\n", name)
		}
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
	Context.DefineMethod("getApplicationInfo", "()Landroid/content/pm/ApplicationInfo;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return ApplicationInfo.NewObject(pkg)
	})
	Context.DefineMethod("registerReceiver", "(Landroid/content/BroadcastReceiver;Landroid/content/IntentFilter;)Landroid/content/Intent;", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		action := args[1].(gava.FakeObject).Value().(string)
		if intent, ok := ex.intent.Load(action); ok {
			return ContentIntent.NewObject(&intentInfo{
				action:  action,
				handler: intent.(Intent),
			})
		} else if ex.debug {
			log.Printf("[%s] Intent undefined\n", action)
		}
		return nil
	})
	Context.DefineMethod("unregisterReceiver", "(Landroid/content/BroadcastReceiver;)V", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		return nil
	})
	Context.DefineMethod("checkCallingOrSelfPermission", "(Ljava/lang/String;)I", gava.Modifier_PUBLIC|gava.Modifier_ABSTRACT).BindCall(func(obj java.IObject, args ...any) any {
		const (
			PERMISSION_GRANTED = 0
			PERMISSION_DENIED  = -1
		)

		if ex.debug {
			permission := args[0].(java.IString).String()
			log.Printf("[%s] Denied permission\n", permission)
		}
		return java.JInt(PERMISSION_DENIED)
	})

	ContextWrapper := ex.cf.DefineClass("android.content.ContextWrapper", Context)
	_ = ContextWrapper
}
