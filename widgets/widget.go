//go:build js && wasm
// +build js,wasm


package widgets

type WidgetContext struct {
}

type Widget interface {
	Render(ctx WidgetContext) error
}
