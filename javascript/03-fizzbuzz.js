/**
 * Exercise 03: FizzBuzz Challenge
 * Iterates from 1 to 100.
 * Prints "Fizz" for multiples of 3, "Buzz" for multiples of 5,
 * and "FizzBuzz" for multiples of both.
 * @returns {void}
 */
function runFizzBuzz() {
  for (let i = 1; i <= 100; i++) {
    // We check both conditions first using the AND (&&) operator
    if (i % 3 === 0 && i % 5 === 0) {
      console.log("FizzBuzz");
    } else if (i % 3 === 0) {
      console.log("Fizz");
    } else if (i % 5 === 0) {
      console.log("Buzz");
    } else {
      console.log(i);
    }
  }
}

// Executing the challenge
runFizzBuzz();

module.exports = { runFizzBuzz };