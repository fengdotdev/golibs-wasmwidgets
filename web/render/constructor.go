//go:build js && wasm
// +build js,wasm

package render

import "syscall/js"

func NewDOM() DOM {
	return DOM{
		self: js.Global().Get("document"),
	}
}
