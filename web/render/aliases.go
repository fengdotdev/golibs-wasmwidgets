//go:build js && wasm
// +build js,wasm

package render

import (
	"github.com/fengdotdev/golibs-wasmwidgets/web"
	"github.com/fengdotdev/golibs-wasmwidgets/web/interfaces"
)

type Events = web.Events

type ElementInterface = interfaces.ElementInterface
type DomInterface = interfaces.DomInterface
