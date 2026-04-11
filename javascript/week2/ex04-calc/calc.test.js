// @ts-nocheck
const { calculate } = require('./calc');

describe('Calculator - Full Test Coverage', () => {
  // Happy Paths
  test('should handle addition: "10 + 5" -> 15', () => {
    expect(calculate("10 + 5")).toBe(15);
  });

  test('should handle subtraction: "20 - 8" -> 12', () => {
    expect(calculate("20 - 8")).toBe(12);
  });

  test('should handle multiplication: "4 * 3" -> 12', () => {
    expect(calculate("4 * 3")).toBe(12);
  });

  test('should handle division: "10 / 2" -> 5', () => {
    expect(calculate("10 / 2")).toBe(5);
  });

  // Sad Paths (Error Handling)
  test('should error on division by zero', () => {
    expect(() => calculate("10 / 0")).toThrow("Division by zero not allowed");
  });

  test('should error on invalid format (missing parts)', () => {
    expect(() => calculate("5 +")).toThrow("Invalid format");
  });

  test('should error on unknown operator', () => {
    expect(() => calculate("5 % 2")).toThrow("Unknown operator");
  });

  test('should error on non-numeric inputs', () => {
    expect(() => calculate("abc + 2")).toThrow("Invalid numbers");
  });
});