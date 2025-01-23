package java

import (
	"strings"

	java "github.com/wnxd/microdbg-java"
)

type Modifier java.JInt

const (
	Modifier_PUBLIC Modifier = 1 << iota
	Modifier_PRIVATE
	Modifier_PROTECTED
	Modifier_STATIC
	Modifier_FINAL
	Modifier_SYNCHRONIZED
	Modifier_VOLATILE
	Modifier_TRANSIENT
	Modifier_NATIVE
	Modifier_INTERFACE
	Modifier_ABSTRACT
	Modifier_STRICT
)

func (mod Modifier) String() string {
	var arr []string
	if (mod & Modifier_PUBLIC) != 0 {
		arr = append(arr, "public")
	}
	if (mod & Modifier_PROTECTED) != 0 {
		arr = append(arr, "protected")
	}
	if (mod & Modifier_PRIVATE) != 0 {
		arr = append(arr, "private")
	}
	if (mod & Modifier_ABSTRACT) != 0 {
		arr = append(arr, "abstract")
	}
	if (mod & Modifier_STATIC) != 0 {
		arr = append(arr, "static")
	}
	if (mod & Modifier_FINAL) != 0 {
		arr = append(arr, "final")
	}
	if (mod & Modifier_TRANSIENT) != 0 {
		arr = append(arr, "transient")
	}
	if (mod & Modifier_VOLATILE) != 0 {
		arr = append(arr, "volatile")
	}
	if (mod & Modifier_SYNCHRONIZED) != 0 {
		arr = append(arr, "synchronized")
	}
	if (mod & Modifier_NATIVE) != 0 {
		arr = append(arr, "native")
	}
	if (mod & Modifier_STRICT) != 0 {
		arr = append(arr, "strictfp")
	}
	if (mod & Modifier_INTERFACE) != 0 {
		arr = append(arr, "interface")
	}
	return strings.Join(arr, " ")
}

func IsStatic(v interface{ GetModifiers() java.JInt }) bool {
	return Modifier(v.GetModifiers())&Modifier_STATIC != 0
}

func IsNative(v interface{ GetModifiers() java.JInt }) bool {
	return Modifier(v.GetModifiers())&Modifier_NATIVE != 0
}
