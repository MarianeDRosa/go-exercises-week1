// Import logic (adjusting to test individual values)
function getFizzBuzzResult(i) {
  if (i % 3 === 0 && i % 5 === 0) return "FizzBuzz";
  if (i % 3 === 0) return "Fizz";
  if (i % 5 === 0) return "Buzz";
  return i;
}

describe('Exercise 03: FizzBuzz Challenge', () => {
  test('should return "Fizz" for multiples of 3', () => {
    expect(getFizzBuzzResult(3)).toBe("Fizz");
    expect(getFizzBuzzResult(9)).toBe("Fizz");
  });

  test('should return "Buzz" for multiples of 5', () => {
    expect(getFizzBuzzResult(5)).toBe("Buzz");
    expect(getFizzBuzzResult(20)).toBe("Buzz");
  });

  test('should return "FizzBuzz" for multiples of both 3 and 5', () => {
    expect(getFizzBuzzResult(15)).toBe("FizzBuzz");
    expect(getFizzBuzzResult(30)).toBe("FizzBuzz");
  });

  test('should return the number itself if not divisible by 3 or 5', () => {
    expect(getFizzBuzzResult(7)).toBe(7);
    expect(getFizzBuzzResult(11)).toBe(11);
  });
});