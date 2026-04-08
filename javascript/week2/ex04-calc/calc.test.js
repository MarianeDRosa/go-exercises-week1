// @ts-nocheck
const { calculate } = require('./calc');

describe('JS Calculator', () => {
  test('should add numbers correctly', () => {
    expect(calculate("10 + 20")).toBe(30);
  });

  test('should throw error on division by zero', () => {
    expect(() => calculate("10 / 0")).toThrow("Divisão por zero não permitida");
  });

  test('should handle invalid format', () => {
    expect(() => calculate("5 +")).toThrow("Formato inválido!");
  });
});