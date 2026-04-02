/**
 * Exercise 04: Temperature Converter
 * Converts Celsius to Fahrenheit.
 * Formula: (Celsius * 9/5) + 32
 * @param {number} celsius - Temperature in Celsius degrees.
 * @returns {number} Temperature in Fahrenheit.
 */
function convertToFahrenheit(celsius) {
  const fahrenheit = (celsius * 9) / 5 + 32;
  return fahrenheit;
}

/**
 * Displays the converted result in the console.
 * @param {number} temperatureC - Temperature in Celsius to be displayed.
 * @returns {void}
 */
function showConversion(temperatureC) {
  const resultF = convertToFahrenheit(temperatureC);
  
  // We use .toFixed(1) to show only one decimal place
  console.log(`${temperatureC.toFixed(1)}°C is equal to ${resultF.toFixed(1)}°F`);
}

// Testing with different values
showConversion(25);    // Room temperature
showConversion(0);     // Freezing point
showConversion(37.5);  // Fever temperature

module.exports = { convertToFahrenheit };