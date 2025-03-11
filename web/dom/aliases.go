//go:build js && wasm
// +build js,wasm

package dom

import (
	"github.com/fengdotdev/golibs-wasmwidgets/web"
	"github.com/fengdotdev/golibs-wasmwidgets/web/interfaces"
)

type Events = web.Events

type ElementInterface = interfaces.ElementInterface
type DomInterface = interfaces.DomInterface
