package registration

import (
	"syscall/js"
	"webasm/cmd/wasm/internal/polish"
)

func RegistrHandler() {
	js.Global().Set("calculatePolish", js.FuncOf(calculatePolish))
}

func calculatePolish(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}

	result, err := polish.CalculateExpression(args[0].String())
	if err != nil {
		print("error")
		return err
	}
	return result
}
