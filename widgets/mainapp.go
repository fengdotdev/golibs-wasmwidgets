//go:build js && wasm
// +build js,wasm

package widgets

type MainAppWidget struct {
	buildContext BuildContext
	child        Widget
}

func MainApp(child Widget) MainAppWidget {
	ctx := NewWidgetContext()
	return MainAppWidget{
		buildContext: *ctx,
		child:        child}
}

func NewMainAppWidget(child Widget) MainAppWidget {
	return MainApp(child)
}

func (w *MainAppWidget) Run() {
	w.Build(w.buildContext)
	w.Loop()
}

func (w *MainAppWidget) Build(buildContext BuildContext) Widget {

	return w.child.Build(buildContext)
}

func (w *MainAppWidget) Loop() {
	select {}
}
