package android

import (
	android "github.com/wnxd/microdbg-android"
	android10 "github.com/wnxd/microdbg-android/internal/10"
	"github.com/wnxd/microdbg/emulator"
)

func NewRuntime(emu emulator.Emulator, options ...android.Option) (android.Runtime, error) {
	return android10.NewRuntime(emu, options...)
}
