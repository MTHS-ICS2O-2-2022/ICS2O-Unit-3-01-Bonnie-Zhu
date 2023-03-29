// Created by: Bonnie Zhu
// Created on: Mar 2023
// This file contains the JS functions for index.html

'use strict'
/**
* This function calculates your pay depending on the amount of hours you work (including tax).
*/
function calculate() {
  // input
  const Base1 = parseInt(document.getElementById('Base1').value)
  const Base2 = parseInt(document.getElementById('Base2').value)
  const Height = parseInt(document.getElementById('Height').value)

  // process
  const Area = (Base1 + Base2) /2 * Height

  // output
  document.getElementById('Area').innerHTML = 'Area is: ' + Area + 'mm²'
}

