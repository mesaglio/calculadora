package calculadora

import (
	"testing"
)

func TestSuma(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"suma positivos", 5, 3, 8},
		{"suma negativos", -2, -3, -5},
		{"suma mixta", 10, -4, 6},
		{"suma con zero", 0, 5, 5},
		{"suma decimales", 2.5, 3.7, 6.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultado := Suma(tt.a, tt.b)
			if resultado != tt.expected {
				t.Errorf("Suma(%v, %v) = %v; esperado %v", tt.a, tt.b, resultado, tt.expected)
			}
		})
	}
}

func TestResta(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
	}{
		{"resta positivos", 5, 3, 2},
		{"resta negativos", -2, -3, 1},
		{"resta mixta", 10, -4, 14},
		{"resta con zero", 5, 0, 5},
		{"resta decimales", 5.5, 2.3, 3.2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultado := Resta(tt.a, tt.b)
			if resultado != tt.expected {
				t.Errorf("Resta(%v, %v) = %v; esperado %v", tt.a, tt.b, resultado, tt.expected)
			}
		})
	}
}

func TestProcesarOperacion(t *testing.T) {
	tests := []struct {
		name        string
		operador    string
		a, b        float64
		expected    float64
		esperaError bool
	}{
		{"suma válida", "+", 5, 3, 8, false},
		{"resta válida", "-", 5, 3, 2, false},
		{"operador inválido", "*", 5, 3, 0, true},
		{"operador vacío", "", 5, 3, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resultado, err := ProcesarOperacion(tt.operador, tt.a, tt.b)

			if tt.esperaError && err == nil {
				t.Errorf("ProcesarOperacion(%s, %v, %v) esperaba error pero no lo obtuvo", tt.operador, tt.a, tt.b)
			}

			if !tt.esperaError && err != nil {
				t.Errorf("ProcesarOperacion(%s, %v, %v) obtuvo error inesperado: %v", tt.operador, tt.a, tt.b, err)
			}

			if !tt.esperaError && resultado != tt.expected {
				t.Errorf("ProcesarOperacion(%s, %v, %v) = %v; esperado %v", tt.operador, tt.a, tt.b, resultado, tt.expected)
			}
		})
	}
}
