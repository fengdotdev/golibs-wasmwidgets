//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

func NewDOM() DOM {
	return DOM{
		self: js.Global().Get("document"),
	}
}
