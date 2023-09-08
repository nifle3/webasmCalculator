package registration

type Calculator interface {
	CalculateExpression(expression string) (float64, error)
}

type WasmHandler struct {
	calculator Calculator
}

func NewWasmHandler(calc Calculator) *WasmHandler {
	return &WasmHandler{
		calculator: calc,
	}
}
