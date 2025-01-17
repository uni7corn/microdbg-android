package android

import (
	android "github.com/wnxd/microdbg-android"
	internal "github.com/wnxd/microdbg-android/internal/10"
)

func WithApkPath(name string) android.Option {
	return internal.WithApkPath(name)
}

func WithRuntimeDir(name string) android.Option {
	return internal.WithRuntimeDir(name)
}

func WithRootDir(name string) android.Option {
	return internal.WithRootDir(name)
}

func WithJNIEnv(env android.JNIEnv) android.Option {
	return internal.WithJNIEnv(env)
}
