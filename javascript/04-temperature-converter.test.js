const { convertToFahrenheit } = require('./04-temperature-converter');

describe('Exercise 04: Temperature Converter', () => {
  test('should convert 0°C to 32°F (freezing point)', () => {
    expect(convertToFahrenheit(0)).toBe(32);
  });

  test('should convert 100°C to 212°F (boiling point)', () => {
    expect(convertToFahrenheit(100)).toBe(212);
  });

  test('should convert 25°C to 77°F (room temperature)', () => {
    expect(convertToFahrenheit(25)).toBe(77);
  });

  test('should convert 37.5°C to 99.5°F (body temperature)', () => {
    // Como QA, testamos números com vírgula (ponto)
    expect(convertToFahrenheit(37.5)).toBe(99.5);
  });
});