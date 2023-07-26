package registration

import (
	"syscall/js"
)

func RegistrHandler() {
	js.Global().Set("calculatePolish", js.FuncOf(calculatePolish))
}

func calculatePolish(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return nil
	}

	return nil
}
