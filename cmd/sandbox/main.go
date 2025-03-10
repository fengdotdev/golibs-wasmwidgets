//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/fengdotdev/golibs-wasmwidgets/actions"
	"github.com/fengdotdev/golibs-wasmwidgets/widgets"
)

type context = widgets.WidgetContext

func main() {
	widgets.MainApp(
		widgets.Container(
			widgets.Column(
				widgets.Text("Hello, World!"),
				widgets.Button(widgets.Text("Press me!"), AlertBTNPress),
			),
		),
	)
}

type homepage struct {
}

func (h *homepage) Render(ctx context) error {
	return nil
}

func AlertBTNPress() {
	actions.Alert("Button pressed!")
}
