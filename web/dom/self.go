//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

func (e DOM) GetSelf() js.Value {
	return e.self // Return the element itself
}


func (e DOM) GetChildren() []ElementInterface {
	return e.children
}

