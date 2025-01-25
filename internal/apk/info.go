package apk

import (
	"context"
	"crypto/x509"
	"encoding/base64"
	"io"
	"io/fs"
	"strings"

	"github.com/google/uuid"
	android "github.com/wnxd/microdbg-android"
	"github.com/wnxd/microdbg-android/internal/virtual"
	"github.com/wnxd/microdbg/emulator"
	"go.mozilla.org/pkcs7"
)

type version struct {
	name string
	code string
}

type info struct {
	fs interface {
		io.Closer
		fs.FS
	}
	name       string
	label      string
	version    version
	permission []string
	rdn        string
	code       string
	lib        string
	files      string
}

func (info *info) init() *info {
	b := [16]byte(uuid.New())
	b64 := base64.URLEncoding.EncodeToString(b[:])
	info.rdn = info.name + "-" + b64
	info.code = "/data/app/" + info.rdn + "/base.apk"
	info.files = "/data/data/" + info.name + "/files"
	return info
}

func (info *info) Close() error {
	return info.fs.Close()
}

func (info *info) Link(art android.Runtime) {
	switch art.Debugger().Arch() {
	case emulator.ARCH_ARM:
		info.lib = "/data/app/" + info.rdn + "/lib/arm"
	case emulator.ARCH_ARM64:
		info.lib = "/data/app/" + info.rdn + "/lib/arm64"
	case emulator.ARCH_X86:
		info.lib = "/data/app/" + info.rdn + "/lib/x86"
	case emulator.ARCH_X86_64:
		info.lib = "/data/app/" + info.rdn + "/lib/x86_64"
	}
	if lib, err := fs.Sub(info.fs, "lib/armeabi-v7a"); err == nil {
		art.LinkFS("/data/app/"+info.rdn+"/lib/arm", virtual.FS(lib))
	}
	if lib, err := fs.Sub(info.fs, "lib/arm64-v8a"); err == nil {
		art.LinkFS("/data/app/"+info.rdn+"/lib/arm64", virtual.FS(lib))
	}
	if lib, err := fs.Sub(info.fs, "lib/x86"); err == nil {
		art.LinkFS("/data/app/"+info.rdn+"/lib/x86", virtual.FS(lib))
	}
	if lib, err := fs.Sub(info.fs, "lib/x86_64"); err == nil {
		art.LinkFS("/data/app/"+info.rdn+"/lib/x86_64", virtual.FS(lib))
	}
}

func (info *info) Name() string {
	return info.name
}

func (info *info) Label() string {
	return info.label
}

func (info *info) Version() (string, string) {
	return info.version.name, info.version.code
}

func (info *info) Permission() []string {
	return info.permission
}

func (info *info) CodePath() string {
	return info.code
}

func (info *info) LibraryDir() string {
	return info.lib
}

func (info *info) FilesDir() string {
	return info.files
}

func (info *info) Certificate() []*x509.Certificate {
	var certs []*x509.Certificate
	fs.WalkDir(info.fs, "META-INF", func(path string, d fs.DirEntry, err error) error {
		if !strings.HasSuffix(path, ".RSA") && !strings.HasSuffix(path, ".DSA") {
			return err
		}
		file, err := info.fs.Open(path)
		if err != nil {
			return nil
		}
		data, err := io.ReadAll(file)
		if err != nil {
			return nil
		}
		p7, err := pkcs7.Parse(data)
		if err != nil {
			return nil
		}
		certs = append(certs, p7.Certificates...)
		return nil
	})
	return certs
}

func (info *info) LoadModule(ctx context.Context, art android.Runtime, name string) (android.Module, error) {
	var path string
	switch art.Debugger().Arch() {
	case emulator.ARCH_ARM:
		path = "lib/armeabi-v7a/lib" + name + ".so"
	case emulator.ARCH_ARM64:
		path = "lib/arm64-v8a/lib" + name + ".so"
	case emulator.ARCH_X86:
		path = "lib/x86/lib" + name + ".so"
	case emulator.ARCH_X86_64:
		path = "lib/x86_64/lib" + name + ".so"
	default:
		return nil, emulator.ErrArchUnsupported
	}
	file, err := info.fs.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return art.LoadModule(ctx, file)
}
