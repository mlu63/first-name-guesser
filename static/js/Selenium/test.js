// Selenium test to account for the case where the application breaks
// in the UI (when submitting strings that lead with escape characters)
var webdriverio = require('webdriverio');

var backslashInput = "\\";
var backslashPrefixInput = "\\Mr";
var backslashFirstNameInput = "Mr \\FirstName LastName";
var backslashLastNameInput = "Mr FirstName \\LastName";
var backslashFirstName = "\\FirstName";
var firstName = "FirstName";
var testName = "Test";

var options = {
    desiredCapabilities: {
        browserName: 'chrome'
    }
};

// single "test" run, as I had difficulty integrating mocha with node.js
webdriverio
  .remote(options)
  .init()
  .url('http://first-name-guesser.herokuapp.com/').then(function() {
      console.log('Launched First Name Guesser');
   })

  // Test the scenario where valid input is submitted
  .setValue('[name="name"]', testName)
  .pause(1000)
  .submitForm('#guesser')
  .pause(1000)
  .getText('span').then(function(value) {
      console.log('Input: ' + testName + '\nExpected: ' + testName + '\nReceived: ' + value + '\n-----');
   })
  .pause(1000)

  // Test the scenario where backslash-leading prefix is submitted
  .setValue('[name="name"]', backslashInput)
  .pause(1000)
  .submitForm('#guesser')
  .pause(1000)
  .getText('span').then(function(value) {
      console.log('Input: ' + backslashInput + '\nExpected: ' + backslashInput + '\nReceived: ' + value + '\n-----');
   })
  .pause(1000)

  // Test the scenario where backslash is submitted
  .setValue('[name="name"]', backslashPrefixInput)
  .pause(1000)
  .submitForm('#guesser')
  .pause(1000)
  .getText('span').then(function(value) {
      console.log('Input: ' + backslashPrefixInput + '\nExpected: ' + backslashPrefixInput + '\nReceived: ' + value + '\n-----');
   })
  .pause(1000)


  // Test the scenario where a backslash-leading first name is submitted
  .setValue('[name="name"]', backslashFirstNameInput)
  .pause(1000)
  .submitForm('#guesser')
  .pause(1000)
  .getText('span').then(function(value) {
      console.log('Input: ' + backslashFirstNameInput + '\nExpected:' + backslashFirstName + '\nReceived: ' + value + '\n-----');
   })
  .pause(1000)

  // Test the scenario where a backslash-leading last name is submitted
  .setValue('[name="name"]', backslashLastNameInput)
  .pause(1000)
  .submitForm('#guesser')
  .pause(1000)
  .getText('span').then(function(value) {
      console.log('Input: ' + backslashLastNameInput + '\nExpected:' + firstName + '\nReceived: ' + value + '\n-----');
   })
  .pause(1000)

  .end();