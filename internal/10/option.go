package android

import (
	android "github.com/wnxd/microdbg-android"
)

type ApkPathOption interface {
	setApkPath(string) error
}

type RuntimeDirOption interface {
	setRuntimeDir(string) error
}

type RootDirOption interface {
	setRootDir(string) error
}

type JNIEnvOption interface {
	setJNIEnv(android.JNIEnv) error
}

func WithApkPath(name string) android.Option {
	return func(art android.Runtime) error {
		if option, ok := art.(ApkPathOption); ok {
			return option.setApkPath(name)
		}
		return android.ErrOptionUnsupported
	}
}

func WithRuntimeDir(name string) android.Option {
	return func(art android.Runtime) error {
		if option, ok := art.(RuntimeDirOption); ok {
			return option.setRuntimeDir(name)
		}
		return android.ErrOptionUnsupported
	}
}

func WithRootDir(name string) android.Option {
	return func(art android.Runtime) error {
		if option, ok := art.(RootDirOption); ok {
			return option.setRootDir(name)
		}
		return android.ErrOptionUnsupported
	}
}

func WithJNIEnv(env android.JNIEnv) android.Option {
	return func(art android.Runtime) error {
		if option, ok := art.(JNIEnvOption); ok {
			return option.setJNIEnv(env)
		}
		return android.ErrOptionUnsupported
	}
}
