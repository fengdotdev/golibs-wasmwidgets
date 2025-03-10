//go:build js && wasm
// +build js,wasm

package widgets

type MainAppWidget struct {
	child Widget
}

func MainApp(child Widget) MainAppWidget {
	return MainAppWidget{child: child}
}

func NewMainAppWidget(child Widget) MainAppWidget {
	return MainApp(child)
}

func (w MainAppWidget) Run() {
	w.Loop()
}

func (w MainAppWidget) Render(ctx WidgetContext) error {

	return w.child.Render(ctx)
}

func (w *MainAppWidget) Loop() {
	select {}
}
