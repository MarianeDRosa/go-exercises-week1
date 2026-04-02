const { greetDeveloper } = require('./01-hello');

describe('Exercise 01: Hello World', () => {
  test('should return the correct greeting message', () => {
    const name = "Mariane";
    const expected = `Hello, Mariane! Your JavaScript environment with ESLint and JSDoc is 100% ready.`;
    
    expect(greetDeveloper(name)).toBe(expected);
  });
});