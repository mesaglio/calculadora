function suma(a, b) {
  return a + b;
}

function resta(a, b) {
  return a - b;
}

function procesarOperacion(operador, a, b) {
  if (operador === '+') {
    return suma(a, b);
  }

  if (operador === '-') {
    return resta(a, b);
  }

  throw new Error(`operador desconocido: ${operador}`);
}

module.exports = { suma, resta, procesarOperacion };
