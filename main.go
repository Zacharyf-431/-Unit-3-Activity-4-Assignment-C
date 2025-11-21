/*
	* Author: Zachary Fowler
	* Version: 1.0.0
	* Date: 2025-11-20
	* This file determines the age group of someone  
	*/
	package main

import (
	"fmt"
)

func main() {
	// INPUT
	var userName string
	var age int
	var studentInput string
	var student bool

	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)

	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Print("Are you a student (true/false)? ")
	fmt.Scanln(&studentInput)

	// PROCESS
	if studentInput == "true" {
		student = true
	} else {
		student = false
	}

	// DECISION
	if student && age >= 13 && age <= 19 {
		fmt.Println("Student", userName, "is a teenager.")
	} else if student && age >= 5 && age <= 12 {
		fmt.Println("Student", userName, "is a child.")
	} else if !student && age >= 20 && age <= 30 {
		fmt.Println(userName, "is a young adult.")
	} else {
		fmt.Println(userName, "is in a different life stage.")
	}

	fmt.Println("Done.")
}