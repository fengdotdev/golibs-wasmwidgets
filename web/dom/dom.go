//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

var _ DomInterface = &DOM{}

type DOM struct {
	self       js.Value
	references map[string]js.Value
	children   []ElementInterface
}

func (w *DOM) CreateElement(tag string) js.Value {
	return w.self.Call("createElement", tag)
}
