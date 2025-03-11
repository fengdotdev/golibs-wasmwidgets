//go:build js && wasm
// +build js,wasm

package widgets





type Widget interface {
	Render(ctx WidgetContext) Widget
}
