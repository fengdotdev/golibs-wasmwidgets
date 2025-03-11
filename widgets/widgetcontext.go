//go:build js && wasm
// +build js,wasm

package widgets

import (
	"github.com/fengdotdev/golibs-wasmwidgets/web/dom"
	"github.com/fengdotdev/golibs-wasmwidgets/web/interfaces"
)

type DomInterface = interfaces.DomInterface

type WidgetContext struct {
	webDom DomInterface
}

func NewWidgetContext() *WidgetContext {
	dom := dom.NewDOM()
	return &WidgetContext{
		webDom: &dom,
	}
}
