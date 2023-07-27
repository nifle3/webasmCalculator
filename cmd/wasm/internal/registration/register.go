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

	result := w.calculator.CalculateExpression(args[0].String())

	return result
}
