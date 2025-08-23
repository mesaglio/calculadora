package calculadora

import (
	"fmt"

	"go.uber.org/zap"
)

// Suma dos números y retorna el resultado
func Suma(a, b float64) float64 {
	zap.L().Info("Suma", zap.Float64("a", a), zap.Float64("b", b))
	return a + b
}

// Resta dos números y retorna el resultado
func Resta(a, b float64) float64 {
	zap.L().Info("Resta", zap.Float64("a", a), zap.Float64("b", b))
	return a - b
}

// Procesa la operación basada en el operador
func ProcesarOperacion(operador string, a, b float64) (float64, error) {
	switch operador {
	case "+":
		return Suma(a, b), nil
	case "-":
		return Resta(a, b), nil
	default:
		return 0, fmt.Errorf("operador no soportado: %s", operador)
	}
}
