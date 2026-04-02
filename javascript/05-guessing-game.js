/**
 * Exercise 05: Guessing Game
 * Simulates a simple guessing game where the computer picks a random number.
 * @param {number} playerGuess - The number chosen by the player.
 * @param {number} secretValue - The hidden target number.
 * @returns {void}
 */
function playGuessingGame(playerGuess, secretValue) {
  if (playerGuess < secretValue) {
    console.log(`Your guess (${playerGuess}) was too LOW! The secret number was ${secretValue}.`);
  } else if (playerGuess > secretValue) {
    console.log(`Your guess (${playerGuess}) was too HIGH! The secret number was ${secretValue}.`);
  } else {
    console.log("🎉 CONGRATULATIONS! You nailed it! You guessed the right number!");
  }
}

// Generating a random number between 1 and 10
// Math.random() gives a decimal between 0 and 1
// Math.floor() rounds it down to the nearest whole number
const secretTarget = Math.floor(Math.random() * 10) + 1;
const myGuess = 5;

// Running the game
playGuessingGame(myGuess, secretTarget);

/**
 * Logic to check the player's guess (Testable version)
 * @param {number} guess - The player's number
 * @param {number} secret - The target number
 * @returns {string} Result keyword for testing
 */
function checkGuess(guess, secret) {
  if (guess < secret) return "LOW";
  if (guess > secret) return "HIGH";
  return "WIN";
}

// Isso é o que permite que o Jest "enxergue" seu código
module.exports = { checkGuess };