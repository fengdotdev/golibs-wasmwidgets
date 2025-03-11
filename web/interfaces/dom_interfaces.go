//go:build js && wasm
// +build js,wasm

package interfaces

import "syscall/js"

type DomInterface interface {
	CreateElement(tag string) js.Value
	ElementSelf
}
