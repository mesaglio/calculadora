package main

import (
	"fmt"
	"os"
	"strconv"

	"calculadora/pkg/calculadora"
)

func main() {
	// Verificar que se proporcionen exactamente 3 argumentos
	if len(os.Args) != 4 {
		fmt.Println("Uso: calculadora <número1> <operador> <número2>")
		fmt.Println("Operadores soportados: + (suma), - (resta)")
		fmt.Println("Ejemplo: calculadora 5 + 3")
		os.Exit(1)
	}

	// Obtener argumentos
	num1Str := os.Args[1]
	operador := os.Args[2]
	num2Str := os.Args[3]

	// Convertir strings a números
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Printf("Error: '%s' no es un número válido\n", num1Str)
		os.Exit(1)
	}

	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Printf("Error: '%s' no es un número válido\n", num2Str)
		os.Exit(1)
	}

	// Procesar la operación
	resultado, err := calculadora.ProcesarOperacion(operador, num1, num2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Mostrar el resultado
	fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operador, num2, resultado)
}
