// @ts-nocheck
function calculate(expression) {
  const parts = expression.trim().split(/\s+/);
  
  if (parts.length !== 3) {
    throw new Error("Formato inválido! Use: '3 + 5'");
  }

  const n1 = parseFloat(parts[0]);
  const operator = parts[1];
  const n2 = parseFloat(parts[2]);

  if (isNaN(n1) || isNaN(n2)) {
    throw new Error("Números inválidos");
  }

  switch (operator) {
    case "+": return n1 + n2;
    case "-": return n1 - n2;
    case "*": return n1 * n2;
    case "/":
      if (n2 === 0) throw new Error("Divisão por zero não permitida");
      return n1 / n2;
    default:
      throw new Error(`Operador desconhecido: ${operator}`);
  }
}

module.exports = { calculate };