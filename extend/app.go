package extend

import (
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

func (ex *extend) defineApp() {
	pkg := ex.art.Package()

	Application := ex.cf.DefineClass("android.app.Application", ex.cf.GetClass("android.content.ContextWrapper"))

	ActivityThread := ex.cf.DefineClass("android.app.ActivityThread")
	activityThread := ActivityThread.NewInstance()
	app := Application.NewInstance()
	ActivityThread.DefineMethod("currentActivityThread", "()Landroid/app/ActivityThread;", gava.Modifier_PUBLIC|gava.Modifier_STATIC).BindCall(func(obj java.IObject, args ...any) any {
		return activityThread
	})
	ActivityThread.DefineMethod("currentPackageName", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_STATIC).BindCall(func(obj java.IObject, args ...any) any {
		return gava.FakeString(pkg.Name())
	})
	ActivityThread.DefineMethod("currentProcessName", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_STATIC).BindCall(func(obj java.IObject, args ...any) any {
		return gava.FakeString(pkg.Name())
	})
	ActivityThread.DefineMethod("currentApplication", "()Ljava/lang/String;", gava.Modifier_PUBLIC|gava.Modifier_STATIC).BindCall(func(obj java.IObject, args ...any) any {
		return app
	})
	ActivityThread.DefineMethod("getApplication", "()Landroid/app/Application;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		return app
	})
	ActivityThread.DefineMethod("getProcessName", "()Ljava/lang/String;", gava.Modifier_PUBLIC).BindCall(func(obj java.IObject, args ...any) any {
		return gava.FakeString(pkg.Name())
	})
}
