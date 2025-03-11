//go:build js && wasm
// +build js,wasm

package web

type Events string

const (
	Click  Events = "click"
	Change Events = "change"
	Submit Events = "submit"
)
