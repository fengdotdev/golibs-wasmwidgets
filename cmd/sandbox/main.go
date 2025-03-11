//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/fengdotdev/golibs-wasmwidgets/web"
	"github.com/fengdotdev/golibs-wasmwidgets/widgets"
)

type context = widgets.WidgetContext

func main() {
	app:= widgets.MainApp(

		widgets.Text("Hello, World!"),
	)
	AlertBTNPress()
	app.Run()
}

type homepage struct {
}

func (h *homepage) Render(ctx context) error {
	return nil
}

func AlertBTNPress() {
	web.Alert("Button pressed!")
}
