//go:build js && wasm
// +build js,wasm

package element

import (
	"syscall/js"

	"github.com/fengdotdev/golibs-wasmwidgets/web/interfaces"
)

type ElementInterface = interfaces.ElementInterface
type DomInterface = interfaces.DomInterface

var _ ElementInterface = Element{}

type Element struct {
	id         string
	class      string
	attr       map[string]string
	rootParent DomInterface       // Reference to the root DOM interface
	parent     ElementInterface   // Parent element in the hierarchy
	self       js.Value           // JavaScript value representing this element
	children   []ElementInterface // Child elements
}

