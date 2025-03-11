//go:build js && wasm
// +build js,wasm

package element


// NewElement creates a new DOM element with the specified tag
func NewElement(tag string, dom DomInterface, parent ElementInterface) Element {
	element := Element{
		rootParent: dom,
		self:       dom.CreateElement(tag).Get("self"),
		children:   make([]ElementInterface, 0),
	}
	return element
}
