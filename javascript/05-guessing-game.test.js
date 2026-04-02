const { checkGuess } = require('./05-guessing-game');

describe('Exercise 05: Guessing Game Logic', () => {
  test('should return "LOW" when guess is smaller', () => {
    expect(checkGuess(2, 8)).toBe("LOW");
  });

  test('should return "HIGH" when guess is larger', () => {
    expect(checkGuess(9, 4)).toBe("HIGH");
  });

  test('should return "WIN" when numbers match', () => {
    expect(checkGuess(5, 5)).toBe("WIN");
  });
});