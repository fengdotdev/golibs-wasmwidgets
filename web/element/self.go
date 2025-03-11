//go:build js && wasm
// +build js,wasm

package element

import "syscall/js"

func (e Element) GetSelf() js.Value {
	return e.self // Return the element itself
}

func (e Element) GetParent() ElementInterface {
	return e.parent
}

func (e Element) GetChildren() []ElementInterface {
	return e.children
}

func (e Element) GetRoot() DomInterface {
	return e.rootParent
}
