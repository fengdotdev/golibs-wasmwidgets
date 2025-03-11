//go:build js && wasm
// +build js,wasm

package web

import "fmt"

type Console struct {
}

func NewConsole() Console {
	return Console{}
}

func (w *Console) Log(v ...interface{}) {
	fmt.Println(v...)
}
