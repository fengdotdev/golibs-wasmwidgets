//go:build js && wasm
// +build js,wasm

package element

func (e *Element) SetTextContent(text string) {
	e.self.Set("textContent", text)
}

func (e *Element) GetTextContent() string {
	return e.self.Get("textContent").String()
}
