// @ts-nocheck
const { countWords } = require('./wordcount');

describe('Word Counter', () => {
  test('should count word occurrences correctly', () => {
    const text = "JS is great and JS is fun";
    const result = countWords(text);
    
    expect(result['js']).toBe(2);
    expect(result['great']).toBe(1);
  });

  test('should handle empty strings', () => {
    const result = countWords("");
    expect(result).toEqual({});
  });
});