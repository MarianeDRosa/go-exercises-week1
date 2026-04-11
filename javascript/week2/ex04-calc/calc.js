// @ts-nocheck
function calculate(expression) {
  const parts = expression.trim().split(/\s+/);
  
  if (parts.length !== 3) {
    throw new Error("Invalid format"); // Ajustado para bater com o teste
  }

  const n1 = parseFloat(parts[0]);
  const operator = parts[1];
  const n2 = parseFloat(parts[2]);

  if (isNaN(n1) || isNaN(n2)) {
    throw new Error("Invalid numbers"); // Ajustado para bater com o teste
  }

  switch (operator) {
    case "+": return n1 + n2;
    case "-": return n1 - n2;
    case "*": return n1 * n2;
    case "/":
      if (n2 === 0) throw new Error("Division by zero not allowed");
      return n1 / n2;
    default:
      throw new Error("Unknown operator"); // Ajustado para bater com o teste
  }
}

module.exports = { calculate };