/*
Package main demonstrates the fundamental concepts of variable scoping in Go.

Scope Types Demonstrated:
1. Package Scope:
	- Variables declared outside any function
	- Accessible throughout the entire package
	- Declared using 'var' keyword
	- go mod init learnPackageImport -> to create a module named learnPackageImport

2. Function/Local Scope:
	- Variables declared inside a function
	- Only accessible within that function
	- Usually declared using short declaration ':='

3. Block Scope:
	- Variables declared inside a code block (if, for, etc.)
	- Only accessible within that specific block
	- Limited lifetime to the block's execution

This is a theoretical demonstration of Go's lexical scoping rules.
*/
package main

import (
	"fmt"
	"learnPackageImport/mathLib"
)

// Package scope variables
var (
	 globalVar1 = "I am package scoped"
	 globalVar2 = 100
)

func main() {
	 // Function scope variable
	 localVar := "I am function scoped"

	 fmt.Println("Package scope access:", globalVar1)
	 fmt.Println("Function scope access:", localVar)

	 // Demonstrating block scope
	 {
		  blockVar := "I am block scoped"
		  fmt.Println("Block scope access:", blockVar)
	 }
	 // blockVar is not accessible here

	 if true {
		  conditionalVar := "I am also block scoped"
		  fmt.Println("Conditional block scope:", conditionalVar)
	 }
	 // conditionalVar is not accessible here

	// MathLib function calls -> Package scope
	 fmt.Println("MathLib Add:", mathLib.Add(2, 3))
	 fmt.Println("MathLib Subtract:", mathLib.Subtract(5, 2))
	 fmt.Println("MathLib Multiply:", mathLib.Multiply(3, 4))
	 fmt.Println("MathLib Divide:", mathLib.Divide(10, 2))
}