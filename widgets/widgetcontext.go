//go:build js && wasm
// +build js,wasm

package widgets

import (
	"github.com/fengdotdev/golibs-wasmwidgets/web/dom"
	"github.com/fengdotdev/golibs-wasmwidgets/web/interfaces"
)

type DomInterface = interfaces.DomInterface



func NewWidgetContext() *BuildContext {
	dom := dom.NewDOM()
	return &BuildContext{
		webDom: &dom,
	}
}
