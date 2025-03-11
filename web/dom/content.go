//go:build js && wasm
// +build js,wasm

package dom

func (e *DOM) SetTextContent(text string) {
	e.self.Set("textContent", text)
}

func (e *DOM) GetTextContent() string {
	return e.self.Get("textContent").String()
}
