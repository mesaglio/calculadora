import unittest
from calculadora.calculator import suma, resta, procesar_operacion


class TestSuma(unittest.TestCase):
    cases = [
        (2.0, 3.0, 5.0),
        (-1.0, -1.0, -2.0),
        (-1.0, 1.0, 0.0),
        (0.0, 0.0, 0.0),
        (1.5, 2.5, 4.0),
    ]

    def test_suma(self):
        for a, b, expected in self.cases:
            with self.subTest(a=a, b=b):
                self.assertEqual(suma(a, b), expected)


class TestResta(unittest.TestCase):
    cases = [
        (5.0, 3.0, 2.0),
        (-1.0, -1.0, 0.0),
        (-1.0, 1.0, -2.0),
        (0.0, 0.0, 0.0),
        (4.0, 1.5, 2.5),
    ]

    def test_resta(self):
        for a, b, expected in self.cases:
            with self.subTest(a=a, b=b):
                self.assertEqual(resta(a, b), expected)


class TestProcesarOperacion(unittest.TestCase):
    def test_suma(self):
        self.assertEqual(procesar_operacion("+", 2.0, 3.0), 5.0)

    def test_resta(self):
        self.assertEqual(procesar_operacion("-", 5.0, 3.0), 2.0)

    def test_operador_invalido(self):
        with self.assertRaises(ValueError):
            procesar_operacion("*", 2.0, 3.0)

    def test_operador_vacio(self):
        with self.assertRaises(ValueError):
            procesar_operacion("", 2.0, 3.0)


if __name__ == "__main__":
    unittest.main()
