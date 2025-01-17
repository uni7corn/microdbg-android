package internal

import (
	"io"

	android "github.com/wnxd/microdbg-android"
)

type Package interface {
	io.Closer
	Link(android.Runtime)
	android.Package
}
