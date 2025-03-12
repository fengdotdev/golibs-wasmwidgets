//go:build js && wasm
// +build js,wasm

package widgets

var _ Widget = &ColumnWidget{}

type ColumnWidget struct {
	Children []Widget
}

func Column(children ...Widget) ColumnWidget {
	return ColumnWidget{Children: children}
}

func NewColumnWidget(children ...Widget) ColumnWidget {
	return Column(children...)
}

func (w ColumnWidget) Build(buildContext BuildContext) Widget {
	return nil
}
