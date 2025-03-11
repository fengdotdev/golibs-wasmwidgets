//go:build js && wasm
// +build js,wasm

package widgets

import "github.com/fengdotdev/golibs-wasmwidgets/web/element"

var _ Widget = &TextWidget{}

type TextWidget struct {
	Text string
}

func Text(text string) TextWidget {
	return TextWidget{Text: text}
}

func NewTextWidget(text string) TextWidget {
	return Text(text)
}

//methods

func (w TextWidget) Render(ctx WidgetContext) Widget {

	e := element.NewElement("p", ctx.webDom)
	e.SetTextContent(w.Text)
	return w
}
