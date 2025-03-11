//go:build js && wasm
// +build js,wasm

package web

import (
	"syscall/js"

	"github.com/fengdotdev/golibs-wasmwidgets/constants"
)

func Alert(msg string) {
	js.Global().Call(constants.Alert, msg)
}
