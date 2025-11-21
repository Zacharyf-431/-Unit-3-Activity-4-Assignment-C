/**
 * @author: Zachary Fowler
 * @version: 1.0.0
 * @date: 2025-11-20
 * @fileoverview: This file calculates interest 
 */

// INPUT
let userName: string | null = prompt("Enter your name: ");
let ageInput: string | null = prompt("Enter your age: ");
let studentInput: string | null = prompt("Are you a student (true/false)? ");

// Checks for misinputs
if (userName !== null && ageInput !== null && studentInput !== null) {

  // PROCESS
  let age: number = Number(ageInput);
  let student: boolean = studentInput.toLowerCase() === "true";

  // OUTPUT
  // outputs based of of comparisons
  if (student && age >= 13 && age <= 19) {
    console.log("Student " + userName + " is a teenager.");
  } else if (student && age >= 5 && age <= 12) {
    console.log("Student " + userName + " is a child.");
  } else if (!student && age >= 20 && age <= 30) {
    console.log(userName + " is a young adult.");
  } else {
    console.log(userName + " is in a different life stage.");
  }
} else {
  console.log("Invalid input");
}

console.log("Done.");