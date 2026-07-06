import sys
import logging
from calculadora.calculator import procesar_operacion

logging.basicConfig(level=logging.INFO)


def main():
    if len(sys.argv) != 4:
        print(f"Uso: {sys.argv[0]} <num1> <operador> <num2>", file=sys.stderr)
        sys.exit(1)

    try:
        a = float(sys.argv[1])
        operador = sys.argv[2]
        b = float(sys.argv[3])
    except ValueError as e:
        print(f"Error al parsear argumentos: {e}", file=sys.stderr)
        sys.exit(1)

    try:
        resultado = procesar_operacion(operador, a, b)
    except ValueError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

    print(f"{a:.2f} {operador} {b:.2f} = {resultado:.2f}")


if __name__ == "__main__":
    main()
