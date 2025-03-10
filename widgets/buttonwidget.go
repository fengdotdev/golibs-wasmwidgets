//go:build js && wasm
// +build js,wasm

package widgets

var _ Widget = &ButtonWidget{}

type ButtonWidget struct {
	content Widget
	onclick func()
}

func Button(content Widget, onclick func()) ButtonWidget {
	return ButtonWidget{}
}

func NewButtonWidget(content Widget, onclick func()) ButtonWidget {
	return Button(content, onclick)
}

func (w ButtonWidget) Render(ctx WidgetContext) error {
	return nil
}
