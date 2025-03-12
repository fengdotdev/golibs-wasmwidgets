//go:build js && wasm
// +build js,wasm

package widgets

type Widget interface {
	Build(buildContext BuildContext) Widget
	CreateElement() Element
}
