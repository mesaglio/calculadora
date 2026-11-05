const test = require('node:test');
const assert = require('node:assert/strict');
const { spawnSync } = require('node:child_process');
const path = require('node:path');

const mainPath = path.join(__dirname, 'main.js');

function runCalculator(args) {
  return spawnSync(process.execPath, [mainPath, ...args], {
    encoding: 'utf8',
  });
}

test('imprime el resultado de una operación válida', () => {
  const result = runCalculator(['2', '+', '3']);

  assert.equal(result.status, 0);
  assert.equal(result.stdout, '2.00 + 3.00 = 5.00\n');
  assert.equal(result.stderr, '');
});

test('rechaza una cantidad incorrecta de argumentos', () => {
  const result = runCalculator(['2', '+']);

  assert.equal(result.status, 1);
  assert.match(result.stderr, /^Uso: .*main\.js <num1> <operador> <num2>\n$/);
});

test('rechaza argumentos que no son números', () => {
  const result = runCalculator(['abc', '+', '3']);

  assert.equal(result.status, 1);
  assert.match(result.stderr, /^Error al parsear argumentos:/);
});

test('rechaza operadores desconocidos', () => {
  const result = runCalculator(['2', '*', '3']);

  assert.equal(result.status, 1);
  assert.equal(result.stderr, 'Error: operador desconocido: *\n');
});
