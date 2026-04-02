const { analyzeNumber } = require('./02-even-odd');

describe('Exercise 02: Even or Odd and Sign Analysis', () => {
  
  test('should correctly identify a positive even number', () => {
    const input = 10;
    const isEven = input % 2 === 0;
    const isPositive = input > 0;
    
    expect(isEven).toBe(true);
    expect(isPositive).toBe(true);
  });

  test('should correctly identify a negative odd number', () => {
    const input = -7;
    const isEven = input % 2 === 0;
    const isNegative = input < 0;
    
    expect(isEven).toBe(false);
    expect(isNegative).toBe(true);
  });

  test('should correctly identify zero', () => {
    const input = 0;
    const isZero = input === 0;
    
    expect(isZero).toBe(true);
  });
});