/**
 * Exercise 01: Hello World
 * @param {string} developerName
 * @returns {string} The full greeting message
 */
function greetDeveloper(developerName) {
  return `Hello, ${developerName}! Your JavaScript environment with ESLint and JSDoc is 100% ready.`;
}

// Para continuar funcionando no terminal sozinho:
console.log(greetDeveloper("WORLD"));

// Exportação para o teste
module.exports = { greetDeveloper };
