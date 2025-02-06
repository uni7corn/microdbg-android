package android

import (
	"context"
	"crypto/x509"
)

type Package interface {
	Name() string
	Label() string
	Version() (name string, code int)
	UsesSdk() (min, target int)
	Permission() []string
	CodePath() string
	LibraryDir() string
	FilesDir() string
	Certificate() []*x509.Certificate
	LoadModule(ctx context.Context, art Runtime, name string) (Module, error)
}
