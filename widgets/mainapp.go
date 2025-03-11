//go:build js && wasm
// +build js,wasm

package widgets

type MainAppWidget struct {
	ctx   WidgetContext
	child Widget
}

func MainApp(child Widget) MainAppWidget {
	ctx := NewWidgetContext()
	return MainAppWidget{
		ctx:   *ctx,
		child: child}
}

func NewMainAppWidget(child Widget) MainAppWidget {
	return MainApp(child)
}

func (w *MainAppWidget) Run() {
	w.Render(w.ctx)
	w.Loop()
}

func (w *MainAppWidget) Render(ctx WidgetContext) Widget {

	return w.child.Render(ctx)
}

func (w *MainAppWidget) Loop() {
	select {}
}
