package registration

import (
	"syscall/js"
)

func (w *WasmHandler) Handler() {
	js.Global().Set("calculatePolish", js.FuncOf(w.calculatePolish))
}

func (w *WasmHandler) calculatePolish(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}

	result, err := w.calculator.CalculateExpression(args[0].String())
	if err != nil {
		js.Global().Call("alert", err.Error())

		return 0
	}
	return result
}
