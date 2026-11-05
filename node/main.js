const { procesarOperacion } = require('./calculadora/calculator');

function main() {
  if (process.argv.length !== 5) {
    console.error(`Uso: ${process.argv[1]} <num1> <operador> <num2>`);
    process.exitCode = 1;
    return;
  }

  const a = Number(process.argv[2]);
  const operador = process.argv[3];
  const b = Number(process.argv[4]);

  if (Number.isNaN(a) || Number.isNaN(b)) {
    console.error('Error al parsear argumentos: se esperaban dos números');
    process.exitCode = 1;
    return;
  }

  try {
    const resultado = procesarOperacion(operador, a, b);
    console.log(`${a.toFixed(2)} ${operador} ${b.toFixed(2)} = ${resultado.toFixed(2)}`);
  } catch (error) {
    console.error(`Error: ${error.message}`);
    process.exitCode = 1;
  }
}

main();
