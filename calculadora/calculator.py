import logging

logger = logging.getLogger(__name__)


def suma(a: float, b: float) -> float:
    result = a + b
    logger.debug("suma: %f + %f = %f", a, b, result)
    return result


def resta(a: float, b: float) -> float:
    result = a - b
    logger.debug("resta: %f - %f = %f", a, b, result)
    return result


def procesar_operacion(operador: str, a: float, b: float) -> float:
    if operador == "+":
        return suma(a, b)
    if operador == "-":
        return resta(a, b)
    raise ValueError(f"operador desconocido: {operador}")
