//go:build js && wasm
// +build js,wasm

package widgets

var _ Widget = &ContainerWidget{}

type ContainerWidget struct {
	Children []Widget
}

func Container(child Widget) ContainerWidget {
	return ContainerWidget{Children: []Widget{child}}
}

func NewContainerWidget(child Widget) ContainerWidget {
	return Container(child)
}

//methods

func (w ContainerWidget) Build(buildContext BuildContext) Widget {
	return nil
}
