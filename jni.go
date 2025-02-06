package android

import (
	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
)

type JNIContext interface {
	ClassFactory() gava.ClassFactory
	Throw(java.IThrowable) java.JInt
	ExceptionOccurred() java.IThrowable
	ExceptionDescribe()
	ExceptionClear()
	FatalError(string)
}

type JNIEnv interface {
	DefineClass(JNIContext, string, java.IObject, []java.JByte) (java.IClass, error)
	FindClass(JNIContext, string) (java.IClass, error)
	ThrowNew(JNIContext, java.IClass, string) (java.JInt, error)
	GetMethod(JNIContext, java.IClass, string, string) (java.IMethod, error)
	GetField(JNIContext, java.IClass, string, string) (java.IField, error)
	GetStaticMethod(JNIContext, java.IClass, string, string) (java.IMethod, error)
	GetStaticField(JNIContext, java.IClass, string, string) (java.IField, error)
	NewString(JNIContext, []java.JChar) (java.IString, error)
	NewStringUTF(JNIContext, string) (java.IString, error)
	NewObjectArray(JNIContext, java.JSize, java.IClass, java.IObject) (java.IObjectArray, error)
	NewBooleanArray(JNIContext, java.JSize) (java.IBooleanArray, error)
	NewByteArray(JNIContext, java.JSize) (java.IByteArray, error)
	NewCharArray(JNIContext, java.JSize) (java.ICharArray, error)
	NewShortArray(JNIContext, java.JSize) (java.IShortArray, error)
	NewIntArray(JNIContext, java.JSize) (java.IIntArray, error)
	NewLongArray(JNIContext, java.JSize) (java.ILongArray, error)
	NewFloatArray(JNIContext, java.JSize) (java.IFloatArray, error)
	NewDoubleArray(JNIContext, java.JSize) (java.IDoubleArray, error)
}
