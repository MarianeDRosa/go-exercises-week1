/**
 * * exercise 02: even or odd
 * this function checks if a number is even or odd and determines its sign.
 * @param {number} inputNumber - the number to be analyzed.
 * @return {void} 
 */
function analyzeNumber(inputNumber) {
  // we use the modulo operator (%) to check paraty
  const isEven = inputNumber % 2 === 0 ? "even" : "odd";
 
  /** @type {string} */
    let sign;

    if (inputNumber > 0) {
        sign = "positive";
    } else if (inputNumber < 0) {
        sign = "negative";
    } else {
        sign = "zero";
    }       

    // displaying the result using template literals
    console.log(`Number ${inputNumber} Parity: ${isEven} Sign: ${sign}.`);
}

// testing the function with different values
analyzeNumber(10); 
analyzeNumber(-7);
analyzeNumber(0);

// Exporting for Jest
module.exports = { analyzeNumber };