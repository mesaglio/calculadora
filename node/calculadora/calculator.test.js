const test = require('node:test');
const assert = require('node:assert/strict');

const { suma, resta, procesarOperacion } = require('./calculator');

test('suma dos números', () => {
  const cases = [
    [2, 3, 5],
    [-1, -1, -2],
    [-1, 1, 0],
    [0, 0, 0],
    [1.5, 2.5, 4],
  ];

  for (const [a, b, expected] of cases) {
    assert.equal(suma(a, b), expected);
  }
});

test('resta dos números', () => {
  const cases = [
    [5, 3, 2],
    [-1, -1, 0],
    [-1, 1, -2],
    [0, 0, 0],
    [4, 1.5, 2.5],
  ];

  for (const [a, b, expected] of cases) {
    assert.equal(resta(a, b), expected);
  }
});

test('procesa la suma', () => {
  assert.equal(procesarOperacion('+', 2, 3), 5);
});

test('procesa la resta', () => {
  assert.equal(procesarOperacion('-', 5, 3), 2);
});

test('rechaza un operador inválido', () => {
  assert.throws(
    () => procesarOperacion('*', 2, 3),
    { message: 'operador desconocido: *' },
  );
});

test('rechaza un operador vacío', () => {
  assert.throws(
    () => procesarOperacion('', 2, 3),
    { message: 'operador desconocido: ' },
  );
});
