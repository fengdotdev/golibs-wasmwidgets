//go:build js && wasm
// +build js,wasm

package widgets

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

func (w TextWidget) Build(buildContext BuildContext) Widget {

	panic("implement me")
	return w
}
