//go:build js && wasm
// +build js,wasm

package interfaces

import "syscall/js"

type ElementInterface interface {
	ElementSelf
	SelfPrime
	ElementChilds
	// ElementContent
}

type ElementSelf interface {
	GetSelf() js.Value
}

type SelfPrime interface {
	GetParent() ElementInterface
	GetChildren() []ElementInterface
	GetRoot() DomInterface
}

type ElementChilds interface {
	AppendChild(child ElementInterface)
	RemoveChild(child ElementInterface)
	ReplaceChild(newChild, oldChild ElementInterface)
}

type ElementContent interface {
	SetTextContent(text string)
	GetTextContent() string
}
