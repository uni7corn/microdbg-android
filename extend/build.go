package extend

import (
	"context"
	"strconv"
	"strings"
	"unsafe"

	gava "github.com/wnxd/microdbg-android/java"
	java "github.com/wnxd/microdbg-java"
	"github.com/wnxd/microdbg/debugger"
)

func (ex *extend) defineBuild() {
	var getprop = func(string) string { return "" }
	if libc, err := ex.art.FindModule("libc.so"); err != nil {
	} else if propGet, err := libc.FindSymbol("__system_property_get"); err != nil {
	} else if addr, err := ex.art.Debugger().MemAlloc(0x100); err == nil {
		defer ex.art.Debugger().MemFree(addr)
		var buf [0x100]byte
		emu := ex.art.Emulator()
		getprop = func(name string) string {
			var len int32
			err := propGet.Call(context.TODO(), debugger.Calling_Default, &len, name, addr)
			if err != nil {
				return ""
			}
			err = emu.MemReadPtr(addr, uint64(len), unsafe.Pointer(unsafe.SliceData(buf[:])))
			if err != nil {
				return ""
			}
			return string(buf[:len])
		}
	}

	Build := ex.cf.DefineClass("android.os.Build")
	Build.Set("BOARD", gava.FakeString(getprop("ro.product.board")))
	Build.Set("BOOTLOADER", gava.FakeString(getprop("ro.bootloader")))
	Build.Set("BRAND", gava.FakeString(getprop("ro.product.brand")))
	Build.Set("SUPPORTED_32_BIT_ABIS", gava.ArrayOf(gava.FakeStringArrayClass, strings.Split(getprop("ro.product.cpu.abilist32"), ",")))
	Build.Set("SUPPORTED_64_BIT_ABIS", gava.ArrayOf(gava.FakeStringArrayClass, strings.Split(getprop("ro.product.cpu.abilist64"), ",")))
	abilist := strings.Split(getprop("ro.product.cpu.abilist"), ",")
	Build.Set("SUPPORTED_ABIS", gava.ArrayOf(gava.FakeStringArrayClass, abilist))
	Build.Set("CPU_ABI", gava.FakeString(abilist[0]))
	if len(abilist) > 1 {
		Build.Set("CPU_ABI2", gava.FakeString(abilist[1]))
	} else {
		Build.Set("CPU_ABI2", gava.FakeString(""))
	}
	Build.Set("DEVICE", gava.FakeString(getprop("ro.product.device")))
	Build.Set("DISPLAY", gava.FakeString(getprop("ro.build.display.id")))
	Build.Set("FINGERPRINT", gava.FakeString(func() string {
		finger := getprop("ro.build.fingerprint")
		if finger == "" {
			finger = getprop("ro.product.brand") + "/" +
				getprop("ro.product.name") + "/" +
				getprop("ro.product.device") + ":" +
				getprop("ro.build.version.release") + "/" +
				getprop("ro.build.id") + "/" +
				getprop("ro.build.version.incremental") + ":" +
				getprop("ro.build.type") + "/" +
				getprop("ro.build.tags")
		}
		return finger
	}()))
	Build.Set("HARDWARE", gava.FakeString(getprop("ro.hardware")))
	Build.Set("HOST", gava.FakeString(getprop("ro.build.host")))
	Build.Set("ID", gava.FakeString(getprop("ro.build.id")))
	Build.Set("MANUFACTURER", gava.FakeString(getprop("ro.product.manufacturer")))
	Build.Set("MODEL", gava.FakeString(getprop("ro.product.model")))
	Build.Set("ODM_SKU", gava.FakeString(getprop("ro.boot.product.hardware.sku")))
	Build.Set("PRODUCT", gava.FakeString(getprop("ro.product.name")))
	Build.Set("RADIO", gava.FakeString("unknown"))
	Build.Set("SERIAL", gava.FakeString(getprop("no.such.thing")))
	Build.Set("SKU", gava.FakeString(getprop("ro.boot.hardware.sku")))
	Build.Set("SOC_MANUFACTURER", gava.FakeString("unknown"))
	Build.Set("SOC_MODEL", gava.FakeString("unknown"))
	Build.Set("TAGS", gava.FakeString(getprop("ro.build.tags")))
	Build.Set("TIME", java.JLong(must(strconv.Atoi(getprop("ro.build.date.utc")))*1000))
	Build.Set("TYPE", gava.FakeString(getprop("ro.build.type")))
	Build.Set("USER", gava.FakeString(getprop("ro.build.user")))

	Build_VERSION := ex.cf.DefineClass("android.os.Build$VERSION")
	Build_VERSION.Set("BASE_OS", gava.FakeString(getprop("ro.build.version.base_os")))
	Build_VERSION.Set("CODENAME", gava.FakeString(getprop("ro.build.version.codename")))
	Build_VERSION.Set("INCREMENTAL", gava.FakeString(getprop("ro.build.version.incremental")))
	Build_VERSION.Set("MEDIA_PERFORMANCE_CLASS", java.JInt(0))
	Build_VERSION.Set("PREVIEW_SDK_INT", java.JInt(0))
	Build_VERSION.Set("RELEASE", gava.FakeString(getprop("ro.build.version.release")))
	Build_VERSION.Set("RELEASE_OR_CODENAME", gava.FakeString(getprop("ro.build.version.release_or_codename")))
	Build_VERSION.Set("RELEASE_OR_PREVIEW_DISPLAY", gava.FakeString(getprop("ro.build.version.release_or_preview_display")))
	sdk := getprop("ro.build.version.sdk")
	Build_VERSION.Set("SDK", gava.FakeString(sdk))
	Build_VERSION.Set("SDK_INT", java.JInt(must(strconv.Atoi(sdk))))
	Build_VERSION.Set("SDK_INT_FULL", java.JInt(must(strconv.Atoi(sdk))))
	Build_VERSION.Set("SECURITY_PATCH", gava.FakeString(getprop("ro.build.version.security_patch")))
}

func must[V, E any](r V, _ E) V {
	return r
}
