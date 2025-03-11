//go:build js && wasm
// +build js,wasm

package element

func (e Element) AppendChild(child ElementInterface) {
	e.self.Call("appendChild", child.GetSelf())
}

func (e Element) RemoveChild(child ElementInterface) {
	e.self.Call("removeChild", child.GetSelf())
}

func (e Element) ReplaceChild(newChild, oldChild ElementInterface) {
	e.self.Call("replaceChild", newChild.GetSelf(), oldChild.GetSelf())
}
