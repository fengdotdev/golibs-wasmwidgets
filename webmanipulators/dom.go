//go:build js && wasm
// +build js,wasm

package webmanipulators

import "syscall/js"

type DomInterface interface {
	AddEventListener(event Events, listener func())
	RemoveEventListener(event Events, listener func())
	CreatElement(tag string) js.Value
}

type DOM struct {
	dom        js.Value
	references map[string]js.Value
}

func NewDOM() DOM {
	return DOM{
		dom: js.Global().Get("document"),
	}
}

func (w *DOM) AddEventListener(event Events, listener func()) {
	w.dom.Call("addEventListener", event, listener)
}

func (w *DOM) RemoveEventListener(event Events, listener func()) {
	w.dom.Call("removeEventListener", event, listener)
}
